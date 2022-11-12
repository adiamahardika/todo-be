package router

import (
	"svc-todo/controller"
	"svc-todo/repository"
	"svc-todo/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(db *gorm.DB) {

	router := gin.Default()

	repository := repository.Repository(db)

	activityService := service.ActivityService(repository)
	activityController := controller.ActivityController(activityService)

	todoService := service.TodoService(repository)
	todoController := controller.TodoController(todoService)

	activity := router.Group("/activity-groups")
	{
		activity.POST("/", activityController.CreateActivity)
		activity.GET("/", activityController.GetActivity)
		activity.GET("/:id", activityController.GetOneActivity)
		activity.DELETE("/:id", activityController.DeleteActivity)
		activity.PATCH("/:id", activityController.UpdateActivity)
	}

	todo := router.Group("/todo-items")
	{
		todo.POST("/", todoController.CreateTodo)
	}
	router.Run(":3030")
}
