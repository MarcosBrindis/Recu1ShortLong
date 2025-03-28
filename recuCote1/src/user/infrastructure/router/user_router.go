package router

import (
	"recuCorte1/src/core/middleware"
	"recuCorte1/src/user/infrastructure/http/controller"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine,
	create *controller.CreateUserController,
	get *controller.GetUserController,
	getAll *controller.GetAllUsersController,
	userPolling *controller.UserPollingController,
	updates *chan bool,
) {

	userGroup := r.Group("/user")
	{
		userGroup.POST("/", middleware.NotifyUpdatesMiddleware(updates), create.HandleCreate)
		userGroup.GET("/:id", get.HandleGet)
		userGroup.GET("/", getAll.HandleGetAll)

		// Endpoints para polling
		userGroup.GET("/shortpoll", userPolling.HandleShortPoll)
		userGroup.GET("/longpoll", userPolling.HandleLongPoll)
		userGroup.GET("/poll/count", userPolling.HandleCountShortPollStreaming)
	}
}
