package models

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseSupport struct {
	DataSupport ContentSupport
	Status      string `json:"status"`
}
type HeaderSupport struct {
	Title string `json:"title"`
}
type Spec struct {
	Title     string `json:"title"`
	DecWindow string `json:"dec_window"`
	DecRAM    string `json:"dec_ram"`
	DecVram   string `json:"dec_vram"`
}
type MinMaxFps struct {
	MinimumSpecs30Fps     []string `json:"Minimum specs / 30 fps"`
	RecommendedSpecs60Fps []string `json:"Recommended specs / 60 fps"`
	HIGHENDSpecs144Fps    []string `json:"HIGH END specs / 144+ fps"`
}

type ContentSupport struct {
	content MinMaxFps
}

func GetSupport(c echo.Context) error {
	return c.JSON(http.StatusOK, BaseSupport{
		DataSupport: ContentSupport{
			content: MinMaxFps{
				MinimumSpecs30Fps: []string{
					"cpu : Intel core 2 Duo E8400",
					"gpu : Intel HD 4000",
				},
				RecommendedSpecs60Fps: []string{
					"cpu : Intel i3-4150",
					"gpu : Geforce GT 730",
				},
				HIGHENDSpecs144Fps: []string{
					"cpu : Intel core i5-4460 3.2GHZ",
					"gpu : GTX 1050 TI",
				},
			},
		},
		Status: "success",
	})
}
