package ginutil

import (
	"fmt"
	"github.com/52lu/go-helpers/errutil"
	"github.com/gin-gonic/gin"
	"reflect"
)

/*
* @Description: 捕获全局painc
* @Author: LiuQHui
* @Param c
* @Date 2024-06-12 11:54:53
 */
func RecoveryMiddleware(ctx *gin.Context) {
	// 捕获错误
	defer func() {
		if err := recover(); err != nil {
			errType := reflect.TypeOf(err).String()
			var customErr error
			if errType == "string" {
				customErr = errutil.ThrowErrorMsg("捕获错误:" + err.(string))
			} else {
				customErr = errutil.ThrowErrorWithPre(err.(error), "捕获错误")
			}
			// 错误响应
			FailMsg(ctx, fmt.Sprintf("%v", customErr))
			ctx.Abort()
			return
		}
	}()
	ctx.Next()
}
