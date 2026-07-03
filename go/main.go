package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Response structure for /demo-service endpoint
type Response struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

// Prometheus metrics
var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
	serviceAvailability = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "service_availability",
			Help: "Service availability (1 = available, 0 = unavailable)",
		},
	)
)

func init() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(requestDuration)
	prometheus.MustRegister(serviceAvailability)
	serviceAvailability.Set(1)
}

func handleDemoService(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	response := Response{
		Nome:    "Observability Demonstration",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	duration := time.Since(start).Seconds()
	requestCounter.WithLabelValues("GET", "/demo-service", "200").Inc()
	requestDuration.WithLabelValues("GET", "/demo-service").Observe(duration)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	requestCounter.WithLabelValues("GET", "/health", "200").Inc()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/demo-service", handleDemoService)
	mux.HandleFunc("/health", handleHealth)
	mux.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Starting http-demo-service on :8080")
	log.Fatal(server.ListenAndServe())
}