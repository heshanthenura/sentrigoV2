package config

import (
	"github.com/cilium/ebpf/link"
	"github.com/heshanthenura/sentrigov2/internal/types"
)

var CaptureConfig types.CaptureConfig
var XDPLink link.Link
