package config

import (
	"sync"
	"time"

	"github.com/cilium/ebpf/link"
	"github.com/heshanthenura/sentrigov2/internal/types"
)

var XDPLink link.Link

type Config struct {
	mu sync.RWMutex

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
		globalConfig = &Config{
			IsCapturing: false,
			BlockedIPs:  make(map[string]bool),
			StartTime:   time.Now(),
		}

	})
	return globalConfig
}

func UpdateIsCapturing(isCapturing bool) {
	cfg := GetConfig()
	cfg.mu.Lock()
	cfg.IsCapturing = isCapturing
	cfg.mu.Unlock()
}

func IsCapturing() bool {
	cfg := GetConfig()
	cfg.mu.RLock()
	defer cfg.mu.RUnlock()
	return cfg.IsCapturing
}
