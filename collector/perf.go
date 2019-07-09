 package collector

import (
	"log"
	"github.com/pepabo/go-netapp/netapp"
	"github.com/prometheus/client_golang/prometheus"
	"strings"
	"fmt"
)

const (
	// Subsystem.
	PerfSubsystem = "perf"
)

// Metric descriptors.
var (
	objects=  []string{"system","system:node","nfsv3","lif","lun","aggregate","disk","workload","processor","processor:node","volume:node","volume:vserver"}
  // we can get all objects via connect to netapp,  set advanced, then issue command statistics catalog object show
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



// Scrape collects data from  netapp system and Perf info 
func (ScrapePerf) Scrape(netappClient *netapp.Client, ch chan<- prometheus.Metric) error {
		for _,object := range objects {
			for _,perfInstanceData := range GetPerfObjectInstanceInfo(netappClient,object) {
				
					var labelName  = []string{"filer"}
					var labelValue  = []string{FilerLabelValue,}

					var metricNamePrefix string 
					var metricNameSuffix string 
				  if strings.Contains(perfInstanceData.Name,"/") {
						instanceNameSlice := strings.Split(perfInstanceData.Name,"/")
						labelValue=append(labelValue,instanceNameSlice[len(instanceNameSlice)-1])
					} else {
						labelValue=append(labelValue,perfInstanceData.Name)
					}

					if strings.Contains(object,":") {
						objectSlice :=strings.Split(object,":")
						labelName = append(labelName,objectSlice[1])
						metricNamePrefix =fmt.Sprintf("%s_",objectSlice[0])
						metricNameSuffix = fmt.Sprintf("_per_%s",objectSlice[1])

					}else{
						labelName = append(labelName,object)
						metricNamePrefix =fmt.Sprintf("%s_",object)
						metricNameSuffix=""
					}
					
					var metricMap = make(map[string]float64)
					perfCounterDataSlice := perfInstanceData.Counters.CounterData  // CounterData is slice which contains all conter-data for one instance
					for _,perfCounterData := range perfCounterDataSlice{
							if  ( perfCounterData.Name =="node_name" ||  perfCounterData.Name == "vserver_name" || perfCounterData.Name == "cpu_name"  ){
									labelName=append(labelName,strings.Split(perfCounterData.Name,"_")[0])
									if perfCounterData.Value == "Multiple_Values" {
											labelValue=append(labelValue,"all")
									}else {
										labelValue=append(labelValue,perfCounterData.Value)
									}
							}else if strings.ContainsAny(perfCounterData.Value,",- ") { // this filter out the histogram data(seperated by ","), value contain space which is actually the  description of the item, value contains "-" which can be a id or uuid, 
									continue 
								}else if  ( strings.Contains(perfCounterData.Name,"_name") ||  strings.Contains(perfCounterData.Name,"_id")  ||  strings.Contains(perfCounterData.Name,"_uuid") ||  strings.ContainsAny(perfCounterData.Value,",- ") || len(perfCounterData.Value) == 0 ){ // this filter out the histogram data(seperated by ","), also filter out non metric items such instance name/uuid/id 
									continue 
							}else if len(perfCounterData.Value) == 0 { // this set the value to 0 when the value is empty 
									metricMap[perfCounterData.Name]=0
								}else{
										if value,ok := parseStatus(perfCounterData.Value);ok {
										metricMap[perfCounterData.Name]=value
								}
							}
					}

					for metricName,metricValue := range  metricMap{
						//metricName := metricNamePrefix+metricName+metricNameSuffix
						metricName := fmt.Sprintf("%s%s%s",metricNamePrefix,metricName,metricNameSuffix)
						metricHelp := fmt.Sprintf("Perf %s %s",labelName[0],metricName)
						desc := prometheus.NewDesc(
							prometheus.BuildFQName(namespace, PerfSubsystem, metricName),
							metricHelp,
						  labelName, nil)
							ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue,metricValue,labelValue...)
					}
		}
		
}
return nil		
}
		




func GetPerfObjectInstanceInfo(netappClient *netapp.Client, objectName string ) (r []netapp.InstanceData){

		var perfInstanceUuids []string 
		perfInstanceList := GetPerfObjectInstanceList(netappClient,objectName)

		for _, perfInstance :=range perfInstanceList {
			perfInstanceUuids=append(perfInstanceUuids,perfInstance.Uuid)
		}

		type newPerfInstanceUuid struct {
			Uuids []string `xml:"instance-uuid"` 

		}

	var newPerfInstanceUuids newPerfInstanceUuid 
	newPerfInstanceUuids.Uuids = perfInstanceUuids

	opts := &netapp.PerfObjectGetInstanceParams{
		InstanceUuids : newPerfInstanceUuids,
		ObjectName: objectName,
	}

	resp,_,err :=netappClient.Perf.PerfObjectGetInstances(opts)

	if err != nil {
		log.Printf("%s", err)
	}
	r = resp.Results.PerfObjectInstanceData.Instances // this return a slice of InstanceData, wich contains multiple instances, each instance contains  arbitrary counts of counter-data 
	 
	return
}
func GetPerfObjectInstanceList(netappClient *netapp.Client, objectName string) (r []netapp.InstanceInfo) {
	

	opts := &netapp.PerfObjectInstanceListInfoIterParams{
	Query: &netapp.InstanceInfoQuery{	},
	DesiredAttributes: &netapp.InstanceInfo{},
	ObjectName: objectName,
 	}
	resp,_,_ := netappClient.Perf.PerfObjectInstanceListInfoIter(opts)
	
  r=resp.Results.AttributesList.InstanceInfo // return a slice of instances
  return 
}
