package logutil

import (
	"fmt"
	"github.com/52lu/go-helpers/fileutil"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"os"
	"time"
)

// DailyLogger 封装了 lumberjack.Logger，并添加按天切割的功能
type DailyLogger struct {
	*lumberjack.Logger
	//targetFile string // 目标文件
}

/*
* @Description:  创建一个新的 DailyLogger
* @Author: LiuQHui
* @Param filename
* @Param lumberJackConfig
* @Return *DailyLogger
* @Date 2025-01-20 16:03:14
 */
func NewDailyLogger(filename string, lumberJackConfig LumberJackConfig) *DailyLogger {
	dl := &DailyLogger{
		Logger: &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    lumberJackConfig.MaxSize,    // 文件最大大小（MB）
			MaxBackups: lumberJackConfig.MaxBackups, // 保留的最大旧文件数量
			MaxAge:     lumberJackConfig.MaxAge,     // 保留的最大天数
			Compress:   lumberJackConfig.Compress,   // 是否压缩
		},
	}
	dl.createSymlink(dl.getTargetFile())
	return dl
}
func (dl *DailyLogger) getTargetFile() string {
	currentDate := getCurrentDate()
	return fmt.Sprintf("%v-%s", dl.Filename, currentDate)
}

// Write 实现 io.Writer 接口，添加按天切割的逻辑
func (dl *DailyLogger) Write(p []byte) (n int, err error) {
	// 判断目录是否存在
	dl.createSymlink(dl.getTargetFile())
	//if dl.createSymlink(dl.getTargetFile()) {
	//	// 如果日期变化，进行日志切割
	//	err := dl.Rotate()
	//	if err != nil {
	//		return 0, err
	//	}
	//}
	return dl.Logger.Write(p)
}

// createSymlink 创建软链接 app.log -> app-YYYY-MM-DD.log
func (dl *DailyLogger) createSymlink(targetFile string) bool {
	// 检查软链接是否已经存在
	_, err := os.Lstat(targetFile)
	if err == nil {
		// 如果软链接已存在，直接返回
		return false
	}
	// 创建目标文件
	file, _ := os.Create(targetFile)
	if file != nil {
		defer file.Close()
	}
	// 存在则删除重建
	if fileutil.ExistFile(dl.Filename) {
		err = os.Remove(dl.Filename)
		if err != nil {
			fmt.Println("日志软链接则删除重建:", err)
		}
	}
	// 创建软链接
	err = os.Symlink(targetFile, dl.Filename)
	if err != nil {
		zap.L().Error("Failed to create symlink", zap.Error(err))
	}
	return true
}

/*
* @Description: 获取日期格式
* @Author: LiuQHui
* @Return string
* @Date 2025-01-20 16:03:41
 */
func getCurrentDate() string {
	return time.Now().Format(time.DateOnly)
	//return time.Now().Format("200601021504")
}
