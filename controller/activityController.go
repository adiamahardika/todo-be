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

type activityController struct {
	activityService service.ActivityServiceInterface
}

func ActivityController(activityService service.ActivityServiceInterface) *activityController {
	return &activityController{activityService}
}

func (controller *activityController) CreateActivity(context *gin.Context) {

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

func (controller *activityController) GetActivity(context *gin.Context) {

	http_status := http.StatusOK
	var response *model.StandardResponse

	result, error := controller.activityService.GetActivity()

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

func (controller *activityController) GetOneActivity(context *gin.Context) {

	id := context.Param("id")
	http_status := http.StatusOK
	var response *model.StandardResponse

	result, error := controller.activityService.GetOneActivity(id)

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
			Message: fmt.Sprintf("Activity with ID %s Not Found", id),
		}
	}

	context.JSON(http_status, response)
}

func (controller *activityController) DeleteActivity(context *gin.Context) {

	id := context.Param("id")
	http_status := http.StatusOK
	var response *model.StandardResponse

	error := controller.activityService.DeleteActivity(id)

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
			Message: fmt.Sprintf("Activity with ID %s Not Found", id),
		}
	}

	context.JSON(http_status, response)
}
