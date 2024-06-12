package ginutil

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type httpServer struct {
	config HttpServerConfig
}

type HttpServerConfig struct {
	Port              int                      // 端口
	RunMode           string                   // 运行默认;Debug 模式 (gin.DebugMode);
	DelayExistSecond  int                      // 延迟多久退出，用于平滑重启
	MiddlewareList    []gin.HandlerFunc        // 中间件
	RouterFunc        []RouterRegisterFunc     // 路由函数
	DefaultMiddleware *DefaultMiddlewareSwitch // 是否开启默认中间件
}
type DefaultMiddlewareSwitch struct {
	EnableRecoveryMiddleware bool // 捕获panic
}

type RouterRegisterFunc func(*gin.Engine)

/*
* @Description: 实例化服务
* @Author: LiuQHui
* @Param conf
* @Return *httpServer
* @Date 2024-06-11 15:46:47
 */
func NewHttpServer(conf HttpServerConfig) *httpServer {
	return &httpServer{
		config: conf,
	}
}

/*
* @Description: 启动服务
* @Author: LiuQHui
* @Receiver h
* @Date 2024-06-11 15:44:43
 */
func (h *httpServer) Start() {
	// 设置运行模式
	gin.SetMode(gin.DebugMode)
	if h.config.RunMode != "" {
		gin.SetMode(h.config.RunMode)
	}
	// 创建从操作系统中断信号的上下文
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	// 创建gin服务实例
	engine := gin.Default()
	// 追加信息到上下文中间件
	engine.Use(AdditionalMiddleware)
	// 判断是否开启默认中间件
	h.enableDefaultMiddle(engine)
	// 注册路由
	for _, registerFunc := range h.config.RouterFunc {
		registerFunc(engine)
	}
	// 注册中间件
	if len(h.config.MiddlewareList) > 0 {
		engine.Use(h.config.MiddlewareList...)
	}
	// 创建服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", h.config.Port),
		Handler: engine,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	// 在 goroutine 中初始化服务器，这样它就不会阻止下面的正常关闭处理
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal. 监听中断信号
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	// 恢复中断信号的默认行为并通知用户关闭
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	// 上下文用于通知服务器有 5 秒的时间来完成当前正在处理的请求;平滑重启防止直接终端服务

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}

/*
* @Description: 判断是否开启默认中间件
* @Author: LiuQHui
* @Receiver h
* @Param engine
* @Date 2024-06-12 12:29:15
 */
func (h *httpServer) enableDefaultMiddle(engine *gin.Engine) {
	// 是否添加默认中间件
	if h.config.DefaultMiddleware != nil {
		defaultMiddlewareSwitch := h.config.DefaultMiddleware
		// 捕获异常中间件
		if defaultMiddlewareSwitch.EnableRecoveryMiddleware {
			engine.Use(RecoveryMiddleware)
		}
	}
}
