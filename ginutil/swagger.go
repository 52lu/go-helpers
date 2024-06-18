package ginutil

import (
	"fmt"
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
	swaggerInfo := h.config.SwaggerConfig.SwaggerSpec
	swaggerConfig := h.config.SwaggerConfig
	// 动态设置Swagger
	swaggerInfo.Title = swaggerConfig.Title
	swaggerInfo.Description = swaggerConfig.Description
	swaggerInfo.Version = swaggerConfig.Version
	swaggerInfo.Host = fmt.Sprintf("0.0.0.0:%v", h.config.Port)
	if swaggerConfig.Host != "" {
		swaggerInfo.Host = swaggerConfig.Host
	}
	swaggerInfo.BasePath = swaggerConfig.BasePath
	// Serve Swagger UI
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
