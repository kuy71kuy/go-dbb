package model

import (
	"gorm.io/gorm"
)

const (
	AdminRole  string = "admin"
	UserRole   string = "user"
	DummyToken string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6IiIsInJvbGUiOiIiLCJleHAiOjEwMDAwMDAwMDB9.VMb4PNNJ_CUq7oom1P1kx5_6VSLOZIP1ZVg9T4rJIMA"
)

type User struct {
	*gorm.Model
	Name        string `json:"name" form:"name"`
	TableNumber int    `json:"table_number" form:"table_number"`
}
