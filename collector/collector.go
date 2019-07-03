package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Exporter namespace.
	namespace = "netapp"
	// Math constant for picoseconds to seconds.
	picoSeconds = 1e12
)


func newDesc(subsystem, name, help string) *prometheus.Desc {
	return prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystem, name),
		help, nil, nil,
	)
}
