package routes

import (
	"demo04/internal/handler"
	"demo04/internal/service"
	"demo04/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.ErrorMiddleware())

	api := ginRouter.Group("api")
	{
		api.GET("ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, "pong")
		})
		user := api.Group("user")
		{
			user.POST("register", handler.UserRegister)
			user.POST("login", handler.UserLogin)
		}
		//api.Use(middleware.JWT())
		friend := api.Group("friend")
		{
			friend.POST("", handler.AddFriend)
			friend.PUT("", handler.DeleteFriend)
			friend.GET("")
		}
		chat := api.Group("chat")
		{
			chat.GET("private", service.PrivateChatHandler)
			chat.GET("room", service.RoomChatHandler)
		}
	}

	return ginRouter
}
