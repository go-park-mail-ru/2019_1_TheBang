package test

import (
	"2019_1_TheBang/pkg/game-service-pkg/app"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRoomListHandle(t *testing.T) {
	app.InitAppInst()

	path := "/room"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	rr := httptest.NewRecorder()

	router := gin.Default()
	router.GET(path, app.RoomsListHandle)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("TestRoomListHandle: expected %v, have %v!\n", http.StatusOK, rr.Code)
	}
}
