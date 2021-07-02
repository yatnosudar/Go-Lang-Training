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

## Update data dengan method PUT
menggunakan method ini juga butuh aplikasi eksternal, seperti kemarin saya menggunakan postman. tambahkan func ini di **pegawai.model.go**
```go
func UpdatePegawai(id int, nama string, alamat string, telepon string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE pegawai SET nama = ?, alamat = ?, telepon = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telepon, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}
	return res, nil
}
```
jika sebelumnya di method POST kita tidak menggunakan parameter id, disini kita menggunakannya agar kita bisa menentukan rows mana yang akan kita Update nantinya. lalu menggunakan sqlStatement dengan value "UPDATE pegawai SET nama = ?, alamat = ?, telepon = ? WHERE id = ?" yaitu untuk mengubah data dari tabel pegawai dari kolom mana dan dengan id berapa dan meng set kolom dari database dengan masing masing menggunakan **"?"** sebagai place holder seperti diatas.

di prepare dengan variabel stmt. Lalu di execute dengan variabel result, err menggunakan stmt.Exec dan diisi dengan parameter yang berurutan seperti pada sqlStatement nya. Jika saat di method POST kita menggunakan **LastInsertId** ,disini kita menggunakan **RowsAffected** untuk mengetahui berapa kolom yang sudah kita ubah. Jika tidak ada error langsung set up response nya lalu di return.

tambahkan func ini di **pegawai.controlers.go**
```go
func UpdatePegawai(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	telepon := c.FormValue("telepon")

	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.UpdatePegawai(conv_id, nama, alamat, telepon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
```
Disini kita menampung 4 parameter yang akan dikirim melalui postman, dengan id untuk menentukan row mana yang akan kita update, karena id dari form value itu berupa string, maka kita akan mengkonversinya ke int dengan menggunakan **strconv.Atoi** , lalu menampung hasil dari func di package models tadi kedalam variabel result dan error, jika tidak ada error maka langsung dikembalikkan dengan format json.

lalu tambahkan ini di file routes kita:
> e.PUT("/pegawai", controllers.UpdatePegawai)

disini kita menggunakan method PUT dan dengan routes yang sama, jika sudah langsung saja ke aplikasi eksternal untuk mengubah data. di Postman kamu hanya perlu Add request kemudian dengan method PUT dan di bagian body bisa kita tambahkan sama seperti pada method POST sebelumnya, hanya saja saat ini kita menambah 1 parameter lagi yaitu id yang bervalue id yang mana akan kita ubah. value yang lain kita isi sesuai yang diinginkan, dan akan menampilkan response rows_affected untuk mengetahui jumlah rows yang diubah.

## Menghapus data dengan method DELETE
menggunakan method ini juga butuh aplikasi eksternal, seperti kemarin saya menggunakan postman. tambahkan func ini di **pegawai.model.go**
```go
func DeletePegawai(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM pegawai WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
```
Kali ini kita hanya menggunakan satu parameter yaitu id, yang berfungsi untuk menentukan menghapus data berdasarkan id. Lalu seperti biasa kita membuat variabel res dan inisiasi database. lalu sqlStatement untuk menghapus dari tabel pegawai dengan id ke placeholdernya, lalu di prepare, jika tidak ada error maka akan langsung meng eksekusi parameter id yang ditampung di variabel result. Sama seperti sebelumnya, disni kita menggunakan RowsAffected juga.

tambahkan func ini di **pegawai.controlers.go**
```go
func DeletePegawai(c echo.Context) error {
	id := c.FormValue("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.DeletePegawai(conv_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}
```
Disini kita hanya menampung 1 parameter yaitu id  untuk menjadi key saat kita menggunakannya di postman nanti. Lalu di konversi menggunakan strconv.Atoi, jika tidak ada error maka langsung saja memanggil fungsi dari package models dan memasukkan parameter id yang sudah di konversi tadi, jika tidak ada error langsung di return saja dengan format json.

lalu tambahkan ini di file routes kita:
> e.DELETE("/pegawai", controllers.DeletePegawai)

disini kita menggunakan method DELETE dan routes yang sama, jika sudah di run langsung saja ke postman dan kita bisa melihat data dengan method get dan menentukan data mana yang akan dihapus, jika sudah langsung ganti methodnya dengan DELETE lalu di bagian body kita masukan key yaitu id dengan value id dari data yang kita ingin hapus. dan jika sukses maka di dalam database nya pun ikut terhapus.

