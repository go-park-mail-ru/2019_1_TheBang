package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetIconHandlerSuccess(t *testing.T) {
	path := "/icon/default_img"

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"filename": "default_img",
	})

	rr := httptest.NewRecorder()
	GetIconHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("TestLogoutHandler, have not cookie: expected %v, have %v!\n",
			http.StatusOK, rr.Code)
	}
}

func TestGetIconHandlerFAIL(t *testing.T) {
	path := "/icon/blablabla"

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"filename": "blablabla",
	})

	rr := httptest.NewRecorder()
	GetIconHandler(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("TestLogoutHandler, have not cookie: expected %v, have %v!\n",
			http.StatusInternalServerError, rr.Code)
	}
}
