package initializers

import (
	"math/rand"
	"net/http"
	"strings"

	prometheusMetrics "github.com/Abiskar-Timsina/prometheus-showcase/core/metrics"
	"github.com/Abiskar-Timsina/prometheus-showcase/core/utils"
)

type Handler struct {
	totalRequest uint64
}

func (s Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch strings.Trim(r.URL.String(), " /") {
	case "":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1>Welcome to the Monitoring and Logging Demo </h1>"))
		return
	case "endpoint-1":
		prometheusMetrics.HeapMemory.Set(float64(utils.GetHeapBytes()))
		prometheusMetrics.Endpoint1Counter.Inc()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1>Endpoint 1</h1>"))
		return

	case "endpoint-2":
		prometheusMetrics.HeapMemory.Set(float64(utils.GetHeapBytes()))
		prometheusMetrics.Endpoint2Counter.Inc()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1>Endpoint 2</h1>"))
		return

	case "random":
		if rand.Int()%2 == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1>Successful!</h1>"))
		return

	default:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strings.Trim(r.URL.String(), " /")))
		return
	}
}
