package metrics

import "github.com/prometheus/client_golang/prometheus"

var Endpoint1Counter prometheus.Counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "Endpoint-1 Counter",
		Help: "The counter to store the total number of requests made to the endpoint-1",
	},
)

var Endpoint2Counter prometheus.Counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "Endpoint-2 Counter",
		Help: "The counter to store the total number of requests made to the endpoint-2",
	},
)

var HeapMemory prometheus.Gauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "CPU Usage",
		Help: "Used to calculate the CPU usuage at any time.",
	},
)
