package collector

import (
	"log"
	"github.com/pepabo/go-netapp/netapp"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

const (
	// Subsystem.
	VolumeSubsystem = "volume"
)

// Metric descriptors.
var (
	VolumeSizeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VolumeSubsystem, "size"),
		"Size of the volume.",
		[]string{"volume","vserver","aggr","node"}, nil)
	VolumeSizeAvailableDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VolumeSubsystem, "size_available"),
		"Available Size of the volume.",
		[]string{"volume","vserver","aggr","node"}, nil)
	VolumeSizeTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VolumeSubsystem, "size_total"),
		"Total Size   of the volume.",
		[]string{"volume","vserver","aggr","node"}, nil)
	VolumeSizeUsedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VolumeSubsystem, "size_used"),
		"Used Size of the volume.",
		[]string{"volume","vserver","aggr","node"}, nil)
	VolumeSizeUsedBySnapshotsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VolumeSubsystem, "size_used_by_snapshots"),
		"Used Size By Snapshots of the volume.",
		[]string{"volume","vserver","aggr","node"}, nil)		
	VolumeSizeReservedBySnapshotDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VolumeSubsystem, "snapshot_reserve_size"),
		"Reserve Size By Snapshots of the volume.",
		[]string{"volume","vserver","aggr","node"}, nil)		
	)


// Scrapesystem collects system Volume info
type ScrapeVolume struct{}

// Name of the Scraper. Should be unique.
func (ScrapeVolume) Name() string {
	return VolumeSubsystem
}

// Help describes the role of the Scraper.
func (ScrapeVolume) Help() string {
	return "Collect Netapp Volume info;"
}


type Volume struct {
	Name                     string
	Vserver                  string
	Aggr                     string
	Node 										 string
	Size                     int
	SizeAvailable            string
	SizeTotal                string
	SizeUsed                 string
	SizeUsedBySnapshots      string
	SizeReservedBySnapshot      string

}


// Scrape collects data from  netapp system and Volume info 
func (ScrapeVolume) Scrape(netappClient *netapp.Client, ch chan<- prometheus.Metric) error {

		
	for _, VolumeInfo := range GetVolumeData(netappClient) {
		ch <- prometheus.MustNewConstMetric(VolumeSizeDesc, prometheus.GaugeValue,float64(VolumeInfo.Size), VolumeInfo.Name,VolumeInfo.Vserver,VolumeInfo.Aggr,VolumeInfo.Node)
		if sizeAvailable,err :=  strconv.ParseFloat(VolumeInfo.SizeAvailable, 64); err == nil{
			ch <- prometheus.MustNewConstMetric(VolumeSizeAvailableDesc, prometheus.GaugeValue,sizeAvailable, VolumeInfo.Name,VolumeInfo.Vserver,VolumeInfo.Aggr,VolumeInfo.Node)
		}
		if sizeTotal,err :=  strconv.ParseFloat(VolumeInfo.SizeTotal, 64); err == nil{
			ch <- prometheus.MustNewConstMetric(VolumeSizeTotalDesc, prometheus.GaugeValue,sizeTotal, VolumeInfo.Name,VolumeInfo.Vserver,VolumeInfo.Aggr,VolumeInfo.Node)
		}
		if sizeUsed,err :=  strconv.ParseFloat(VolumeInfo.SizeUsed, 64); err == nil{
			ch <- prometheus.MustNewConstMetric(VolumeSizeUsedDesc, prometheus.GaugeValue,sizeUsed, VolumeInfo.Name,VolumeInfo.Vserver,VolumeInfo.Aggr,VolumeInfo.Node)
		}		
		if sizeUsedBySnapshots,err :=  strconv.ParseFloat(VolumeInfo.SizeUsedBySnapshots, 64); err == nil{
			ch <- prometheus.MustNewConstMetric(VolumeSizeUsedBySnapshotsDesc, prometheus.GaugeValue,sizeUsedBySnapshots, VolumeInfo.Name,VolumeInfo.Vserver,VolumeInfo.Aggr,VolumeInfo.Node)
		}			
		if sizeReservedBySnapshot,err :=  strconv.ParseFloat(VolumeInfo.SizeReservedBySnapshot, 64); err == nil{
			ch <- prometheus.MustNewConstMetric(VolumeSizeReservedBySnapshotDesc, prometheus.GaugeValue,sizeReservedBySnapshot, VolumeInfo.Name,VolumeInfo.Vserver,VolumeInfo.Aggr,VolumeInfo.Node)
		}					
	}
	return nil
}



func GetVolumeData(netappClient *netapp.Client) (r []*Volume) {
	ff := new(bool)
	*ff = false

	opts := &netapp.VolumeOptions {
		Query: &netapp.VolumeQuery{
			VolumeInfo : &netapp.VolumeInfo {},
		},
		DesiredAttributes: &netapp.VolumeQuery{
			VolumeInfo: &netapp.VolumeInfo{
				VolumeSpaceAttributes: &netapp.VolumeSpaceAttributes{},
				VolumeIDAttributes :  &netapp.VolumeIDAttributes {},
				VolumeStateAttributes: &netapp.VolumeStateAttributes{},
			},
		},
	}

	l := getVolumeList(netappClient,opts)

	for _, n := range l {
		r = append(r, &Volume{
			Name:                n.VolumeIDAttributes.Name,
			Vserver:             n.VolumeIDAttributes.OwningVserverName,
			Aggr: 							 n.VolumeIDAttributes.ContainingAggregateName,
			Node: 		 					 n.VolumeIDAttributes.Node,
			Size:                n.VolumeSpaceAttributes.Size,
			SizeAvailable:       n.VolumeSpaceAttributes.SizeAvailable,
			SizeTotal:           n.VolumeSpaceAttributes.SizeTotal,
			SizeUsed:            n.VolumeSpaceAttributes.SizeUsed,
			SizeUsedBySnapshots: n.VolumeSpaceAttributes.SizeUsedBySnapshots,
			SizeReservedBySnapshot: n.VolumeSpaceAttributes.SnapshotReserveSize,
		})
	}
	return
}

func getVolumeList(netappClient *netapp.Client,opts *netapp.VolumeOptions ) (r []netapp.VolumeInfo) {

	var pages []*netapp.VolumeListResponse 
	handler := func(r netapp.VolumeListPagesResponse) bool {
		if r.Error != nil {
				log.Printf("%s", r.Error)
			return false
		}
		pages = append(pages, r.Response)
		return true
	}

	netappClient.Volume.ListPages(opts, handler)

	for _, p := range pages {
		r = append(r, p.Results.AttributesList...)
	}

	return
}
