package service

import (
	"svc-todo/entity"
	"svc-todo/repository"
)

type TodoServiceInterface interface {
	CreateTodo(request *entity.Todo) (entity.Todo, error)
}

type todoService struct {
	todoRepository repository.TodoRepositoryInterface
}

func TodoService(todoRepository repository.TodoRepositoryInterface) *todoService {
	return &todoService{todoRepository}
}

func (service *todoService) CreateTodo(request *entity.Todo) (entity.Todo, error) {

	return service.todoRepository.CreateTodo(request)
}
