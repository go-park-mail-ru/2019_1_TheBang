package gamemonitoring

import (
	"2019_1_TheBang/config/gameconfig"
	"2019_1_TheBang/pkg/game-service-pkg/app"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var GameMonitoringMutex = &sync.Mutex{}

var label = "count"

func init() {
	prometheus.MustRegister(RoomsCountMetric)

	go RoomCounter()
}

var (
	RoomsCountMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "rooms_in_game",
			Help: "How many rooms in game at the moment",
		},
		[]string{label},
	)
)

func RoomCounter() {
	ticker := time.NewTicker(gameconfig.MonitoringTick)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:

			if app.AppInst != nil {
				app.AppInst.Locker.Lock()
				count := float64(app.AppInst.RoomsCount)
				app.AppInst.Locker.Unlock()

				RoomsCountMetric.WithLabelValues(label).Set(count)
			}
		}
	}
}
