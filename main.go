package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"xianrail_exporter/metrics"
)

func main() {
	metrics.RecordMetrics()
	go initMetricsServer()
	go initAppServer()
	infiniteLoop()
}

func initAppServer() {
	appServeMux := http.NewServeMux()
	log.Printf("listening port 2113 ...")
	appServeMux.Handle("/metrics", http.RedirectHandler("https://www.baidu.com", 302))
	err := http.ListenAndServe(":2114", appServeMux)
	if err != nil {
		panic(err)
	}
}

func initMetricsServer() {
	metricsServeMux := http.NewServeMux()
	log.Printf("listening port 2112 ...")
	metricsServeMux.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2115", metricsServeMux)
	if err != nil {
		panic(err)
	}
}

func infiniteLoop() {
	select {}
}
