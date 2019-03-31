package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"2019_1_TheBang/pkg/server/middlewares"

	"github.com/gorilla/mux"
)

func TestLogoutHandlerFAIL(t *testing.T) {
	path := "/auth"
	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, middlewares.AuthMiddleware(LogoutHandler))
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("TestLogoutHandler, have not cookie: expected %v, have %v!\n", http.StatusUnauthorized, rr.Code)
	}
}

func TestLogoutHandlerSUCCESS(t *testing.T) {
	cookie, err := GetTESTAdminCookie()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer DeleteTESTAdmin()

	path := "/auth"
	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, middlewares.AuthMiddleware(LogoutHandler))
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("TestLogoutHandler, have not cookie: expected %v, have %v!\n", http.StatusOK, rr.Code)
	}
}
