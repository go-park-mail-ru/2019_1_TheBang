package handlers

import (
	"github.com/go-park-mail-ru/2019_1_TheBang/pkg/server/models"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMyProfileCreateHandler(t *testing.T) {
	path := "/user"

	tt := []TestCase{{`{
			"nickname": "man",
			"name": "Ivan",
			"surname": "Ivanov",
			"dob": "01.11.1968",
			"passwd": "tvoe_kakoe_delo"
		}`, http.StatusCreated},
		{`{
			"nickname": "man",
			"name": "Ivan",
			"surname": "Ivanov",
			"dob": "01.11.1968",
			"passwd": "tvoe_kakoe_delo"
		}`, http.StatusConflict},
	}

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

	if ok := models.DeleteUser("man"); !ok {
		t.Fatal("User was not deleted")
	}
}