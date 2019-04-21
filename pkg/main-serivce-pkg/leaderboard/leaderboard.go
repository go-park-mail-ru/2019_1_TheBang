package leaderboard

import (
	"net/http"

	"2019_1_TheBang/api"
	"2019_1_TheBang/config"
	"2019_1_TheBang/config/mainconfig"
)

func LeaderPage(number uint) (profs []api.Profile, status int) {
	offset := mainconfig.RowsOnLeaderPage * (number - 1)
	rows, err := mainconfig.DB.Query(SQLGetPage, mainconfig.RowsOnLeaderPage, offset)
	if err != nil {
		config.Logger.Warnw("LeaderPage",
			"warn", err.Error())

		return profs, http.StatusInternalServerError
	}

	profs = []api.Profile{}
	for rows.Next() {
		p := api.Profile{}
		if err := rows.Scan(&p.Nickname,
			&p.Name,
			&p.Surname,
			&p.DOB,
			&p.Photo,
			&p.Score); err != nil {
			config.Logger.Warnw("LeaderPage",
				"warn", err.Error())

			return profs, http.StatusInternalServerError
		}

		profs = append(profs, p)
	}

	if len(profs) == 0 {
		return profs, http.StatusNotFound
	}

	return profs, http.StatusOK
}

var SQLGetPage = `select nickname, name, surname, dob, photo, score from ` + mainconfig.DBSCHEMA + `users
					order by score desc
					limit $1 offset $2`
