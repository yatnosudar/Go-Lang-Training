package models

import (
	"fmt"
	"net/http"
	"strconv"
	"valorant/db"

	"github.com/labstack/echo/v4"
)

type Weapon struct {
	Status string `json:"status"`
	Data   Gaya
}

type Data struct {
	Title string `json:"title"`
}

type ArsenalDetail struct {
	Id              int    `json:"id"`
	ArsenalName     string `json:"arsenal_name"`
	CategoryArsenal string `json:"Category_arsenal"`
	Description     string `json:"description"`
	Image           string `json:"image"`
}

type Gaya struct {
	Gaya []ArsenalDetail
}

func GetArsenal(c echo.Context) error {
	var senjata []ArsenalDetail
	for i := 0; i < 10; i++ {
		senjata = append(senjata, ArsenalDetail{
			Id:              i,
			ArsenalName:     fmt.Sprintf("Arsenal name %d", i),
			CategoryArsenal: fmt.Sprintf("Maps %d", i),
			Description:     fmt.Sprintf("image %d", i),
			Image:           fmt.Sprintf("desc %d", i),
		})
	}
	return c.JSON(http.StatusOK, Weapon{
		Data: Gaya{
			Gaya: senjata,
		},
		Status: "success",
	})
}

func DetailArenal(c echo.Context) error {
	ArsenalDetailnid := c.QueryParam("id")
	id, _ := strconv.Atoi(ArsenalDetailnid)
	Arsenal := ArsenalDetail{
		Id:              id,
		ArsenalName:     fmt.Sprintf("arsenal name %d", 1),
		CategoryArsenal: fmt.Sprintf("CategoryArsenal %d", 1),
		Description:     fmt.Sprintf("description %d", 1),
		Image:           fmt.Sprintf("image %d", 1),
	}
	return c.JSON(http.StatusOK, Arsenal)

}

func FetchAllArsenal() (Response, error) {
	var obj ArsenalDetail
	var arrobj []ArsenalDetail
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM arsenal"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.ArsenalName, &obj.CategoryArsenal, &obj.Description, &obj.Image)
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