## Login
Kali ini kita akan mencoba membuat API untuk login. Hal pertama yang akan kita lakukan adalah membuat sebuah tabel yaitu tabel users dan diberi kolom id, username dan password. Kali ini kita menggunakan sebuah hash method bernama bycrypt. Penjelasan selengkapnya [Lihat disini](https://pkg.go.dev/golang.org/x/crypto/bcrypt). 

Langsung saja kita install di direktori project kita:
> go get golang.org/x/crypto/bcrypt

**password.helper.go**
```go
package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
```
Kita membuat sebuah file yang memiliki 2 fungsi, yang pertama HashPassword memiliki parameter password dan return sebuah hash yang dikembalikan dalam bentuk string. Membuat variabel bytes untuk menampung hash, dengan menggunakan package bycrypt untuk men Generate password dalam tipe []byte dan parameter kedua yaitu cost-nya. lalu di return.

func yang kedua yaitu CheckPasswordHash memiliki parameter password dan hash, dengan return value boolean. membuat variabel err dan menggunakan package bycrypt untuk meng CompareHashAndPassword, parameter pertama yaitu hash dalam tipe byte dan yang kedua yaitu password jika ada error langsung return false, err jika tidak ada maka akan return true.

**login.controller.go**
```go
func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helper.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}
```
membuat sebuah function di package controller untuk men-GenerateHashPassword, lallu menerima input dari user ke parameter password, membuat variabel hash lalu mengimport dari func HashPassword untuk mengubah input dari user tadi menjadi hash password, di return dalam bentuk json.

tambahkan ini di file routes:
> e.GET("/generate-hash/:password", controllers.GenerateHashPassword)

disini kita menggunakan method GET dan routes generate-hash dengan parameter password. Lalu di postman kita langsung saja ke localhost kita dengan routes "/generate-hash/" dan setelah "/" kita bisa menambahkan password apapun itu dan saat di send kita akan melihat hasil hash dari bycrypt. Lalu hasil hash tersebut kita copy dan masukkan ke dalam tabel users untuk password.

**login.model.go**
```go
package models

import (
	"database/sql"
	"echo/db"
	"echo/helper"
	"fmt"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func CheckLogin(username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE username = ?"

	err := con.QueryRow(sqlStatement, username).Scan(
		&obj.Id, &obj.Username, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Username not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query error")
		return false, err
	}

	match, err := helper.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and Password doesn't match!")
		return false, err
	}

	return true, nil
}
```
buat di package models, membuat struct User dengan isi Id dan Username, kemudian membuat func CheckLogin dan kita menerima 2 input yaitu username dan password dengan return boolean untuk menentukan apakah terotentikasi atau tidaknya. membuat var obj dari struct user, dan pwd untuk menampung hash. langsung kita inisiasi database, dan sqlStatement untuk mencari data dari tabel users berdasarkan username dan memasukkan place holder"?". langsung eksekusi sqlStatement tadi dengan variabel err menggunakan con.QueryRow, jika ketemu maka kita masukkan ke dalam obj nya. 

Jika error nya adalah sql.ErrNoRows maka akan menampilkan pesan Username no found, dan jika error karena hal lain maka menampilkan pesan Query error. Jika tidak ada error maka kita tampung dalam variabel match dan bisa langsung di check menggunakan fungsi CheckPasswordHash tadi dengan parameter pertama adalah plain password dan kedua adalah untuk hash nya, jika tidak match maka akan mengembalikkan pesan seperti di atas.

**login.controller.go**
```go
func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	return c.String(http.StatusOK, "Berhasil login!")
}
```
kita membuat sebuah func lagi di package controller untuk meng Check apakah username dan passwordnya sudah terdaftar di database. kita akan menerima inputan dari user, yaitu username dan password. ditampung di variabel res dan mengambil dari fungsi di package models dan mem passing username dan password. Jika ada error maka akan menampilkan pesan error nya, jika tidak ada error maka langsung return dalam bentuk string dan menampilkan pesan Berhasil login.

tambahkan ini di file routes:
> e.POST("/login", controllers.CheckLogin)

Menggunakan method POST karena karena kita akan menerima inputan dari user yang berupa username dan password, langsung saja coba di postman masuk ke bagian body, lalu ke form data dan memasukkan key username dan password dan memiliki value sesuai pada database, dan dibagian password kita bisa memasukkan plain password kita.

## Implementasi JWT
JWT adalah singkatan dari JSON web token, untuk penjelasan lengkapnya bisa [lihat disini](https://github.com/golang-jwt/jwt). Yang merupakan sebuah package dari golang untuk membuat sebuah token yang memiliki sebuah payload atau data, sehingga token ini bisa dikirimkan di masing-masing request. Langsung saja install di direktori:
> go get github.com/golang-jwt/jwt

kita ganti tambahkan di func **CheckLogin** di **login.controler.go**
```go
func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}
	// Create token

	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})

}
```
Kita mengganti return dengan pesan "Berhasil login!" kemarin dengan token. Ditampung di variabel token untuk membuat sebuah token dengan jwt. Kemudian membuat variabel claims untuk meng set MapClaims, lalu men-set up payload atau informasi yang akan diberikan, disnin saya memberikan info username, level, dan expire nya di set selama 72 jam.

lalu kita kasih SignedString yang diberi nama "secret", jika ada error maka tampilkan pesannya, jika tidak ada maka kita return tokennya. Lalu test di postman sama ketika login seperti yang kemarin. Jika berhasil maka akan menampilkan token dari jwt.

## Middleware
Sebelumnya kita sudah membuat token untuk login, kali ini kita akan menentukan routes mana yang akan kita lindungi oleh jwt tersebut yang akan kita berikan autentikasi saat ingin mengaksesnya. Kita bisa membaca dokumentasi lengkapnya [disini](https://echo.labstack.com/middleware/). 
**middleware.go**
```go
package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})
```
membuat file baru di package midddleware, mengimport package dari echo middleware. Kemudian membuat variabel IsAuthenticated dengan menggunakan JWTWithConfig dari middleware dan JWTConfig yang diberi SigningKey yang sama seperti di func CheckLogin di package controller.

lalu kita tambahkan di routes yang ingin kita berikan autentikasi:
> e.GET("/pegawai", controllers.FetchAllPegawai, middleware.IsAuthenticated)

diatas saya memberi autentikasi dari routes pegawai dengan method GET yang bertujuan saat kita ingin melihat data pegawai, maka harus melakukan autentikasi terlebih dahulu.

Jika kita test menggunakan postman maka akan menampilkan pesan "missing or malformed jwt", dimana kita harus memberikan sebuah token pada saat kita login tadi. masuk ke bagian auhorization lalu ganti tipe nya menjadi Bearer token, lalu copy token tadi, jika sudah di send maka akan menampilkan data pegawai.

Middleware tadi bisa kita pasang juga untuk routes yang lain.

## Validasi Input User
Pada dasarnya inputan dari user itu tidak selalu sesuai dengan yang kita harapkan, maka dari itu golang menyediakan sebuah package untuk mem-validasi nya yang bernama Validator. Untuk dokumentasi lengkapnya bisa [lihat disini](https://github.com/go-playground/validator). 

**Instalasi**
```bash
go get github.com/go-playground/validator
```
**validation.controller.go**
```go
package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Customer struct {
	Nama   string `validate:"required"`
	Email  string `validate:"required,email"`
	Alamat string `validate:"required"`
	Umur   int    `validate:"gte=13,lte=40"`
}

func TestStructValidation(c echo.Context) error {
	v := validator.New()

	customer := Customer{
		Nama:   "Dani",
		Email:  "danikurnia@gmail.com",
		Alamat: "Jl.Pasirluyu",
		Umur:   17,
	}

	err := v.Struct(customer)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Succes"})
}

func TestVariableValidation(c echo.Context) error {
	v := validator.New()

	email := "dani@gmail.com"

	err := v.Var(email, "required,email")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Email not valid",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Succes"})
}
```
Kita akan mencoba untuk mem-validasi input Customer, Membuat struct Customer lalu diberi struct text validate untuk menentukan apa saja yang diperlukan agar bisa tervalidasi, required artinya text tersebut harus diisi, email untuk menyesuaikan dengan format email, gte artinya di bagian Umur yang dimasukkan harus lebih dari 13 dan kurang dari 40. Lalu di func **TestStructValidation** adalah untuk mem validasi dari sebuah struct. membuat sebuah objek v lalu membuat variabel customer untuk mengisi data dari struct Customer. ditampung di variabel err dengan value v.Struct(customer) yaitu validator dalam bentuk struct dengan struct customer. Jika ada error atau isi dari struct tersebut tidak sesuai dengan rules yang kita tetapkan maka akan menampilkan pesan error, jika tidak maka akan tampil pesan succes.

tambah di file routes:
> e.GET("/test-struct-validation", controllers.TestStructValidation)

Lalu bisa di test di postman

di func yang kedua kita mencoba mem validasi dari variabel. membuat objek v terlebih dahulu lalu membuat variabel email lalu ditampung di variabel err dengan value v.Var(email, "required,email") yaitu validator dalam bentuk variabel dengan variabel email dan  rules nya required,email. Jika tidak variabel email tidak sesuai rule, maka akan tampil pesan Email not valid jika tidak maka menampilkan pesan succes.

tambah di file routes:
> e.GET("/test-variable-validation", controllers.TestVariableValidation)

## Implementasi validasi ke REST API
Kita akan meng implementasikannya ke RESTful API yang telah kita buat sebelumnya. 
```go
type Pegawai struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama" validate:"required"`
	Alamat  string `json:"alamat" validate:"required"`
	Telepon string `json:"telepon" validate:"required"`
}
```
yang pertama kita lakukan adalah menambah struct text validate ke dalam struct Pegawai, disini saya hanya memasukkan required. 
```go
func AddPegawai(nama string, alamat string, telepon string) (Response, error) {
	var res Response

	v := validator.New()

	pegawai := Pegawai{
		Nama:    nama,
		Alamat:  alamat,
		Telepon: telepon,
	}
	err := v.Struct(pegawai)
	if err != nil {
		return res, err
	}
```
Lalu func AddPegawai ini akan menerima variabel dari controller, disini kita memasukkan nilai struct Pegawai dari parameter diatas. dan jika sudah maka tinggal dicoba di postman dan coba untuk tidak mengisi salah satu field, maka yang terjadi adalah error. jika semua field terisi maka akan succes.
