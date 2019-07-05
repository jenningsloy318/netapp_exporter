package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"bytes"
	"strconv"
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



func parseStatus(data string) (float64, bool) {

	// vserver state
	if bytes.Equal([]byte(data), []byte("running")) {
		return 1, true
	}
	if bytes.Equal([]byte(data), []byte("stopped")) {
		return 0, true
	}
	if bytes.Equal([]byte(data), []byte("starting")) {
		return 2, true
	}
	if bytes.Equal([]byte(data), []byte("stopping")) {
		return 3, true
	}
	if bytes.Equal([]byte(data), []byte("initializing")) {
		return 4, true
	}
	if bytes.Equal([]byte(data), []byte("deleting")) {
		return 5, true
	}
//volume state
	if bytes.Equal([]byte(data), []byte("online")) {
		return 1, true
	}
	if bytes.Equal([]byte(data), []byte("offline")) {
		return 0, true
	}
	if bytes.Equal([]byte(data), []byte("restricted")) {
		return 2, true
	}
	if bytes.Equal([]byte(data), []byte("mixed")) {
		return 3, true
	}

	value, err := strconv.ParseFloat(string(data), 64)
	return value, err == nil
}
