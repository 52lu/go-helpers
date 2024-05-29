package gencore

import (
	"52lu/gin-api-template/generate/gormutil/gencore/tmpl"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"
)

/*
* @Description: 生成dao父类
* @Author: LiuQHui
* @Receiver g
* @Date 2024-05-28 19:21:55
 */
func (g genUtilClient) generateBaseDao() error {
	// 获取路径
	filePath := fmt.Sprintf("%s/baseDao.gen.go", g.conf.OutPath)
	pathSplit := strings.Split(g.conf.OutPath, "/")
	return g.generateFileByTemplate(tmpl.DefaultBaseDaoTemplate, tmpl.BaseDaoTemplateVar{
		PackageName: pathSplit[len(pathSplit)-1],
		DateTime:    time.Now().Format(time.DateTime),
	}, filePath)
}

/*
* @Description: 根据模版创建文件
* @Author: LiuQHui
* @Param tmpl
* @Param data
* @Param filePath
* @Date 2024-05-28 17:45:37
 */
func (g genUtilClient) generateFileByTemplate(tmpl string, data any, filePath string) error {
	// 解析模板
	t, err := template.New(tmpl).Parse(tmpl)
	if err != nil {
		return err
	}
	// 判断文件是否存在
	if existFile(filePath) {
		fmt.Println("File exists: ", filePath)
		return nil
	}

	// 创建一个文件来写入生成的 Go 代码
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 执行模板并将输出写入文件
	return t.Execute(file, data)
}

/*
* @Description: 判断文件是否存在
* @Author: LiuQHui
* @Param path
* @Date 2024-05-28 18:56:20
 */
func existFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
