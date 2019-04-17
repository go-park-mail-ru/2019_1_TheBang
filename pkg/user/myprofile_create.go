package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"2019_1_TheBang/api"
	"2019_1_TheBang/config"
)

func MyProfileCreateHandler(w http.ResponseWriter, r *http.Request) {
	signup := api.Signup{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("MyProfileCreateHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError,
			"warn", "can not read body")

		return
	}

	err = json.Unmarshal(body, &signup)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("MyProfileCreateHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError,
			"warn", "can not marshal json")

		return
	}

	signup.DOB = "2018-01-01"

	status := CreateUser(&signup)
	if status != http.StatusCreated {
		w.WriteHeader(status)

		return
	}

	w.WriteHeader(http.StatusCreated)
}
