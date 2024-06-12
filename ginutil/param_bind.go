package ginutil

import (
	"bytes"
	"fmt"
	"github.com/52lu/go-helpers/jsonutil"
	"github.com/52lu/go-helpers/verifyutil"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io"
	"net/http"
	"strings"
)

/*
* @Description: 绑定和验证参数
* @Author: LiuQHui
* @Param ctx
* @Param obj
* @Return error
* @Date 2024-06-12 11:33:28
 */
func BindParamWithValidate(ctx *gin.Context, obj interface{}) error {
	// 绑定url中的参数
	_ = ctx.ShouldBindQuery(obj)
	// 绑定表单参数
	err := ctx.ShouldBindWith(obj, binding.Form)
	if err != nil {
		return err
	}
	// 绑定json
	if ctx.Request.Method == http.MethodPost {
		contentType := ctx.Request.Header.Get("Content-Type")
		if strings.Contains(contentType, binding.MIMEJSON) {
			// 绑定json
			var b customBindJson
			return ctx.ShouldBindWith(obj, b)
		}
	}
	// 验证
	return validateNew(obj)
}

type customBindJson struct{}

func (customBindJson) Name() string {
	return "json"
}

func (customBindJson) Bind(req *http.Request, obj interface{}) error {
	if req == nil || req.Body == nil {
		return fmt.Errorf("invalid request")
	}
	return decodeJSON(req.Body, obj)
}

func (customBindJson) BindBody(body []byte, obj interface{}) error {
	return decodeJSON(bytes.NewReader(body), obj)
}

// 解析参数参数
func decodeJSON(r io.Reader, obj interface{}) error {
	decoder := jsonutil.Json.NewDecoder(r)
	if binding.EnableDecoderUseNumber {
		decoder.UseNumber()
	}
	if binding.EnableDecoderDisallowUnknownFields {
		decoder.DisallowUnknownFields()
	}
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return validateNew(obj)
}

// 验证参数
func validateNew(obj interface{}) error {
	if binding.Validator == nil {
		return nil
	}
	return verifyutil.ValidateStruct(obj)
}
