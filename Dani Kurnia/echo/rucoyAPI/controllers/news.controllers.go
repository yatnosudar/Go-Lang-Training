package controllers

import (
	"echo/rucoyAPI/models"
	"net/http"
	"strconv"

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

func AddNews(c echo.Context) error {
	title := c.FormValue("title")
	reddit := c.FormValue("reddit")
	facebook := c.FormValue("facebook")
	date := c.FormValue("date")
	image := c.FormValue("image")
	description := c.FormValue("description")

	result, err := models.AddNews(title, reddit, facebook, date, image, description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateNews(c echo.Context) error {
	id := c.FormValue("id")
	title := c.FormValue("title")
	reddit := c.FormValue("reddit")
	facebook := c.FormValue("facebook")
	date := c.FormValue("date")
	image := c.FormValue("image")
	description := c.FormValue("description")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateNews(conv_id, title, reddit, facebook, date, image, description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
