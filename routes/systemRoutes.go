package routes

import (
	"net/http"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func SystemRoutes(db *gorm.DB, e *echo.Echo) {
		
	echo.NotFoundHandler = func(c echo.Context) error {
		return c.Render(http.StatusNotFound, "404.html", nil)
	}

}