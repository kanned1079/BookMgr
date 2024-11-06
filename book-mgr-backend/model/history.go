package model

import (
	"gorm.io/gorm"
	"time"
)

type History struct {
	gorm.Model
	Id         int64          `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	BorrowId   string         `json:"borrow_id"`
	UserId     int64          `json:"user_id"`
	BookId     int64          `json:"book_id"`
	BorrowedAt *time.Time     `json:"borrowed_at"`
	IsBack     bool           `json:"is_back"`           // 是否归还
	Book       Book           `gorm:"foreignKey:BookId"` // 添加 Book 字段，并标注外键
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func (History) TableName() string {
	return "t_history"
}
