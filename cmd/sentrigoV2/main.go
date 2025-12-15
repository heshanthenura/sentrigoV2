package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
	iface := "docker0"
	snaplen := int32(96)

	handle, err := pcap.OpenLive(iface, snaplen, false, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packetSource.DecodeOptions = gopacket.DecodeOptions{
		Lazy:   true,
		NoCopy: true,
	}

	fmt.Println("Listening on", iface)

	var count uint64

	for {
		packet, err := packetSource.NextPacket()
		if err != nil {
			continue
		}
		count++
		if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
			ip := ipLayer.(*layers.IPv4)
			if count%10000 == 0 {
				fmt.Printf(
					"Packets: %d | Src: %s | Dst: %s | Proto: %s\n",
					count,
					ip.SrcIP,
					ip.DstIP,
					ip.Protocol,
				)
			}
		}
	}
}
