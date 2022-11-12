package service

import (
	"svc-todo/entity"
	"svc-todo/repository"
	"time"
)

type ActivityServiceInterface interface {
	CreateActivity(request *entity.Activity) (entity.Activity, error)
	GetActivity() ([]entity.Activity, error)
}

type activityService struct {
	activityRepository repository.ActivityRepositoryInterface
}

func ActivityService(activityRepository repository.ActivityRepositoryInterface) *activityService {
	return &activityService{activityRepository}
}

func (service *activityService) CreateActivity(request *entity.Activity) (entity.Activity, error) {
	now := time.Now()

	request.CreatedAt = now
	request.UpdatedAt = now

	return service.activityRepository.CreateActivity(request)
}

func (service *activityService) GetActivity() ([]entity.Activity, error) {

	return service.activityRepository.GetActivity()

}
