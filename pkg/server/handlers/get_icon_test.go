package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetIconHandler(t *testing.T) {
	path := "/icon/default_img"

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, GetIconHandler)
	router.ServeHTTP(rr, req)

	//не отправляет из этой директории изображение
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("TestLogoutHandler, have not cookie: expected %v, have %v!\n",
			http.StatusInternalServerError , rr.Code)
	}
}
