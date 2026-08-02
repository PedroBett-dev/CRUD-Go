package routes

import (
	"github.com/PedroBett-dev/CRUD-Go.git/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserByID/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userMail", controller.FindUserByMail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser", controller.UpdateUser)
	r.DELETE("/deleteUser/:UserId", controller.DeleteUser)

}
