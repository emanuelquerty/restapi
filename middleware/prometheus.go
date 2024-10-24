package middleware

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Generic metrics
var (
	totalRequestCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_count",
			Help: "Total number of http requests per endpoint",
		},
		[]string{"endpoint"},
	)

	requestDurationObserver = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_mili_seconds",
			Help:    "Duration of http requests per endpoint",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"endpoint"},
	)
)
