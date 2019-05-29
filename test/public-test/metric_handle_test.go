package test

import (
	"2019_1_TheBang/pkg/public/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMetricHandle(t *testing.T) {
	router := gin.Default()
	router.Use(middleware.MetricMiddleware)
	router.OPTIONS("/options", func(c *gin.Context) {})

	req, err := http.NewRequest("OPTIONS", "/options", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("TestCorsHeaders: expected %v, have %v!\n",
			http.StatusOK, rr.Code)
	}
}
