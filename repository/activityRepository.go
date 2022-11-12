package repository

import "svc-todo/entity"

type ActivityRepositoryInterface interface {
	CreateActivity(request *entity.Activity) (entity.Activity, error)
}

func (repo *repository) CreateActivity(request *entity.Activity) (entity.Activity, error) {
	var activity entity.Activity

	error := repo.db.Create(&request).Find(&activity).Error

	return activity, error
}
