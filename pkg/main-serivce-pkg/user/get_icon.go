package user

import (
	"io/ioutil"
	"net/http"

	"2019_1_TheBang/config"
	"fmt"
	"os"

	"github.com/gorilla/mux"
)

func GetIconHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/jpeg")

	vars := mux.Vars(r)
	filename := vars["filename"]

	root, _ := os.Getwd()
	fmt.Println(root)

	filepath := root + "/tmp/" + filename
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("GetIconHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError,
			"warn", "GetIconHandler: we can not read image")

		return
	}

	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		config.Logger.Warnw("GetIconHandler",
			"RemoteAddr", r.RemoteAddr,
			"status", http.StatusInternalServerError,
			"warn", "GetIconHandler: we can not write image")

		return
	}
}