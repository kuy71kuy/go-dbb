package controller

import (
	"app/config"
	"app/model"
	"app/utils"
	"app/utils/res"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

func IndexOrderItem(c echo.Context) error {
	var menus []model.OrderItem

	err := config.DB.Find(&menus).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve menu"))
	}

	if len(menus) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndexOrderItem(menus)

	return c.JSON(http.StatusOK, utils.SuccessResponse("OrderItem data successfully retrieved", response))
}

func ShowOrderItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var menu model.OrderItem

	if err := config.DB.First(&menu, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve menu"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("OrderItem data successfully retrieved", menu))
}

func StoreOrderItem(c echo.Context) error {
	var menu model.OrderItem
	if err := c.Bind(&menu); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	if err := config.DB.Create(&menu).Error; err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to store menu data"))
	}

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", menu))
}

func UpdateOrderItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}
	var updatedOrderItem model.OrderItem

	if err := c.Bind(&updatedOrderItem); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	var existingOrderItem model.OrderItem
	result := config.DB.First(&existingOrderItem, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve menu"))
	}
	config.DB.Model(&existingOrderItem).Updates(updatedOrderItem)

	return c.JSON(http.StatusOK, utils.SuccessResponse("OrderItem data successfully updated", nil))
}

func DeleteOrderItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}
	var existingOrderItem model.OrderItem
	result := config.DB.First(&existingOrderItem, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve menu"))
	}
	config.DB.Delete(&existingOrderItem)

	return c.JSON(http.StatusOK, utils.SuccessResponse("OrderItem data successfully deleted", nil))
}
