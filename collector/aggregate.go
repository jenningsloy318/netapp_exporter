package collector

import (
	"log"
	"github.com/pepabo/go-netapp/netapp"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)





const (
	// Subsystem.
	AggrSubsystem = "aggr"
)

// Metric descriptors.
var (
	aggrSizeUsedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, AggrSubsystem, "size_used"),
		"Used size of aggr.",
		[]string{"Aggr","cluster","Node"}, nil)
	aggrSizeTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, AggrSubsystem, "size_total"),
		"Total size of aggr.",
		[]string{"Aggr","cluster","Node"}, nil)
  aggrSizeAvailableDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, AggrSubsystem, "size_available"),
		"Available size of aggr.",
		[]string{"Aggr","cluster","Node"}, nil)			
	aggrTotalReservedSpaceDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, AggrSubsystem, "total_reserved_space"),
		"Total Reserved Space  of aggr.",
		[]string{"Aggr","cluster","Node"}, nil)		
	aggrPercentUsedCapacityDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, AggrSubsystem, "percent_used_capacity"),
		"Percent Used Capacity of aggr.",
		[]string{"Aggr","cluster","Node"}, nil)							
	aggrPhysicalUsedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, AggrSubsystem, "physical_used"),
		"Physical Used size of aggr.",
		[]string{"Aggr","cluster","Node"}, nil)		
	aggrPhysicalUsedPercentDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, AggrSubsystem, "physical_used_percent"),
		"Physical Used Percent of aggr.",
		[]string{"Aggr","cluster","Node"}, nil)
	aggrSnapSizeTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, AggrSubsystem, "snap_size_total"),
		"Snap Size Total of aggr.",
		[]string{"Aggr"}, nil)					
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
	Cluseter            string
	SizeUsed            int
	SizeTotal           int
	SizeAvailable       int
	TotalReservedSpace  int
	PercentUsedCapacity string
	PhysicalUsed        int
	PhysicalUsedPercent int

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
		ch <- prometheus.MustNewConstMetric(aggrSizeUsedDesc, prometheus.GaugeValue,float64(AggrInfo.SizeUsed), AggrInfo.Name, AggrInfo.Cluseter,AggrInfo.OwnerName)
		ch <- prometheus.MustNewConstMetric(aggrSizeTotalDesc, prometheus.GaugeValue,float64(AggrInfo.SizeTotal), AggrInfo.Name, AggrInfo.Cluseter,AggrInfo.OwnerName)
		ch <- prometheus.MustNewConstMetric(aggrSizeAvailableDesc, prometheus.GaugeValue,float64(AggrInfo.SizeAvailable), AggrInfo.Name, AggrInfo.Cluseter,AggrInfo.OwnerName)
		ch <- prometheus.MustNewConstMetric(aggrTotalReservedSpaceDesc, prometheus.GaugeValue,float64(AggrInfo.TotalReservedSpace), AggrInfo.Name, AggrInfo.Cluseter,AggrInfo.OwnerName)
		PercentUsedCapacity,_ := strconv.ParseFloat(AggrInfo.PercentUsedCapacity, 64);
		ch <- prometheus.MustNewConstMetric(aggrPercentUsedCapacityDesc, prometheus.GaugeValue,PercentUsedCapacity, AggrInfo.Name, AggrInfo.Cluseter,AggrInfo.OwnerName)
		ch <- prometheus.MustNewConstMetric(aggrPhysicalUsedDesc, prometheus.GaugeValue,float64(AggrInfo.PhysicalUsed), AggrInfo.Name, AggrInfo.Cluseter,AggrInfo.OwnerName)
		ch <- prometheus.MustNewConstMetric(aggrPhysicalUsedPercentDesc, prometheus.GaugeValue,float64(AggrInfo.PhysicalUsedPercent), AggrInfo.Name, AggrInfo.Cluseter,AggrInfo.OwnerName)
  	}

		for _, AggrSpaceInfo := range GetAggrSpaceData(netappClient) {
			SnapSizeTotal,_ := strconv.ParseFloat(AggrSpaceInfo.SnapSizeTotal, 64)

		ch <- prometheus.MustNewConstMetric(aggrSnapSizeTotalDesc, prometheus.GaugeValue,SnapSizeTotal, AggrSpaceInfo.Aggregate)
		
		}

	
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
			AggrOwnershipAttributes: &netapp.AggrOwnershipAttributes{},
			AggrSpaceAttributes:     &netapp.AggrSpaceAttributes{},
		},
	}

	l := getAggrList(netappClient,opts)

	for _, n := range l {
		r = append(r, &Aggregate{
			Name:                n.AggregateName,
			OwnerName:           n.AggrOwnershipAttributes.OwnerName,
			Cluseter:						 n.AggrOwnershipAttributes.Cluster,
			SizeUsed:            n.AggrSpaceAttributes.SizeUsed,
			SizeTotal:           n.AggrSpaceAttributes.SizeTotal,
			SizeAvailable:       n.AggrSpaceAttributes.SizeAvailable,
			TotalReservedSpace:  n.AggrSpaceAttributes.TotalReservedSpace,
			PercentUsedCapacity: n.AggrSpaceAttributes.PercentUsedCapacity,
			PhysicalUsed:        n.AggrSpaceAttributes.PhysicalUsed,
			PhysicalUsedPercent: n.AggrSpaceAttributes.PhysicalUsedPercent,
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

func GetAggrSpaceData(netappClient *netapp.Client) (r []*AggregateSpace) {
	opts := &netapp.AggrSpaceOptions  {
		Query: &netapp.AggrSpaceInfoQuery  {
			AggrSpaceInfo   : &netapp.AggrSpaceInfo{},
		},
		DesiredAttributes: &netapp.AggrSpaceInfoQuery  {
			AggrSpaceInfo  : &netapp.AggrSpaceInfo{},
		},
	}
	l,_,_ := netappClient.AggregateSpace.List(opts)
	for _, n := range l.Results.AttributesList.AggrAttributes {
		r = append(r, &AggregateSpace{
			Aggregate:                n.Aggregate,
			SnapSizeTotal:            n.SnapSizeTotal,
			PercentSnapshotSpace: 	  n.PercentSnapshotSpace,
			PhysicalUsed:             n.PhysicalUsed,
			PhysicalUsedPercent:      n.PhysicalUsedPercent,
		})
	}
	return
}