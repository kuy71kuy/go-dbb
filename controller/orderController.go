package controller

import (
	"app/config"
	"app/model"
	"app/utils"
	"app/utils/res"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func IndexOrder(c echo.Context) error {
	var orders []model.Order

	err := config.DB.Find(&orders).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve order"))
	}

	if len(orders) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndexOrder(orders)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Order data successfully retrieved", response))
}

func ShowOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var order model.Order

	if err := config.DB.First(&order, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve order"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Order data successfully retrieved", order))
}

func UpdateOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}
	var updatedOrder model.Order

	if err := c.Bind(&updatedOrder); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	var existingOrder model.Order
	result := config.DB.First(&existingOrder, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve order"))
	}
	config.DB.Model(&existingOrder).Updates(updatedOrder)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Order data successfully updated", nil))
}
