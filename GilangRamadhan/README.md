<h1>#Echo framework Golang</h1>
Echo adalah framework bahasa golang untuk pengembangan aplikasi web.

Instalasi
Untuk menginstal Echo memerlukan versi Go v1.13 atau lebih tinggi. > Gunakanlah go get di directory project anda seperti dibawah ini

	$ cd <PROJECT IN $GOPATH>
	$ go get -u github.com/labstack/echo/...
Hello World
buat file server.go

	package main

	import (
		"net/http"
	
		"github.com/labstack/echo/v4"
	)

	func main() {
		e := echo.New()
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		})
		e.Logger.Fatal(e.Start(":2004"))
	}