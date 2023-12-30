package routes

import (
	"app/controller"
	"app/middleware"
	"app/utils"
	m "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func customJWTErrorHandler(c echo.Context, err error) error {
	return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid credentials"))
}
func Init() *echo.Echo {

	e := echo.New()

	e.Use(middleware.NotFoundHandler)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to RESTful API Services test")
	})
	eJwt := e.Group("")

	eJwt.Use(m.WithConfig(m.Config{
		SigningKey:   []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler: customJWTErrorHandler, // Set the custom error handler
	}))

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

	return e

}
