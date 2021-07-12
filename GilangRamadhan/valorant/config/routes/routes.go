package routes

import (
	"valorant/controllers"
	"valorant/middleware"

	"valorant/models"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/maps", models.GetMaps, middleware.IsAuthenticated)
	e.GET("/agent", models.GetAgent, middleware.IsAuthenticated)
	e.GET("/arsenal", models.GetArsenal, middleware.IsAuthenticated)
	e.GET("/home", models.GetHome, middleware.IsAuthenticated)
	e.GET("/news", models.GetNews, middleware.IsAuthenticated)
	e.GET("/leaderboards", models.GetLeaderboards, middleware.IsAuthenticated)

	e.GET("/agent/", models.DetailAgent, middleware.IsAuthenticated)
	e.GET("/arsenal/", models.DetailArenal, middleware.IsAuthenticated)
	e.GET("/leaderboards/", models.DetailLeaderboards, middleware.IsAuthenticated)
	e.GET("/maps/", models.DetailMaps, middleware.IsAuthenticated)

	e.POST("/detailagent", models.SearchDetailAgent)
	e.POST("/leaderboards", models.SearchDetailLeaderboards)

	//CRUD

	e.GET("/agent/data", controllers.FetchAllAgent, middleware.IsAuthenticated)
	e.GET("/arsenal/data", controllers.FetchAllArsenal, middleware.IsAuthenticated)

	e.POST("/agent/tambahagent", controllers.StoreAgent, middleware.IsAuthenticated)
	e.POST("/arsenal/tambaharsenal", controllers.StoreArsenal, middleware.IsAuthenticated)

	e.PUT("/agent/update", controllers.UpdateAgent, middleware.IsAuthenticated)

	e.DELETE("/agent/delete", controllers.DeleteAgent, middleware.IsAuthenticated)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)

	e.GET("/test-struct-validation", controllers.TestStructValidation)
	e.GET("/test-variabel-validation", controllers.TestVariableValidation)

	return e
}
