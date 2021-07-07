package routes

import (
	"echo/rucoyAPI/controllers"
	"echo/rucoyAPI/middleware"

	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	e := echo.New()
	// Database
	e.GET("highscore/experience", controllers.FetchAllTopExp)
	e.GET("highscore/melee", controllers.FetchAllTopMelee)
	e.GET("highscore/distance", controllers.FetchAllTopDistance)
	e.GET("highscore/magic", controllers.FetchAllTopMagic)
	e.GET("highscore/defense", controllers.FetchAllTopDefense)
	e.GET("news", controllers.GetNews)
	e.GET("guilds", controllers.GetListGuild)
	e.GET("characters/:name", controllers.GetDetailCharacters)
	e.GET("guilds/:name", controllers.GetDetailGuild)
	e.GET("characters/", controllers.SearchCharacters)
	e.GET("guilds/", controllers.SearchGuild)

	e.GET("mycharacters/:name", controllers.GetDetailCharactersPrivate, middleware.IsAuthenticated)

	return e
}
