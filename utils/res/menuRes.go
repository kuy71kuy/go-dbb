package res

import (
	"app/model"
	"app/model/web"
)

func ConvertIndexMenu(menus []model.Menu) []web.GetMenuResponse {
	var results []web.GetMenuResponse
	for _, menu := range menus {
		menuResponse := web.GetMenuResponse{
			Id:       int(menu.ID),
			Name:     menu.Name,
			Price:    menu.Price,
			Category: menu.Category,
			Stock:    menu.Stock,
		}
		results = append(results, menuResponse)
	}

	return results
}
