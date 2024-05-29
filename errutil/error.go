package errutil

import (
	"fmt"
	cockroachdbError "github.com/cockroachdb/errors"
)

/**
*  ThrowErrorMsg
*  @Desc：抛出错误信息
*  @Author LiuQingHui
*  @param err
*  @return error
*  @Date 2023-05-19 17:10:54
**/
func ThrowErrorMsg(err string) error {
	return cockroachdbError.New(err)
}

/*
* @Description: 抛出错误信息
* @Author: LiuQHui
* @Param format
* @Param arg
* @Date 2023-05-19 14:37:22
 */
func ThrowErrorMsgF(format string, arg ...interface{}) error {
	return cockroachdbError.New(fmt.Sprintf(format, arg...))
}

/*
 * @Description: 抛出error
 * @Author: LiuQHui
 * @Param err
 * @Return error
 */
func ThrowError(err error, depth ...int) error {
	if len(depth) == 0 {
		depth = []int{0}
	}
	return cockroachdbError.WithStackDepth(err, depth[0]+1)
}

/**
*  @Desc：携带关键字的
*  @Author LiuQingHui
*  @param err
*  @param pre
*  @return error
*  @Date 2021-11-23 10:27:24
**/
func ThrowErrorWithPre(err error, pre string) error {
	return cockroachdbError.Wrap(err, pre)
}
