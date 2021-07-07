package controllers

import (
	"echo/rucoyAPI/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetDetailCharactersPrivate(c echo.Context) error {
	name := c.Param("name")

	characters, _ := models.GetMyChar(name)

	return c.JSON(http.StatusOK, characters)
}
