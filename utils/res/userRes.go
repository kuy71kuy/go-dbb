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
