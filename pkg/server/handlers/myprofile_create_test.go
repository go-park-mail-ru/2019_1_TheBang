package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"2019_1_TheBang/pkg/server/models"

	"github.com/gorilla/mux"
)

func TestMyProfileCreateHandler(t *testing.T) {
	path := "/user"

	tt := []TestCase{{`{"nickname": "lil", "passwd": "man"}`, http.StatusCreated}}

	for _, tc := range tt {
		req, err := http.NewRequest("POST", path, strings.NewReader(tc.Body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc(path, MyProfileCreateHandler)
		router.ServeHTTP(rr, req)

		if rr.Code != tc.Status {
			t.Errorf("Error: expected %v, have %v!\n", tc.Status, rr.Code)
		}
	}

	if ok := models.DeleteUser("lil"); !ok {
		t.Fatal("User was not deleted")
	}
}
