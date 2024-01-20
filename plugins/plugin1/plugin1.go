package main

type Plugin1 struct {
}

func (p Plugin1) GetName() string {
	return "Plugin1"
}

var PluginInstance = Plugin1{}
