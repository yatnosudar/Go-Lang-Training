package models

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type LeaderboardsBase struct {
	DataLeaderboards LeaderboardsContent
	Status           string `json:"status"`
}

type Mode struct {
	Title string `json:"title"`
	Date  string `json:"date"`
	Page  int    `json:"page"`
}

type Player struct {
	ID       int    `json:"id"`
	Rating   int    `json:"rating"`
	Username string `json:"username"`
	Ongoing  int    `json:"ongoing"`
}

type Next struct {
	Login    string `json:"login"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
}

type LeaderboardsContent struct {
	LeaderboardsContent []Player
}

func GetLeaderboards(c echo.Context) error {
	var pickleaderboards []Player
	for i := 0; i < 5; i++ {
		pickleaderboards = append(pickleaderboards, Player{
			ID:       i,
			Rating:   i,
			Username: fmt.Sprintf("Usernmae player %d", i),
			Ongoing:  i,
		})
	}
	return c.JSON(http.StatusOK, LeaderboardsBase{
		DataLeaderboards: LeaderboardsContent{
			LeaderboardsContent: pickleaderboards,
		},
		Status: "success",
	})
}

func DetailLeaderboards(c echo.Context) error {
	Playerid := c.QueryParam("id")
	id, _ := strconv.Atoi(Playerid)
	leaderboards := Player{
		ID:       id,
		Rating:   12,
		Username: fmt.Sprintf("Usernameplayer %d", 1),
		Ongoing:  100,
	}
	return c.JSON(http.StatusOK, leaderboards)

}

func SearchDetailLeaderboards(c echo.Context) error {
	var value []Player
	SearchLeaderboards := c.QueryParam("search")
	id, _ := strconv.Atoi(SearchLeaderboards)
	for _, v := range value {
		value = append(value, Player{
			ID:       id,
			Rating:   133,
			Username: fmt.Sprintln(v, SearchLeaderboards),
			Ongoing:  111,
		})
	}
	return c.JSON(http.StatusOK, &Player{
		ID:       id,
		Username: fmt.Sprint(SearchLeaderboards),
	})

}
