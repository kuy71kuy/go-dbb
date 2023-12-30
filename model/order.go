package model

import "gorm.io/gorm"

type Order struct {
	*gorm.Model
	UserId int    `json:"user_id" form:"user_id"`
	Status string `json:"status" form:"status"`
}
