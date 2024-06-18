package middlewareutil

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"net/http"
)

/*
* @Description: 跨域中间件
* @Author: LiuQHui
* @Return gin.HandlerFunc
* @Date 2024-06-18 16:18:31
 */
func CORSMiddleware() gin.HandlerFunc {
	// 创建cors实例
	corsInstance := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // 允许所有源访问
		AllowedMethods: []string{
			http.MethodGet, http.MethodPost, http.MethodPut,
			http.MethodDelete, http.MethodOptions}, // 允许的方法
		AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization"}, // 允许的头信息
		ExposedHeaders:   []string{"Content-Length"},                          // 暴露的头信息
		AllowCredentials: true,                                                // 允许发送凭据
		MaxAge:           86400,                                               // 预检请求的缓存时间(秒)
	})
	// 返回一个 Gin 中间件
	return func(c *gin.Context) {
		// 将 CORS 处理器作为处理程序函数附加到请求处理程序链中
		corsInstance.HandlerFunc(c.Writer, c.Request)
		// 继续处理链中的下一个处理器
		c.Next()
	}
}
