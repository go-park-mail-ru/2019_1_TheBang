package middleware

import (
	"2019_1_TheBang/pkg/public/monitoring"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MetricMiddleware(c *gin.Context) {
	c.Next()
	if !(c.Writer.Status() == http.StatusOK &&
		c.Request.URL.Path == "/metrics") {

		monitoring.MonitoringMutex.Lock()
		defer monitoring.MonitoringMutex.Unlock()

		monitoring.HttpReqs.WithLabelValues(strconv.Itoa(c.Writer.Status())).Add(1)
	}
}
