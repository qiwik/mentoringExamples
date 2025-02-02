package main

import (
	"net/http"
	"strconv"
)

func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		prometheusCounter.Add(2)
		prometheusCounterVec.WithLabelValues("LOL").Inc()

		w.Write([]byte("Success!"))

		return
	}

	w.Write([]byte("Not GET method"))
}

func histogramHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		v, _ := strconv.ParseFloat(r.Header.Get("value"), 64)
		prometheusHistogram.Observe(v)

		prometheusHistogramVec.WithLabelValues("testing").Observe(v)

		w.Write([]byte("Success!"))

		return
	}

	w.Write([]byte("Not GET method"))
}

func gaugeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		v, _ := strconv.ParseFloat(r.Header.Get("value"), 64)
		prometheusGauge.Add(v)

		w.Write([]byte("Success!"))

		return
	}

	w.Write([]byte("Not GET method"))
}

func summaryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		v, _ := strconv.ParseFloat(r.Header.Get("value"), 64)
		prometheusSummary.Observe(v)

		w.Write([]byte("Success!"))

		return
	}

	w.Write([]byte("Not GET method"))
}
