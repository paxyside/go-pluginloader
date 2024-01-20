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
	pluginLoader = loader.NewPluginLoader()
}

func GetPluginsList(c *gin.Context) {
	pluginLoader.PluginsList()
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
	}

	err := pluginLoader.LoadPlugin(path)
	if err != nil {
		log.Println("Error loading plugin:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to load plugin: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "plugin loaded successfully"})

}

func PostUnloadPlugin(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing 'name' parameter"})
	}

	pluginLoader.UnloadPlugin(name)
	c.JSON(http.StatusOK, gin.H{"message": "plugin unloaded successfully"})
}
