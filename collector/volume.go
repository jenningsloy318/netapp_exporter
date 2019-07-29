package collector

import (
	"log"
	"github.com/pepabo/go-netapp/netapp"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Subsystem.
	VolumeSubsystem = "volume"
)

// Metric descriptors.
var (
	volumeLabels   = append(BaseLabelNames, "volume", "vserver", "aggr", "node")
	VolumeSizeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VolumeSubsystem, "size"),
		"Size of the volume.",
		volumeLabels, nil)
	VolumeSizeAvailableDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VolumeSubsystem, "size_available"),
		"Available Size of the volume.",
		volumeLabels, nil)
	VolumeSizeTotalDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VolumeSubsystem, "size_total"),
		"Total Size   of the volume.",
		volumeLabels, nil)
	VolumeSizeUsedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VolumeSubsystem, "size_used"),
		"Used Size of the volume.",
		volumeLabels, nil)
	VolumeSizeUsedBySnapshotsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VolumeSubsystem, "size_used_by_snapshots"),
		"Used Size By Snapshots of the volume.",
		volumeLabels, nil)
	VolumeSizeReservedBySnapshotDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VolumeSubsystem, "snapshot_reserve_size"),
		"Reserve Size By Snapshots of the volume.",
		volumeLabels, nil)
	VolumeStateDesc = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, VolumeSubsystem, "state"),
		"State of the volume, 1 (online), 0(offline), 2(restricted), or 3(mixed).",
		volumeLabels, nil)
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
	Name                   string
	Vserver                string
	Aggr                   string
	Node                   string
	Size                   int
	SizeAvailable          string
	SizeTotal              string
	SizeUsed               string
	SizeUsedBySnapshots    string
	SizeReservedBySnapshot string
	State                  string
}

// Scrape collects data from  netapp system and Volume info
func (ScrapeVolume) Scrape(netappClient *netapp.Client, ch chan<- prometheus.Metric) error {

	for _, VolumeInfo := range GetVolumeData(netappClient) {
		vserverLabelValues := append(BaseLabelValues, VolumeInfo.Name, VolumeInfo.Vserver, VolumeInfo.Aggr, VolumeInfo.Node)
		ch <- prometheus.MustNewConstMetric(VolumeSizeDesc, prometheus.GaugeValue, float64(VolumeInfo.Size), vserverLabelValues...)
		if sizeAvailable, ok := parseStatus(VolumeInfo.SizeAvailable); ok {
			ch <- prometheus.MustNewConstMetric(VolumeSizeAvailableDesc, prometheus.GaugeValue, sizeAvailable, vserverLabelValues...)
		}
		if sizeTotal, ok := parseStatus(VolumeInfo.SizeTotal); ok {
			ch <- prometheus.MustNewConstMetric(VolumeSizeTotalDesc, prometheus.GaugeValue, sizeTotal, vserverLabelValues...)
		}
		if sizeUsed, ok := parseStatus(VolumeInfo.SizeUsed); ok {
			ch <- prometheus.MustNewConstMetric(VolumeSizeUsedDesc, prometheus.GaugeValue, sizeUsed, vserverLabelValues...)
		}
		if sizeUsedBySnapshots, ok := parseStatus(VolumeInfo.SizeUsedBySnapshots); ok {
			ch <- prometheus.MustNewConstMetric(VolumeSizeUsedBySnapshotsDesc, prometheus.GaugeValue, sizeUsedBySnapshots, vserverLabelValues...)
		}
		if sizeReservedBySnapshot, ok := parseStatus(VolumeInfo.SizeReservedBySnapshot); ok {
			ch <- prometheus.MustNewConstMetric(VolumeSizeReservedBySnapshotDesc, prometheus.GaugeValue, sizeReservedBySnapshot, vserverLabelValues...)
		}
		if stateVal, ok := parseStatus(VolumeInfo.State); ok {
			ch <- prometheus.MustNewConstMetric(VolumeStateDesc, prometheus.GaugeValue, stateVal, vserverLabelValues...)
		}

	}
	return nil
}

func GetVolumeData(netappClient *netapp.Client) (r []*Volume) {
	ff := new(bool)
	*ff = false

	opts := &netapp.VolumeOptions{
		Query: &netapp.VolumeQuery{
			VolumeInfo: &netapp.VolumeInfo{},
		},
		DesiredAttributes: &netapp.VolumeQuery{
			VolumeInfo: &netapp.VolumeInfo{
				VolumeIDAttributes: &netapp.VolumeIDAttributes{
					Name:                    "x",
					OwningVserverName:       "x",
					ContainingAggregateName: "x",
					Node:                    "x",
				},
				VolumeSpaceAttributes: &netapp.VolumeSpaceAttributes{
					Size:                1,
					SizeAvailable:       "x",
					SizeTotal:           "x",
					SizeUsed:            "x",
					SizeUsedBySnapshots: "x",
					SnapshotReserveSize: "x",
				},
				VolumeStateAttributes: &netapp.VolumeStateAttributes{
					State: "x",
				},
			},
		},
	}

	l := getVolumeList(netappClient, opts)

	for _, n := range l {
		r = append(r, &Volume{
			Name:                   n.VolumeIDAttributes.Name,
			Vserver:                n.VolumeIDAttributes.OwningVserverName,
			Aggr:                   n.VolumeIDAttributes.ContainingAggregateName,
			Node:                   n.VolumeIDAttributes.Node,
			Size:                   n.VolumeSpaceAttributes.Size,
			SizeAvailable:          n.VolumeSpaceAttributes.SizeAvailable,
			SizeTotal:              n.VolumeSpaceAttributes.SizeTotal,
			SizeUsed:               n.VolumeSpaceAttributes.SizeUsed,
			SizeUsedBySnapshots:    n.VolumeSpaceAttributes.SizeUsedBySnapshots,
			SizeReservedBySnapshot: n.VolumeSpaceAttributes.SnapshotReserveSize,
			State:                  n.VolumeStateAttributes.State,
		})
	}
	return
}

func getVolumeList(netappClient *netapp.Client, opts *netapp.VolumeOptions) (r []netapp.VolumeInfo) {

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
