package allocsample

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metric struct {
	name       string
	help       string
	value      float64
	metricType prometheus.ValueType
}

func uniqueDeclaration() {
	var m *metric
	for range [4 * 60]int{} {
		m = nil
		m = new(metric)
		m.name = "Dummy op"
	}
}

func declarationInsideLoop() {
	for range [4 * 60]int{} {
		var m *metric
		m = new(metric)
		m.name = "Dummy op"
	}
}
