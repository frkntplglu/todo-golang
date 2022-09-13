package services

import (
	"ToDo/models"
	"ToDo/repositories"
)

type TodoService interface {
	Get() ([]models.TodoModel, error)
	Create(model models.TodoModel) (uint, error)
	Edit(id uint) (models.TodoModel, error)
	Delete(id uint)
}

type todoService struct {
	serviceRepo repositories.TodoRepository
}

var _ TodoService = todoService{}

func NewTodoService(todoRepo repositories.TodoRepository) TodoService {
	return todoService{todoRepo}
}

func (todo todoService) Get() ([]models.TodoModel, error) {
	return todo.serviceRepo.Get()
}

func (todo todoService) Create(model models.TodoModel) (uint, error) {
	return todo.serviceRepo.Create(model)
}

func (todo todoService) Edit(id uint) (models.TodoModel, error) {
	//TODO implement me
	panic("implement me")
}

func (todo todoService) Delete(id uint) {
	//TODO implement me
	panic("implement me")
}
