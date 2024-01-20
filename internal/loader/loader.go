package loader

import (
	"sync"
)

type Plugin interface {
	GetName() string
}

type PluginLoader struct {
	LoadedPlugins map[string]Plugin
	mu            sync.Mutex
}
