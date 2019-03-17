package main

import (
	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

//
//import (
//	"github.com/dgrijalva/jwt-go"
//	"github.com/go-park-mail-ru/2019_1_TheBang/config"
//	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/auth"
//	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/handlers"
//	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/models"
//	"github.com/gorilla/mux"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//	"time"
//)
//
//

//
//
//var cookie http.Cookie
//
//
func TestLeaderbordHandler(t *testing.T) {
	path := "/leaderboard/1"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, handlers.LeaderbordHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("TestLeaderbordHandler: expected %v, have %v!\n", http.StatusOK, rr.Code)
	}
}
//
//func TestCheckTocken(t *testing.T) {
//	claims := models.CustomClaims{
//		"admin",
//		jwt.StandardClaims{
//			Issuer: config.ServerName,
//		},
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	ss, _ := token.SignedString("NOT SECRET")
//
//	expiration := time.Now().Add(10 * time.Hour)
//	cookie := http.Cookie{
//		Name:     config.CookieName,
//		Value:    ss,
//		Expires:  expiration,
//		HttpOnly: true,
//	}
//
//	req, _ := http.NewRequest("GET", "/", nil)
//	req.AddCookie(&cookie)
//
//	_, ok := auth.CheckTocken(req)
//	if ok {
//		t.Errorf("TestCheckTocken: faked token was accepted!")
//	}
//}
//
////func TestNicknameFromCookie(t *testing.T) {
////	cookie, err := getAdminCookie()
////	if err != nil {
////		t.Fatal(err.Error())
////	}
////
////	rr := httptest.NewRecorder()
////	req, _ := http.NewRequest("GET", "/", nil)
////	req.AddCookie(cookie)
////
////	nickname, _ := handlers.NicknameFromCookie(rr, req)
////	if nickname != "admin" {
////		t.Errorf("TestNicknameFromCookie: admin's cookie was not recognized!")
////	}
////}
//
//func TestGetIconHandler(t *testing.T) {
//	filename := "the is no this image"
//	path := "/icon" + filename
//	req, _ := http.NewRequest("GET", path, nil)
//
//	rr := httptest.NewRecorder()
//	router := mux.NewRouter()
//	router.HandleFunc(path, handlers.GetIconHandler)
//	router.ServeHTTP(rr, req)
//
//	if rr.Code != http.StatusInternalServerError {
//		t.Errorf("TestGetIconHandler: there is no this image (%v)!", filename)
//	}
//}