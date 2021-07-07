package routes

import (
	"valorant/controllers"
	"valorant/models"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/maps", models.GetMaps)
	e.GET("/agent", models.GetAgent)
	e.GET("/arsenal", models.GetArsenal)
	e.GET("/home", models.GetHome)
	e.GET("/news", models.GetNews)
	e.GET("/leaderboards", models.GetLeaderboards)

	e.GET("/maps", models.DetailAgent)
	e.GET("/arsenal", models.DetailArenal)
	e.GET("/leaderboards", models.DetailLeaderboards)
	e.GET("/maps", models.DetailMaps)

	e.GET("/agent/tambah", controllers.FetchAllAgent)
	e.GET("/arsenal/tambah", controllers.FetchAllArsenal)

	e.POST("/Leaderboards/", models.SearchDetailLeaderboards)

	return e
}