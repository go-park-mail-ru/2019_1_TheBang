package test

import (
	"2019_1_TheBang/pkg/main-serivce-pkg/user"
	"github.com/gin-gonic/gin"

	"2019_1_TheBang/pkg/public/middleware"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMyProfileInfoUpdateHandlerSUCCESS(t *testing.T) {
	cookie, err := GetTESTAdminCookie()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer DeleteTESTAdmin()

	body := `{"name": "Bob","surname": "Bobov","dob": "01.11.1968"}`
	path := "/user"
	req, err := http.NewRequest("PUT", path, strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(cookie)

	rr := httptest.NewRecorder()

	router := gin.Default()
	router.PUT(path, user.MyProfileInfoUpdateHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("MyProfileInfoUpdateHandler, have cookie: expected %v, have %v!\n", http.StatusOK, rr.Code)
	}
}

func TestMyProfileInfoUpdateHandlerFAIL(t *testing.T) {
	body := `{"name": "Bob","surname": "Bobov","dob": "01.11.1968"}`
	path := "/user"
	req, err := http.NewRequest("PUT", path, strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := gin.Default()
	router.Use(middleware.AuthMiddlewareGin)

	router.PUT(path, user.MyProfileInfoUpdateHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("MyProfileInfoUpdateHandler, have cookie: expected %v, have %v!\n", http.StatusUnauthorized, rr.Code)
	}
}
