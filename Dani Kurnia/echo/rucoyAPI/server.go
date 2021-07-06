package main

import (
	"echo/RucoyAPI/db"
	"echo/RucoyAPI/routes"
)

func main() {
	db.Init()

	e := routes.Routes()

	e.Logger.Fatal(e.Start(":2004"))
}
