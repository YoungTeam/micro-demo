package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"mgo/api-srv/router"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	router.Init(g)

	/* 	s := http.Server{
		Addr:           ":8080",
		Handler:        g,
		ReadTimeout:    time.Second * 30,
		WriteTimeout:   time.Second * 30,
		MaxHeaderBytes: 1 << 20,
	} */

	// 修改consul地址，如果是本机，这段代码和后面的那行使用代码都是可以不用的
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"10.152.113.166:8500",
		}
	})

	// Create service
	service := web.NewService(
		web.Name("go.micro.api.user"),
		web.Registry(reg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
	)

	service.Init()
	// Register Handler
	service.Handle("/", g)

	/* 	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("run gin error: %s", err)
	} */

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
