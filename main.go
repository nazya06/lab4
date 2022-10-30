package main

import (
	"lab4module/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Naz  Hello, World!")
	})
	e.POST("/users", routes.SaveUser)
	e.GET("/users/:id", routes.GetUser)
	e.PUT("/users/:id", routes.UpdateUser)
	e.DELETE("/users/:id", routes.DeleteUser)
	e.GET("/users/filter", routes.FilterUsers)

	e.Logger.Fatal(e.Start(":8080"))
}
