package controller

import (
	"net/http"

	rest_err "github.com/PedroBett-dev/CRUD-Go.git/src/configuration/err"
	"github.com/PedroBett-dev/CRUD-Go.git/src/model/repository"
	"github.com/PedroBett-dev/CRUD-Go.git/src/model/response"
	"github.com/gin-gonic/gin"
)

func FindUserById(c *gin.Context) {
	userId := c.Param("userId")

	user, err := userRepo.FindByID(userId)
	if err != nil {
		if repository.IsNoRows(err) {
			errRest := rest_err.NewNotFoundError("User not found")

			c.JSON(errRest.Code, errRest)
			return
		}

		errRest := rest_err.NewInternalServerError("Error trying to find user")

		c.JSON(errRest.Code, errRest)
		return
	}

	userResponse := response.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Age:   user.Age,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, userResponse)
}

func FindUserByMail(c *gin.Context) {
	userMail := c.Param("userMail")

	user, err := userRepo.FindByEmail(userMail)
	if err != nil {
		if repository.IsNoRows(err) {
			errRest := rest_err.NewNotFoundError("User not found")

			c.JSON(errRest.Code, errRest)
			return
		}

		errRest := rest_err.NewInternalServerError("Error trying to find user")

		c.JSON(errRest.Code, errRest)
		return
	}

	userResponse := response.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Age:   user.Age,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, userResponse)
}

func ListAllUsers(c *gin.Context) {
	users, err := userRepo.List()

	if err != nil {
		errRest := rest_err.NewInternalServerError("Error trying to list users")

		c.JSON(errRest.Code, errRest)
		return
	}

	usersResponse := []response.UserResponse{}

	for _, user := range users {
		usersResponse = append(usersResponse, response.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Age:   user.Age,
			Email: user.Email,
		})
	}

	c.JSON(http.StatusOK, usersResponse)
}
