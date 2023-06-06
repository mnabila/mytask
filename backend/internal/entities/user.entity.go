package entities

type User struct {
	Id       string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id,omitempty"`
	Name     string `gorm:"type:varchar(30);not null" json:"name,omitempty"`
	Email    string `gorm:"type:varchar(50);unique;not null" json:"email,omitempty"`
	Password string `gorm:"type:varchar(60)" json:"-"`
}

type UserRequest struct {
	Name     string `json:"name,omitempty" validate:"required,alpha"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

type UpdatePasswordRequest struct {
	NewPassword     string `validate:"required" json:"newPassword,omitempty"`
	OldPassword     string `validate:"required" json:"oldPassword,omitempty"`
	ConfirmPassword string `validate:"required" json:"confirmPassword,omitempty"`
}
