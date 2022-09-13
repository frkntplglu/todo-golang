package services

import (
	"ToDo/models"
	"ToDo/repositories"
)

type TodoService interface {
	Get() ([]models.TodoModel, error)
	Create(model models.TodoModel) (uint, error)
	Edit(id uint64, newItem models.TodoModel) (models.TodoModel, error)
	Delete(id uint64)
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

func (todo todoService) Edit(id uint64, newItem models.TodoModel) (models.TodoModel, error) {
	return todo.serviceRepo.Edit(id, newItem)
}

func (todo todoService) Delete(id uint64) {
	todo.serviceRepo.Delete(id)
}
