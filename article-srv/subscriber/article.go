package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	article "mgo/article-srv/proto/article"
)

type Article struct{}

func (e *Article) Handle(ctx context.Context, msg *article.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *article.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
