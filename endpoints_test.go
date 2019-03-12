package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserCreateConflictHandler(t *testing.T) {

	path := "/user"

	tt := []struct {
		body   string
		status int
	}{{`{
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

func TestLoginAcount(t *testing.T) {

	tt := []struct {
		body   string
		status int
	}{
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
