package main

import (
	"net/http"

	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/web"
)

func main() {

	// 修改consul地址，如果是本机，这段代码和后面的那行使用代码都是可以不用的
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"10.152.113.166:8500",
		}
	})

	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.docs"),
		web.Registry(reg),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	//println(http.FileServer(http.Dir("html")))
	service.Handle("/", http.FileServer(http.Dir("docs")))

	// register call handler
	//service.HandleFunc("/web/call", handler.WebCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
