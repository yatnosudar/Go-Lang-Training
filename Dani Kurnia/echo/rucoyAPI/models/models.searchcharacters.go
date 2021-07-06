package models

import (
	"database/sql"
	"echo/RucoyAPI/db"
	"fmt"
	"net/http"
)

type SearchCharacter struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Level      int    `json:"level"`
	Guild      string `json:"guild"`
	LastOnline string `json:"last_online"`
	Born       string `json:"born"`
}

func SearchCharacters(value string) (Response, error) {
	var obj SearchCharacter
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM characters WHERE name = ?"

	err := con.QueryRow(sqlStatement, value).Scan(
		&obj.Id, &obj.Name, &obj.Level, &obj.Guild, &obj.LastOnline, &obj.Born,
	)
	if err == sql.ErrNoRows {
		fmt.Println("Characters not found")
		return res, err
	}
	if err != nil {
		return res, err
	}

	response := SearchCharacter{
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
	}

	return res, nil
}
