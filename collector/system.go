package collector

import (
	"log"
	"github.com/pepabo/go-netapp/netapp"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Subsystem.
	SystemSubsystem = "system"
)

// Metric descriptors.
var (
	systemLabels         = append(BaseLabelNames, "node", "location")
	systemNodeUptimeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, SystemSubsystem, "uptime"),
		"uptime of the node.",
		systemLabels, nil)
	systemNodeFailedFanCountDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, SystemSubsystem, "failed_fan_count"),
		"Failed Fan Count of the node.",
		systemLabels, nil)
	systemNodeFailedPowerSupplyCountDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, SystemSubsystem, "failed_powersupply_count"),
		"Failed PowerSupply Count of the node.",
		systemLabels, nil)
	systemNodeOverTemperatureDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, SystemSubsystem, "over_temperature"),
		"Over Temperature of the node.",
		systemLabels, nil)
)

// Scrapesystem collects system node info
type ScrapeSystem struct{}

// Name of the Scraper. Should be unique.
func (ScrapeSystem) Name() string {
	return SystemSubsystem
}

// Help describes the role of the Scraper.
func (ScrapeSystem) Help() string {
	return "Collect Netapp System and Node info;"
}

type Node struct {
	Name                      string
	OwnerName                 string
	Model                     string
	Location                  string
	Uuid                      string
	Uptime                    string
	EnvFailedFanCount         int
	EnvFailedPowerSupplyCount int
	EnvOverTemperature        bool
}

// Scrape collects data from  netapp system and node info
func (ScrapeSystem) Scrape(netappClient *netapp.Client, ch chan<- prometheus.Metric) error {

	for _, NodeInfo := range GetNodeData(netappClient) {
		systemLabelValues := append(BaseLabelValues, NodeInfo.Name, NodeInfo.Location)
		if uptime, ok := parseStatus(NodeInfo.Uptime); ok {
			ch <- prometheus.MustNewConstMetric(systemNodeUptimeDesc, prometheus.GaugeValue, uptime, systemLabelValues...)
		}
		ch <- prometheus.MustNewConstMetric(systemNodeFailedFanCountDesc, prometheus.GaugeValue, float64(NodeInfo.EnvFailedFanCount), systemLabelValues...)
		ch <- prometheus.MustNewConstMetric(systemNodeFailedPowerSupplyCountDesc, prometheus.GaugeValue, float64(NodeInfo.EnvFailedPowerSupplyCount), systemLabelValues...)
		ch <- prometheus.MustNewConstMetric(systemNodeOverTemperatureDesc, prometheus.GaugeValue, boolToFloat64(NodeInfo.EnvOverTemperature), systemLabelValues...)

	}
	return nil
}

func GetNodeData(netappClient *netapp.Client) (r []*Node) {
	opts := &netapp.NodeDetailOptions{
		Query: &netapp.NodeDetailsQuery{},
		DesiredAttributes: &netapp.NodeDetailsQuery{
			NodeDetails: &netapp.NodeDetails{
				Name:                      "x",
				NodeOwner:                 "x",
				NodeModel:                 "x",
				NodeLocation:              "x",
				NodeUuid:                  "x",
				NodeUptime:                "x",
				EnvFailedFanCount:         1,
				EnvFailedPowerSupplyCount: 1,
				EnvOverTemperature:        false,
			},
		},
	}

	l := getNodeList(netappClient, opts)

	for _, n := range l {
		r = append(r, &Node{
			Name:                      n.Name,
			OwnerName:                 n.NodeOwner,
			Model:                     n.NodeModel,
			Location:                  n.NodeLocation,
			Uuid:                      n.NodeUuid,
			Uptime:                    n.NodeUptime,
			EnvFailedFanCount:         n.EnvFailedFanCount,
			EnvFailedPowerSupplyCount: n.EnvFailedPowerSupplyCount,
			EnvOverTemperature:        n.EnvOverTemperature,
		})
	}
	return
}

func getNodeList(netappClient *netapp.Client, opts *netapp.NodeDetailOptions) (r []netapp.NodeDetails) {

	var pages []*netapp.NodeDetailsResponse
	handler := func(r netapp.NodeDetailsPagesResponse) bool {
		if r.Error != nil {
			log.Printf("%s", r.Error)
			return false
		}
		pages = append(pages, r.Response)
		return true
	}

	netappClient.System.ListPages(opts, handler)

	for _, p := range pages {
		r = append(r, p.Results.NodeDetails...)
	}

	return
}
