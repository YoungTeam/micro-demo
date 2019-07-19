package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	article "mgo/article-srv/proto/article"
)

type Article struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Article) Call(ctx context.Context, req *article.Request, rsp *article.Response) error {
	log.Log("Received Article.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Article) Stream(ctx context.Context, req *article.StreamingRequest, stream article.Article_StreamStream) error {
	log.Logf("Received Article.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&article.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Article) PingPong(ctx context.Context, stream article.Article_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&article.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
