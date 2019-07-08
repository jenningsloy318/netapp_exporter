package collector

import (
	"log"
	"github.com/pepabo/go-netapp/netapp"
	"github.com/prometheus/client_golang/prometheus"
)





const (
	// Subsystem.
	AggrSubsystem = "aggr"
	
)

// Metric descriptors.
var (
	aggrLabels = []string{"filer","aggr","cluster","node"}
	aggrSizeUsedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, AggrSubsystem, "size_used"),
		"Used size of aggr.",
		aggrLabels, nil)
	aggrSizeTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, AggrSubsystem, "size_total"),
		"Total size of aggr.",
		aggrLabels, nil)
  aggrSizeAvailableDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, AggrSubsystem, "size_available"),
		"Available size of aggr.",
		aggrLabels, nil)			
	aggrTotalReservedSpaceDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, AggrSubsystem, "total_reserved_space"),
		"Total Reserved Space  of aggr.",
		aggrLabels, nil)		
	aggrPercentUsedCapacityDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, AggrSubsystem, "percent_used_capacity"),
		"Percent Used Capacity of aggr.",
		aggrLabels, nil)							
	aggrPhysicalUsedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, AggrSubsystem, "physical_used"),
		"Physical Used size of aggr.",
		aggrLabels, nil)		
	aggrPhysicalUsedPercentDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, AggrSubsystem, "physical_used_percent"),
		"Physical Used Percent of aggr.",
		aggrLabels, nil)
	aggrSnapSizeTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, AggrSubsystem, "snap_size_total"),
		"Snap Size Total of aggr.",
		aggrLabels, nil)					
	)
// Scrapesystem collects system node info
type ScrapeAggr struct{}

// Name of the Scraper. Should be unique.
func (ScrapeAggr) Name() string {
	return AggrSubsystem
}

// Help describes the role of the Scraper.
func (ScrapeAggr) Help() string {
	return "Collect Netapp aggr info;"
}


type Aggregate struct {
	Name                string
	OwnerName           string
	Cluster            string
	SizeUsed            int
	SizeTotal           int
	SizeAvailable       int
	TotalReservedSpace  int
	PercentUsedCapacity string
	PhysicalUsed        int
	PhysicalUsedPercent int
	SnapSizeTotal       string

}


type AggregateSpace struct {
	Aggregate              string
	SnapSizeTotal          string
	PercentSnapshotSpace   string
	PhysicalUsed           string
	PhysicalUsedPercent    string

}
// Scrape collects data from  netapp aggregate info 
func (ScrapeAggr) Scrape(netappClient *netapp.Client, ch chan<- prometheus.Metric) error {

	for _, AggrInfo := range GetAggrData(netappClient) {
		aggrLabelValues :=[]string{FilerLabelValue,AggrInfo.Name, AggrInfo.Cluster,AggrInfo.OwnerName}
		ch <- prometheus.MustNewConstMetric(aggrSizeUsedDesc, prometheus.GaugeValue,float64(AggrInfo.SizeUsed),aggrLabelValues...)
		ch <- prometheus.MustNewConstMetric(aggrSizeTotalDesc, prometheus.GaugeValue,float64(AggrInfo.SizeTotal), aggrLabelValues...)
		ch <- prometheus.MustNewConstMetric(aggrSizeAvailableDesc, prometheus.GaugeValue,float64(AggrInfo.SizeAvailable), aggrLabelValues...)
		ch <- prometheus.MustNewConstMetric(aggrTotalReservedSpaceDesc, prometheus.GaugeValue,float64(AggrInfo.TotalReservedSpace), aggrLabelValues...)
		
		if PercentUsedCapacity,ok := parseStatus(AggrInfo.PercentUsedCapacity);ok{
			ch <- prometheus.MustNewConstMetric(aggrPercentUsedCapacityDesc, prometheus.GaugeValue,PercentUsedCapacity, aggrLabelValues...)
		 }
		ch <- prometheus.MustNewConstMetric(aggrPhysicalUsedDesc, prometheus.GaugeValue,float64(AggrInfo.PhysicalUsed), aggrLabelValues...)
		ch <- prometheus.MustNewConstMetric(aggrPhysicalUsedPercentDesc, prometheus.GaugeValue,float64(AggrInfo.PhysicalUsedPercent), aggrLabelValues...)			
		if SnapSizeTotal,ok := parseStatus(AggrInfo.SnapSizeTotal);ok{
				ch <- prometheus.MustNewConstMetric(aggrSnapSizeTotalDesc, prometheus.GaugeValue,SnapSizeTotal,  aggrLabelValues...)
  	}
	}
//		for _, AggrSpaceInfo := range GetAggrSpaceData(netappClient) {
//			if SnapSizeTotal,ok := parseStatus(AggrSpaceInfo.SnapSizeTotal);ok{
//
//				ch <- prometheus.MustNewConstMetric(aggrSnapSizeTotalDesc, prometheus.GaugeValue,SnapSizeTotal, FilerLabelValue,AggrSpaceInfo.Aggregate)
//			}
//		}

	
	return nil
}


