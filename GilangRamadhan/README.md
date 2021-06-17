𝑬𝑪𝑯𝑶 𝑭𝑹𝑨𝑴𝑬𝑾𝑶𝑹𝑲

Echo adalah framework bahasa golang untuk pengembangan aplikasi web. Framework ini cukup terkenal di komunitas. 
Echo merupakan framework besar, didalamnya terdapat banyak sekali dependensi.

terdapat banyak cara di :
 https://echo.labstack.com/

𝙄𝙣𝙨𝙩𝙖𝙡𝙡𝙖𝙩𝙞𝙤𝙣

 	mkdir myapp && cd myapp
 	go mod init myapp
	go get github.com/labstack/echo/v4

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
	
package net/http untuk berbagai macam fitur untuk keperluan aplikasi berbasis web
package github.com/labstack/echo/v4 untuk meng-import echo kedalam file project kita.
untuk penjelasan selengkapnya bisa dilihat disini
Sebuah objek router e dicetak lewat echo.New(). Lalu lewat objek router tersebut, dilakukan registrasi rute untuk / dengan method GET dan handler adalah closure handler. Lalu, mencetak output berupa String apabila tidak ada error akan menampilkan Hello, World. Terakhir, dari objek router di-start-lah sebuah web server pada port 2004.

Cobalah run file nya:

go run server.go

Lalu jalankan di web browser anda:

https://localhost:2004
