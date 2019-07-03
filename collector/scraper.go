package collector

import (

	"github.com/prometheus/client_golang/prometheus"
	"github.com/pepabo/go-netapp/netapp"
)

// Scraper is minimal interface that let's you add new prometheus metrics to mysqld_exporter.
type Scraper interface {
	// Name of the Scraper. Should be unique.
	Name() string
	// Help describes the role of the Scraper.
	// Example: "Collect  node metrics"
	Help() string
	// Scrape collects data from netappClient connection.
	Scrape(netappClient *netapp.Client, ch chan<- prometheus.Metric) error
}
