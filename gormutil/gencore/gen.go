package gencore

import (
	"errors"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/tools/go/packages"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
	"strings"
)

type genUtilClient struct {
	conf GenConfig
}

type GenConfig struct {
	// 数据库连接-> 用户:密码@tcp(ip:port)/db?charset=utf8mb4&parseTime=true&loc=Local
	MysqlDsn    string
	OutPath     string // 输出目录
	TablePre    string // 表前缀
	ModelSuffix string // 生成的mode后缀名
	GenConf     *gen.Config
}

const (
	DefaultOutPath = "./app/dao/query"
)

/*
* @Description: 生成客户端
* @Author: LiuQHui
* @Param cf
* @Date 2024-05-27 18:07:12
 */
func NewGenUtilClient(cf GenConfig) genUtilClient {
	if cf.OutPath == "" {
		cf.OutPath = DefaultOutPath
	}
	return genUtilClient{
		conf: cf,
	}
}

/*
* @Description: 获取默认实例
* @Author: LiuQHui
* @Receiver g
* @Date 2024-05-27 18:38:21
 */
func (g genUtilClient) getDefaultGenerator() *gen.Generator {
	// 创建实例
	generator := gen.NewGenerator(gen.Config{
		// 相对执行`go run`时的路径, 会自动创建目录
		OutPath: g.conf.OutPath,
		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true, // generate pointer when field is nullable
		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values
		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false, // detect integer field's unsigned type, adjust generated data type
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false, // generate with gorm index tag
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true, // generate with gorm column type tag
		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	return generator
}

/*
* @Description: 获取gen实例
* @Author: LiuQHui
* @Receiver g
* @Date 2024-05-27 18:39:53
 */
func (g genUtilClient) getGeneratorInstance() *gen.Generator {
	if g.conf.GenConf == nil {
		return g.getDefaultGenerator()
	}
	return gen.NewGenerator(*g.conf.GenConf)
}

/*
* @Description: 配置检查
* @Author: LiuQHui
* @Receiver g
* @Date 2024-05-27 18:42:42
 */
func (g genUtilClient) checkConf() error {
	if g.conf.MysqlDsn == "" {
		return errors.New("没有配置MySQL[MysqlDsn]连接信息")
	}
	return nil
}

/*
* @Description: 运行生成
* @Author: LiuQHui
* @Receiver g
* @Date 2024-05-28 17:17:19
 */
func (g genUtilClient) Run() error {
	//// 生成model和query
	//modelList, err := g._runGormGen()
	//if err != nil {
	//	return err
	//}
	//fmt.Println(modelList)
	////  生成基类Dao
	//if err := g.generateBaseDao(); err != nil {
	//	return err
	//}

	g.parse()
	// todo 为每个model生成Dao
	return nil
}

func (g genUtilClient) parse() {
	// 加载包
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes,
		Dir:  g.conf.OutPath, // 当前目录
	})
	if err != nil {
		log.Fatalf("Error loading package: %v", err)
	}

	// 处理加载的包信息
	for _, pkg := range pkgs {
		fmt.Printf("Package: %s\n", pkg.Name)
		fmt.Printf("Files:\n")
		for _, file := range pkg.GoFiles {
			fmt.Printf("  %s\n", file)
		}
		fmt.Printf("Syntax:\n")
		for _, syntax := range pkg.Syntax {
			fmt.Printf("  %s\n", syntax.Name.Name)
		}
	}
}

/*
* @Description: 生成model和query
* @Author: LiuQHui
* @Receiver g
* @Date 2024-05-27 19:28:37
 */
func (g genUtilClient) _runGormGen() ([]interface{}, error) {
	var modelList []interface{}
	// 配置检查
	if err := g.checkConf(); err != nil {
		return modelList, err
	}
	// 获取gen实例
	genInstance := g.getGeneratorInstance()
	// 连接数据库
	db, err := gorm.Open(mysql.Open(g.conf.MysqlDsn))
	if err != nil {
		panic(fmt.Errorf("连接数据失败:%w", err))
	}
	// 设置DB
	genInstance.UseDB(db)
	// 模型结构体的后缀命名规则
	genInstance.WithModelNameStrategy(func(tableName string) (modelName string) {
		if g.conf.ModelSuffix != "" {
			return g.removeTablePre(tableName) + g.conf.ModelSuffix
		}
		return g.removeTablePre(tableName)
	})
	// 文件命名规则
	genInstance.WithFileNameStrategy(func(tableName string) (fileName string) {
		return firstLower(g.removeTablePre(tableName))
	})

	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf
	dataTypeCustomMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"tinyint":   func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"smallint":  func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"mediumint": func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"bigint":    func(columnType gorm.ColumnType) (dataType string) { return "int64" },
		"int":       func(columnType gorm.ColumnType) (dataType string) { return "int64" },
	}
	genInstance.WithDataTypeMap(dataTypeCustomMap)
	// 创建全部模型文件, 并覆盖前面创建的同名模型
	modelList = genInstance.GenerateAllTable()
	// 生成基础函数
	genInstance.ApplyBasic(modelList...)
	// 填充常用SQL
	genInstance.ApplyInterface(func(QueryMethodInterface) {}, modelList...)
	// 执行
	genInstance.Execute()
	return modelList, nil

}

/*
* @Description: 删除表前缀: msc_primary_pool -> PrimaryPool
* @Author: LiuQHui
* @Receiver g
* @Param tableName
* @Date 2024-05-27 18:45:37
 */
func (g genUtilClient) removeTablePre(tableName string) string {
	// 去除表前缀
	if g.conf.TablePre == "" {
		return tableName
	}
	newStr := strings.TrimPrefix(tableName, g.conf.TablePre)
	newStr = strings.ReplaceAll(newStr, "_", " ")
	return strings.ReplaceAll(cases.Title(language.Und).String(newStr), " ", "")
}

/*
* @Description: 首位改成小写
* @Author: LiuQHui
* @Param s
* @Date 2024-05-27 18:46:28
 */
func firstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
