 package collector

import (
	"log"
	"github.com/pepabo/go-netapp/netapp"
	"github.com/prometheus/client_golang/prometheus"
	"strings"
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
		for _,obj := range objects {
			for _,instanceData := range GetPerfObjectInstanceInfo(netappClient,obj) {
				
					var labelName   []string
					var labelValue  []string

					var metricNamePrefix string 
					var metricNameSufix string 
				  if strings.Contains(instanceData.Name,"/") {
						instanceNameSlice := strings.Split(instanceData.Name,"/")
						labelValue=append(labelValue,instanceNameSlice[len(instanceNameSlice)-1])
					} else {
						labelValue=append(labelValue,instanceData.Name)
					}

					if strings.Contains(obj,":") {
						objSlice :=strings.Split(obj,":")
						labelName = append(labelName,objSlice[1])
						metricNamePrefix =  objSlice[0]+"_"
						metricNameSufix = "_per_"+objSlice[1]
					}else{
						labelName = append(labelName,obj)
						metricNamePrefix = obj+"_"
						metricNameSufix=""
					}
					
					var metricMap = make(map[string]float64)

					for _,perfCounterData := range instanceData.Counters.CounterData{
							if  ( perfCounterData.Name =="node_name" ||  perfCounterData.Name == "vserver_name" || perfCounterData.Name == "cpu_name"  ){
									labelName=append(labelName,strings.Split(perfCounterData.Name,"_")[0])
									if perfCounterData.Value == "Multiple_Values" {
											labelValue=append(labelValue,"all")
									}else {
										labelValue=append(labelValue,perfCounterData.Value)
									}
							}else if  ( strings.Contains(perfCounterData.Name,"_name") ||  strings.Contains(perfCounterData.Name,"_id")  ||  strings.Contains(perfCounterData.Name,"_uuid") ||  strings.ContainsAny(perfCounterData.Value,",- ") || len(perfCounterData.Value) == 0 ){ // this filter out the histogram data(seperated by ","), also filter out non metric items such instance name/uuid/id 
									continue 
								}else{
										if value,ok := parseStatus(perfCounterData.Value);ok {
										metricMap[perfCounterData.Name]=value
								}
							}
					}

					for metricName,metricValue := range  metricMap{
						metricName := metricNamePrefix+metricName+metricNameSufix
						desc := prometheus.NewDesc(
							prometheus.BuildFQName(namespace, PerfSubsystem, metricName),
							"Perf "+labelName[0]+" "+metricName,
						  labelName, nil)
							ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue,metricValue,labelValue...)
					}
//
//					for _,perfCounterData := range instanceData.Counters.CounterData{
//						metricName := metricNamePrefix+perfCounterData.Name+metricNameSufix
//						desc := prometheus.NewDesc(
//							prometheus.BuildFQName(namespace, PerfSubsystem, metricName),
//							"Perf "+labelName+" "+metricName,
//							[]string{labelName}, nil)
//						if  (strings.Contains(perfCounterData.Value,"name")  || strings.Contains(perfCounterData.Value,"id") |//| strings.Contains(perfCounterData.Value,"uuid")) {// this will filter out the items  which are not //metrics, but instance name, id or uuid
//							continue 
//						}
//						if  strings.ContainsAny(perfCounterData.Value,",- ")  { // this will filter out the items whose value //contains "," or "-", which mayb histogram data, or not even metrics value(contains "-")
//								continue
//							}
//            if value,ok := parseStatus(perfCounterData.Value);ok {
//								ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue,value, instanceName)
//						}					
//			}
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
	r = resp.Results.PerfObjectInstanceData.Instances
	 
	return
}
func GetPerfObjectInstanceList(netappClient *netapp.Client, objectName string) (r []netapp.InstanceInfo) {
	

	opts := &netapp.PerfObjectInstanceListInfoIterParams{
	Query: &netapp.InstanceInfoQuery{	},
	DesiredAttributes: &netapp.InstanceInfo{},
	ObjectName: objectName,
 	}
	resp,_,_ := netappClient.Perf.PerfObjectInstanceListInfoIter(opts)
	
  r=resp.Results.AttributesList.InstanceInfo
  return 
}
