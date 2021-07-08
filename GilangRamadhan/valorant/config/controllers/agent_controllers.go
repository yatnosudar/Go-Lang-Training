package controllers

import (
	"net/http"
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
