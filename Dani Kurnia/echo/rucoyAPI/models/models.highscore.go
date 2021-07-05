package models

import (
	"echo/Rucoy/db"
	"net/http"
)

type TopExperience struct {
	Rank       int    `json:"rank"`
	Name       string `json:"name"`
	Level      int    `json:"level"`
	Experience string `json:"experience"`
}

type TopMelee struct {
	Rank  int    `json:"rank"`
	Name  string `json:"name"`
	Melee string `json:"meele"`
}

type TopDist struct {
	Rank     int    `json:"rank"`
	Name     string `json:"name"`
	Distance string `json:"meele"`
}

type TopMage struct {
	Rank  int    `json:"rank"`
	Name  string `json:"name"`
	Magic string `json:"magic"`
}

type TopDef struct {
	Rank    int    `json:"rank"`
	Name    string `json:"name"`
	Defense string `json:"defense"`
}

func FetchAllTopExp() (Response, error) {
	var obj TopExperience
	var arrobj []TopExperience
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM topexp"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Rank, &obj.Name, &obj.Level, &obj.Experience)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]interface{}{
		"1 title": "Top 10 - Experience",
		"content": arrobj,
	}

	return res, nil
}

func FetchAllTopMelee() (Response, error) {
	var obj TopMelee
	var arrobj []TopMelee
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM topmelee"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Rank, &obj.Name, &obj.Melee)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]interface{}{
		"1 title": "Top 10 - Melee",
		"content": arrobj,
	}

	return res, nil
}

func FetchAllTopDistance() (Response, error) {
	var obj TopDist
	var arrobj []TopDist
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM topdistance"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Rank, &obj.Name, &obj.Distance)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]interface{}{
		"1 title": "Top 10 - Distance",
		"content": arrobj,
	}

	return res, nil
}

func FetchAllTopMagic() (Response, error) {
	var obj TopMage
	var arrobj []TopMage
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM topmagic"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Rank, &obj.Name, &obj.Magic)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]interface{}{
		"1 title": "Top 10 - Magic",
		"content": arrobj,
	}

	return res, nil
}

func FetchAllTopDefense() (Response, error) {
	var obj TopDef
	var arrobj []TopDef
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM topdefense"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Rank, &obj.Name, &obj.Defense)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]interface{}{
		"1 title": "Top 10 - Defense",
		"content": arrobj,
	}

	return res, nil
}
