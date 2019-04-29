package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"2019_1_TheBang/api"
	"2019_1_TheBang/pkg/main-serivce-pkg/leaderboard"
	"2019_1_TheBang/pkg/main-serivce-pkg/user"
	"github.com/gin-gonic/gin"
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

	path := "/leaderbord/:page"
	req, err := http.NewRequest("GET", "/leaderbord/1", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	rr := httptest.NewRecorder()

	router := gin.Default()
	router.GET(path, leaderboard.LeaderbordHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("TestLeaderbordHandler: expected %v, have %v!\n", http.StatusOK, rr.Code)
	}
}

func TestLeaderbordHandlerFAIL(t *testing.T) {
	path := "/leaderboard/:page"
	req, err := http.NewRequest("GET", "/leaderboard/0", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	rr := httptest.NewRecorder()
	router := gin.Default()
	router.GET(path, leaderboard.LeaderbordHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("TestLeaderbordHandler: expected %v, have %v!\n", http.StatusBadRequest, rr.Code)
	}
}
