package capture

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func openHandle(cfg CaptureConfig) (*pcap.Handle, error) {
	return pcap.OpenLive(
		cfg.Interface,
		cfg.SnapLen,
		cfg.Promisc,
		pcap.BlockForever,
	)
}

func newPacketSource(handle *pcap.Handle) *gopacket.PacketSource {
	ps := gopacket.NewPacketSource(handle, handle.LinkType())
	ps.DecodeOptions = gopacket.DecodeOptions{
		Lazy:   true,
		NoCopy: true,
	}
	return ps
}
