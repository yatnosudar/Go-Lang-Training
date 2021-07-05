package controllers

import (
	"echo/Rucoy/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetNews(c echo.Context) error {
	result, err := models.GetNewsDetail()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
