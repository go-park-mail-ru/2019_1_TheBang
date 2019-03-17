package handlers

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func GetIconHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")

	vars := mux.Vars(r)
	filename := vars["filename"]

	filepath := "tmp/" + filename
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("GetIconHandler: we can not read image")

		return
	}

	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("GetIconHandler: we can not write image")

		return
	}
}
