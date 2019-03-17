package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

//func TestLeaderbordHandlerSUCCESS(t *testing.T) {
//	path := "/leaderbord/1"
//	req, err := http.NewRequest("GET", path, nil)
//	if err != nil {
//		t.Fatal(err.Error())
//	}
//
//	rr := httptest.NewRecorder()
//	router := mux.NewRouter()
//	router.HandleFunc(path, LeaderbordHandler)
//	router.ServeHTTP(rr, req)
//
//	if rr.Code != http.StatusBadRequest {
//		t.Errorf("TestLeaderbordHandler: expected %v, have %v!\n", http.StatusOK, rr.Code)
//	}
//}

func TestLeaderbordHandlerFAIL(t *testing.T) {
	path := "/leaderboard/0"
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(path, LeaderbordHandler)
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("TestLeaderbordHandler: expected %v, have %v!\n", http.StatusBadRequest, rr.Code)
	}
}
