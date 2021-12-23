package infrastructure

import (
	"net/http"

	"github.com/ga-tcho/go-rest-api/interfaces/controllers"
	"github.com/labstack/echo/v4"
)

func Init() {
	e := echo.New()
	userController := controllers.NewUserController(NewMySqlDb())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users", func(c echo.Context) error { return userController.GetUsers(c) })
	e.Logger.Fatal(e.Start(":8080"))
}
