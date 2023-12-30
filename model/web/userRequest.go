package web

type UserRequest struct {
	Name        string `json:"name" form:"name"`
	TableNumber string `json:"table_number" form:"table_number"`
}
