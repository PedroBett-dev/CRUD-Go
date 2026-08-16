package controller

import (
	"net/http"

	rest_err "github.com/PedroBett-dev/CRUD-Go.git/src/configuration/err"
	"github.com/PedroBett-dev/CRUD-Go.git/src/configuration/validation"
	"github.com/PedroBett-dev/CRUD-Go.git/src/model/entity"
	"github.com/PedroBett-dev/CRUD-Go.git/src/model/request"
	"github.com/PedroBett-dev/CRUD-Go.git/src/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	user := entity.UserEntity{
		Name:     userRequest.Name,
		Age:      userRequest.Age,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	id, err := userRepo.Create(user)
	if err != nil {
		errRest := rest_err.NewInternalServerError("Error trying to create user")

		c.JSON(errRest.Code, errRest)
		return
	}

	userResponse := response.UserResponse{
		ID:    id,
		Name:  user.Name,
		Age:   user.Age,
		Email: user.Email,
	}

	c.JSON(http.StatusCreated, userResponse)
}