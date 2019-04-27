package test

import (
	"2019_1_TheBang/pkg/main-serivce-pkg/logout"
	"2019_1_TheBang/pkg/public/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLogoutHandlerFAIL(t *testing.T) {
	path := "/auth"
	req, err := http.NewRequest("DELETE", path, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := gin.Default()
	router.Use(middleware.AuthMiddlewareGin)
	router.DELETE(path, logout.LogoutHandler)

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

	router := gin.Default()
	router.Use(middleware.AuthMiddlewareGin)
	router.DELETE(path, logout.LogoutHandler)

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("TestLogoutHandler, have not cookie: expected %v, have %v!\n", http.StatusOK, rr.Code)
	}
}
