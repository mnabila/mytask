package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mytask/common"
	"github.com/mnabila/mytask/internal/entities"
	"github.com/mnabila/mytask/internal/services"
	"gorm.io/gorm"
)

type TodoController struct {
	Todo *services.TodoService
}

func NewTodoController(db *gorm.DB) *TodoController {
	return &TodoController{
		Todo: services.NewTodoService(db),
	}
}

func (ctrl TodoController) GetTodo(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*entities.UserClaims)

	todo, fail := ctrl.Todo.GetTodo(claims.Id)
	if fail != nil {
		return c.Status(fail.StatusCode).JSON(entities.ApiResponse{
			Success: false,
			Message: fail.Message,
		})
	}
	return c.JSON(entities.ApiResponse{
		Success: true,
		Data:    todo,
	})
}

func (ctrl TodoController) GetTodoById(c *fiber.Ctx) error {
	todoId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: fiber.ErrBadRequest.Message,
		})
	}

	todo, fail := ctrl.Todo.GetTodoById(todoId)
	if fail != nil {
		return c.Status(fail.StatusCode).JSON(entities.ApiResponse{
			Success: false,
			Message: fail.Message,
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Data:    todo,
	})
}
func (ctrl TodoController) CreateTodo(c *fiber.Ctx) error {
	var req entities.TodoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(entities.ApiResponse{
			Success: false,
			Message: fiber.ErrUnprocessableEntity.Message,
		})
	}

	if err := common.ValidateRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: fiber.ErrBadRequest.Message,
			Data:    err,
		})
	}

	claims := c.Locals("claims").(*entities.UserClaims)
	todo, fail := ctrl.Todo.CreateTodo(claims.Id, req)
	if fail != nil {
		return c.Status(fail.StatusCode).JSON(entities.ApiResponse{
			Success: false,
			Message: fail.Message,
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "success",
		Data:    todo,
	})
}

func (ctrl TodoController) UpdateTodoById(c *fiber.Ctx) error {
	todoId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: fiber.ErrBadRequest.Message,
		})
	}

	var req entities.TodoRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: fiber.ErrBadRequest.Message,
		})
	}

	if err := common.ValidateRequest(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: fiber.ErrBadRequest.Message,
			Data:    err,
		})
	}

	claims := c.Locals("claims").(*entities.UserClaims)
	todo, fail := ctrl.Todo.UpdateTodo(claims.Id, todoId, req)
	if fail != nil {
		return c.Status(fail.StatusCode).JSON(entities.ApiResponse{
			Success: false,
			Message: fail.Message,
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "success",
		Data:    todo,
	})
}

func (ctrl TodoController) DeleteTodoById(c *fiber.Ctx) error {
	todoId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(entities.ApiResponse{
			Success: false,
			Message: fiber.ErrBadRequest.Message,
		})
	}

	if fail := ctrl.Todo.DeleteTodoById(todoId); fail != nil {
		return c.Status(fail.StatusCode).JSON(entities.ApiResponse{
			Success: false,
			Message: fail.Message,
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "success",
	})
}
