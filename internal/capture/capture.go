package capture

import (
	"context"
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/heshanthenura/sentrigov2/internal/types"
	"github.com/heshanthenura/sentrigov2/internal/utils"
)

func StartCapture(ctx context.Context, captureConfig types.CaptureConfig) error {
	log.Printf("starting capture on interface: %s", captureConfig.IfaceName)

	handle, err := pcap.OpenLive(
		captureConfig.IfaceName,
		captureConfig.SnapshotLen,
		captureConfig.Promiscuous,
		captureConfig.Timeout,
	)
	if err != nil {
		log.Printf("failed to open interface %s: %v", captureConfig.IfaceName, err)
		return fmt.Errorf("pcap.OpenLive: %w", err)
	}

	shutdownDone := make(chan struct{})
	go func() {
		<-ctx.Done()
		log.Println("closing packet source...")
		handle.Close()
		close(shutdownDone)
	}()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	for {
		select {
		case <-ctx.Done():
			log.Println("capture stopped")
			<-shutdownDone
			return nil

		case packet, ok := <-packetSource.Packets():
			if !ok {
				log.Println("packet source closed")
				return nil
			}
			func() {
				defer func() {
					if r := recover(); r != nil {
						log.Printf("recovered from panic in ProcessPacket: %v", r)
					}
				}()
				utils.ProcessPacket(packet)
			}()
		}
	}
}
