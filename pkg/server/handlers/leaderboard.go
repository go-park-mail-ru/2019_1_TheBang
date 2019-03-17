package handlers

import (
	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func LeaderbordHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page := vars["page"]
	number, err := strconv.Atoi(page)
	if err != nil {
		log.Printf("LeaderbordHandler: %v\n", err)
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
