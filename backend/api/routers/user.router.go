package routers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mytask/api/controllers"
	"github.com/mnabila/mytask/api/middlewares"
	"gorm.io/gorm"
)

func UseUserRouter(db *gorm.DB, r fiber.Router) {
	secret := os.Getenv("JWT_SECRET")
	ctrl := controllers.NewUserController(db, secret)
	auth := middlewares.UseJWTMiddleware(secret)

	r.Post("/user/login", ctrl.LoginUser)
	r.Post("/user/register", ctrl.RegisterNewUser)
	r.Get("/user/profile", auth, ctrl.GetProfileUser)
	r.Put("/user/password", auth, ctrl.UpdatePassword)

}
