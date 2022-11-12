package repository

import "svc-todo/entity"

type TodoRepositoryInterface interface {
	CreateTodo(request *entity.Todo) (entity.Todo, error)
}

func (repo *repository) CreateTodo(request *entity.Todo) (entity.Todo, error) {
	var todo entity.Todo

	result := repo.db.Create(&request).Last(&todo)

	return todo, result.Error
}
