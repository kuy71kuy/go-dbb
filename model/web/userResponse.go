package web

type UserResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type GetUserResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name" form:"name"`
	TableNumber int    `json:"table_number" form:"table_number"`
}
type GetMenuResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name" form:"name"`
	Price    int    `json:"price" form:"price"`
	Category string `json:"category" form:"category"`
	Stock    int    `json:"stock" form:"stock"`
}

type UserLoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
