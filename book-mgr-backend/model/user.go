package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id        int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Role      string         `json:"role"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (User) TableName() string {
	return "t_user"
}
