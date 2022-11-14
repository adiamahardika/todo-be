package repository

import (
	"svc-todo/entity"
)

type TodoRepositoryInterface interface {
	CreateTodo(request *entity.Todo) (entity.Todo, error)
	GetTodo(query *entity.Todo) ([]entity.Todo, error)
	GetOneTodo(id *string) (entity.Todo, error)
	DeleteTodo(id *string) error
	UpdateTodo(request *entity.Todo, id *string) (entity.Todo, error)
}

func (repo *repository) CreateTodo(request *entity.Todo) (entity.Todo, error) {
	var todo entity.Todo

	result := repo.db.Create(&request).Last(&todo)

	return todo, result.Error
}

func (repo *repository) GetTodo(query *entity.Todo) ([]entity.Todo, error) {
	var todos []entity.Todo

	error := repo.db.Where(query).Find(&todos).Error

	return todos, error
}

func (repo *repository) GetOneTodo(id *string) (entity.Todo, error) {
	var todos entity.Todo

	error := repo.db.First(&todos, id).Error

	return todos, error
}

func (repo *repository) DeleteTodo(id *string) error {
	var todos entity.Todo

	error := repo.db.Delete(&todos, id).Error

	return error
}

func (repo *repository) UpdateTodo(request *entity.Todo, id *string) (entity.Todo, error) {
	var todo entity.Todo

	result := repo.db.First(&todo).Save(request)

	return todo, result.Error
}
