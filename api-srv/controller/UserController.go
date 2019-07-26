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

/**
 * @api {get} /user/:id Get User information
 * @apiVersion 0.1.0
 * @apiName GetUser
 * @apiGroup User
 *
 * @apiParam {Number} id Users unique ID.
 *
 * @apiSuccess {String} firstname Firstname of the User.
 * @apiSuccess {String} lastname  Lastname of the User.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "firstname": "John",
 *       "lastname": "Doe"
 *     }
 *
 * @apiError UserNotFound The id of the User was not found.
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 *     {
 *       "error": "UserNotFound"
 *     }
 */
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
