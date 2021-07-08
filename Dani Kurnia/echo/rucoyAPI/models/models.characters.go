package models

import (
	"database/sql"
	"echo/RucoyAPI/db"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Characters struct {
	Id         int    `json:"id"`
	Name       string `json:"name" validation:"required"`
	Level      int    `json:"level"`
	Guild      string `json:"guild"`
	LastOnline string `json:"last_online"`
	Born       string `json:"born" validation:"required,datetime"`
}

type Log struct {
	Killer   []Characters
	KillTime string `json:"killtime"`
	Killed   Characters
}

func GetCharacters(nickname string) (Response, error) {
	var log []Log
	var obj Characters
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM characters WHERE name = ?"

	err := con.QueryRow(sqlStatement, nickname).Scan(
		&obj.Id, &obj.Name, &obj.Level, &obj.Guild, &obj.LastOnline, &obj.Born,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Nickname not found")
		return res, err
	}

	if err != nil {
		return res, err
	}
	for i := 0; i < 10; i++ {
		log = append(log, Log{
			Killer:   []Characters{},
			KillTime: fmt.Sprintf("About %d hours ago", i),
			Killed:   Characters{},
		})
	}

	response := Characters{
		Id:         obj.Id,
		Name:       obj.Name,
		Level:      obj.Level,
		Guild:      obj.Guild,
		LastOnline: obj.LastOnline,
		Born:       obj.Born,
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = map[string]interface{}{
		"content": response,
		"log":     log,
	}

	return res, nil
}

func AddCharacters(name, born string) (Response, error) {
	var res Response
	v := validator.New()

	var characters = Characters{
		Name: name,
		Born: born,
	}

	err := v.Struct(characters)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()
	sqlStatement := "INSERT INTO characters (name, born) VALUES (?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, born)
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

func UpdateCharacters(id int, name string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE characters SET name = ?, description = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, id)
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
