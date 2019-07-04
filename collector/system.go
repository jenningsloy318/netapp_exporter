package collector

import (
	"log"
	"github.com/pepabo/go-netapp/netapp"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)





const (
	// Subsystem.
	SystemSubsystem = "system"
)

// Metric descriptors.
var (
	systemNodeUptimeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, SystemSubsystem, "uptime"),
		"uptime of the node.",
		[]string{"name", "model", "location", "productversion", "uuid"}, nil)
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
	Name                string
	OwnerName           string
	Uptime							string
	Model               string
	Location 						string
	Productversion      string
	Uuid								string
}


// Scrape collects data from  netapp system and node info 
func (ScrapeSystem) Scrape(netappClient *netapp.Client, ch chan<- prometheus.Metric) error {

	for _, NodeInfo := range GetNodeData(netappClient) {
		if uptime,err :=  strconv.ParseFloat(NodeInfo.Uptime, 64); err == nil{
			ch <- prometheus.MustNewConstMetric(systemNodeUptimeDesc, prometheus.GaugeValue,uptime, NodeInfo.Name, NodeInfo.Model, NodeInfo.Location, NodeInfo.Productversion, NodeInfo.Uuid)
	 }
	}
	return nil
}





func GetNodeData(netappClient *netapp.Client) (r []*Node) {
	opts := &netapp.NodeDetailOptions {
		Query: &netapp.NodeDetailsQuery{},
		DesiredAttributes: &netapp.NodeDetailsQuery{
			NodeDetails : &netapp.NodeDetails{},
		},
	}

	l := getNodeList(netappClient,opts)

	for _, n := range l {
		r = append(r, &Node{
			Name:                n.Name,
			OwnerName:           n.NodeOwner,
			Uptime:						   n.NodeUptime,
			Model:               n.NodeModel,
			Location: 					 n.NodeLocation,
			Productversion:      n.ProductVersion,
			Uuid:								 n.NodeUuid,
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