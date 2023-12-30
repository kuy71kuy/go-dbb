package model

import "gorm.io/gorm"

type Payment struct {
	*gorm.Model
	UserId      int    `json:"user_id" form:"user_id"`
	OrderId     int    `json:"order_id" form:"order_id"`
	Amount      int    `json:"amount" form:"amount"`
	Status      string `json:"status" form:"status"`
	ReferenceNo string `json:"reference_no" form:"reference_no"`
}
