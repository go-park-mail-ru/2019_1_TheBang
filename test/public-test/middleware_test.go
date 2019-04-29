package middleware

import (
	"2019_1_TheBang/pkg/public/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCorsHeaders(t *testing.T) {
	router := gin.Default()
	router.Use(middleware.CorsMiddlewareGin)
	router.OPTIONS("/options", func(c *gin.Context) {})

	req, err := http.NewRequest("OPTIONS", "/options", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("TestCorsHeaders: expected %v, have %v!\n",
			http.StatusNoContent, rr.Code)
	}
}
