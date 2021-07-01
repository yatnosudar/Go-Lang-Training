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
**maps.go**
```go
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
```
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

membuat function **Routes** membuat object e yang akan dicetak dengan **echo.New()** lalu object e di registrasi dengan rute **/maps** dengan method GET dan mencetak dari function **models.Getmaps** yang berada di package models, lalu mengembalikan e

**main.go**
```go
package main

import (
	"valorant/routes"
)

func main() {
	e := routes.Routes()

	e.Logger.Fatal(e.Start(":1234"))
}
```
membuat object e dengan menggunakan function **Routes** di package routes dan men-Start web server pada port 1234.
Run file main.go:
> go run main.go
Lalu jalankan di web browser anda:
> https://localhost:1234/maps

## Query parameter
Perhatikan source code di bawah
**maps.go**
```go
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

func DetailMaps(c echo.Context) error {
	Mapsid := c.QueryParam("id")
	id, _ := strconv.Atoi(Mapsid)
	Maps := MapsName{
		ID:    id,
		Title: fmt.Sprintf("Maps %d", 1),
		Image: fmt.Sprintf("Maps %d", 1),
		Desc:  fmt.Sprintf("Maps %d", 1),
		View:  fmt.Sprintf("Maps %d", 1),
	}
	return c.JSON(http.StatusOK, Maps)

}
```

membuat function **DetailMaps**, saya membuat variabel dengan nama MapsId yang berisi query parameter dengan nilai id, lalu menngunakan **strconv.Atoi** untuk mengubah MapsId yang bernilai int agar menjadi string

lalu membuat variabel news yang berisi struct MapsName dengan nilai Id: id yang sudah di deklarasikan diatas, untuk yang lainnya saya hanya mengisinya dengan dummy, lalu dikembalikan dengan format JSON, dan jika tidak ada error makan mengembalikkan **Maps**

Selanjutnya kita tambahkan seperti dibawah ke file routes.go:
> e.GET("/maps", models.DetailMaps)

Lalu jalankan di web browser dengan:
> https://localhost:1234/maps?id=30

**?** berarti parameternya dan **=30** adalah value dari parameter tersebut, disini kita bisa mengisi nilai berapapun

# Koneksi Database mysql

kali ini kita akan mengisi data dengan yang ada pada database.

pertama-tama kita harus membuat sebuah database terlebih dahulu, lalu membuat tabel. disini saya akan membuat database pegawai dengan kolom id, nama, alamat, dan telepon. dengan mengisi **id** yang dijadikan primary key dan auto increment pastinya, untuk yang lainnya bisa mengisi tipenya dengan format masing-masing, seperti nama dengan varchar, dan seterusnya dan diberi lenght untuk masing-masing kolom.

kita memerlukan sebuah package dari go untuk me manage konfigurasi file dan environment namanya yaitu **gonfig**, untuk penjelasan selengkapnya bisa dilihat [KLIK](https:github.com/tkanos/gonfig). Langsung bisa di install di direktori project anda dengan

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

## Menambah data dengan method POST
dengan method ini dibutuhkan aplikasi bantuan menggunakan postman untuk bisa menambahkan data kali ini. tambahkan func di **pegawai.model.go**
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

menggunakan method POST dan dengan routes yang sama dengan  menampilkan data pegawai, jika sudah langsung saja ke aplikasi eksternal untuk menambah data. di Postman kamu hanya perlu Add request kemudian dengan method POST dan di bagian body bisa tambahkan dengan parameter dengan key yang sudah kita buat pada file diatas dengan value kita isi sesuai yang diinginkan

## Update data dengan method PUT

menggunakan method ini juga butuh aplikasi eksternal, seperti kemarin menggunakan postman. tambahkan func ini di **pegawai.model.go**
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
jika sebelumnya di method POST tidak menggunakan parameter id, disini  menggunakannya agar  bisa menentukan rows mana yang akan Update nantinya. lalu menggunakan sqlStatement dengan value "UPDATE pegawai SET nama = ?, alamat = ?, telepon = ? WHERE id = ?" yaitu untuk mengubah data dari tabel pegawai dari kolom mana dan dengan id berapa dan meng set kolom dari database dengan masing masing menggunakan **"?"** sebagai place holder seperti diatas.

di prepare dengan variabel stmt. Lalu di execute dengan variabel result, err menggunakan stmt.Exec dan diisi dengan parameter yang berurutan seperti pada sqlStatement nya. Jika saat di method POST menggunakan **LastInsertId** ,disini menggunakan **RowsAffected** untuk mengetahui berapa kolom yang sudah diubah. Jika tidak ada error langsung set up response nya lalu di return.

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
 menampung 4 parameter yang akan dikirim melalui postman, dengan id untuk menentukan row mana yang akan  update, karena id dari form value itu berupa string, maka  akan mengkonversinya ke int dengan menggunakan **strconv.Atoi** , lalu menampung hasil dari func di package models tadi kedalam variabel result dan error, jika tidak ada error maka langsung dikembalikkan dengan format json.

lalu tambahkan ini di file routes kita:
> e.PUT("/pegawai", controllers.UpdatePegawai)

 menggunakan method PUT dan dengan routes yang sama, jika sudah langsung saja ke aplikasi eksternal untuk mengubah data. di Postman kamu hanya perlu Add request kemudian dengan method PUT dan di bagian body bisa kita tambahkan sama seperti pada method POST sebelumnya, hanya saja saat ini kita menambah 1 parameter lagi yaitu id yang bervalue id yang mana akan kita ubah. value yang lain  isi sesuai yang diinginkan, dan akan menampilkan response rows_affected untuk mengetahui jumlah rows yang diubah.
