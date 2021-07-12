package models

import (
	"fmt"
	"net/http"
	"strconv"

	"valorant/db"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type Agent struct {
	Data   Media
	Status string `json:"status"`
}

type SelectAgent struct {
	No               int    `json:"no"`
	NameAgent        string `json:"name_agent" validate:"required"`
	Role             string `json:"role" validate:"required"`
	DescriptionAgent string `json:"description_agent" validate:"required"`
}

type SpecialAbilities struct {
	ID     int    `json:"id"`
	Image  string `json:"image"`
	Video  string `json:"video"`
	Skill1 string `json:"skill_1,omitempty"`
	Skill2 string `json:"skill_2,omitempty"`
	Skill3 string `json:"skill_3,omitempty"`
	Skill4 string `json:"skill_4,omitempty"`
}

type Media struct {
	Media []SpecialAbilities
	yoru  SelectAgent
}

func GetAgent(c echo.Context) error {
	var pilih []SpecialAbilities
	for i := 0; i < 3; i++ {
		pilih = append(pilih, SpecialAbilities{
			ID:     i,
			Image:  fmt.Sprintf("image %d", i),
			Video:  fmt.Sprintf("video %d", i),
			Skill1: fmt.Sprintf("skill1 %d", i),
			Skill2: fmt.Sprintf("skill2 %d", i),
			Skill3: fmt.Sprintf("skill3 %d", i),
			Skill4: fmt.Sprintf("skill4 %d", i),
		})
	}
	return c.JSON(http.StatusOK, Agent{
		Data: Media{
			Media: pilih,
			yoru: SelectAgent{
				NameAgent:        "name_agent",
				Role:             "role_agent",
				DescriptionAgent: "description_agent",
			},
		},
		Status: "success",
	})
}

func DetailAgent(c echo.Context) error {
	SelectAgentno := c.QueryParam("no")
	no, _ := strconv.Atoi(SelectAgentno)
	Agent := SelectAgent{
		No:               no,
		NameAgent:        fmt.Sprintf("Namaagent%d", 1),
		Role:             fmt.Sprintf("role%d", 1),
		DescriptionAgent: fmt.Sprintf("desc%d", 1),
	}
	return c.JSON(http.StatusOK, Agent)

}
func SearchDetailAgent(c echo.Context) error {
	var value []SelectAgent
	SearchAgent := c.QueryParam("search")
	no, _ := strconv.Atoi(SearchAgent)
	for _, v := range value {
		value = append(value, SelectAgent{
			No:               12,
			NameAgent:        fmt.Sprintln(v, SearchAgent),
			Role:             fmt.Sprintln(v, SearchAgent),
			DescriptionAgent: fmt.Sprintln(v, SearchAgent),
		})
	}
	return c.JSON(http.StatusOK, &SelectAgent{
		No:        no,
		NameAgent: fmt.Sprint(SearchAgent),
	})

}

func FetchAllAgent() (Response, error) {
	var obj SelectAgent
	var arrobj []SelectAgent
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM agent"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.No, &obj.NameAgent, &obj.Role, &obj.DescriptionAgent)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StoreAgent(nameagent string, role string, descriptionagent string) (Response, error) {
	var res Response

	v := validator.New()

	peg := SelectAgent{
		NameAgent:        nameagent,
		Role:             role,
		DescriptionAgent: descriptionagent,
	}

	err := v.Struct(peg)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := "INSERT agent (nameagent, role, descriptionagent) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nameagent, role, descriptionagent)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateAgent(No int, NameAgent string, Role string, DescriptionAgent string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE agent SET nameagent = ?, role = ?, descriptionagent = ? WHERE no = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(NameAgent, Role, DescriptionAgent, No)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteAgent(No int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM agent WHERE no = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(No)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
