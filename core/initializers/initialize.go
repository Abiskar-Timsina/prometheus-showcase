package initializers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Abiskar-Timsina/prometheus-showcase/core/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// register the prometheus metrics
func init() {
	prometheus.MustRegister(metrics.Endpoint1Counter)
	prometheus.MustRegister(metrics.Endpoint2Counter)
	prometheus.MustRegister(metrics.HeapMemory)
}

func InitializeServer(port int64) {
	portString := strconv.FormatInt(port, 10)
	// Create a server multiplexer to have use multiple handlers.
	multiplexServer := http.NewServeMux()

	// Let the default /metrics be used by the prometheus handler.
	multiplexServer.Handle("/metrics", promhttp.Handler())

	// Let everything else be handled by the custom handler.
	multiplexServer.Handle("/", Handler{})

	log.Fatal(http.ListenAndServe("0.0.0.0:"+portString, multiplexServer))
}
