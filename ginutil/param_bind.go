package wmutil

import (
	"bytes"
	"github.com/52lu/go-helpers/jsonutil"
	"github.com/52lu/go-helpers/verifyutil"
	"strings"

	//"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io"
	"net/http"
)

// 可以自定义jsoniter配置或者添加插件
//var json = jsoniter.ConfigCompatibleWithStandardLibrary

type customBindJson struct{}

/**
*  BindParamWithValidate
*  @Desc：绑定和验证参数(json和表单)
* // 可以将json中的类型自动转换
*  @Author LiuQingHui
*  @param ctx
*  @param obj
*  @return error
*  @Date 2021-12-04 16:03:39
**/
func BindParamWithValidate(ctx *gin.Context, obj interface{}) error {
	// 绑定url中的参数
	_ = ctx.ShouldBindQuery(obj)
	// 绑定json
	if ctx.Request.Method == http.MethodPost {
		contentType := ctx.Request.Header.Get("Content-Type")
		if strings.Contains(contentType, binding.MIMEJSON) {
			// 绑定json
			var b customBindJson
			return ctx.ShouldBindWith(obj, b)
		}
	}
	// 绑定POST表单参数
	err := ctx.ShouldBindWith(obj, binding.Form)
	if err != nil {
		return err
	}
	// 验证
	return validateNew(obj)
}

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
	// 自动适应类型
	//extra.RegisterFuzzyDecoders()
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
