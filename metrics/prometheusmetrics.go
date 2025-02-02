package main

import "github.com/prometheus/client_golang/prometheus"

var prometheusCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Namespace:   "study_namespace",
	Name:        "prometheusCounter",
	Help:        "study counter metric",
	ConstLabels: map[string]string{"constant_label": "test_label"},
})

var prometheusCounterVec = prometheus.NewCounterVec(prometheus.CounterOpts{
	Namespace:   "study_namespace",
	Name:        "prometheusCounterVec",
	Help:        "study counter vec metric",
	ConstLabels: map[string]string{"constant_label": "test_vec_label"},
}, []string{"superLabel"})

var prometheusHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
	Namespace: "study_namespace",
	Name:      "prometheusHistogram",
	Help:      "study histogram metric",
	Buckets:   []float64{5, 10, 50, 100, 500, 1000},
})

var prometheusHistogramVec = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Namespace: "study_namespace",
	Name:      "prometheusHistogramVec",
	Help:      "study histogram vec metric",
	Buckets:   []float64{5, 10, 50, 100, 500, 1000},
}, []string{"superLabel"})

var prometheusGauge = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace:   "study_namespace",
	Name:        "prometheusGauge",
	Help:        "study gauge metric",
	ConstLabels: map[string]string{"gauge_label": "test_label"},
})

var prometheusSummary = prometheus.NewSummary(prometheus.SummaryOpts{
	Namespace:   "study_namespace",
	Name:        "prometheusSummary",
	Help:        "study summary metric",
	ConstLabels: map[string]string{"summary_label": "test_label"},
	Objectives:  map[float64]float64{0.5: 0.5, 0.9: 0.2, 0.99: 0.1},
})

func registerPrometheus() {
	prometheus.MustRegister(prometheusCounter)
	prometheus.MustRegister(prometheusCounterVec)
	prometheus.MustRegister(prometheusHistogram)
	prometheus.MustRegister(prometheusHistogramVec)
	prometheus.MustRegister(prometheusGauge)
	prometheus.MustRegister(prometheusSummary)
}
