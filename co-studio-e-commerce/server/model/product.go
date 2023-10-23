package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid();column:product_id" json:"product_id"`
	ProductName string     `gorm:"unique" json:"product_name"`
	Price       int        `json:"price"`
	Description string     `json:"description"`
	ImageURL    string     `json:"image_url"`
	Enable      int        `json:"enable"`
	IsUpdate    int        `json:"is_update"`
	WhoUpdate   string     `json:"who_update"`
	UpdateDate  time.Time  `json:"update_date"`
	IsDelete    int        `json:"is_delete"`
	CategoryID  uuid.UUID  `json:"category_id" gorm:"type:uuid"`
	Categories  Categories `gorm:"foreignKey:CategoryID"`
}

type Image struct {
	gorm.Model
	Name        string
	Description string
	Data        []byte
}

func (Product) TableName() string {
	return "product"
}
