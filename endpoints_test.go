package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type TestCase struct {
body   string
status int
}

func getAdminCookie() (*http.Cookie, error) {
	cookieBody := `{"nickname": "admin", "passwd": "admin"}`
	cookiePath := "/auth"

	req, err := http.NewRequest("POST", cookiePath, strings.NewReader(cookieBody))
	if err != nil {
		return nil, errors.New("request was failed")
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(cookiePath, LogInHandler)
	router.ServeHTTP(rr, req)

	cookie := rr.Result().Cookies()[0]

	return cookie, nil
}

func TestMyProfileCreateHandler(t *testing.T) {

	path := "/user"

	tt := []TestCase{{`{
			"nickname": "Nikita",
			"name": "Ivan",
			"surname": "Ivanov",
			"dob": "01.11.1968",
			"passwd": "tvoe_kakoe_delo"
		}`, http.StatusCreated},
		{`{
			"nickname": "Nikita",
			"name": "Ivan",
			"surname": "Ivanov",
			"dob": "01.11.1968",
			"passwd": "tvoe_kakoe_delo"
		}`, http.StatusConflict},
	}

	for _, tc := range tt {
		req, err := http.NewRequest("POST", path, strings.NewReader(tc.body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc(path, MyProfileCreateHandler)
		router.ServeHTTP(rr, req)

		if rr.Code != tc.status {
			t.Errorf("Error: expected %v, have %v!\n", tc.status, rr.Code)
		}
	}
}

var cookie http.Cookie

func TestLogInHandler(t *testing.T) {

	tt := []TestCase{
		{`{
  					"nickname": "admin",
  					"passwd": "admin"
				}`, http.StatusOK},
		{
			`{
  					"nickname": "admin",
  					"passwd": "odmin"
				}`, http.StatusUnauthorized},
	}

	path := "/auth"

	for _, tc := range tt {
		req, err := http.NewRequest("POST", path, strings.NewReader(tc.body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc(path, LogInHandler)
		router.ServeHTTP(rr, req)

		if rr.Code != tc.status {
			t.Errorf("Error: expected %v, have %v!\n", tc.status, rr.Code)
		}
	}


}

func TestMyProfileInfoHandler(t *testing.T) {
	cookie, err := getAdminCookie()
	if err != nil {
		t.Fatal(err.Error())
	}

	path := "/user"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, MyProfileInfoHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusForbidden {
		t.Errorf("Error: expected %v, have %v!\n", http.StatusOK, rr.Code)
	}

	req.AddCookie(cookie)

	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Error: expected %v, have %v!\n", http.StatusOK, rr.Code)
	}
}

func TestMyProfileInfoUpdateHandler(t *testing.T) {
	cookie, err := getAdminCookie()
	if err != nil {
		t.Fatal(err.Error())
	}

	body := `{"name": "Bob","surname": "Bobov","dob": "01.11.1968"}`
	path := "/user"
	req, err := http.NewRequest("PUT", path, strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, MyProfileInfoUpdateHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusForbidden {
		t.Errorf("MyProfileInfoUpdateHandler, have not cookie: expected %v, have %v!\n", http.StatusOK, rr.Code)
	}

	req.AddCookie(cookie)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusAccepted {
		t.Errorf("MyProfileInfoUpdateHandler, have cookie: expected %v, have %v!\n", http.StatusAccepted, rr.Code)
	}
}

func TestLogoutHandler(t *testing.T) {
	cookie, err := getAdminCookie()
	if err != nil {
		t.Fatal(err.Error())
	}

	path := "/auth"
	req, err := http.NewRequest("DELETE", path, nil)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, LogoutHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("TestLogoutHandler, have not cookie: expected %v, have %v!\n", http.StatusOK, rr.Code)
	}

	rr = httptest.NewRecorder()
	req.AddCookie(cookie)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("TestLogoutHandler, have cookie: expected %v, have %v!\n", http.StatusOK, rr.Code)
	}
}

func TestLeaderbordHandler(t *testing.T) {
	path := "/leaderboard"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, LeaderbordHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("TestLeaderbordHandler: expected %v, have %v!\n", http.StatusOK, rr.Code)
	}
}

func TestCheckTocken(t *testing.T) {
	claims := customClaims{
		"admin",
		jwt.StandardClaims{
			Issuer: ServerName,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString("NOT SECRET")

	expiration := time.Now().Add(10 * time.Hour)
	cookie := http.Cookie{
		Name:     CookieName,
		Value:    ss,
		Expires:  expiration,
		HttpOnly: true,
	}

	req, _ := http.NewRequest("GET", "/", nil)
	req.AddCookie(&cookie)

	_, ok := CheckTocken(req)
	if ok {
		t.Errorf("TestCheckTocken: faked token was accepted!")
	}
}

func TestNicknameFromCookie(t *testing.T) {
	cookie, err := getAdminCookie()
	if err != nil {
		t.Fatal(err.Error())
	}

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.AddCookie(cookie)

	nickname, _ := NicknameFromCookie(rr, req)
	if nickname != "admin" {
		t.Errorf("TestNicknameFromCookie: admin's cookie was not recognized!")
	}
}

func TestGetIconHandler(t *testing.T) {
	filename := "the is no this image"
	path := "/icon" + filename
	req, _ := http.NewRequest("GET", path, nil)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, GetIconHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("TestGetIconHandler: there is no this image (%v)!", filename)
	}
}