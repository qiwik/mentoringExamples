package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/count", countHandler)
	http.HandleFunc("/histogram", histogramHandler)
	http.HandleFunc("/gauge", gaugeHandler)
	http.HandleFunc("/summary", summaryHandler)

	registerPrometheus()

	log.Println("Listening on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("can't run server")
	}
}
