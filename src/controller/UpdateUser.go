package controller

import (
	"net/http"

	rest_err "github.com/PedroBett-dev/CRUD-Go.git/src/configuration/err"
	"github.com/PedroBett-dev/CRUD-Go.git/src/configuration/validation"
	"github.com/PedroBett-dev/CRUD-Go.git/src/model/entity"
	"github.com/PedroBett-dev/CRUD-Go.git/src/model/repository"
	"github.com/PedroBett-dev/CRUD-Go.git/src/model/request"
	"github.com/PedroBett-dev/CRUD-Go.git/src/model/response"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userId := c.Query("userId")

	user := entity.UserEntity{
		Name:     userRequest.Name,
		Age:      userRequest.Age,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	err := userRepo.Update(userId, user)
	if err != nil {
		if repository.IsNoRows(err) {
			errRest := rest_err.NewNotFoundError("User not found")

			c.JSON(errRest.Code, errRest)
			return
		}

		errRest := rest_err.NewInternalServerError("Error trying to update user")

		c.JSON(errRest.Code, errRest)
		return
	}

	userResponse := response.UserResponse{
		ID:    userId,
		Name:  user.Name,
		Age:   user.Age,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, userResponse)
}