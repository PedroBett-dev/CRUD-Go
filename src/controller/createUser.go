package controller

import (
	"fmt"
	"log"

	"github.com/PedroBett-dev/CRUD-Go.git/src/configuration/validation"
	"github.com/PedroBett-dev/CRUD-Go.git/src/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error tryind to marshal object, error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
