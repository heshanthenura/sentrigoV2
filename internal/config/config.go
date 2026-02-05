package config

import (
	"sync"
	"time"

	"github.com/cilium/ebpf/link"
	"github.com/heshanthenura/sentrigov2/internal/types"
)

var XDPLink link.Link

type Config struct {
	CaptureConfig types.CaptureConfig
	XDPLink       link.Link

	IsCapturing bool
	BlockedIPs  map[string]bool
	StartTime   time.Time
}

var (
	globalConfig *Config
	configOnce   sync.Once
)

func GetConfig() *Config {
	configOnce.Do(func() {
		globalConfig = &Config{}
	})

	return globalConfig
}
