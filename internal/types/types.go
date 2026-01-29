package types

import "time"

type CaptureConfig struct {
	IfaceName   string        `json:"iface_name"`
	SnapshotLen int32         `json:"snapshot_len"`
	Promiscuous bool          `json:"promiscuous"`
	Timeout     time.Duration `json:"timeout"`
}
