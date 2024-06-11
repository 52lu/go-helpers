package ginutil

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

func TestRunHttpServer(t *testing.T) {
	server := NewHttpServer(HttpServerConfig{
		Port:           8800,
		MiddlewareList: nil,
		RouterFunc: []RouterRegisterFunc{
			addRoute,
		},
	})
	server.Start()
}

func addRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})
}
