package collector

import (
	"log"
	"github.com/pepabo/go-netapp/netapp"
	"github.com/prometheus/client_golang/prometheus"
)





const (
	// Subsystem.
	SnapshotSubsystem = "snapshot"
)

// Metric descriptors.
var (
	snapshotLabels = []string{"filer","snapshot","volume","vserver"}
	snapshotTotalSizeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, SnapshotSubsystem, "total_size"),
		"Size of the snapshot.",
		snapshotLabels, nil)
	snapshotAdminStateDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, SnapshotSubsystem, "state"),
		"The state of  the snapshot.",
		snapshotLabels, nil)
	snapshotBusyDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, SnapshotSubsystem, "is_busy"),
		"The busy state of  the snapshot.",
		snapshotLabels, nil)

)


// Scrapesystem collects system node info
type ScrapeSnapshot struct{}

// Name of the Scraper. Should be unique.
func (ScrapeSnapshot) Name() string {
	return SnapshotSubsystem
}

// Help describes the role of the Scraper.
func (ScrapeSnapshot) Help() string {
	return "Collect Netapp Snapshot info;"
}


type Snapshot struct {
	Name                              string
	Busy                              bool
	State                             string  
	Total                             int     
	Volume                            string  
	Vserver                           string  

}


// Scrape collects data from  netapp system and node info 
func (ScrapeSnapshot) Scrape(netappClient *netapp.Client, ch chan<- prometheus.Metric) error {

	for _, SnapshotInfo := range GetSnapshotData(netappClient) {
		snapshotLabelValues := []string{FilerLabelValue,SnapshotInfo.Name, SnapshotInfo.Volume, SnapshotInfo.Vserver}
		ch <- prometheus.MustNewConstMetric(snapshotTotalSizeDesc, prometheus.GaugeValue,float64(SnapshotInfo.Total), snapshotLabelValues...)
		ch <- prometheus.MustNewConstMetric(snapshotBusyDesc, prometheus.GaugeValue,boolToFloat64(SnapshotInfo.Busy), snapshotLabelValues...)
		if value,ok :=parseStatus(SnapshotInfo.State);ok {
			ch <- prometheus.MustNewConstMetric(snapshotAdminStateDesc, prometheus.GaugeValue,value, snapshotLabelValues...)
		}	 
	}
	return nil
}





func GetSnapshotData(netappClient *netapp.Client) (r []*Snapshot) {
	opts := &netapp.SnapshotOptions {
		Query: &netapp.SnapshotQuery{},
		DesiredAttributes: &netapp.SnapshotQuery{
			SnapshotInfo : &netapp.SnapshotInfo{
				Name                      :"x",
				Volume                    :"x",   
				Vserver                   :"x",
				Busy                      :true,
				State                     :"x", 
				Total                     :1,
			},
		},
	}

	l := getSnapshotList(netappClient,opts)

	for _, n := range l {
		r = append(r, &Snapshot{
			Name                      :n.Name,
			Volume                    :n.Volume,
			Vserver                   :n.Vserver,
			Busy                      :n.Busy,
			State                     :n.State,
			Total                     :n.Total,
			})
	}
	return
}

func getSnapshotList(netappClient *netapp.Client, opts *netapp.SnapshotOptions) (r []netapp.SnapshotInfo) {

	var pages []*netapp.SnapshotListResponse 
	handler := func(r netapp.SnapshotListPagesResponse) bool {
		if r.Error != nil {
				log.Printf("%s", r.Error)
			return false
		}
		pages = append(pages, r.Response)
		return true
	}

	netappClient.Snapshot.ListPages(opts, handler)

	for _, p := range pages {
		r = append(r, p.Results.AttributesList.SnapshotAttributes...)
	}

	return
}