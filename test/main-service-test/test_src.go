package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"

	"2019_1_TheBang/api"
	"2019_1_TheBang/pkg/main-serivce-pkg/login"
	"2019_1_TheBang/pkg/main-serivce-pkg/user"

	"github.com/gin-gonic/gin"
)

// type TestCase struct {
// 	Body   string
// 	Status int
// }

var testAdminNick = "testadmin"

func GetTESTAdminCookie() (*http.Cookie, error) {
	prof := api.Signup{
		Nickname: testAdminNick,
		Name:     testAdminNick,
		Surname:  testAdminNick,
		DOB:      "2017-01-01",
		Passwd:   testAdminNick,
	}

	body, _ := json.Marshal(prof)
	path := "/user"

	req, err := http.NewRequest("POST", path, bytes.NewReader(body))
	if err != nil {
		return nil, errors.New("request was failed")
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router := gin.Default()
	router.POST(path, user.MyProfileCreateHandler)

	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusCreated {
		log.Fatal("Admin was not created!")
	}

	logApi := api.Login{
		Nickname: testAdminNick,
		Passwd:   testAdminNick,
	}

	body, _ = json.Marshal(logApi)
	path = "/auth"

	req, err = http.NewRequest("POST", path, bytes.NewReader(body))
	if err != nil {
		return nil, errors.New("request was failed")
	}
	req.Header.Set("Content-Type", "application/json")

	rr = httptest.NewRecorder()

	router = gin.Default()
	router.POST(path, login.LogInHandler)

	router.ServeHTTP(rr, req)

	// toDO сделать проверку на этот момент (есть ли кука)
	cookie := rr.Result().Cookies()[0]

	return cookie, nil
}

func DeleteTESTAdmin() {
	ok := user.DeleteUser(testAdminNick)
	if !ok {
		log.Fatal("Can not delete testAdmin!")
	}
}
