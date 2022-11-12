package service

import (
	"svc-todo/entity"
	"svc-todo/repository"
)

type TodoServiceInterface interface {
	CreateTodo(request *entity.Todo) (entity.Todo, error)
	GetTodo(query *entity.Todo) ([]entity.Todo, error)
	GetOneTodo(id *string) (entity.Todo, error)
	DeleteTodo(id *string) error
	UpdateTodo(request *entity.Todo, id *string) (entity.Todo, error)
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

func (service *todoService) GetTodo(query *entity.Todo) ([]entity.Todo, error) {

	return service.todoRepository.GetTodo(query)

}

func (service *todoService) GetOneTodo(id *string) (entity.Todo, error) {

	return service.todoRepository.GetOneTodo(id)

}

func (service *todoService) DeleteTodo(id *string) error {

	return service.todoRepository.DeleteTodo(id)

}

func (service *todoService) UpdateTodo(request *entity.Todo, id *string) (entity.Todo, error) {

	return service.todoRepository.UpdateTodo(request, id)
}
