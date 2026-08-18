package engine

import (
	"fmt"
	"sentrigoV2/engine/internal/pcap"
)

func StartEngine() error {
	capture, err := pcap.NewCapture("wlp4s0")
	if err != nil {
		return err
	}
	defer capture.Close()

	for packet := range capture.Packets() {
		fmt.Println(packet)
	}

	return nil
}
