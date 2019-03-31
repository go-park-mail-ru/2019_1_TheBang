package handlers

import (
	"net/http"
	"strconv"

	"2019_1_TheBang/pkg/server/models"

	"github.com/gorilla/mux"
)

type VarsHandler func(w http.ResponseWriter, r *http.Request, vars map[string]string)

func (vh VarsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vh(w, r, vars)
}

func LeaderbordHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := vars["page"]
	number, err := strconv.Atoi(page)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	if number == 0 {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	json, status := models.LeaderPage(uint(number))
	if status != http.StatusOK {
		w.WriteHeader(status)

		return
	}

	w.Write(json)
}
