package controller

import (
	"context"
	"log"
	"mgo/api-srv/client"
	user "mgo/user-srv/proto/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	us = client.UserClient()
)

func GetUser(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	//.Param("id")
	//fmt.Print(id)
	log.Printf("=====%d", id)
	response, err := us.GetUserById(context.TODO(),
		&user.GetUserByIdRequest{
			Id: id,
		})

	if err != nil {
		c.JSON(500, err)
	}

	c.JSON(200, response)
}
