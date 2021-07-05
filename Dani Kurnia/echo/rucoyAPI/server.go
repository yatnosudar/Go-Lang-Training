package main

import (
	"echo/Rucoy/db"
	"echo/Rucoy/routes"
)

func main() {
	db.Init()

	e := routes.Routes()

	e.Logger.Fatal(e.Start(":2004"))
}
