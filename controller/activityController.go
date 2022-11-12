package controller

import (
	"net/http"
	"svc-todo/entity"
	"svc-todo/general"
	"svc-todo/model"
	"svc-todo/service"

	"github.com/gin-gonic/gin"
)

type activityController struct {
	activityService service.ActivityServiceInterface
}

func ActivityController(activityService service.ActivityServiceInterface) *activityController {
	return &activityController{activityService}
}

func (controller *activityController) Create(context *gin.Context) {

	var request entity.Activity

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
		request, error = controller.activityService.CreateActivity(&request)

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
