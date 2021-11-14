package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"paper.com/paper/model"
	"paper.com/paper/repository"
	"paper.com/paper/service"
)

func GetAllColors(ctx echo.Context) error {
	assetService := service.NewAssetService(repository.NewAssetRepository())
	allColors, err := assetService.GetAllColors()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ColorResponse{
			Error:  err.Error(),
			Colors: allColors,
		})
	}

	return ctx.JSON(http.StatusOK, model.ColorResponse{Colors: allColors})
}

func GetAllImages(ctx echo.Context) error {
	assetService := service.NewAssetService(repository.NewAssetRepository())
	allImages, err := assetService.GetAllImages()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.ImageResponse{
			Error:  err.Error(),
			Images: allImages,
		})
	}

	return ctx.JSON(http.StatusOK, model.ImageResponse{Images: allImages})
}