func GetAggrData(netappClient *netapp.Client) (r []*Aggregate) {
	ff := new(bool)
	*ff = false

	opts := &netapp.AggrOptions {
		Query: &netapp.AggrInfo{
			AggrRaidAttributes: &netapp.AggrRaidAttributes{
				IsRootAggregate: ff,
			},
		},
		DesiredAttributes: &netapp.AggrInfo{
			AggrOwnershipAttributes: &netapp.AggrOwnershipAttributes{
				OwnerName           :"x",
				Cluster            :"x",
			
			},
			AggrSpaceAttributes:     &netapp.AggrSpaceAttributes{
				SizeUsed            : 1,
				SizeTotal           : 1,
				SizeAvailable       : 1,
				TotalReservedSpace  : 1,
				PercentUsedCapacity : "x",
				PhysicalUsed        : 1,
				PhysicalUsedPercent : 1,
			},
		},
	}

	l := getAggrList(netappClient,opts)

	for _, n := range l {
		SnapSizeTotalValue := GetAggrSpaceData(netappClient,n.AggregateName)["SnapSizeTotal"]
		r = append(r, &Aggregate{
			Name:                n.AggregateName,
			OwnerName:           n.AggrOwnershipAttributes.OwnerName,
			Cluster:						 n.AggrOwnershipAttributes.Cluster,
			SizeUsed:            n.AggrSpaceAttributes.SizeUsed,
			SizeTotal:           n.AggrSpaceAttributes.SizeTotal,
			SizeAvailable:       n.AggrSpaceAttributes.SizeAvailable,
			TotalReservedSpace:  n.AggrSpaceAttributes.TotalReservedSpace,
			PercentUsedCapacity: n.AggrSpaceAttributes.PercentUsedCapacity,
			PhysicalUsed:        n.AggrSpaceAttributes.PhysicalUsed,
			PhysicalUsedPercent: n.AggrSpaceAttributes.PhysicalUsedPercent,
			SnapSizeTotal:       SnapSizeTotalValue,
		})
	}
	return
}

func getAggrList(netappClient *netapp.Client,opts *netapp.AggrOptions ) (r []netapp.AggrInfo) {

	var pages []*netapp.AggrListResponse 
	handler := func(r netapp.AggrListPagesResponse) bool {
		if r.Error != nil {
				log.Printf("%s", r.Error)
			return false
		}
		pages = append(pages, r.Response)
		return true
	}

	netappClient.Aggregate.ListPages(opts, handler)

	for _, p := range pages {
		r = append(r, p.Results.AggrAttributes...)
	}

	return
}

var AggrSpaceData = make(map[string]string)

func GetAggrSpaceData(netappClient *netapp.Client, aggrName string) (AggrSpaceData map[string]string) {
	opts := &netapp.AggrSpaceOptions{
		Query:  &netapp.AggrSpaceInfoQuery{
			AggrSpaceInfo: &netapp.AggrSpaceInfo{
				Aggregate:  aggrName, 
			},
		},
		DesiredAttributes: &netapp.AggrSpaceInfoQuery{
			AggrSpaceInfo: &netapp.AggrSpaceInfo{
				SnapSizeTotal : "x",
			},
		},
	}
	l,_,_ := netappClient.AggregateSpace.List(opts)
	AggrSpaceData["SnapSizeTotal"]=l.Results.AttributesList.AggrAttributes[0].SnapSizeTotal
	return
}