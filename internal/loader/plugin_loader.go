package loader

import (
	"errors"
	"log"
	"plugin"
)

func NewPluginLoader() *PluginLoader {
	return &PluginLoader{
		LoadedPlugins: make(map[string]Plugin),
	}
}

func (pm *PluginLoader) withLock(fn func() error) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	return fn()
}

func (pm *PluginLoader) LoadPlugin(path string) error {
	return pm.withLock(func() error {
		p, err := plugin.Open(path)
		if err != nil {
			log.Printf("Error opening plugin: %v\n", err)
			return err
		}

		sym, err := p.Lookup("PluginInstance")
		if err != nil {
			log.Printf("Error looking up symbol: %v\n", err)
			return err
		}

		pluginInstance, ok := sym.(Plugin)
		if !ok {
			log.Println("Unexpected type from module symbol")
			return errors.New("Unexpected type from module symbol")
		}

		name := pluginInstance.GetName()

		if _, exists := pm.LoadedPlugins[name]; exists {
			log.Printf("Plugin with name '%s' already loaded\n", name)
			return errors.New("plugin with the same name already loaded")
		}

		pm.LoadedPlugins[name] = pluginInstance

		return nil
	})
}

func (pm *PluginLoader) PluginsList() error {
	return pm.withLock(func() error {
		log.Println("Loaded plugins:")
		for name := range pm.LoadedPlugins {
			log.Println(name)
		}
		return nil
	})
}

func (pm *PluginLoader) UnloadPlugin(name string) error {
	return pm.withLock(func() error {
		if _, exists := pm.LoadedPlugins[name]; exists {
			delete(pm.LoadedPlugins, name)
			log.Printf("Plugin '%s' unloaded\n", name)
			return nil
		} else {
			log.Printf("Plugin with name '%s' not found\n\n", name)
			return errors.New("plugin not found")
		}
	})
}
