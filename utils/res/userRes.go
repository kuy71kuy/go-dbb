package res

import (
	"app/model"
	"app/model/web"
)

func ConvertIndex(users []model.User) []web.GetUserResponse {
	var results []web.GetUserResponse
	for _, user := range users {
		userResponse := web.GetUserResponse{
			Id:          int(user.ID),
			Name:        user.Name,
			TableNumber: user.TableNumber,
		}
		results = append(results, userResponse)
	}

	return results
}

func ConvertIndexMenu(menus []model.Menu) []web.GetMenuResponse {
	var results []web.GetMenuResponse
	for _, menu := range menus {
		menuRespose := web.GetMenuResponse{
			Id:       int(menu.ID),
			Name:     menu.Name,
			Price:    menu.Price,
			Category: menu.Category,
			Stock:    menu.Stock,
		}
		results = append(results, menuRespose)
	}

	return results
}
