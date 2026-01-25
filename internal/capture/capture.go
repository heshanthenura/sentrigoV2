package capture

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type IP struct {
	Source string
	Count  int
}

var IPs []IP

func StartCapture() {
	fmt.Println("capture sttarts")
	device := "docker0"
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
		printPacketInfo(packet)
	}
}

func printPacketInfo(packet gopacket.Packet) {

	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		AppendIP(ipLayer.(*layers.IPv4).SrcIP.String())
	}

	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		AppendIP(tcpLayer.(*layers.TCP).SrcPort.String())
	}
	PrintIP()
	fmt.Println("-----")
}

func AppendIP(ip string) {
	for i, existingIP := range IPs {
		if existingIP.Source == ip {
			IPs[i].Count++
			return
		}
	}
	IPs = append(IPs, IP{Source: ip, Count: 1})
}
func PrintIP() {
	for i, ip := range IPs {
		fmt.Printf("%d: %s - %d\n", i+1, ip.Source, ip.Count)
	}
}
