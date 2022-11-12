package repository

import (
	"svc-todo/entity"
)

type ActivityRepositoryInterface interface {
	CreateActivity(request *entity.Activity) (entity.Activity, error)
	GetActivity() ([]entity.Activity, error)
	GetOneActivity(id *string) (entity.Activity, error)
	DeleteActivity(id *string) error
	UpdateActivity(request *entity.Activity, id *string) (entity.Activity, error)
}

func (repo *repository) CreateActivity(request *entity.Activity) (entity.Activity, error) {
	var activity entity.Activity

	result := repo.db.Create(&request).Last(&activity)

	return activity, result.Error
}

func (repo *repository) GetActivity() ([]entity.Activity, error) {
	var activities []entity.Activity

	error := repo.db.Find(&activities).Error

	return activities, error
}

func (repo *repository) GetOneActivity(id *string) (entity.Activity, error) {
	var activities entity.Activity

	error := repo.db.First(&activities, id).Error

	return activities, error
}

func (repo *repository) DeleteActivity(id *string) error {
	var activities entity.Activity

	error := repo.db.Delete(&activities, id).Error

	return error
}

func (repo *repository) UpdateActivity(request *entity.Activity, id *string) (entity.Activity, error) {
	var activity entity.Activity

	error := repo.db.Model(&activity).Where("id = ?", id).Updates(request).Find(&activity).Error

	return activity, error
}
