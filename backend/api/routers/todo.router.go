package routers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mytask/api/controllers"
	"github.com/mnabila/mytask/api/middlewares"
	"gorm.io/gorm"
)

func UseTodoRouter(db *gorm.DB, r fiber.Router) {
	secret := os.Getenv("JWT_SECRET")
	ctrl := controllers.NewTodoController(db)
	auth := middlewares.UseJWTMiddleware(secret)

	r.Get("/todo", auth, ctrl.GetTodo)
	r.Post("/todo", auth, ctrl.CreateTodo)
	r.Get("/todo/:id", auth, ctrl.GetTodoById)
	r.Put("/todo/:id", auth, ctrl.UpdateTodoById)
	r.Delete("/todo/:id", auth, ctrl.DeleteTodoById)

}
