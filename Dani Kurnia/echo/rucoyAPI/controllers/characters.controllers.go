package controllers

import (
	"echo/rucoyAPI/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetDetailCharacters(c echo.Context) error {
	name := c.Param("name")

	characters, _ := models.GetCharacters(name)

	return c.JSON(http.StatusOK, characters)
}