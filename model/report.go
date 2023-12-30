package model

import (
	"gorm.io/gorm"
	"time"
)

type Report struct {
	*gorm.Model
	StartDate time.Time `json:"start_date" form:"start_date"`
	EndDate   time.Time `json:"end_date" form:"end_date"`
	Amount    int       `json:"amount" form:"amount"`
}
