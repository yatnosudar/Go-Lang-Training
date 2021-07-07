package models

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Base struct {
	Data   Content
	Status string
}

type MapsName struct {
	ID    int
	Title string
	Image string
	Desc  string
	View  string
}

type Content struct {
	Content []MapsName
}

func GetMaps(c echo.Context) error {
	var pick []MapsName
	for i := 0; i < 5; i++ {
		pick = append(pick, MapsName{
			ID:    i,
			Title: fmt.Sprintf("Maps %d", i),
			Image: fmt.Sprintf("image %d", i),
			Desc:  fmt.Sprintf("desc %d", i),
			View:  fmt.Sprintf("View %d", i),
		})
	}
	return c.JSON(http.StatusOK, &Base{
		Data: Content{
			Content: pick,
		},
		Status: "success",
	})
}

func DetailMaps(c echo.Context) error {
	Mapsid := c.QueryParam("id")
	id, _ := strconv.Atoi(Mapsid)
	Maps := MapsName{
		ID:    id,
		Title: fmt.Sprintf("Maps %d", 1),
		Image: fmt.Sprintf("Maps %d", 1),
		Desc:  fmt.Sprintf("Maps %d", 1),
		View:  fmt.Sprintf("Maps %d", 1),
	}
	return c.JSON(http.StatusOK, Maps)

}
