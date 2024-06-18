package ginutil

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	// f9f1185a93874f5396d3cd1cf27d4f47
	for i := 0; i < 20; i++ {
		traceId := uuid.New().String()
		fmt.Println(i, "traceId:", strings.ReplaceAll(traceId, "-", ""))
	}
}

func TestRunHttpServer(t *testing.T) {
	server := NewHttpServer(HttpServerConfig{
		Port:                 8800,
		CommonMiddlewareList: nil,
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
