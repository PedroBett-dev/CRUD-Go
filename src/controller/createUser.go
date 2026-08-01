package controller

import (
	"fmt"

	RestErr "github.com/PedroBett-dev/CRUD-Go.git/src/configuration/err"
	"github.com/PedroBett-dev/CRUD-Go.git/src/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		restErr := RestErr.NewBadRequestError(
			fmt.Sprintf("Incorrect JSON structure, error=%s\n", err.Error()),
		)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
