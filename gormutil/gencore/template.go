package gencore

import (
	"52lu/go-helpers/fileutil"
	"52lu/go-helpers/gormutil/gencore/tmpl"
	"52lu/go-helpers/strutil"
	"fmt"
	"github.com/thoas/go-funk"
	"golang.org/x/tools/go/packages"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
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
	// 判断目录是否存在，不存在则创建
	if err := fileutil.CreatePath(g.conf.OutPath); err != nil {
		return err
	}
	// 获取路径
	filePath := fmt.Sprintf("%s/baseDao.gen.go", g.conf.OutPath)
	pathSplit := strings.Split(g.conf.OutPath, "/")
	return g.generateFileByTemplate(tmpl.DefaultBaseDaoTemplate, tmpl.BaseDaoTemplateVar{
		PackageName: pathSplit[len(pathSplit)-1],
		DateTime:    time.Now().Format(time.DateTime),
	}, filePath)
}

var (
	onceParse sync.Once
)

/*
* @Description: 为每个model生成dao
* @Author: LiuQHui
* @Receiver g
* @Param modelName
* @Date 2024-05-30 11:20:03
 */
func (g genUtilClient) generateModelDao(modelName string, tableColumns []string) error {
	// 判断目录是否存在，不存在则创建
	if err := fileutil.CreatePath(g.conf.OutPath); err != nil {
		return err
	}
	absPath, err := filepath.Abs(g.conf.OutPath)
	if err != nil {
		return err
	}
	// 解析项目信息
	// 加载包
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName,
		Dir:  g.conf.OutPath, // 当前目录
	})
	if err != nil {
		log.Fatalf("Error loading package: %v", err)
		return err
	}
	queryPkgPath := pkgs[0].PkgPath
	// 取model PkgPath
	queryPkgPathSplit := strings.Split(queryPkgPath, "/")
	queryPkgPathSplit = queryPkgPathSplit[:len(queryPkgPathSplit)-1]
	queryPkgPathSplit = append(queryPkgPathSplit, "model")
	modelPkgPath := strings.Join(queryPkgPathSplit, "/")

	// 填充参数
	pathSplit := strings.Split(absPath, "/")
	packageName := pathSplit[len(pathSplit)-2]
	pathSplit = pathSplit[:len(pathSplit)-1]
	daoPath := strings.Join(pathSplit, "/")

	tmplVar := tmpl.ModelDaoVar{
		UseGormHookDataLog: g.conf.UseGormHookDataLog,
		PackageName:        packageName,
		DateTime:           time.Now().Format(time.DateTime),
		ModelPkgPath:       modelPkgPath,
		QueryPkgPath:       queryPkgPath,
		ModelName:          modelName,
		ReceiverPre:        strings.ToLower(modelName[0:1]),
		DaoName:            strings.ReplaceAll(modelName, g.conf.ModelSuffix, "") + "Dao",
	}
	// 获取路径
	filePath := fmt.Sprintf("%s/%v.go", daoPath, strutil.ToLowerFirstEachWord(tmplVar.DaoName))

	// 获取
	tmplStr := tmpl.DefaultDaoTemplate
	if funk.ContainsString(tableColumns, "id") {
		tmplStr += tmpl.DaoCommonByIdMethod
	}
	return g.generateFileByTemplate(tmplStr, tmplVar, filePath)
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
	if fileutil.ExistFile(filePath) {
		// 判断是否覆盖
		if !g.conf.OverDaoFile {
			fmt.Println("File exists: ", filePath)
			return nil
		}
		// 覆盖，删除旧的
		_ = os.Remove(filePath)
	}
	// 创建一个文件来写入生成的 Go 代码
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("File Generate: ", filePath)
	defer file.Close()
	// 执行模板并将输出写入文件
	return t.Execute(file, data)
}
