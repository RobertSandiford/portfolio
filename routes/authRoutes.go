package routes

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func AuthRoutes(db *gorm.DB, e *echo.Echo) {
	
	////////////////////////
	// Auth
	////////////////////////
	e.GET("/sign-up", func(c echo.Context) error {
		fmt.Println("sign-up")
		return c.Render(http.StatusOK, "sign-up.html", nil)
	})

	e.GET("/log-in", func(c echo.Context) error {
		return c.Render(http.StatusOK, "log-in.html", nil)
	})

	e.GET("/log-in-submit", func(c echo.Context) error {
		data := StandardTemplate{
			Header: "Logged in",
			Text:   "Welcome back, sunshine",
		}
		return c.Render(http.StatusOK, "standard.html", data)
	})

	e.GET("/log-out", func(c echo.Context) error {
		data := StandardTemplate{
			Header: "Logged out",
			Text:   "Come back soon, OK?",
		}
		return c.Render(http.StatusOK, "standard.html", data)
	})

}