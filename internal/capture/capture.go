package capture

import (
	"fmt"

	"github.com/google/gopacket"
)

type CaptureConfig struct {
	Interface string
	SnapLen   int32
	Promisc   bool
}

func StartCapture(cfg CaptureConfig) error {
	handle, err := openHandle(cfg)
	if err != nil {
		return err
	}
	defer handle.Close()

	packetSource := newPacketSource(handle)

	fmt.Println("Listening on", cfg.Interface)

	packetChan := make(chan gopacket.Packet, 1000)
	go logPackets(packetChan)

	var count uint64

	for {
		packet, err := packetSource.NextPacket()
		if err != nil {
			continue
		}

		count++
		select {
		case packetChan <- packet:
		default:
		}
	}
}
