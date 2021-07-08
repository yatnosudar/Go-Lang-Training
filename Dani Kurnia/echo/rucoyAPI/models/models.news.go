package models

import (
	"echo/RucoyAPI/db"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Newss struct {
	Id          int    `json:"id"`
	Title       string `json:"title" validate:"required"`
	Reddit      string `json:"reddit" validate:"required"`
	Facebook    string `json:"facebook" validate:"required"`
	Date        string `json:"date" validate:"required"`
	Image       string `json:"image"`
	Description string `json:"description" validate:"required"`
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

func AddNews(title, reddit, facebook, date, image, description string) (Response, error) {
	var res Response
	v := validator.New()

	var news = Newss{
		Title:       title,
		Reddit:      reddit,
		Facebook:    facebook,
		Date:        date,
		Image:       image,
		Description: description,
	}

	err := v.Struct(news)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()
	sqlStatement := "INSERT INTO news (title, reddit, facebook, date, image, description) VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(title, reddit, facebook, date, image, description)
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

func UpdateNews(id int, title, reddit, facebook, date, image, description string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE news SET title = ?, reddit = ?, facebook = ?, date = ?, image = ?, description = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(title, reddit, facebook, date, image, description, id)
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
