package models

import (
	"echo/RucoyAPI/db"
	"net/http"
)

type Newss struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Reddit      string `json:"reddit"`
	Facebook    string `json:"facebook"`
	Date        string `json:"date"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

func GetNewsDetail() (Response, error) {
	var obj Newss
	var arrobj []Newss
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM news"
	rows, _ := con.Query(sqlStatement)

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&obj.Id, &obj.Title, &obj.Reddit, &obj.Facebook, &obj.Date, &obj.Image, &obj.Description)
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = map[string]interface{}{
		"Content": arrobj,
	}

	return res, nil
}
