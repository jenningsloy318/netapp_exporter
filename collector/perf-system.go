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
	objects=  []string{"system",}

	PerfinfoDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, PerfSubsystem, "size"),
		"Size of the perf.",
		[]string{"object", }, nil)
	
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
					for _,perfCounterData := range instanceData.Counters.CounterData{
						if countersValue,ok := parseStatus(perfCounterData.Value);ok {

							metricName := perfCounterData.Name
							desc := prometheus.NewDesc(
								prometheus.BuildFQName(namespace, PerfSubsystem, metricName),
								"Perf "+obj+" "+metricName,
								[]string{"object"}, nil)
							
								ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue,countersValue, obj)
						}
					}
			}
		}
		return nil
}



type PerfSystem struct {
		instance_name              string
		avg_processor_busy         float64
		cifs_ops                   float64
		compile_flags              float64
		cp                         float64
		cp_time                    float64
		cpu_busy                   float64
		cpu_elapsed_time           float64
		cpu_elapsed_time1          float64
		cpu_elapsed_time2          float64
		disk_data_read             float64
		disk_data_written          float64
		domain_busy                string
		domain_shared              string
		dswitchto_cnt              string
		fcp_data_recv              float64
		fcp_data_sent              float64
		fcp_ops                    float64
		hard_switches              float64
		hdd_data_read              float64
		hdd_data_written           float64
		hostname                   float64
		http_ops                   float64
		idle                       float64
		idle_time                  float64
		interrupt                  float64
		interrupt_in_cp            float64
		interrupt_in_cp_time       float64
		interrupt_num              float64
		interrupt_num_in_cp        float64
		interrupt_time             float64
		intr_cnt                   string
		intr_cnt_ipi               float64
		intr_cnt_msec              float64
		intr_cnt_total             float64
		iscsi_data_recv            float64
		iscsi_data_sent            float64
		iscsi_ops                  float64
		memory                     float64
		net_data_recv              float64
		net_data_sent              float64
		nfs_ops                    float64
		non_interrupt              float64
		non_interrupt_time         float64
		num_processors             float64
		ontap_version              float64
		other_data                 float64
		other_latency              float64
		other_ops                  float64
		partner_data_recv          float64
		partner_data_sent          float64
		partner_ops                float64
		processor_plevel           string
		processor_plevel_time      string
		read_data                  float64
		read_latency               float64
		read_ops                   float64
		sk_switches                float64
		ssd_data_read              float64
		ssd_data_written           float64
		sys_read_data              float64
		sys_total_data             float64
		sys_write_data             float64
		system_id                  float64
		system_model               float64
		system_ops                 float64
		tape_data_read             float64
		tape_data_written          float64
		time                       float64
		time_per_interrupt         float64
		time_per_interrupt_in_cp   float64
		total_data                 float64
		total_latency              float64
		total_ops                  float64
		total_processor_busy       float64
		total_processor_busy_time  float64
		uptime                     float64
		wafliron                   float64
		write_data                 float64
		write_latency              float64
		write_ops                  float64
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
