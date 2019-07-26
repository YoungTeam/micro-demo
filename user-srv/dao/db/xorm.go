package db

import (
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

// type MysqlConfig struct {
// 	Master struct {
// 		Dsn string `json:"dsn"`
// 	} `json:"master"`
// 	Slaves []struct {
// 		Dsn string `json:"dsn"`
// 	} `json:"slaves"`
// }

//func GetDbConn(alias string) (*xorm.Engine, error) {

//}

type DbConf struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// X 全局DB
var engine *xorm.Engine

func init() {
	GetConn()
}

func GetConn() {
	var err error
	err = config.LoadFile("D:/workspace/GO/mgo/user-srv/conf/config.yaml")
	if err != nil {
		log.Fatalf("Could not load config file: %s", err.Error())
		return
	}

	//conf := config.Map()
	var dbConf DbConf
	config.Get("mysql").Scan(&dbConf)
	//fmt.Println(dbConf.User)
	//fmt.Printf("%v\n", dbConf)

	source := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8", dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database)

	engine, err = xorm.NewEngine("mysql", source)

	if err != nil {
		log.Fatalf("db error: %#v\n", err.Error())
	}

	/*--------------------------------------------------------------------------------------------------
	1、使用RegisterSqlMap()注册SqlMap配置
	2、RegisterSqlTemplate()方法注册SSqlTemplate模板配置
	3、SqlMap配置文件总根目录和SqlTemplate模板配置文件总根目录可为同一目录
	--------------------------------------------------------------------------------------------------*/

	//注册SqlMap配置，可选功能，如应用中无需使用SqlMap，可无需初始化
	//此处使用xml格式的配置，配置文件根目录为"./sql/oracle"，配置文件后缀为".xml"
	// err = x.RegisterSqlMap(xorm.Xml("./db/sql/xml", ".xml"))
	// if err != nil {
	// 	log.Fatalf("db error: %#v\n", err.Error())
	// }
	//注册动态SQL模板配置，可选功能，如应用中无需使用SqlTemplate，可无需初始化
	//此处注册动态SQL模板配置，使用Pongo2模板引擎，配置文件根目录为"./sql/oracle"，配置文件后缀为".stpl"
	// err = x.RegisterSqlTemplate(xorm.Default("./db/sql/tpl", ".sql"))
	// if err != nil {
	// 	log.Fatalf("db error: %#v\n", err.Error())
	// }

	//开启SqlMap配置文件和SqlTemplate配置文件更新监控功能，将配置文件更新内容实时更新到内存，如无需要可以不调用该方法
	// err = engine.StartFSWatcher()
	// if err != nil {
	// 	log.Printf("sql parse error: %#v\n", err)
	// }

	err = engine.Ping()
	if err != nil {
		log.Fatalf("db connect error: %#v\n", err.Error())
	}

	// 30minute ping db to keep connection
	timer := time.NewTicker(time.Minute * 30)
	go func(x *xorm.Engine) {
		for _ = range timer.C {
			err = x.Ping()
			if err != nil {
				log.Fatalf("db connect error: %#v\n", err.Error())
			}
		}
	}(engine)
	engine.ShowSQL(true)
}
