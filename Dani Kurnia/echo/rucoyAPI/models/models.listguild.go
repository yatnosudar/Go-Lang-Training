package models

import (
	"echo/RucoyAPI/db"
	"net/http"
)

type ListGuild struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

func GetListGuild() (Response, error) {
	var obj ListGuild
	var arrobj []ListGuild
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM guild"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, nil
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Description, &obj.Date)
		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = arrobj

	return res, nil
}
