package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/go-park-mail-ru/2019_1_TheBang/api"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogInHandlerFAIL(t *testing.T) {
	path := "/auth"
	bodyStruct := api.Login{
		Passwd: testAdminNick,
		Nickname: testAdminNick,
	}
	body, _ := json.Marshal(bodyStruct)

	req, err := http.NewRequest("POST", path, bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, LogInHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("TestLogoutHandler, have not cookie: expected %v, have %v!\n",  http.StatusUnauthorized, rr.Code)
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
		Passwd: testAdminNick,
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
	router.HandleFunc(path, LogInHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("TestLogoutHandler, have not cookie: expected %v, have %v!\n",  http.StatusOK, rr.Code)
	}
}


