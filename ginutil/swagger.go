package ginutil

import (
	"fmt"
	"github.com/52lu/go-helpers/confutil"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/*
* @Description: swagger初始化
* @Author: LiuQHui
* @Receiver h
* @Param engine
* @Date 2024-06-18 16:53:45
 */
func (h *httpServer) swaggerInit(engine *gin.Engine) {
	swaggerInfo := h.config.SwaggerSpec
	// 动态设置Swagger
	swaggerInfo.Title = confutil.GetString("app.name")
	swaggerInfo.Description = confutil.GetString("app.description")
	swaggerInfo.Version = confutil.GetString("app.version")
	swaggerInfo.Host = fmt.Sprintf("0.0.0.0:%v", confutil.GetString("app.port"))
	swaggerInfo.BasePath = ""
	// Serve Swagger UI
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
