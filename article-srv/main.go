package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"mgo/article-srv/handler"
	"mgo/article-srv/subscriber"

	article "mgo/article-srv/proto/article"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.article"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	article.RegisterArticleHandler(service.Server(), new(handler.Article))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.article", service.Server(), new(subscriber.Article))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.article", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
