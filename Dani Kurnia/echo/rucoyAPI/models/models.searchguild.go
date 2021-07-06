package models

import (
	"database/sql"
	"echo/RucoyAPI/db"
	"fmt"
	"net/http"
)

type SearchGuilds struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Members     string `json:"members"`
}

func SearchGuild(value string) (Response, error) {
	var obj SearchGuilds
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM guild WHERE name = ?"

	err := con.QueryRow(sqlStatement, value).Scan(
		&obj.Id, &obj.Name, &obj.Description, &obj.Date, &obj.Members,
	)
	if err == sql.ErrNoRows {
		fmt.Println("Guild not found")
		return res, err
	}
	if err != nil {
		return res, err
	}

	response := SearchGuilds{
		Id:          obj.Id,
		Name:        obj.Name,
		Description: obj.Description,
		Date:        obj.Date,
		Members:     obj.Members,
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = map[string]interface{}{
		"content": response,
	}

	return res, nil
}
