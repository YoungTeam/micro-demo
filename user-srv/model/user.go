package model

type User struct {
	Id       int64  `xorm:"not null pk autoincr INT(11)"`
	UserId   string `xorm:"not null VARCHAR(50)"`
	UserName string `xorm:"not null VARCHAR(100)"`
}
