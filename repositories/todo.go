package repositories

import (
	"ToDo/models"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Get() ([]models.TodoModel, error)
	Create(model models.TodoModel) (uint, error)
	Edit(id uint64, newItem models.TodoModel) (models.TodoModel, error)
	Delete(id uint64)
	Migration() error
}
type todoRepository struct {
	Db *gorm.DB
}

var _ TodoRepository = todoRepository{}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return todoRepository{Db: db}
}

func (todo todoRepository) Get() ([]models.TodoModel, error) {
	var todos []models.TodoModel
	todo.Db.Find(&todos)

	return todos, nil
}

func (todo todoRepository) Create(model models.TodoModel) (uint, error) {
	err := todo.Db.Create(&model)
	if err != nil {
		return 0, nil
	}

	return model.ID, nil
}

func (todo todoRepository) Edit(id uint64, newItem models.TodoModel) (models.TodoModel, error) {
	var todoItem models.TodoModel
	//todoItem = todo.Db.First(&todoItem, id)
	todo.Db.Model(todoItem).Where("id = ?", id).Updates(newItem)
	return todoItem, nil
}

func (todo todoRepository) Delete(id uint64) {
	var todoItem models.TodoModel
	todo.Db.Delete(&todoItem, id)

}

func (todo todoRepository) Migration() error {
	return todo.Db.AutoMigrate(&models.TodoModel{})
}
