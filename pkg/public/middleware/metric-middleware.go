package middleware

import (
	"2019_1_TheBang/pkg/public/metric"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func MetricMiddleware(c *gin.Context) {
	c.Next()
	if !(c.Writer.Status() == http.StatusOK &&
		c.Request.URL.Path == "/metrics") {

		metric.MonitoringMutex.Lock()
		defer metric.MonitoringMutex.Unlock()

		metric.HttpReqs.WithLabelValues(strconv.Itoa(c.Writer.Status())).Add(1)
	}
}
