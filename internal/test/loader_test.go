package main

import (
	"PluginsLoader/internal/loader"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadPlugin(t *testing.T) {
	pl := loader.NewPluginLoader()

	err := pl.LoadPlugin("test_plugin.so")
	assert.Nil(t, err)

	fmt.Printf("Loaded plugins: %v\n", pl.LoadedPlugins) // Добавь эту строку

	err = pl.LoadPlugin("test_plugin.so")
	assert.Nil(t, err)
}

func TestUploadPlugin(t *testing.T) {
	pl := loader.NewPluginLoader()

	err := pl.LoadPlugin("test_plugin.so")
	assert.Nil(t, err)

	pl.UnloadPlugin("TestPlugin")
	assert.Empty(t, pl.LoadedPlugins)

}

// go test ./internal/test  ok      PluginsLoader/internal/test     0.004s
