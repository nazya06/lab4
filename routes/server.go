package routes

import (
	"github.com/labstack/echo/v4"
	"lab4module/data"
	"net/http"
)

// e.GET("/users/:id", getUser)
func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	user := data.User{
		Id:          id,
		Name:        "TestName",
		PhoneNumber: "+77776665544",
		Age:         20,
		Surname:     "TestSurname",
	}
	return c.JSON(http.StatusOK, user)
}

// e.POST("/users", routes.SaveUser)
func SaveUser(c echo.Context) error {
	user := new(data.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}

// e.PUT("/users/:id", routes.UpdateUser)
func UpdateUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")

	return c.String(http.StatusOK, id)
}

// e.DELETE("/users/:id", routes.DeleteUser)
func DeleteUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// e.GET("/users/filter?age=20", routes.FilterUsers)
func FilterUsers(c echo.Context) error {
	// Get age from the query
	var request data.FilterUserRequest
	if err := c.Bind(&request); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, request)
}
