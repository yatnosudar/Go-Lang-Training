package models

import (
	"database/sql"
	"echo/RucoyAPI/db"
	"fmt"
	"net/http"
)

type Char struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Level      int    `json:"level"`
	Guild      string `json:"guild"`
	LastOnline string `json:"last_online"`
	Born       string `json:"born"`
}

type PrivateInfo struct {
	Diamond    int    `json:"diamond"`
	Gold       string `json:"gold"`
	Experience string `json:"experience"`
	Melee      int    `json:"melee"`
	Distance   int    `json:"distance"`
	Magic      int    `json:"magic"`
	Defense    int    `json:"defense"`
	TimePlay   string `json:"time_play"`
}

type LogPvp struct {
	Killer   []Characters
	KillTime string `json:"killtime"`
	Killed   Characters
}

func GetMyChar(nickname string) (Response, error) {
	var log []Log
	var obj Characters
	var obj2 PrivateInfo
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
	sqlStatementPv := "SELECT * FROM private"
	con.QueryRow(sqlStatementPv, nickname).Scan(
		&obj2.Diamond, &obj2.Gold, &obj2.Experience, &obj2.Melee, &obj2.Distance, &obj2.Magic, &obj2.Defense, &obj2.TimePlay,
	)

	response := Characters{
		Id:         obj.Id,
		Name:       obj.Name,
		Level:      obj.Level,
		Guild:      obj.Guild,
		LastOnline: obj.LastOnline,
		Born:       obj.Born,
	}
	privateResponse := PrivateInfo{
		Diamond:    obj2.Diamond,
		Gold:       obj2.Gold,
		Experience: obj2.Experience,
		Melee:      obj2.Melee,
		Distance:   obj2.Distance,
		Magic:      obj2.Magic,
		Defense:    obj2.Defense,
		TimePlay:   obj2.TimePlay,
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = map[string]interface{}{
		"content":             response,
		"log":                 log,
		"private_information": privateResponse,
	}

	return res, nil
}
