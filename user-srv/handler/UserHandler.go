package handler

import (
	"context"
	"mgo/tools/golog"
	"mgo/user-srv/dao"
	pb "mgo/user-srv/proto/user"
)

type UserHandler struct {
	UserDao dao.UserDao
}

// 获取一个用户
func (u *UserHandler) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest, resp *pb.GetUserByIdResponse) error {
	user, err := u.UserDao.GetUser(req.Id)
	if err != nil {
		return nil
	}

	resp.Header = make(map[string]*pb.Pair)
	resp.Header["name"] = &pb.Pair{Key: 1, Values: "abc"}

	resp.Code = "200"
	resp.Msg = "成功"
	resp.User = &pb.User{
		Id:       user.Id,
		UserId:   user.UserId,
		UserName: user.UserName,
	}

	golog.Debugf("%v", user)
	//fmt.Printf("%v", user)
	return nil

}
