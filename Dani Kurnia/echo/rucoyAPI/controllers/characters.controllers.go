package controllers

// import (
// 	"echo/db"

// 	"github.com/labstack/echo/v4"
// )

// type BaseDetailCharacters struct {
// 	Status string `json:"status"`
// 	Data   Head
// }
// type Head struct {
// 	ID      int    `json:"id"`
// 	Name    string `json:"name"`
// 	Status  string `json:"status"`
// 	Role    string `json:"role"`
// 	Content TableCharacters
// 	Log     []Log
// }
// type TableCharacters struct {
// 	Header     string `json:"header"`
// 	Name       string `json:"name"`
// 	Level      int    `json:"level"`
// 	Guild      string `json:"guild"`
// 	LastOnline string `json:"last online"`
// 	Born       string `json:"born"`
// }
// type Log struct {
// 	Killer []BaseDetailCharacters
// 	Killed BaseDetailCharacters
// 	Time   string `json:"time"`
// }

// func GetDetailCharacters(c echo.Context) {
// 	con := db.CreateCon()
// 	sqlStatement := "S"

// }
