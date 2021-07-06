package models

import (
	"database/sql"
	"echo/RucoyAPI/db"
	"fmt"
	"net/http"
)

type Characters struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Level      int    `json:"level"`
	Guild      string `json:"guild"`
	LastOnline string `json:"last_online"`
	Born       string `json:"born"`
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
