package test

import (
	"2019_1_TheBang/config/gameconfig"
	"2019_1_TheBang/pkg/game-service-pkg/app"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// router.POST("/room", app.CreateRoomHandle)

func TestCreateRoomHandleEasy(t *testing.T) {
	gameconfig.CONFIGPATH = "../../config/gameconfig"
	gameconfig.InitGameConfig()
	app.InitAppInst()
	app.AppInst.NewRoom()

	path := "/room"

	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	rr := httptest.NewRecorder()

	router := gin.Default()
	router.POST(path, app.CreateRoomHandle)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("TestCreateRoomHandleEasy: expected %v, have %v!\n", http.StatusOK, rr.Code)
	}
}

func TestCreateRoomHandle(t *testing.T) {
	gameconfig.CONFIGPATH = "../../config/gameconfig"
	gameconfig.InitGameConfig()
	app.InitAppInst()
	app.AppInst.NewRoom()

	path := "/room"

	req, err := http.NewRequest("POST", path, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	router := gin.Default()
	router.POST(path, app.CreateRoomHandle)

	for i := 0; i < int(gameconfig.MaxRoomsInGame); i++ {
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusConflict {
		t.Errorf("TestCreateRoomHandle: expected %v, have %v!\n", http.StatusOK, rr.Code)
	}
}
