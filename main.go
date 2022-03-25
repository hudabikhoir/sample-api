package main

import (
	"echo-sample-api/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		result := map[string]string{
			"response_code": "200",
			"message":       "Success!",
			"version":       "1.0.0-alpha-1",
		}
		return c.JSON(http.StatusOK, result)
	})

	e.GET("/user/list", controllers.GetUser)
	e.POST("/user", controllers.CreateUser)
	e.GET("/user/:id", controllers.GetUserByID)
	e.PUT("/user/:id", controllers.UpdateUserByID)
	e.DELETE("/user/:id", controllers.DeleteUserByID)

	e.Logger.Fatal(e.Start(":4003"))
}
