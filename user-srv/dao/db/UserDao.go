package db

import (
	"errors"
	"mgo/user-srv/model"
)

type UserDaoImpl struct {
}

func (dao *UserDaoImpl) GetUser(id int64) (*model.User, error) {
	user := &model.User{}
	user.Id = id
	has, err := engine.Get(user)
	if err != nil {
		return user, err
	}

	if !has {
		return user, errors.New("not found")
	}

	return user, nil
}
