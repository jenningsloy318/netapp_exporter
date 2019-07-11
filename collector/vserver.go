package collector

import (
	//	"log"
	"github.com/pepabo/go-netapp/netapp"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Subsystem.
	VserverSubsystem = "vserver"
)

// Metric descriptors.
var (
	vserverLabels                         = append(BaseLabelNames, "vserver", "type")
	VServerVolumeDeleteRetentionHoursDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VserverSubsystem, "volume_delete_retention_hours"),
		"Volume Delete Retention Hours of the vserver.",
		vserverLabels, nil)
	VServerAdminStateDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VserverSubsystem, "state"),
		"Admin State of the vserver,1(running), 0(stopped), 2(starting),3(stopping), 4(initializing), or 5(deleting).",
		vserverLabels, nil)
	VServerOperationalStateDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VserverSubsystem, "operational_state"),
		"Operational State of the vserver, 1(running), 0(stopped).",
		vserverLabels, nil)
)

// Scrapesystem collects system vserver info
type ScrapeVserver struct{}

// Name of the Scraper. Should be unique.
func (ScrapeVserver) Name() string {
	return VserverSubsystem
}

// Help describes the role of the Scraper.
func (ScrapeVserver) Help() string {
	return "Collect Netapp Vserver info;"
}

type VServer struct {
	VserverName                string
	VserverType                string
	VolumeDeleteRetentionHours int
	State                      string
	OperationalState           string
}

// Scrape collects data from  netapp system and vserver info
func (ScrapeVserver) Scrape(netappClient *netapp.Client, ch chan<- prometheus.Metric) error {

	for _, VserverInfo := range GetVserverData(netappClient) {
		vserverLabelValues := append(BaseLabelValues, VserverInfo.VserverName, VserverInfo.VserverType)
		ch <- prometheus.MustNewConstMetric(VServerVolumeDeleteRetentionHoursDesc, prometheus.GaugeValue, float64(VserverInfo.VolumeDeleteRetentionHours), vserverLabelValues...)
		if len(VserverInfo.State) > 0 {
			if stateVal, ok := parseStatus(VserverInfo.State); ok {
				ch <- prometheus.MustNewConstMetric(VServerAdminStateDesc, prometheus.GaugeValue, stateVal, vserverLabelValues...)
			}
		}
		if len(VserverInfo.OperationalState) > 0 {
			if opsStateVal, ok := parseStatus(VserverInfo.OperationalState); ok {
				ch <- prometheus.MustNewConstMetric(VServerOperationalStateDesc, prometheus.GaugeValue, opsStateVal, vserverLabelValues...)
			}
		}

	}
	return nil
}

func GetVserverData(netappClient *netapp.Client) (r []*VServer) {
	opts := &netapp.VServerOptions{
		Query: &netapp.VServerQuery{},
		DesiredAttributes: &netapp.VServerQuery{
			VServerInfo: &netapp.VServerInfo{
				VserverName:                "x",
				VserverType:                "x",
				VolumeDeleteRetentionHours: 1,
				State:                      "x",
				OperationalState:           "x",
			},
		},
	}
	l, _, _ := netappClient.VServer.List(opts)
	for _, n := range l.Results.AttributesList.VserverInfo {
		r = append(r, &VServer{
			VserverName:                n.VserverName,
			VserverType:                n.VserverType,
			VolumeDeleteRetentionHours: n.VolumeDeleteRetentionHours,
			State:                      n.State,
			OperationalState:           n.OperationalState,
		})
	}
	return
}
