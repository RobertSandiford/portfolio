package routes

import (
	"net/http"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func InvestorRoutes(db *gorm.DB, e *echo.Echo) {
	
	////////////////////////
	// Investor pages
	////////////////////////

	e.GET("/portfolio", func(c echo.Context) error {
		return c.Render(http.StatusOK, "portfolio.html", nil)
	})

	e.GET("/funding", func(c echo.Context) error {
		return c.Render(http.StatusOK, "funding.html", nil)
	})

	e.GET("/explore", func(c echo.Context) error {
		return c.Render(http.StatusOK, "explore.html", nil)
	})

	e.GET("/managers", func(c echo.Context) error {
		return c.Render(http.StatusOK, "managers.html", nil)
	})

}