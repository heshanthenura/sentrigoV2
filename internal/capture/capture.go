package capture

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
)

type IP struct {
	Source string
	Count  int
}

var IPs []IP
var blockedMap *ebpf.Map
var xdpLink link.Link
var ifaceName = "docker0"

func StartCapture() {

	blockedMap = LinkEBPF()

	fmt.Println("capture starts")
	device := ifaceName
	snaphotLen := int32(1600)
	promiscuous := false
	timeout := pcap.BlockForever

	handle, err := pcap.OpenLive(device, snaphotLen, promiscuous, timeout)

	if err != nil {
		log.Fatal(err)
	}

	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for packet := range packetSource.Packets() {
		processPacket(packet)
	}
}

func LinkEBPF() *ebpf.Map {
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
		var ipKey uint32 = binary.LittleEndian.Uint32(blockedIP)
		var value uint8 = 1
		fmt.Printf("%d\n", ipKey)
		blocked.Put(ipKey, value)
	}

	iface, err := net.InterfaceByName(ifaceName)
	if err != nil {
		log.Fatalf("iface %s: %v", ifaceName, err)
	}
	lk, err := link.AttachXDP(link.XDPOptions{
		Program:   prog,
		Interface: iface.Index,
	})
	if err != nil {
		log.Fatalf("attach xdp: %v", err)
	}
	xdpLink = lk
	log.Printf("XDP attached to %s", iface.Name)
	return blocked
}

func processPacket(packet gopacket.Packet) {

	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		fmt.Println(ipLayer.(*layers.IPv4).SrcIP.String())
	}

}
