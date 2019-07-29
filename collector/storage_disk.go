package collector

import (
	"log"
	"github.com/pepabo/go-netapp/netapp"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Subsystem.
	StorageDiskSubsystem = "storage_disk"
)

// Metric descriptors.
var (
	storageDiskLabels       = append(BaseLabelNames, "disk","node", "type","model")
	storageDiskHealthStateDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, StorageDiskSubsystem, "is_failed"),
		"if this disk is failed.",
		storageDiskLabels, nil)
)

// Scrapesystem collects system node info
type ScrapeStorageDisk struct{}

// Name of the Scraper. Should be unique.
func (ScrapeStorageDisk) Name() string {
	return StorageDiskSubsystem
}

// Help describes the role of the Scraper.
func (ScrapeStorageDisk) Help() string {
	return "Collect Netapp storage disk info;"
}

type StorageDisk struct {
	DiskName            string
	DiskType            string
	Model               string
	IsFailed            *bool
	HomeNodeName        string 
}


// Scrape collects data from  netapp StorageDisk info
func (ScrapeStorageDisk) Scrape(netappClient *netapp.Client, ch chan<- prometheus.Metric) error {

	for _, storageDiskInfo := range GetStorageDiskData(netappClient) {
		storageDiskLabelValues := append(BaseLabelValues, storageDiskInfo.DiskName,storageDiskInfo.HomeNodeName,storageDiskInfo.DiskType,storageDiskInfo.Model)
		ch <- prometheus.MustNewConstMetric(storageDiskHealthStateDesc, prometheus.GaugeValue, boolToFloat64(*storageDiskInfo.IsFailed), storageDiskLabelValues...)
	
	}

	return nil
}

func GetStorageDiskData(netappClient *netapp.Client) (r []*StorageDisk) {
	
	ff := new(bool)
	*ff = false

	opts := &netapp.StorageDiskOptions{
		Query: &netapp.StorageDiskInfo{},
		DesiredAttributes: &netapp.StorageDiskInfo{
			DiskName: "x",
			DiskInventoryInfo:  &netapp.DiskInventoryInfo{
				DiskType: "x",
				Model: "x",
			},
			DiskOwnershipInfo: &netapp.DiskOwnershipInfo{
				HomeNodeName: "x",
				IsFailed: ff,
			},
		},
	}




	res,_,err := netappClient.StorageDisk.StorageDiskGetIter(opts) 
	if err !=nil  {
		log.Fatalf("error when getting storage disks, %s",err)
	}

	for _, n := range res.Results.AttributesList.StorageDiskInfo {
		r = append(r, &StorageDisk{
			DiskName:                n.DiskName,
			DiskType:							   n.DiskInventoryInfo.DiskType,
			Model:									 n.DiskInventoryInfo.Model,
			HomeNodeName:						 n.DiskOwnershipInfo.HomeNodeName,
			IsFailed:								 n.DiskOwnershipInfo.IsFailed,
		})
	}
	return
}


