package entities

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt   time.Time      `json:"createdAt,omitempty"`
	UpdatedAt   time.Time      `json:"updatedAt,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	Task        string         `gorm:"type:varchar(150);not null" json:"task,omitempty"`
	Description string         `gorm:"type:varchar(500)" json:"description,omitempty"`
	UserId      string         `gorm:"type:uuid;foreignKey:UserId;constraint:OnDelete:CASCADE" json:"userId,omitempty"`
}

type TodoRequest struct {
	Task        string `json:"task,omitempty" validate:"required"`
	Description string `json:"description,omitempty"`
}
