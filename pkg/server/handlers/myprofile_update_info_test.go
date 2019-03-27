package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/middlewares"
	"github.com/gorilla/mux"
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
	router := mux.NewRouter()
	router.HandleFunc(path, middlewares.AuthMiddleware(MyProfileInfoUpdateHandler))
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
	router := mux.NewRouter()
	router.HandleFunc(path, middlewares.AuthMiddleware(MyProfileInfoUpdateHandler))
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("MyProfileInfoUpdateHandler, have cookie: expected %v, have %v!\n", http.StatusUnauthorized, rr.Code)
	}
}
