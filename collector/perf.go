package collector

import (
	"log"
	"github.com/pepabo/go-netapp/netapp"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Subsystem.
	PerfSubsystem = "perf"
)

// Metric descriptors.
var (
	PerfinfoDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, PerfSubsystem, "size"),
		"Size of the perf.",
		[]string{"name", }, nil)
	
	)


// Scrapesystem collects system Perf info
type ScrapePerf struct{}

// Name of the Scraper. Should be unique.
func (ScrapePerf) Name() string {
	return PerfSubsystem
}

// Help describes the role of the Scraper.
func (ScrapePerf) Help() string {
	return "Collect Netapp Perf info;"
}


type Perf struct {
	Name                     string

}









// Scrape collects data from  netapp system and Perf info 
func (ScrapePerf) Scrape(netappClient *netapp.Client, ch chan<- prometheus.Metric) error {
 	  for _,instanceData := range GetPerfObjectInstanceInfo(netappClient) {
				for _,perfCounterData := range instanceData.Counters.CounterData{
			 		if countersValue,ok := parseStatus(perfCounterData.Value);ok {
						ch <- prometheus.MustNewConstMetric(PerfinfoDesc, prometheus.GaugeValue,countersValue, perfCounterData.Name)
		 			}
 				}
		}
	
 	return nil
}

func GetPerfObjectInstanceInfo(netappClient *netapp.Client) (r []netapp.InstanceData){

		var perfInstanceUuids []string 
		perfInstanceList := GetPerfObjectInstanceList(netappClient)

		for _, perfInstance :=range perfInstanceList {
			perfInstanceUuids=append(perfInstanceUuids,perfInstance.Uuid)
		}

		type newPerfInstanceUuid struct {
			Uuids []string `xml:"instance-uuid"` 

		}

	var newPerfInstanceUuids newPerfInstanceUuid 
	newPerfInstanceUuids.Uuids = perfInstanceUuids

	ops := &netapp.PerfObjectGetInstanceParams{
		InstanceUuids : newPerfInstanceUuids,
		ObjectName: "volume",
	}

	resp,_,err :=netappClient.Perf.PerfObjectGetInstances(ops)

	if err != nil {
		log.Printf("%s", err)
	}
 r = resp.Results.PerfObjectInstanceData.Instances
	return 
}

func GetPerfObjectInstanceList(netappClient *netapp.Client) (r []netapp.InstanceInfo) {
	

 ops := &netapp.PerfObjectInstanceListInfoIterParams{
	Query: &netapp.InstanceInfoQuery{	},
	DesiredAttributes: &netapp.InstanceInfo{},
	ObjectName: "volume",
 }
	resp,_,_ := netappClient.Perf.PerfObjectInstanceListInfoIter(ops)
	
  r=resp.Results.AttributesList.InstanceInfo
  return 
}
