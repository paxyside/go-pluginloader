package main

import (
	"PluginsLoader/internal/loader"
	"time"
)

type Plugin1 struct {
	loadTime time.Time
}

func NewPlugin1(timeProvider loader.TimeLoad) *Plugin1 {
	return &Plugin1{
		loadTime: timeProvider.GetLoadTime(),
	}
}

func (p *Plugin1) GetName() string {
	return "Plugin1"
}

var PluginInstance Plugin1
