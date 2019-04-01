package leaderboard

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"2019_1_TheBang/api"
	"2019_1_TheBang/pkg/user"
	

	"github.com/gorilla/mux"
)

func TestLeaderbordHandlerSUCCESS(t *testing.T) {
	tom := api.Signup{
		Nickname: "tom",
		Passwd:   "tom",
	}
	bob := api.Signup{
		Nickname: "bob",
		Passwd:   "bob",
	}
	user.CreateUser(&tom)
	user.CreateUser(&bob)
	defer user.DeleteUser("bob")
	defer user.DeleteUser("tom")

	req, err := http.NewRequest("GET", "/leaderbord/1", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	req = mux.SetURLVars(req, map[string]string{
		"page": "1",
	})

	rr := httptest.NewRecorder()
	LeaderbordHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("TestLeaderbordHandler: expected %v, have %v!\n", http.StatusOK, rr.Code)
	}
}

func TestLeaderbordHandlerFAIL(t *testing.T) {
	path := "/leaderboard/0"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	req = mux.SetURLVars(req, map[string]string{
		"page": "0",
	})

	rr := httptest.NewRecorder()
	LeaderbordHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("TestLeaderbordHandler: expected %v, have %v!\n", http.StatusBadRequest, rr.Code)
	}
}
