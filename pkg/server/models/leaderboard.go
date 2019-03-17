package models

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2019_1_TheBang/api"
	"github.com/go-park-mail-ru/2019_1_TheBang/config"
	"log"
	"net/http"
)

func LeaderPage(number uint) (jsonData []byte, status int) {
	offset := config.RowsOnLeaderPage * (number - 1)
	rows, err := config.DB.Query(SQLGetPage, config.RowsOnLeaderPage, offset)
	if err != nil {
		log.Printf("LeaderPage: %v\n", )
		return jsonData, http.StatusInternalServerError
	}

	profs := []api.Profile{}
	for rows.Next() {
		p := api.Profile{}
		if err := rows.Scan(&p.Nickname,
			&p.Name,
			&p.Surname,
			&p.DOB,
			&p.Photo,
			&p.Score);
		err != nil {
			log.Printf("LeaderPage: %v\n", err.Error())

			return jsonData ,http.StatusInternalServerError
		}

		profs = append(profs, p)
	}

	if len(profs) == 0 {
		return jsonData, http.StatusNotFound
	}

	jsonData, err = json.Marshal(profs)
	if err != nil {
		log.Printf("LeaderPage: %v", err.Error())

		return jsonData ,http.StatusInternalServerError
	}

	return jsonData, http.StatusOK
}

var SQLGetPage = `select nickname, name, surname, dob, photo, score from project_bang.users
					order by score
					limit $1 offset $2`