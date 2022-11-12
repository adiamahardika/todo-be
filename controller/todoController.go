package controller

import (
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
