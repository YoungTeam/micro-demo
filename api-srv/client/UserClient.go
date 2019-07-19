package client

import (
	proto "mgo/user-srv/proto/user"

	"github.com/micro/go-micro/client"
)

// user 服务客户端
func UserClient() proto.UserService {
	return proto.NewUserService("go.micro.srv.user", client.DefaultClient)
}
