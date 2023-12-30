package model

import (
	"gorm.io/gorm"
)

type Menu struct {
	*gorm.Model
	Name     string `json:"name" form:"name"`
	Price    int    `json:"price" form:"price"`
	Category string `json:"category" form:"category"`
	Stock    int    `json:"stock" form:"stock"`
}
