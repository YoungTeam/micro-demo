package main

import (
	"mgo/tools/golog"
	"mgo/tools/golog/conf"
	"mgo/user-srv/dao/db"
	"time"

	"github.com/micro/go-micro"
	"mgo/user-srv/handler"
	user "mgo/user-srv/proto/user"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
)

func init() {
	golog.SetLogger(golog.ZAPLOG,
		conf.WithLogType(conf.LogJsontype),
		conf.WithProjectName("go_xxx"),
		conf.WithFilename("D:/workspace/GO/mgo/user-srv/test.log"),
	)

	//fmt.Println(conf.DebugLevel)
	//golog.SetLogLevel(conf.DebugLevel)
	golog.SetLogLevel(conf.DebugLevel)
}

func main() {
	// 修改consul地址，如果是本机，这段代码和后面的那行使用代码都是可以不用的
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"10.152.113.166:8500",
		}
	})

	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	userDao := &db.UserDaoImpl{}
	// Register Handlers
	user.RegisterUserServiceHandler(service.Server(), &handler.UserHandler{userDao})

	// Run server
	if err := service.Run(); err != nil {
		golog.Fatal("server run error", "err", err)
		//MainLogger.Fatal("server error", zap.stringerr)
	}

}
