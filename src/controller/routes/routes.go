package routes

import (
	"github.com/PedroBett-dev/CRUD-Go.git/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/user/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userMail", controller.FindUserByMail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user", controller.UpdateUser)
	r.DELETE("/user/:UserId", controller.DeleteUser)

}
