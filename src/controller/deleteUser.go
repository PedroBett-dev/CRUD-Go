package controller

import (
	"net/http"

	rest_err "github.com/PedroBett-dev/CRUD-Go.git/src/configuration/err"
	"github.com/PedroBett-dev/CRUD-Go.git/src/model/repository"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {

	userId := c.Param("UserId")

	err := userRepo.Delete(userId)
	if err != nil {
		if repository.IsNoRows(err) {
			errRest := rest_err.NewNotFoundError("User not found")

			c.JSON(errRest.Code, errRest)
			return
		}

		errRest := rest_err.NewInternalServerError("Error trying to delete user")

		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}