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

func IndexMenu(c echo.Context) error {
	var menus []model.Menu

	err := config.DB.Find(&menus).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve menu"))
	}

	if len(menus) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	response := res.ConvertIndexMenu(menus)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Menu data successfully retrieved", response))
}

func ShowMenu(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var menu model.Menu

	if err := config.DB.First(&menu, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve menu"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Menu data successfully retrieved", menu))
}

func StoreMenu(c echo.Context) error {
	var menu model.Menu
	if err := c.Bind(&menu); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	if err := config.DB.Create(&menu).Error; err != nil {
		log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to store menu data"))
	}

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", menu))
}

func UpdateMenu(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}
	var updatedMenu model.Menu

	if err := c.Bind(&updatedMenu); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	var existingMenu model.Menu
	result := config.DB.First(&existingMenu, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve menu"))
	}
	config.DB.Model(&existingMenu).Updates(updatedMenu)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Menu data successfully updated", nil))
}

func DeleteMenu(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}
	var existingMenu model.Menu
	result := config.DB.First(&existingMenu, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve menu"))
	}
	config.DB.Delete(&existingMenu)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Menu data successfully deleted", nil))
}
