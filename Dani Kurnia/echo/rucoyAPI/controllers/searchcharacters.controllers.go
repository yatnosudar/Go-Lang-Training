package controllers

import (
	"echo/rucoyAPI/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ValueCharacters struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func SearchCharacters(c echo.Context) error {
	// var value []ValueCharacters
	search := c.QueryParam("search")

	// for _, v := range value {
	// 	value = append(value, ValueCharacters{
	// 		Name: fmt.Sprintln(v, models.SearchCharacters()),
	// 	})
	// }
	result, _ := models.SearchCharacters(search)

	return c.JSON(http.StatusOK, result)
}
