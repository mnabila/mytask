package services

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mnabila/mytask/common"
	"github.com/mnabila/mytask/internal/entities"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	DB        *gorm.DB
	SecretKey string
}

func NewUserService(db *gorm.DB, secretKey string) *UserService {
	db.AutoMigrate(&entities.User{})
	return &UserService{
		DB:        db,
		SecretKey: secretKey,
	}
}

func (s UserService) Authentication(in entities.AuthenticationRequest) (*entities.User, *entities.ApiError) {
	var user entities.User
	if err := s.DB.First(&user, "email", in.Email).Error; err != nil {
		return nil, &entities.ApiError{
			StatusCode: http.StatusNotFound,
			Message:    "User tidak ditemukan",
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return nil, &entities.ApiError{
			StatusCode: http.StatusNotFound,
			Message:    "Email/Sandi tidak cocok",
		}
	}

	return &user, nil
}

func (s UserService) GetUserById(userId string) (*entities.User, *entities.ApiError) {
	var user entities.User

	if err := s.DB.First(&user, "id", userId).Error; err != nil {
		return nil, &entities.ApiError{
			StatusCode: http.StatusNotFound,
			Message:    "User tidak ditemukan",
		}
	}
	return &user, nil
}

func (s UserService) RegisterUser(in entities.UserRequest) (*entities.User, *entities.ApiError) {
	password, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.MinCost)
	if err != nil {
		return nil, &entities.ApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	user := entities.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: string(password),
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return nil, &entities.ApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	return &user, nil
}

func (s UserService) UpdatePassword(userId string, in entities.UpdatePasswordRequest) *entities.ApiError {
	if in.NewPassword != in.ConfirmPassword {
		return &entities.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    "Sandi baru tidak cocok",
		}
	}

	var user entities.User
	if err := s.DB.Find(&user, "id", userId).Error; err != nil {
		return &entities.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    "User tidak ditemukan",
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.OldPassword)); err != nil {
		return &entities.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    "Sandi tidak cocok",
		}
	}

	password, err := bcrypt.GenerateFromPassword([]byte(in.NewPassword), bcrypt.MinCost)
	if err != nil {
		return &entities.ApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Silahkan coba lagi",
		}
	}

	err = s.DB.Model(&user).Update("password", string(password)).Error
	if err != nil {
		return &entities.ApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Silahkan coba lagi",
		}
	}
	return nil
}

func (s *UserService) GenerateAccessToken(user *entities.User) (string, error) {
	claims := &entities.UserClaims{
		Id: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(10 * time.Minute),
			},
		},
	}

	token, err := common.MarshalClaims(s.SecretKey, claims)
	if err != nil {
		return "", err
	}

	return token, nil
}
