package controllers

import (
	"echo/rucoyAPI/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetDetailGuild(c echo.Context) error {
	name := c.Param("name")

	guild, _ := models.GetDetailGuild(name)

	return c.JSON(http.StatusOK, guild)
}
