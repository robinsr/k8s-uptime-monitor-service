// metrics.go
// This file defines Prometheus metrics for the uptime monitoring service.

package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	targetUp = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "uptime_monitor_target_up",
			Help: "Indicates if the target is up (1) or down (0)",
		},
		[]string{"target"},
	)

	responseDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "uptime_monitor_response_duration_seconds",
			Help:    "Duration of HTTP responses in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"target"},
	)

	targetFailures = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "uptime_monitor_target_failures_total",
			Help: "Total number of failures for each target",
		},
		[]string{"target"},
	)
)

// RecordSuccess creates metric datapoints on target UP
func RecordSuccess(target string, duration float64) {
	targetUp.WithLabelValues(target).Set(1)
	responseDuration.WithLabelValues(target).Observe(duration)
}

// RecordFailed creates metric datapoints on target DOWN
func RecordFailed(target string) {
	targetUp.WithLabelValues(target).Set(0)
	targetFailures.WithLabelValues(target).Inc()
}
