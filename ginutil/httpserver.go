package ginutil

import (
	"context"
	"errors"
	"fmt"
	"github.com/52lu/go-helpers/middlewareutil"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag"
	"github.com/thoas/go-funk"
	"log"
	"net/http"
	"os/signal"
	"sort"
	"strings"
	"syscall"
	"time"
)

type httpServer struct {
	config HttpServerConfig
}

type HttpServerConfig struct {
	Port                 int                     // 端口
	RunMode              string                  // 运行默认;Debug 模式 (gin.DebugMode);
	RunEnv               string                  // 运行环境
	OpenFavicon          bool                    // 打开favicon.ico请求
	DelayExistSecond     int                     // 延迟多久退出，用于平滑重启
	CommonMiddlewareList []gin.HandlerFunc       // 公共中间件
	RouterFunc           []RouterRegisterFunc    // 路由函数
	DefaultMiddleware    DefaultMiddlewareSwitch // 是否开启默认中间件
	ExtendedMap          map[string]interface{}  // 扩展信息，用于启动时打印
	SwaggerConfig        *SwaggerConfig
}
type SwaggerConfig struct {
	Version     string // 版本
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
	SwaggerSpec *swag.Spec
}
type DefaultMiddlewareSwitch struct {
	DisableRecoveryMiddleware bool // 禁用panic中间件,默认false
	DisableCorsMiddleware     bool // 禁用跨域cors，默认false
}

type RouterRegisterFunc func(group *gin.Engine)

var _httpServer *httpServer

/*
* @Description: 实例化服务
* @Author: LiuQHui
* @Param conf
* @Return *httpServer
* @Date 2024-06-11 15:46:47
 */
func NewHttpServer(conf HttpServerConfig) *httpServer {
	_httpServer = &httpServer{
		config: conf,
	}
	return _httpServer
}

/*
* @Description: 设置默认参数
* @Author: LiuQHui
* @Receiver h
* @Date 2024-06-18 18:28:05
 */
func (h *httpServer) setDefaultConfig() {
	// 运行模式
	if h.config.RunMode == "" {
		h.config.RunMode = gin.DebugMode
	}
	// 延迟多久退出，用于平滑重启
	if h.config.DelayExistSecond == 0 {
		h.config.DelayExistSecond = 5
	}
	// 路由不存在时处理
	h.config.RouterFunc = append(h.config.RouterFunc, noRoute)
	// swagger
	if h.config.SwaggerConfig != nil {
		if h.config.SwaggerConfig.Host == "" {
			h.config.SwaggerConfig.Host = fmt.Sprintf("0.0.0.0:%v", h.config.Port)
		}
		if h.config.SwaggerConfig.Version == "" {
			h.config.SwaggerConfig.Version = "v1.0.0"
		}
	}
}

/*
* @Description: 启动服务
* @Author: LiuQHui
* @Receiver h
* @Date 2024-06-11 15:44:43
 */
func (h *httpServer) Start() {
	// 设置默认参数
	h.setDefaultConfig()
	// 设置运行模式
	gin.SetMode(h.config.RunMode)
	// 创建从操作系统中断信号的上下文
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	// 创建gin服务实例
	engine := gin.Default()
	// 判断是否开启默认中间件
	h.enableDefaultMiddle(engine)
	// 注册中间件
	if len(h.config.CommonMiddlewareList) > 0 {
		engine.Use(h.config.CommonMiddlewareList...)
	}
	// 默认不处理 favicon.ico请求
	if !h.config.OpenFavicon {
		engine.GET("/favicon.ico", func(c *gin.Context) {
			c.Status(http.StatusNoContent)
		})
	}
	// 注册路由
	for _, registerFunc := range h.config.RouterFunc {
		registerFunc(engine)
	}
	// swagger初始化
	if h.config.SwaggerConfig != nil {
		h.swaggerInit(engine)
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

	// 打印启动信息
	h.printRunInfo()

	// Listen for the interrupt signal. 监听中断信号
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	// 恢复中断信号的默认行为并通知用户关闭
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	// 上下文用于通知服务器有n秒的时间来完成当前正在处理的请求;平滑重启防止直接终端服务
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.DelayExistSecond)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}

/*
* @Description: 打印启动信息
* @Author: LiuQHui
* @Receiver h
* @Date 2024-06-18 10:35:06
 */
func (h *httpServer) printRunInfo() {
	fmt.Println("========================================> 服务信息 <======================================== ")
	fmt.Printf("运行环境: %v\n", h.config.RunEnv)
	fmt.Printf("GinMode: %v\n", h.config.RunMode)
	fmt.Printf("延迟处理(平滑重启): %d秒 \n", h.config.DelayExistSecond)
	fmt.Printf("服务地址: http://0.0.0.0:%v\n", h.config.Port)
	if h.config.SwaggerConfig != nil {
		fmt.Printf("服务版本: %s\n", h.config.SwaggerConfig.Version)
		fmt.Printf("SwaggerUI: http://%v/swagger/index.html\n", h.config.SwaggerConfig.Host)
	}
	if len(h.config.ExtendedMap) > 0 {
		keys := funk.Keys(h.config.ExtendedMap).([]string)
		// 对键进行排序
		sort.Strings(keys)
		// 根据排序后的键，遍历并输出 map 的内容
		fmt.Println("\n扩展信息(ExtendedMap):")
		for _, key := range keys {
			fmt.Printf("  %s: %v \n", key, h.config.ExtendedMap[key])
		}
	}
	fmt.Println("==========================================>  END <========================================== ")
}

/*
* @Description: 判断是否开启默认中间件
* @Author: LiuQHui
* @Receiver h
* @Param engine
* @Date 2024-06-12 12:29:15
 */
func (h *httpServer) enableDefaultMiddle(engine *gin.Engine) {
	// 追加信息到上下文中间件
	engine.Use(AdditionalMiddleware)
	if h.config.RunMode == gin.DebugMode {
		fmt.Printf("[ginutil]   MiddlewareRegister %v --> github.com/52lu/go-helpers/ginutil.RequestEntranceLogMiddleware\n", strings.Repeat(" ", 13))
		fmt.Printf("[ginutil]   MiddlewareRegister %v --> github.com/52lu/go-helpers/ginutil.AdditionalMiddleware\n", strings.Repeat(" ", 13))
	}
	// 是否关闭默认中间件
	defaultMiddlewareSwitch := h.config.DefaultMiddleware
	// 捕获异常中间件
	if !defaultMiddlewareSwitch.DisableRecoveryMiddleware {
		if h.config.RunMode == gin.DebugMode {
			fmt.Printf("[ginutil]   MiddlewareRegister %v --> github.com/52lu/go-helpers/ginutil.RecoveryMiddleware\n", strings.Repeat(" ", 13))
		}
		engine.Use(RecoveryMiddleware)
	}
	// 跨域cors
	if !defaultMiddlewareSwitch.DisableCorsMiddleware {
		if h.config.RunMode == gin.DebugMode {
			fmt.Printf("[ginutil]   MiddlewareRegister %v --> github.com/52lu/go-helpers/middlewareutil.CORSMiddleware\n", strings.Repeat(" ", 13))
		}
		engine.Use(middlewareutil.CORSMiddleware())
	}
}

/*
* @Description: 不存在的路由进行报错
* @Author: LiuQHui
* @Param group
* @Date 2024-06-18 18:33:47
 */
func noRoute(group *gin.Engine) {
	// 不存在路由
	group.NoRoute(func(ctx *gin.Context) {
		FailRespEnum(ctx, RespEnumErrorNotFound)
	})
}
