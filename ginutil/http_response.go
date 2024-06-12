package ginutil

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

// 定义统一返回接口格式
type Response struct {
	Code       int         `json:"code"`
	Msg        string      `json:"msg"`
	Data       interface{} `json:"data"`
	Additional Additional  `json:"additional"`
}

// 额外信息
type Additional struct {
	Time    string `json:"time"`
	TraceId string `json:"trace_id"`
}

var (
	SUCCESS = 200
	ERROR   = 0
)

func setAdditional(ctx context.Context, res *Response) {
	res.Additional = Additional{
		Time: time.Now().Format(time.DateTime),
		//TraceId: utils.GetTraceId(ctx),
	}
}

// 请求响应
func resultJson(ctx *gin.Context, code int, msg string, data interface{}) {
	response := Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	setAdditional(ctx, &response)
	ctx.JSON(SUCCESS, response)
}

/**
*  @Desc：错误处理
*  @Author LiuQingHui
*  @param ctx
*  @param code
*  @param err
*  @param data
*  @Date 2021-11-22 12:38:33
**/
func resultErrorJson(ctx *gin.Context, code int, errMsg string) {
	//logger.Warning(ctx, "resultErrorJson", map[string]interface{}{
	//	"data":     data,
	//	"method":   ctx.Request.Method,
	//	"url":      ctx.Request.URL,
	//	"PostForm": ctx.Request.PostForm,
	//	"Header":   ctx.Request.Header,
	//	"Form":     ctx.Request.Form,
	//	"Query":    ctx.Request.URL.Query(),
	//	"error":    fmt.Sprintf("%+v", err),
	//})
	response := Response{
		Code: code,
		Msg:  errMsg,
	}
	setAdditional(ctx, &response)
	ctx.JSON(SUCCESS, response)
}

/*
* @Description: 成功响应
* @Author: LiuQHui
* @Param ctx
* @Date 2024-06-11 18:27:29
 */
func Success(ctx *gin.Context) {
	resultJson(ctx, SUCCESS, "success", nil)
}

/*
* @Description: 返回固定消息和数据
* @Author: LiuQHui
* @Param ctx
* @Param data
* @Date 2024-06-11 17:50:08
 */
func SuccessWithData(ctx *gin.Context, data interface{}) {
	resultJson(ctx, SUCCESS, "success", data)
}

/*
* @Description: 成功响应(指定code、msg)
* @Author: LiuQHui
* @Param ctx
* @Param code
* @Param msg
* @Param data
* @Date 2024-06-11 17:49:55
 */
func SuccessResp(ctx *gin.Context, code int, msg string, data interface{}) {
	resultJson(ctx, code, msg, data)
}

/*
* @Description: 自定义错误code
* @Author: LiuQHui
* @Param ctx
* @Param code
* @Param errMsg
* @Date 2024-06-12 09:40:47
 */
func FailResp(ctx *gin.Context, code int, errMsg string) {
	resultErrorJson(ctx, code, errMsg)
}

/*
* @Description: 默认code返回
* @Author: LiuQHui
* @Param ctx
* @Param errMsg
* @Date 2024-06-12 09:37:00
 */
func Fail(ctx *gin.Context, errMsg string) {
	resultErrorJson(ctx, ERROR, errMsg)
}
