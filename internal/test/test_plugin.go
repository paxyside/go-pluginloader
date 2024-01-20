package main

type TestPlugin struct {
}

func (p TestPlugin) GetName() string {
	return "TestPlugin"
}

var PluginInstance = TestPlugin{}
