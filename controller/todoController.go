package controller

import (
	"fmt"
	"net/http"
	"svc-todo/entity"
	"svc-todo/general"
	"svc-todo/model"
	"svc-todo/service"

	"github.com/gin-gonic/gin"
)

type todoController struct {
	todoService service.TodoServiceInterface
}

func TodoController(todoService service.TodoServiceInterface) *todoController {
	return &todoController{todoService}
}

func (controller *todoController) CreateTodo(context *gin.Context) {

	var request entity.Todo

	error := context.ShouldBindJSON(&request)
	http_status := http.StatusOK
	var response *model.StandardResponse

	if error != nil {
		http_status = http.StatusBadRequest
		response = &model.StandardResponse{
			Status:  error.Error(),
			Message: error.Error(),
		}
	} else {
		request, error = controller.todoService.CreateTodo(&request)

		if error == nil {
			http_status = http.StatusCreated
			response = &model.StandardResponse{
				Status:  general.SuccessMessage,
				Message: general.SuccessMessage,
				Data:    request,
			}
		} else {
			http_status = http.StatusBadRequest
			response = &model.StandardResponse{
				Status:  error.Error(),
				Message: error.Error(),
			}
		}
	}

	context.JSON(http_status, response)
}

func (controller *todoController) GetTodo(context *gin.Context) {
	var query entity.Todo
	context.ShouldBind(&query)
	http_status := http.StatusOK
	var response *model.StandardResponse

	result, error := controller.todoService.GetTodo(&query)

	if error == nil {
		http_status = http.StatusOK
		response = &model.StandardResponse{
			Status:  general.SuccessMessage,
			Message: general.SuccessMessage,
			Data:    result,
		}
	} else {
		http_status = http.StatusBadRequest
		response = &model.StandardResponse{
			Status:  error.Error(),
			Message: error.Error(),
		}
	}

	context.JSON(http_status, response)
}

func (controller *todoController) GetOneTodo(context *gin.Context) {

	id := context.Param("id")
	http_status := http.StatusOK
	var response *model.StandardResponse

	result, error := controller.todoService.GetOneTodo(&id)

	if error == nil {
		http_status = http.StatusOK
		response = &model.StandardResponse{
			Status:  general.SuccessMessage,
			Message: general.SuccessMessage,
			Data:    result,
		}
	} else {
		http_status = http.StatusNotFound
		response = &model.StandardResponse{
			Status:  general.NotFoundMessage,
			Message: fmt.Sprintf("Todo with ID %s Not Found", id),
		}
	}

	context.JSON(http_status, response)
}

func (controller *todoController) DeleteTodo(context *gin.Context) {

	id := context.Param("id")
	http_status := http.StatusOK
	var response *model.StandardResponse

	error := controller.todoService.DeleteTodo(&id)

	if error == nil {
		http_status = http.StatusOK
		response = &model.StandardResponse{
			Status:  general.SuccessMessage,
			Message: general.SuccessMessage,
		}
	} else {
		http_status = http.StatusNotFound
		response = &model.StandardResponse{
			Status:  general.NotFoundMessage,
			Message: fmt.Sprintf("Todo with ID %s Not Found", id),
		}
	}

	context.JSON(http_status, response)
}

func (controller *todoController) UpdateTodo(context *gin.Context) {

	var request entity.Todo
	id := context.Param("id")

	error := context.ShouldBindJSON(&request)
	http_status := http.StatusOK
	var response *model.StandardResponse

	if error != nil {
		http_status = http.StatusBadRequest
		response = &model.StandardResponse{
			Status:  error.Error(),
			Message: error.Error(),
		}
	} else {
		request, error = controller.todoService.UpdateTodo(&request, &id)

		if error == nil {
			http_status = http.StatusOK
			response = &model.StandardResponse{
				Status:  general.SuccessMessage,
				Message: general.SuccessMessage,
				Data:    request,
			}
		} else {
			http_status = http.StatusBadRequest
			response = &model.StandardResponse{
				Status:  error.Error(),
				Message: error.Error(),
			}
		}
	}

	context.JSON(http_status, response)
}
