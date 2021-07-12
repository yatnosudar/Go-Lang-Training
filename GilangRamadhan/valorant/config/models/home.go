package models

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type Tampilan struct {
	Status string `json:"status"`
	Biru   Merah
}

type Aku struct {
	Downloading string `json:"downloading"`
	Videos      string `json:"videos"`
}

type LatestNews struct {
	Page     int    `json:"page"`
	Image    string `json:"image"`
	Date     string `json:"date"`
	Title    string `json:"title"`
	Category string `json:"category"`
}

type DescriptionGame struct {
	Video       string `json:"video"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AgentHome struct {
	Video       string `json:"video"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Page        string `json:"page"`
}

type MapsHome struct {
	Video       string `json:"video"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Page        string `json:"page"`
}

type Merah struct {
	Merah   []LatestNews
	KamuAku Aku
	Game    DescriptionGame
	Agent   AgentHome
	Maps    AgentHome
}

func GetHome(c echo.Context) error {
	var rumah []LatestNews
	for i := 0; i < 4; i++ {
		rumah = append(rumah, LatestNews{
			Page:     i,
			Image:    fmt.Sprintf("image %d", i),
			Date:     fmt.Sprintf("date %d", i),
			Title:    fmt.Sprintf("title %d", i),
			Category: fmt.Sprintf("category %d", i),
		})
	}
	return c.JSON(http.StatusOK, Tampilan{
		Biru: Merah{
			Merah: rumah,
			KamuAku: Aku{
				Downloading: "http:playvalorant/downloading",
				Videos:      "http:playvalorant/videos",
			},
			Game: DescriptionGame{
				Video:       "http:playvalorant/videos",
				Title:       "masukan title",
				Description: "masukan description",
			},
		},
		Status: "success",
	})
}
