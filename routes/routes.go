package routes

import (
	"app/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	// user
	e.POST("/register", controllers.CreateUser) // register user
	e.POST("/login", controllers.LoginUser)     // login user
	// room
	e.GET("/listrooms/:userid", controllers.ListRooms) // list room user with another user
	// chat
	e.GET("/listchats/:roomid", controllers.ListChats) // list chat user with another user
	e.POST("/sendchat", controllers.SendChat)          // send new chat from user to another user
	return e
}
