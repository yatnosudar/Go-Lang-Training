package controllers

import (
	"echo/rucoyAPI/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SearchGuild(c echo.Context) error {
	search := c.QueryParam("search")
	result, _ := models.SearchGuild(search)

	return c.JSON(http.StatusOK, result)
}
