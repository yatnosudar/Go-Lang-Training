			𝑬𝑪𝑯𝑶 𝑭𝑹𝑨𝑴𝑬𝑾𝑶𝑹𝑲

Instalasi 

terdapat banyak cara di :
	https://echo.labstack.com/

pertama-tama

 	mkdir myapp && cd myapp
 	go mod init myapp
"go get github.com/labstack/echo/v4"

1.Hello, World!

  	Create Main.go

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
		e.Logger.Fatal(e.Start(":1323"))
  	}