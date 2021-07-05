package routes

import (
	"echo/rucoyAPI/controllers"

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
	e.GET("news/detail", controllers.GetNews)

	return e
}
