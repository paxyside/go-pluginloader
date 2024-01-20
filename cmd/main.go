package main

import (
	"PluginsLoader/api/server"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/plugins/list", server.GetPluginsList)
	r.POST("/plugins/load", server.PostLoadPlugin)
	r.POST("/plugins/unload", server.PostUnloadPlugin)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
