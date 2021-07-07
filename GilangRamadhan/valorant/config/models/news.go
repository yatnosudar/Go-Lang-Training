package models

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseNews struct {
	DataNews ContentNews
	Status   string `json:"status"`
}

// type NewsData struct {
// 	Title string `json:"title"`
// }

type NewsFeatured struct {
	ID          int    `json:"id"`
	Image       string `json:"image"`
	Title       string `json:"title"`
	Descraption string `json:"descraption"`
	Date        string `json:"date"`
}

type ContentNews struct {
	Judul       string `json:"judul"`
	ContentNews []NewsFeatured
}

func GetNews(c echo.Context) error {
	var pickNews []NewsFeatured
	for i := 0; i < 5; i++ {
		pickNews = append(pickNews, NewsFeatured{
			ID:          i,
			Image:       fmt.Sprintf("image %d", i),
			Title:       fmt.Sprintf("title %d", i),
			Descraption: fmt.Sprintf("descraption %d", i),
			Date:        fmt.Sprintf("date %d", i),
		})
	}
	return c.JSON(http.StatusOK, BaseNews{
		DataNews: ContentNews{
			ContentNews: pickNews,
			Judul:       "MASUKAN TITLE",
		},
		Status: "success",
	})
}
