package pcap

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type Capture struct {
	handle *pcap.Handle
	source *gopacket.PacketSource
}

func NewCapture(device string) (*Capture, error) {
	handle, err := pcap.OpenLive(
		device,
		65535,
		true,
		pcap.BlockForever,
	)
	if err != nil {
		return nil, err
	}

	source := gopacket.NewPacketSource(
		handle,
		handle.LinkType(),
	)

	return &Capture{
		handle: handle,
		source: source,
	}, nil
}

func (c *Capture) Packets() chan gopacket.Packet {
	return c.source.Packets()
}

func (c *Capture) Close() {
	c.handle.Close()
}
