package server

import (
	"PluginsLoader/api/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPluginsList(t *testing.T) {
	router := gin.Default()
	router.GET("/plugins/list", server.GetPluginsList)

	req, err := http.NewRequest("GET", "/plugins/list", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedJSON := `{"loaded_plugins":[]}`
	assert.Equal(t, expectedJSON, w.Body.String())
}

// go test ./api/test  ok  PluginsLoader/api/test  0.006s
