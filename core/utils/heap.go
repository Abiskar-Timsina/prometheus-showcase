package utils

import "runtime/metrics"

func GetHeapBytes() uint64 {
	runtimeMetrics := make([]metrics.Sample, 1)
	runtimeMetrics[0].Name = "/gc/heap/allocs:bytes"
	metrics.Read(runtimeMetrics)
	return runtimeMetrics[0].Value.Uint64()
}
