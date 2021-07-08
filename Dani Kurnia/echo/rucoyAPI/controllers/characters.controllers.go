package controllers

import (
	"echo/rucoyAPI/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetDetailCharacters(c echo.Context) error {
	name := c.Param("name")

	characters, _ := models.GetCharacters(name)

	return c.JSON(http.StatusOK, characters)
}

func AddCharacters(c echo.Context) error {
	name := c.FormValue("name")
	born := c.FormValue("born")

	result, err := models.AddCharacters(name, born)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCharacters(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateCharacters(conv_id, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
