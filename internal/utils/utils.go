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

func LinkEBPF(captureConfig types.CaptureConfig) *ebpf.Map {
	spec, err := ebpf.LoadCollectionSpec("ebpf/interceptor.o")

	if err != nil {
		log.Fatalf("load spec: %v", err)
	}
	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		log.Fatalf("new collection: %v", err)
	}

	prog := coll.Programs["block_ips"]
	if prog == nil {
		log.Fatalf("program block_ips not found")
	}

	blocked := coll.Maps["blocked_ips"]
	if blocked == nil {
		log.Fatalf("map blocked_ips not found")
	}

	blockedIP := net.ParseIP("172.17.0.2").To4()
	if blockedIP != nil {
		var ipKey uint32 = binary.BigEndian.Uint32(blockedIP)
		var value uint8 = 1
		fmt.Printf("%d\n", ipKey)
		blocked.Put(ipKey, value)
	}

	iface, err := net.InterfaceByName(captureConfig.IfaceName)
	if err != nil {
		log.Fatalf("iface %s: %v", captureConfig.IfaceName, err)
	}
	lk, err := link.AttachXDP(link.XDPOptions{
		Program:   prog,
		Interface: iface.Index,
	})
	if err != nil {
		log.Fatalf("attach xdp: %v", err)
	}
	config.XDPLink = lk
	log.Printf("XDP attached to %s", iface.Name)
	return blocked
}

func ProcessPacket(packet gopacket.Packet) {

	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println(ipLayer.(*layers.IPv4).SrcIP.String())
	}

}
