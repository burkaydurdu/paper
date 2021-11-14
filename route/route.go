package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"paper.com/paper/handler"
)

func health(c echo.Context) error {
	var content struct {
		Message string `json:"message"`
	}

	content.Message = "Success"

	return c.JSON(http.StatusOK, content)
}

func InitRoute(e *echo.Echo) {
	e.GET("/health", health)
	e.GET("/colors/all", handler.GetAllColors)
	e.GET("/images/all", handler.GetAllImages)
}
