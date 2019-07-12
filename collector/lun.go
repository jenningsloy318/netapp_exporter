package collector

import (
	"log"

	"github.com/pepabo/go-netapp/netapp"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Subsystem.
	LunSubsystem = "lun"
)

// Metric descriptors.
var (
	lunLabels   = append(BaseLabelNames, "volume", "node", "vserver")
	lunSizeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, LunSubsystem, "size"),
		"Size of the lun.",
		lunLabels, nil)
	lunSizeUsedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, LunSubsystem, "size_used"),
		"Size Used of the lun.",
		lunLabels, nil)
	lunStagingStateDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, LunSubsystem, "is_staging"),
		"whether the lun is  staging state.",
		lunLabels, nil)
	lunOnlineStateDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, LunSubsystem, "is_online"),
		"whether the lun is  online state.",
		lunLabels, nil)
	lunAdminStateDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, LunSubsystem, "state"),
		"the state of  the lun.",
		lunLabels, nil)
)

// Scrapesystem collects system node info
type ScrapeLun struct{}

// Name of the Scraper. Should be unique.
func (ScrapeLun) Name() string {
	return LunSubsystem
}

// Help describes the role of the Scraper.
func (ScrapeLun) Help() string {
	return "Collect Netapp Lun info;"
}

type Lun struct {
	Node     string
	Volume   string
	Vserver  string
	Size     int
	SizeUsed int
	Staging  bool
	Online   bool
	State    string
}

// Scrape collects data from  netapp system and node info
func (ScrapeLun) Scrape(netappClient *netapp.Client, ch chan<- prometheus.Metric) error {

	for _, LunInfo := range GetLunData(netappClient) {
		lunLabelValues := append(BaseLabelValues, LunInfo.Volume, LunInfo.Node, LunInfo.Vserver)
		ch <- prometheus.MustNewConstMetric(lunSizeDesc, prometheus.GaugeValue, float64(LunInfo.Size), lunLabelValues...)
		ch <- prometheus.MustNewConstMetric(lunSizeUsedDesc, prometheus.GaugeValue, float64(LunInfo.SizeUsed), lunLabelValues...)
		ch <- prometheus.MustNewConstMetric(lunStagingStateDesc, prometheus.GaugeValue, boolToFloat64(LunInfo.Staging), lunLabelValues...)
		ch <- prometheus.MustNewConstMetric(lunOnlineStateDesc, prometheus.GaugeValue, boolToFloat64(LunInfo.Online), lunLabelValues...)
		if value, ok := parseStatus(LunInfo.State); ok {
			ch <- prometheus.MustNewConstMetric(lunAdminStateDesc, prometheus.GaugeValue, value, lunLabelValues...)
		}
	}
	return nil
}

func GetLunData(netappClient *netapp.Client) (r []*Lun) {
	opts := &netapp.LunOptions{
		Query: &netapp.LunQuery{},
		DesiredAttributes: &netapp.LunQuery{
			LunInfo: &netapp.LunInfo{
				Node:     "x",
				Volume:   "x",
				Vserver:  "x",
				Size:     1,
				SizeUsed: 1,
				Staging:  false,
				Online:   true,
				State:    "x",
			},
		},
	}

	l := getLunList(netappClient, opts)

	for _, n := range l {
		r = append(r, &Lun{
			Node:     n.Node,
			Volume:   n.Volume,
			Vserver:  n.Vserver,
			Size:     n.Size,
			SizeUsed: n.SizeUsed,
			Staging:  n.Staging,
			Online:   n.Online,
			State:    n.State,
		})
	}
	return
}

func getLunList(netappClient *netapp.Client, opts *netapp.LunOptions) (r []netapp.LunInfo) {

	var pages []*netapp.LunListResponse
	handler := func(r netapp.LunListPagesResponse) bool {
		if r.Error != nil {
			log.Printf("%s", r.Error)
			return false
		}
		pages = append(pages, r.Response)
		return true
	}

	netappClient.Lun.ListPages(opts, handler)

	for _, p := range pages {
		r = append(r, p.Results.AttributesList.LunAttributes...)
	}

	return
}
