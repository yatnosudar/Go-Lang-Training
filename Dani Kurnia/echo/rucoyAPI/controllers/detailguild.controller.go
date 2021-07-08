package controllers

import (
	"echo/rucoyAPI/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetDetailGuild(c echo.Context) error {
	name := c.Param("name")

	guild, _ := models.GetDetailGuild(name)

	return c.JSON(http.StatusOK, guild)
}

func AddGuild(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("reddit")
	date := c.FormValue("facebook")
	members := c.FormValue("date")

	result, err := models.AddGuild(name, description, date, members)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateGuild(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	description := c.FormValue("description")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateGuild(conv_id, name, description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
