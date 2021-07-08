package routes

import (
	"echo/rucoyAPI/controllers"
	"echo/rucoyAPI/middleware"

	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	e := echo.New()
	// Database
	e.GET("/highscore/experience", controllers.FetchAllTopExp)
	e.GET("/highscore/melee", controllers.FetchAllTopMelee)
	e.GET("/highscore/distance", controllers.FetchAllTopDistance)
	e.GET("/highscore/magic", controllers.FetchAllTopMagic)
	e.GET("/highscore/defense", controllers.FetchAllTopDefense)

	e.GET("/news", controllers.GetNews)
	e.POST("/news", controllers.AddNews)
	e.PUT("/news", controllers.UpdateNews)

	e.GET("/guilds", controllers.GetListGuild)
	e.GET("/guilds/:name", controllers.GetDetailGuild)
	e.POST("/guilds", controllers.AddGuild)
	e.PUT("/guilds", controllers.UpdateGuild)
	e.GET("guilds/", controllers.SearchGuild)

	e.GET("/characters/:name", controllers.GetDetailCharacters)
	e.GET("/characters/", controllers.SearchCharacters)
	e.POST("/characters", controllers.AddCharacters)
	e.PUT("/characters", controllers.UpdateCharacters)
	e.GET("/mycharacters/:name", controllers.GetDetailCharactersPrivate, middleware.IsAuthenticated)

	e.GET("/generate-hash-pw/:password", controllers.GenerateHashPassword)
	e.POST("login/", controllers.CheckLogin)

	return e
}
