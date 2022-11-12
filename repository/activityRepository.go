package repository

import (
	"svc-todo/entity"
)

type ActivityRepositoryInterface interface {
	CreateActivity(request *entity.Activity) (entity.Activity, error)
}

func (repo *repository) CreateActivity(request *entity.Activity) (entity.Activity, error) {
	var activity entity.Activity

	result := repo.db.Create(&request).Last(&activity)

	return activity, result.Error
}
