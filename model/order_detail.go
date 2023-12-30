package model

import "gorm.io/gorm"

type OrderDetail struct {
	*gorm.Model
	OrderId int `json:"order_id" form:"order_id"`
	MenuId  int `json:"menu_id" form:"menu_id"`
	Amount  int `json:"amount" form:"amount"`
}
