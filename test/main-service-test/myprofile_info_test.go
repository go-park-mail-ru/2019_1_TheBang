package test

import (
	"2019_1_TheBang/pkg/middleware"
	"2019_1_TheBang/pkg/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestMyProfileInfoHandlerSUCCESS(t *testing.T) {
	cookie, err := GetTESTAdminCookie()
	if err != nil {
		t.Fatal(err.Error())
	}
	defer DeleteTESTAdmin()

	path := "/user"

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	req.AddCookie(cookie)

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, middleware.AuthMiddleware(user.MyProfileInfoHandler))
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("TestMyProfileInfoHandler: can not get valid prof's data! have: %v, expected: %v",
			rr.Code, http.StatusOK)
	}
}

func TestMyProfileInfoHandlerFAIL(t *testing.T) {
	path := "/user"

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, middleware.AuthMiddleware(user.MyProfileInfoHandler))
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("TestMyProfileInfoHandler: we should not get data withput cookie! have: %v, expected: %v",
			rr.Code, http.StatusOK)
	}
}
