package handler

import (
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/magiconair/properties/assert"
)

func TestAssetHandler_GetAllColors(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		res, req := CreateHTTPReq(http.MethodGet, "/colors/all", nil)

		e := echo.New()

		echoTestContext := e.NewContext(req, res)

		err := GetAllColors(echoTestContext)

		assert.Equal(t, err, nil)
	})
}
