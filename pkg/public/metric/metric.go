package metric

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var MonitoringMutex = &sync.Mutex{}

var (
	HttpReqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
		[]string{"code"},
	)
)

func init() {
	prometheus.MustRegister(HttpReqs)
	HttpReqs.WithLabelValues("200")
	HttpReqs.WithLabelValues("201")
	HttpReqs.WithLabelValues("400")
	HttpReqs.WithLabelValues("401")
	HttpReqs.WithLabelValues("404")
	HttpReqs.WithLabelValues("500")
}
