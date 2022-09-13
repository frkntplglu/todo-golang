package handlers

import (
	"ToDo/models"
	"ToDo/services"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler interface {
	Get(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Edit(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type todoHandler struct {
	todoService services.TodoService
}

func NewTodoHandler(todoService services.TodoService) TodoHandler {
	return todoHandler{todoService: todoService}
}

type TodoResponse struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func (todo todoHandler) Get(ctx *fiber.Ctx) error {
	/* id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ctx.Status(400).JSON(TodoResponse{Error: err.Error()})
	} */

	model, err := todo.todoService.Get()
	if err != nil {
		return ctx.Status(400).JSON(TodoResponse{Error: err.Error()})
	}

	return ctx.Status(200).JSON(TodoResponse{Data: model})
}

func (todo todoHandler) Create(ctx *fiber.Ctx) error {
	model := models.TodoModel{}
	err := ctx.BodyParser(&model)
	if err != nil {
		return ctx.Status(400).JSON(TodoResponse{Error: err.Error()})
	}
	_, err = todo.todoService.Create(model)
	if err != nil {
		return ctx.Status(400).JSON(TodoResponse{Error: err.Error()})
	}
	return ctx.SendStatus(201)
}

func (todo todoHandler) Edit(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (todo todoHandler) Delete(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

var _ TodoHandler = todoHandler{}
