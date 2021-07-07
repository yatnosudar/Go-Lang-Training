package controllers

import (
	"net/http"
	"valorant/models"

	"github.com/labstack/echo/v4"
)

func FetchAllArsenal(c echo.Context) error {
	result, err := models.FetchAllArsenal()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}
