package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"2019_1_TheBang/api"
	"2019_1_TheBang/pkg/main-serivce-pkg/login"

	"github.com/gorilla/mux"
)

func TestLogInHandlerFAIL(t *testing.T) {
	path := "/auth"

	fakeNick := "smbdy"
	bodyStruct := api.Login{
		Passwd:   fakeNick,
		Nickname: fakeNick,
	}
	body, _ := json.Marshal(bodyStruct)

	req, err := http.NewRequest("POST", path, bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, login.LogInHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("TestLogoutHandler, have not cookie: expected %v, have %v!\n", http.StatusUnauthorized, rr.Code)
	}
}

func TestLogInHandlerSUCCESS(t *testing.T) {
	cookie, err := GetTESTAdminCookie()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer DeleteTESTAdmin()

	path := "/auth"
	bodyStruct := api.Login{
		Passwd:   testAdminNick,
		Nickname: testAdminNick,
	}
	body, _ := json.Marshal(bodyStruct)

	req, err := http.NewRequest("POST", path, bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, login.LogInHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("TestLogoutHandler, have not cookie: expected %v, have %v!\n", http.StatusOK, rr.Code)
	}
}
