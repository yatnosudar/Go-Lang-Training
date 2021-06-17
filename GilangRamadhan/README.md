# Echo framework Golang
Echo adalah framework bahasa golang untuk pengembangan aplikasi web.

## Instalasi
Untuk menginstal Echo memerlukan versi Go v1.13 atau lebih tinggi. > Gunakanlah **go get** di directory project anda seperti dibawah ini
```bash
$ cd <PROJECT IN $GOPATH>
$ go get -u github.com/labstack/echo/...
```

## Hello World
buat file **main.go**
```go
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
	e.Logger.Fatal(e.Start(":1234"))
}
```
- package **net/http** untuk berbagai macam fitur untuk keperluan aplikasi berbasis web
- package **github.com/labstack/echo/v4** untuk meng-import echo kedalam file project kita.
- untuk penjelasan selengkapnya bisa dilihat [disini](https://pkg.go.dev/github.com/labstack/echo/v4)

Sebuah objek router e dicetak lewat echo.New(). Lalu lewat objek router tersebut, dilakukan registrasi rute untuk / dengan method GET dan handler adalah closure handler. Lalu, mencetak output berupa String apabila tidak ada error akan menampilkan Hello, World. Terakhir, dari objek router di-start-lah sebuah web server pada port 1234.

Cobalah run file nya:
> go run main.go

Lalu jalankan di web browser anda:
> https://localhost:1234

##Membuat API dengan Echo
Perhatikan source code di bawah ini 
**home.go**
'''go
package models

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Base struct {
	Data   Content
	Status string
}

type MapsName struct {
	ID    int
	Title string
	Image string
	Desc  string
	View  string
}

type Content struct {
	Content []MapsName
}

func GetMaps(c echo.Context) error {
	var pick []MapsName
	for i := 0; i < 5; i++ {
		pick = append(pick, MapsName{
			ID:    i,
			Title: fmt.Sprintf("Maps %d", i),
			Image: fmt.Sprintf("image %d", i),
			Desc:  fmt.Sprintf("desc %d", i),
			View:  fmt.Sprintf("View %d", i),
		})
	}
	return c.JSON(http.StatusOK, &Base{
		Data: Content{
			Content: pick,
		},
		Status: "success",
	})
}
'''
saya menyimpan file ini di package models

lalu membuat struct seperti di atas, lalu membuat sebuah function bernama **GetMaps** untuk mengumbalikan nilai di struct yang sudah di buat tersebut dan dikembalikan dengan format JSON, jika tidak ada error maka akan mengembalikan JSON Dengan objek status bervalue string, Data bervalue struct MapsName, dan seterusnya.

**routes.go**
```go
package routes

import (
	"valorant/models"

	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	e := echo.New()

	e.GET("/maps", models.GetMaps)
	
	return e
}
```
saya menyimpan file tersebut di package routes