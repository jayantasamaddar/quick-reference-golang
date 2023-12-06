package main

import (
	"net/http"

	"github.com/jayantasamaddar/quick-reference-golang/echo-CRUD-api/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	api := e.Group("/api")

	api.POST("/users", controllers.User.CreateUser)
	api.GET("/users", controllers.User.GetAllUsers)
	api.GET("/users/:id", controllers.User.GetUser)
	api.PUT("/users/:id", controllers.User.UpdateUser)
	api.DELETE("/users/:id", controllers.User.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
