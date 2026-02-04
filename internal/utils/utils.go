package utils

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/heshanthenura/sentrigov2/internal/config"
	"github.com/heshanthenura/sentrigov2/internal/types"
)

func GetAllInterfaces() ([]pcap.Interface, error) {
	devices, err := pcap.FindAllDevs()
	return devices, err
}

func LinkEBPF(captureConfig types.CaptureConfig) (*ebpf.Map, error) {
	spec, err := ebpf.LoadCollectionSpec("ebpf/interceptor.o")
	if err != nil {
		return nil, fmt.Errorf("load spec: %w", err)
	}

	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		return nil, fmt.Errorf("new collection: %w", err)
	}

	prog := coll.Programs["block_ips"]
	if prog == nil {
		return nil, fmt.Errorf("program block_ips not found")
	}

	blocked := coll.Maps["blocked_ips"]
	if blocked == nil {
		return nil, fmt.Errorf("map blocked_ips not found")
	}

	blockedIP := net.ParseIP("172.17.0.2").To4()
	if blockedIP != nil {
		var ipKey uint32 = binary.BigEndian.Uint32(blockedIP)
		var value uint8 = 1
		fmt.Printf("%d\n", ipKey)
		if err := blocked.Put(ipKey, value); err != nil {
			return nil, fmt.Errorf("failed to put blocked IP: %w", err)
		}
	}

	iface, err := net.InterfaceByName(captureConfig.IfaceName)
	if err != nil {
		return nil, fmt.Errorf("iface %s: %w", captureConfig.IfaceName, err)
	}

	lk, err := link.AttachXDP(link.XDPOptions{
		Program:   prog,
		Interface: iface.Index,
	})
	if err != nil {
		return nil, fmt.Errorf("attach xdp: %w", err)
	}

	config.XDPLink = lk
	log.Printf("XDP attached to %s", iface.Name)
	return blocked, nil
}

func ProcessPacket(packet gopacket.Packet) {

	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println(ipLayer.(*layers.IPv4).SrcIP.String())
	}

}
