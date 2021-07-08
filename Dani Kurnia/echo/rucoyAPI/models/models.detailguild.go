package models

import (
	"database/sql"
	"echo/RucoyAPI/db"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type DetailGuild struct {
	Id          int    `json:"id"`
	Name        string `json:"name" validator:"required"`
	Description string `json:"description" validator:"required"`
	Date        string `json:"date" validator:"required,datetime"`
	Members     string `json:"members"`
}

func GetDetailGuild(name string) (Response, error) {
	var obj DetailGuild
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM guild WHERE name = ?"

	err := con.QueryRow(sqlStatement, name).Scan(
		&obj.Id, &obj.Name, &obj.Description, &obj.Date, &obj.Members,
	)
	if err == sql.ErrNoRows {
		fmt.Println("Guild Not Found")
		return res, err
	}

	if err != nil {
		return res, err
	}

	response := DetailGuild{
		Id:          obj.Id,
		Name:        obj.Name,
		Description: obj.Description,
		Date:        obj.Date,
		Members:     obj.Members,
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = response

	return res, nil
}

func AddGuild(name, description, date, members string) (Response, error) {
	var res Response
	v := validator.New()

	var guild = DetailGuild{
		Name:        name,
		Description: description,
		Date:        date,
		Members:     members,
	}

	err := v.Struct(guild)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()
	sqlStatement := "INSERT INTO guild (name, description, date, members) VALUES (?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, description, date, members)
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

func UpdateGuild(id int, name, description string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE guild SET name = ?, description = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, description, id)
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
