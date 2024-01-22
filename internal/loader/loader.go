package loader

import (
	"sync"
	"time"
)

type Plugin interface {
	GetName() string
}

type PluginLoader struct {
	LoadedPlugins map[string]Plugin
	TimeService   TimeProvider
	mu            sync.Mutex
}

type TimeProvider interface {
	GetLoadTime() time.Time
}

type TimeLoad struct{}

func (rtp *TimeLoad) GetLoadTime() time.Time {
	return time.Now()
}
