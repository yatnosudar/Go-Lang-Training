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


## Koneksi Database mysql
Jika sebelumnya kita hanya mengisi data dengan format dummy, kali ini kita akan mengisi data dengan yang ada pada database.

pertama-tama kita harus membuat sebuah database terlebih dahulu, lalu membuat tabel. disini saya akan membuat database pegawai dengan kolom id, nama, alamat, dan telepon. dengan mengisi **id** yang dijadikan primary key dan auto increment pastinya, untuk yang lainnya bisa mengisi tipenya dengan format masing-masing, seperti nama dengan varchar, dan seterusnya dan diberi lenght untuk masing-masing kolom.

kita memerlukan sebuah package dari go untuk me manage konfigurasi file dan environment namanya yaitu **gonfig**, untuk penjelasan selengkapnya bisa dilihat [disini](https:github.com/tkanos/gonfig). Langsung bisa di install di direktori project anda dengan
```bash
$ go get github.com/tkanos/gonfig
```
lalu kita akan membuat package config untuk lebih mempermudah

**config.json**
```json
{
    "DB_USERNAME" : "root",
    "DB_PASSWORD" : "",
    "DB_PORT" : "3306",
    "DB_HOST" : "127.0.0.1",
    "DB_NAME" : "echo_rest"
}
```
file diatas adalah format default dari localhost kita, dan untuk **DB_NAME** bisa diisi dengan nama database masing-masing

**config.go**
```go
package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Configuration {
	conf := Configuration{}
	gonfig.GetConf("config/config.json", &conf)
	return conf
}
```
diatas kita membuat struct seperti pada file config.json, gunanya agar bisa menampung dan memiliki value dari config.json. Lalu membuat func **GetConfig** yang memiliki return struct **Configuration**, membuat variabel conf dan menggunakan package **gonfing.GetConf** untuk mendapatkan value dari file config.json.

lalu kita perlu menginstall mysql driver dari Golang, dokumentasi lengkap [disini](https://github.com/go-sql-driver/mysql). Langsung saja install di direktori
```bash
$ go get -u github.com/go-sql-driver/mysql
```
lalu buat file **db.go** di package db
```go
package db

import (
	"database/sql"
	"echo/config"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME
	db, err = sql.Open("mysql", connectionString)

	if err != nil {
		panic("connectionString error...")
	}

	err := db.Ping()

	if err != nil {
		panic("DSN invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
```
mambuat variabel db dengan pointer sql.DB dari package dari package database sql golang, dan variabel err yang bervalue error. Lalu membuat func **Init** dan membuat var conf untuk mengambil konfigurasi yang di return dari file **config.go** sebelumnya, membuat **connectionString** atau DSN dengan format diatas. Dan menggunakannya dengan membuat variabel berisi **sql.Open("mysql", connectionString)** dari pkg sql-driver dengan driver mysql dan DSN dari connectionString diatas.
Jika ada error langsung dihentikan dengan panic dan di cek dengan **db.Ping()** jika terdapat error pada DSN. Lalu membuat func **CreateCon** untuk mereturn instans dari db yang sudah di set di atas

lalu memanggil function tersebut di file utama kita, yaitu server.go tambahkan:
> db.Init()

jika tidak terjadi error maka sudah sukses connect ke database

## Menampilkan data dengan method GET
pertama kita akan membuat file **response.go** di package models karena kita akan mengembalikan data dengan format yang sama yaitu json
```go
package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
```
kita membuat struct Response dengan status bertipe int karena akan diisi dengan http.Status..., message dengan string, lalu Data dengan interface{}
**pegawai.model.go**
```go
package models

import (
	"echo/db"
	"net/http"
)

type Pegawai struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}

func FetchAllPegawai() (Response, error) {
	var obj Pegawai
	var arrobj []Pegawai
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM pegawai"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}
```
pertama membuat struct dengan sama dengan masing-masing kolom pada database, lalu membuat func **FetchAllPegawai** yang me return Response yang kita buat sebelumnya dan error. Membuat var obj yang berisi dari struct Pegawai dan membuat var arrobj dari array struct pegawai karena bisa saja datanya lebih dari 1. membuat connection dari func **CreateCon** di package db sebelumnya, lalu membuat sqlStatement untuk menampilkan data dari tabel pegawai.

Membuat variabel rows karena kita mempunyai data pegawai lebih dari 1 dengan value **con.Query(sqlStatement)**. Lalu membuat loop dengan for **rows.Next()**, ketika masih ada data kita scan atau mengisi kolom id, nama, alamat, telepon dengan &obj dan harus sesuai dengan urutan pada database. Jika tidak ada error maka arrobj langsung di append dengan obj. Saat sudah, maka akan langsung me return response nya seperti di atas.

**pegawai.controllers.go**
```go
package controllers

import (
	"echo/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchAllPegawai(c echo.Context) error {
	result, err := models.FetchAllPegawai()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
```
membuat func **FetchAllPegawai** dengan c echo.Context dan return error. Kemudian membuat variabel result dan error yang bervalue dari func **FetchAllPegawai** di package models, jika ada error maka mengembalikan http.StatusInternalServerError dan message error tersebut, jika tidak ada maka langsung mengembalikkan result diatas. Langsung saja masukkan ke routes kita: 
> e.GET("/pegawai", controllers.FetchAllPegawai)

dengan routes /pegawai dan memanggil function dari controllers tadi.

## Menambah data dengan method POST
dengan method ini dibutuhkan aplikasi bantuan, saya menggunakan postman untuk bisa menambahkan data kali ini. tambahkan func di **pegawai.model.go**
```go
func AddPegawai(nama string, alamat string, telepon string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT pegawai (nama, alamat, telepon) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telepon)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}
	return res, nil
}
```
Func **AddPegawai** menerima parameter input sesuai dengan kolom yang ada pada database kita, tetapi tidak dengan Id karena Id akan di set auto increment. Untuk menambah data kita menggunakan sql statement insert ke tabel pegawai dengan kolom seperti di atas dan bervalue dengan parameter sesuai jumlah kolom. Lalu membuat variabel stmt dan err untuk mempersiapkan sqlStatement dengan **con.Prepare** agar siap di eksekusi. Kemudian membuat variabel result untuk mengeksekusi dengan stmt.Exec yang berisi kolom parameter di atas dengan urutan yang sama pada database.

Kemudian membuat lastInsertedID yang bertujuan untuk kita mendapatkan Id terbaru pada data yang telah kita buat. jika tidak ada error maka akan me return response seperti diatas dengan Data yang berisi map seperti di atas.

Tambahkan function ini di file **pegawai.controllers.go**
```go
func AddPegawai(c echo.Context) error {
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	result, err := models.AddPegawai(nama, alamat, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
```
Di function ini kita menampung 3 parameter yang akan dikirim oleh aplikasi eksternal, yaitu ada nama, alamat, dan telepon dengan menggunakan **c.FormValue()**. Kemudian membuat variabel result dan err yang akan menampung apa yang dikembalikkan oleh func **AddPegawai** di package models, jika ada error maka akan mengembalikan pesan error, jika tidak maka akan mengembalikan result tadi.

lalu tambahkan ini di file routes kita:
> e.POST("/pegawai", controllers.AddPegawai)

disini kita menggunakan method POST dan dengan routes yang sama dengan kita menampilkan data pegawai, jika sudah langsung saja ke aplikasi eksternal untuk menambah data. di Postman kamu hanya perlu Add request kemudian dengan method POST dan di bagian body bisa kita tambahkan dengan parameter dengan key yang sudah kita buat pada file diatas dengan value kita isi sesuai yang diinginkan.
