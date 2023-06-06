package services

import (
	"net/http"

	"github.com/mnabila/mytask/internal/entities"
	"gorm.io/gorm"
)

type TodoService struct {
	DB *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
	db.AutoMigrate(&entities.Todo{})
	return &TodoService{
		DB: db,
	}
}

func (s TodoService) GetTodo(userId string) ([]entities.Todo, *entities.ApiError) {
	var todos []entities.Todo
	if err := s.DB.Find(&todos, "user_id", userId).Error; err != nil {
		return nil, &entities.ApiError{
			StatusCode: http.StatusNotFound,
			Message:    "Todo tidak ditemukan",
		}
	}

	return todos, nil
}

func (s TodoService) GetTodoById(todoId int) (*entities.Todo, *entities.ApiError) {
	var todo entities.Todo
	if err := s.DB.First(&todo, "id", todoId).Error; err != nil {
		return nil, &entities.ApiError{
			StatusCode: http.StatusNotFound,
			Message:    "Todo tidak ditemukan",
		}
	}
	return &todo, nil
}

func (s TodoService) DeleteTodoById(todoId int) *entities.ApiError {
	if err := s.DB.Delete(&entities.Todo{}, todoId).Error; err != nil {
		return &entities.ApiError{
			StatusCode: http.StatusNotFound,
			Message:    "Todo tidak ditemukan",
		}
	}
	return nil
}

func (s TodoService) CreateTodo(userId string, in entities.TodoRequest) (*entities.Todo, *entities.ApiError) {
	todo := entities.Todo{
		Task:        in.Task,
		Description: in.Description,
		UserId:      userId,
	}

	if err := s.DB.Create(&todo).Error; err != nil {
		return nil, &entities.ApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return &todo, nil
}

func (s TodoService) UpdateTodo(userId string, todoId int, in entities.TodoRequest) (*entities.Todo, *entities.ApiError) {
	var todo entities.Todo

	if err := s.DB.Where("user_id", userId).First(&todo, todoId).Error; err != nil {
		return nil, &entities.ApiError{
			StatusCode: http.StatusNotFound,
			Message:    "Todo tidak ditemukan",
		}
	}

	if err := s.DB.Model(&todo).Updates(in).Error; err != nil {
		return nil, &entities.ApiError{
			StatusCode: http.StatusNotFound,
			Message:    "Todo tidak ditemukan",
		}
	}
	return &todo, nil
}
