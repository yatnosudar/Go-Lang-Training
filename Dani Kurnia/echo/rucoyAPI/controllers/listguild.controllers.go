package controllers

import (
	"echo/rucoyAPI/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetListGuild(c echo.Context) error {
	result, err := models.GetListGuild()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}
