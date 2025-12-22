package capture

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func logPackets(in <-chan gopacket.Packet) {
	var count uint64
	for packet := range in {
		count++

		if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
			ip := ipLayer.(*layers.IPv4)
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
