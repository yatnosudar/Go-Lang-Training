# Echo framework Golang
Echo adalah framework bahasa golang untuk pengembangan aplikasi web.

## Instalasi
Untuk menginstal Echo memerlukan versi Go v1.13 atau lebih tinggi. > Gunakanlah **go get** di directory project anda seperti dibawah ini
```bash
$ cd <PROJECT IN $GOPATH>
$ go get -u github.com/labstack/echo/...
```

## Hello World
buat file **server.go**
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
	e.Logger.Fatal(e.Start(":2004"))
}
```
- package **net/http** untuk berbagai macam fitur untuk keperluan aplikasi berbasis web
- package **github.com/labstack/echo/v4** untuk meng-import echo kedalam file project kita.
- untuk penjelasan selengkapnya bisa dilihat [disini](https://pkg.go.dev/github.com/labstack/echo/v4)


Sebuah objek router e dicetak lewat echo.New(). Lalu lewat objek router tersebut, dilakukan registrasi rute untuk / dengan method GET dan handler adalah closure handler. Lalu, mencetak output berupa String apabila tidak ada error akan menampilkan Hello, World. Terakhir, dari objek router di-start-lah sebuah web server pada port 2004.

Cobalah run file nya:
> go run server.go

Lalu jalankan di web browser anda:
> https://localhost:2004

## Membuat API dengan Echo
Perhatikan source code di bawah ini
**Welcome.go**
```go
package models

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseWelcome struct {
	Status string
	Data   WelcomeContent
}

type WelcomeContent struct {
	Header   string
	Download string
	Online   int
	Server   int
	Page     []string
}

func GetWelcome(c echo.Context) error {
	return c.JSON(http.StatusOK, &BaseWelcome{
		Status: "succes",
		Data: WelcomeContent{
			Header:   "Welcome to Rucoy Online",
			Download: "https://playstore.com/RucoyOnline",
			Online:   3000,
			Server:   19,
			Page: []string{
				"https://gambar1",
				"https://gambar2",
				"https://gambar3",
				"https://gambar4",
				"https://gambar5",
				"https://gambar6",
				"https://gambar7",
				"https://gambar8",
			},
		},
	})
}
```
saya menyimpan file ini di package models

lalu membuat struct seperti di atas, lalu membuat sebuah function bernama **GetWelcome** untuk mengumbalikan nilai di struct yang sudah di buat tersebut dan dikembalikan dengan format JSON, jika tidak ada error maka akan mengembalikan JSON Dengan objek status bervalue string, Data bervalue struct WelcomeContent, dan seterusnya.

**routes.go**
```go
package routes

import (
	"echo/Rucoy/models"

	"github.com/labstack/echo/v4"
)

func Routes() *echo.Echo {
	e := echo.New()

	e.GET("/welcome", models.GetWelcome)
	
	return e
}
```
saya menyimpan file tersebut di package routes

membuat function **Routes** membuat object e yang akan dicetak dengan **echo.New()** lalu object e di registrasi dengan rute **/welcome** dengan method GET dan mencetak dari function **models.GetWelcome** yang berada di package models, lalu mengembalikan e

**server.go**
```go
package main

import (
	"echo/Rucoy/routes"
)

func main() {
	e := routes.Routes()

	e.Logger.Fatal(e.Start(":2004"))
}
```
membuat object e dengan menggunakan function **Routes** di package routes dan men-Start web server pada port 2004.
Run file server.go:
> go run server.go
Lalu jalankan di web browser anda:
> https://localhost:2004/welcome

## Query parameter
Perhatikan source code di bawah
**News.go**
```go
package models

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)
type BaseNews struct {
	Status string
	Data   Content
}

type News struct {
	Id          int
	Title       string
	Reddit      string
	Facebook    string
	Date        string
	Image       string
	Description string
}

type Content struct {
	Content []News
}
func DetailNews(c echo.Context) error {
	NewsId := c.QueryParam("id")
	id, _ := strconv.Atoi(NewsId)
	news := News{
		Id:          id,
		Title:       fmt.Sprintf("TItle%d", 1),
		Reddit:      "https://reddit",
		Facebook:    "https://fb",
		Date:        "dd-mm-yyyy",
		Image:       "Url",
		Description: fmt.Sprintf("Desc%d", 1),
	}
	return c.JSON(http.StatusOK, news)

}
```
membuat function **DetailNews**, saya membuat variabel dengan nama NewsId yang berisi query parameter dengan nilai id, lalu menngunakan **strconv.Atoi** untuk mengubah NewsId yang bernilai int agar menjadi string

lalu membuat variabel news yang berisi struct News dengan nilai Id: id yang sudah di deklarasikan diatas, untuk yang lainnya saya hanya mengisinya dengan dummy, lalu dikembalikan dengan format JSON, dan jika tidak ada error makan mengembalikkan **news**

Selanjutnya kita tambahkan seperti dibawah ke file routes.go:
> e.GET("/news", models.GetNews)

Lalu jalankan di web browser dengan:
> https://localhost:2004/news?id=30

**?** berarti parameternya dan **=30** adalah value dari parameter tersebut, disini kita bisa mengisi nilai berapapun

## Query parameter search
perhatikan source code di bawah ini
**SearchGuild.go**
```go
package models

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MainSearchGuild struct {
	Status string `json:"status"`
	Data   BodySearchGuild
}
type BodySearchGuild struct {
	Title            string `json:"title"`
	Search           string `json:"search"`
	ValueSearchGuild []ValueSearchGuild
}
type ValueSearchGuild struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
func SearchDetailGuild(c echo.Context) error {
	var value []ValueSearchGuild
	SearchGuild := c.QueryParam("search")
	id, _ := strconv.Atoi(SearchGuild)
	for _, v := range value {
		value = append(value, ValueSearchGuild{
			ID:   id,
			Name: fmt.Sprintln(v, SearchGuild),
		})
	}
	return c.JSON(http.StatusOK, &ValueSearchGuild{
		ID:   id,
		Name: fmt.Sprint(SearchGuild),
	})

}
```
membuat struct seperti di atas, lalu membuat function **SearchDetailGuild** dan membuat variabel value yang berisi array dari struct **ValueSearchGuild**. Kemudian membuat var SearchGuild dengan value query parameter "search" dan mengkonversi var id dengan value dari SearchGuild.

Lalu melakukan perulangan dengan range dari value, kemudian di append dengan **ValueSearchGuild** dengan ID yang berisikan id yang sudah dikonversi, dan Name fmt.Sprintln(v, SearchGuild)

Dan dikembaikan dengan format JSON, jika tidak terjadi error maka akan mencetak seperti di atas

Kita tambahkan baris ini di file routes.go
> e.POST("/guilds/", models.SearchDetailGuild)

kali ini menggunakan method POST karena untuk searching biasanya menggunakan method ini. Dan untuk menjalankan di web browser sama seperti di atas tadi.