package db

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	GetConn()
}

func TestUser(t *testing.T) {

	GetConn()
	u := new(UserDaoImpl)
	user, err := u.GetUser(1)
	if err == nil {
		fmt.Printf("%v", user)
	}

}
