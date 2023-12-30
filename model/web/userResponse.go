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

type GetOrderResponse struct {
	Id     int    `json:"id"`
	UserId int    `json:"user_id" form:"user_id"`
	Status string `json:"status" form:"status"`
}

type GetOrderItemResponse struct {
	Id      int `json:"id"`
	OrderId int `json:"order_id" form:"order_id"`
	MenuId  int `json:"menu_id" form:"menu_id"`
	Amount  int `json:"amount" form:"amount"`
}

type UserLoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
