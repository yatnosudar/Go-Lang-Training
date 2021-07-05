package controllers

import (
	"echo/Rucoy/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchAllTopExp(c echo.Context) error {
	result, err := models.FetchAllTopExp()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchAllTopMelee(c echo.Context) error {
	result, err := models.FetchAllTopMelee()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchAllTopDistance(c echo.Context) error {
	result, err := models.FetchAllTopDistance()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchAllTopMagic(c echo.Context) error {
	result, err := models.FetchAllTopMagic()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchAllTopDefense(c echo.Context) error {
	result, err := models.FetchAllTopDefense()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
