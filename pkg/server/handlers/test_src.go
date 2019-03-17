package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2019_1_TheBang/api"
	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
)

type TestCase struct {
	Body   string
	Status int
}

var testAdminNick = "testadmin"

func GetTESTAdminCookie() (*http.Cookie, error) {
	prof := api.Signup{
		Nickname: testAdminNick,
		Name: testAdminNick,
		Surname: testAdminNick,
		DOB: "2017-01-01",
		Passwd: testAdminNick,
	}

	body, _ := json.Marshal(prof)
	path := "/user"

	req, err := http.NewRequest("POST", path, bytes.NewReader(body))
	if err != nil {
		return nil, errors.New("request was failed")
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, MyProfileCreateHandler)
	router.ServeHTTP(rr, req)

	login := api.Login{
		Nickname: testAdminNick,
		Passwd: testAdminNick,
	}

	body, _ = json.Marshal(login)
	path = "/auth"

	req, err = http.NewRequest("POST", path, bytes.NewReader(body))
	if err != nil {
		return nil, errors.New("request was failed")
	}
	req.Header.Set("Content-Type", "application/json")

	rr = httptest.NewRecorder()
	router = mux.NewRouter()
	router.HandleFunc(path, LogInHandler)
	router.ServeHTTP(rr, req)

	cookie := rr.Result().Cookies()[0]

	return cookie, nil
}

func DeleteTESTAdmin() {
	ok := models.DeleteUser(testAdminNick)
	if !ok  {
		log.Fatal("Can not delete testAdmin!")
	}
}

