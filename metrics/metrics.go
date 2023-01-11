package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log"
	"time"
	"xianrail_exporter/api/xianrail"
)

var (
	total = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "xianrail",
		Subsystem: "passenger",
		Name:      "total",
		Help:      "Total count of passengers of all stations.",
	})
)

func RecordMetrics() {
	go func() {
		for {
			var totalValue int64 = 0
			data, err := xianrail.RequestAlarm()
			if err != nil {
				continue
			}
			for lineIndex, stationReal := range data.PfReal.PfStationReal {
				log.Printf("parsing #%d line: %s", lineIndex, *stationReal.LineName)
				for stationIndex, station := range stationReal.Stations {
					log.Printf("  - parsing #%d station: %s", stationIndex, *station.StationName)
					int64number, err := station.AllQuatity.Int64()
					if station.AllQuatity != nil && err == nil {
						totalValue += int64number
					} else {
						log.Printf("Invalid AllQuantity value(%v) for station(%s): err(%v)", *station.AllQuatity, *station.StationName, err)
					}
				}
			}
			log.Printf("total: %d", totalValue)
			total.Set(float64(totalValue))
			time.Sleep(1 * time.Minute)
		}
	}()
}
