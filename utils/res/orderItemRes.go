package res

import (
	"app/model"
	"app/model/web"
)

func ConvertIndexOrderItem(orderItems []model.OrderItem) []web.GetOrderItemResponse {
	var results []web.GetOrderItemResponse
	for _, orderItem := range orderItems {
		orderItemResponse := web.GetOrderItemResponse{
			Id:      int(orderItem.ID),
			OrderId: orderItem.OrderId,
			MenuId:  orderItem.MenuId,
			Amount:  orderItem.Amount,
		}
		results = append(results, orderItemResponse)
	}

	return results
}
