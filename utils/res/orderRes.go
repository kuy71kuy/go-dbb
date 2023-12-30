package res

import (
	"app/model"
	"app/model/web"
)

func ConvertIndexOrder(orders []model.Order) []web.GetOrderResponse {
	var results []web.GetOrderResponse
	for _, order := range orders {
		orderResponse := web.GetOrderResponse{
			Id:     int(order.ID),
			UserId: order.UserId,
			Status: order.Status,
		}
		results = append(results, orderResponse)
	}

	return results
}
