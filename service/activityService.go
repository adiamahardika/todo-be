package service

import (
	"svc-todo/entity"
	"svc-todo/repository"
)

type ActivityServiceInterface interface {
	CreateActivity(request *entity.Activity) (entity.Activity, error)
	GetActivity() ([]entity.Activity, error)
	GetOneActivity(id *string) (entity.Activity, error)
	DeleteActivity(id *string) error
	UpdateActivity(request *entity.Activity, id *string) (entity.Activity, error)
}

type activityService struct {
	activityRepository repository.ActivityRepositoryInterface
}

func ActivityService(activityRepository repository.ActivityRepositoryInterface) *activityService {
	return &activityService{activityRepository}
}

func (service *activityService) CreateActivity(request *entity.Activity) (entity.Activity, error) {

	return service.activityRepository.CreateActivity(request)
}

func (service *activityService) GetActivity() ([]entity.Activity, error) {

	return service.activityRepository.GetActivity()

}

func (service *activityService) GetOneActivity(id *string) (entity.Activity, error) {

	return service.activityRepository.GetOneActivity(id)

}

func (service *activityService) DeleteActivity(id *string) error {

	return service.activityRepository.DeleteActivity(id)

}

func (service *activityService) UpdateActivity(request *entity.Activity, id *string) (entity.Activity, error) {

	return service.activityRepository.UpdateActivity(request, id)
}
