package dao

import "mgo/user-srv/model"

type UserDao interface {
	GetUser(id int64) (*model.User, error)
}
