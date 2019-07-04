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
	VServerInfoDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VserverSubsystem, "volume_delete_retention_hours"),
		"uptime of the vserver.",
		[]string{"vserver","vserver_subtype"}, nil)
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
	VserverName                   string
	VolumeDeleteRetentionHours    int
	VserverSubtype                string

}


// Scrape collects data from  netapp system and vserver info 
func (ScrapeVserver) Scrape(netappClient *netapp.Client, ch chan<- prometheus.Metric) error {

	for _, VserverInfo := range GetVserverData(netappClient) {
		
			ch <- prometheus.MustNewConstMetric(VServerInfoDesc, prometheus.GaugeValue,float64(VserverInfo.VolumeDeleteRetentionHours), VserverInfo.VserverName,VserverInfo.VserverSubtype)
	 
	}
	return nil
}





func GetVserverData(netappClient *netapp.Client) (r []*VServer) {
	opts := &netapp.VServerOptions  {
		Query: &netapp.VServerQuery {
			VServerInfo  : &netapp.VServerInfo{},
		},
		DesiredAttributes: &netapp.VServerQuery {
			VServerInfo : &netapp.VServerInfo{},
		},
	}
	l,_,_ := netappClient.VServer.List(opts)
	for _, n := range l.Results.AttributesList.VserverInfo {
		r = append(r, &VServer{
			VserverName:                n.VserverName,
			VolumeDeleteRetentionHours: n.VolumeDeleteRetentionHours,
			VserverSubtype: 								n.VserverSubtype,
		})
	}
	return
}
