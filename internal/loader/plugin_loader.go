package loader

import (
	"log"
	"plugin"
)

func NewPluginLoader() *PluginLoader {
	return &PluginLoader{
		LoadedPlugins: make(map[string]Plugin),
	}
}

func (pm *PluginLoader) withLock(fn func()) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	fn()
}

func (pm *PluginLoader) LoadPlugin(path string) error {
	pm.withLock(func() {
		p, err := plugin.Open(path)
		if err != nil {
			log.Printf("Error opening plugin: %v\n", err)
			return
		}

		sym, err := p.Lookup("PluginInstance")
		if err != nil {
			log.Printf("Error looking up symbol: %v\n", err)
			return
		}

		pluginInstance, ok := sym.(Plugin)
		if !ok {
			log.Println("Unexpected type from module symbol")
			return
		}

		name := pluginInstance.GetName()
		if _, exists := pm.LoadedPlugins[name]; exists {
			log.Printf("Plugin with name '%s' already loaded\n", name)
			return
		}

		pm.LoadedPlugins[name] = pluginInstance
	})
	return nil
}

func (pm *PluginLoader) PluginsList() {
	pm.withLock(func() {
		log.Println("Loaded plugins:")
		for name := range pm.LoadedPlugins {
			log.Println(name)
		}
	})
}

func (pm *PluginLoader) UnloadPlugin(name string) {
	pm.withLock(func() {
		if _, exists := pm.LoadedPlugins[name]; exists {
			delete(pm.LoadedPlugins, name)
			log.Printf("Plugin '%s' unloaded\n", name)
		} else {
			log.Printf("Plugin with name '%s' not found\n\n", name)
		}
	})
}
