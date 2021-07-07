package models

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseMedia struct {
	Ungu   ContentMedia
	Status string `json:"status"`
}

type DataMedia struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Video       string `json:"video"`
}

type PageImage struct {
	ID          int    `json:"id"`
	Image       string `json:"image"`
	Previos     string `json:"previos"`
	Next        string `json:"next"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Download    string `json:"Download"`
}

type ContentMedia struct {
	content []PageImage
	Hijau   DataMedia
}

func GetMedia(c echo.Context) error {
	var pickmedia []PageImage
	for i := 0; i < 5; i++ {
		pickmedia = append(pickmedia, PageImage{
			ID:          i,
			Image:       fmt.Sprintf("image %d", i),
			Previos:     fmt.Sprintf("previos %d", i),
			Next:        fmt.Sprintf("next %d", i),
			Title:       fmt.Sprintf("title %d", i),
			Description: fmt.Sprintf("description %d", i),
			Download:    fmt.Sprintf("download %d", i),
		})
	}
	return c.JSON(http.StatusOK, BaseMedia{
		Ungu: ContentMedia{
			content: pickmedia,
		},
		Status: "success",
	})
}
