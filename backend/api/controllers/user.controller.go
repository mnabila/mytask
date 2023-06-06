package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mnabila/mytask/common"
	"github.com/mnabila/mytask/internal/entities"
	"github.com/mnabila/mytask/internal/services"
	"gorm.io/gorm"
)

type UserController struct {
	User *services.UserService
}

func NewUserController(db *gorm.DB, JwtSecret string) *UserController {
	return &UserController{
		User: services.NewUserService(db, JwtSecret),
	}
}

func (ctrl UserController) LoginUser(c *fiber.Ctx) error {
	var req entities.AuthenticationRequest
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

	user, fail := ctrl.User.Authentication(req)
	if fail != nil {
		return c.Status(fail.StatusCode).JSON(entities.ApiResponse{
			Success: false,
			Message: fail.Message,
		})
	}

	token, err := ctrl.User.GenerateAccessToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	c.Set(fiber.HeaderAuthorization, token)

	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "success",
		Data:    user,
	})

}

func (ctrl UserController) RegisterNewUser(c *fiber.Ctx) error {
	var req entities.UserRequest

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

	user, fail := ctrl.User.RegisterUser(req)
	if fail != nil {
		return c.Status(fail.StatusCode).JSON(entities.ApiResponse{
			Success: false,
			Message: fail.Message,
		})
	}

	token, err := ctrl.User.GenerateAccessToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(entities.ApiResponse{
			Success: false,
			Message: err.Error(),
		})
	}
	c.Set(fiber.HeaderAuthorization, token)

	return c.JSON(entities.ApiResponse{
		Success: true,
		Data:    user,
	})
}

func (ctrl UserController) GetProfileUser(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*entities.UserClaims)
	user, fail := ctrl.User.GetUserById(claims.Id)
	if fail != nil {
		return c.Status(fail.StatusCode).JSON(entities.ApiResponse{
			Success: false,
			Message: fail.Message,
		})
	}

	return c.JSON(entities.ApiResponse{
		Success: true,
		Data:    user,
	})
}

func (ctrl UserController) UpdatePassword(c *fiber.Ctx) error {
	var req entities.UpdatePasswordRequest

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
	if fail := ctrl.User.UpdatePassword(claims.Id, req); fail != nil {
		return c.Status(fail.StatusCode).JSON(entities.ApiResponse{
			Success: false,
			Message: fail.Message,
		})
	}
	return c.JSON(entities.ApiResponse{
		Success: true,
		Message: "Sukses",
	})
}
