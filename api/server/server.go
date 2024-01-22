package server

import (
	"PluginsLoader/internal/loader"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var pluginLoader *loader.PluginLoader

func init() {
	timeProvider := &loader.TimeLoad{}
	pluginLoader = loader.NewPluginLoader(timeProvider)
}

func GetPluginsList(c *gin.Context) {
	err := pluginLoader.PluginsList()
	if err != nil {
		return
	}
	loadedPlugins := make([]string, 0, len(pluginLoader.LoadedPlugins))
	for name := range pluginLoader.LoadedPlugins {
		loadedPlugins = append(loadedPlugins, name)
	}
	c.JSON(http.StatusOK, gin.H{"loaded_plugins": loadedPlugins})
}

func PostLoadPlugin(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing 'path' parameter"})
		return
	}

	err := pluginLoader.LoadPlugin(path)
	if err != nil {
		log.Println("Error loading plugin:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to load plugin: %v", err)})
		return
	}

	// Получение времени загрузки плагина
	loadTime := pluginLoader.TimeService.GetLoadTime()

	c.JSON(http.StatusOK, gin.H{
		"message":   "plugin loaded successfully",
		"load_time": loadTime,
	})
}

func PostUnloadPlugin(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing 'name' parameter"})
		return
	}

	err := pluginLoader.UnloadPlugin(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "plugin unloaded successfully"})
}
