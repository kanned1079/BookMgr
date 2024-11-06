package model

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	Id        int64          `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name      string         `json:"name"`
	Publisher string         `json:"publisher"`
	Year      int32          `json:"year"`
	Remark    string         `json:"remark" gorm:"type:TEXT"`
	Author    string         `json:"author"`
	ISBN      string         `json:"isbn"`
	Price     float64        `json:"price"`
	Residue   int64          `json:"residue"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (Book) TableName() string {
	return "t_books"
}
