package controllers

import (
	"net/http"
	"strconv"
	"valorant/models"

	"github.com/labstack/echo/v4"
)

func FetchAllAgent(c echo.Context) error {
	result, err := models.FetchAllAgent()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func StoreAgent(c echo.Context) error {
	NameAgent := c.FormValue("nameagent")
	Role := c.FormValue("role")
	DescriptionAgent := c.FormValue("descriptionagent")

	result, err := models.StoreAgent(NameAgent, Role, DescriptionAgent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateAgent(c echo.Context) error {
	No := c.FormValue("no")
	NameAgent := c.FormValue("nameagent")
	Role := c.FormValue("role")
	DescriptionAgent := c.FormValue("descriptionagent")

	conv_No, err := strconv.Atoi(No)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdateAgent(conv_No, NameAgent, Role, DescriptionAgent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteAgent(c echo.Context) error {
	No := c.FormValue("no")

	conv_No, err := strconv.Atoi(No)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeleteAgent(conv_No)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
