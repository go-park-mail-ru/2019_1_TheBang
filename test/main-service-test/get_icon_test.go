package test

import (
	"2019_1_TheBang/pkg/main-serivce-pkg/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetIconHandlerSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/icon/default_img", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	path := "/icon/:filename"
	router := gin.Default()
	router.GET(path, user.GetIconHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("TestLogoutHandler, have not cookie: expected %v, have %v!\n",
			http.StatusOK, rr.Code)
	}
}

func TestGetIconHandlerFAIL(t *testing.T) {
	req, err := http.NewRequest("GET", "/icon/blablabla", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	path := "/icon/:filename"
	router := gin.Default()
	router.GET(path, user.GetIconHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("TestLogoutHandler, have not cookie: expected %v, have %v!\n",
			http.StatusInternalServerError, rr.Code)
	}
}
