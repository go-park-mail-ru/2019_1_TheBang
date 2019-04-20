package leaderboard

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

type VarsHandler func(w http.ResponseWriter, r *http.Request, vars map[string]string)

func (vh VarsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vh(w, r, vars)
}

func LeaderbordHandler(c *gin.Context) {
	page := c.GetInt("page")
	if page == 0 {
		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	json, status := LeaderPage(uint(page))
	if status != http.StatusOK {
		w.WriteHeader(status)

		return
	}

	w.Write(json)
}
