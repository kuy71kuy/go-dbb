package routes

import (
	"app/controller"
	"app/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.NotFoundHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to RESTful API Services test")
	})

	//Manage User
	e.POST("/users/register", controller.Store)
	e.GET("/users", controller.Index)
	e.GET("/users/:id", controller.Show)
	e.PUT("/users/:id", controller.Update)
	e.DELETE("/users/:id", controller.Delete)

	//Manage Menu
	e.POST("/menus", controller.StoreMenu)
	e.GET("/menus", controller.IndexMenu)
	e.GET("/menus/:id", controller.ShowMenu)
	e.PUT("/menus/:id", controller.UpdateMenu)
	e.DELETE("/menus/:id", controller.DeleteMenu)

	//Manage Order
	e.GET("/orders", controller.IndexOrder)
	e.GET("/orders/:id", controller.ShowOrder)
	e.PUT("/orders/:id", controller.UpdateOrder)

	//Manage OrderItem
	e.POST("/orderitem", controller.StoreOrderItem)
	e.GET("/orderitem", controller.IndexOrderItem)
	e.GET("/orderitem/:id", controller.ShowOrderItem)
	e.PUT("/orderitem/:id", controller.UpdateOrderItem)
	e.DELETE("/orderitem/:id", controller.DeleteOrderItem)

	return e

}
