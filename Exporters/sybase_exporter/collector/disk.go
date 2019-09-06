package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	registerCollector("disk", true, NewDiskCollector)
}

const (
	diskCollectorSubsystem = "disk"
)

var (
	diskReadsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, diskCollectorSubsystem, "reads"),
		"total",
		[]string{"disk"}, nil,
	)
	diskWritesTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, diskCollectorSubsystem, "writes"),
		"total",
		[]string{"disk"}, nil,
	)
	diskErrorsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, diskCollectorSubsystem, "errors"),
		"total",
		[]string{"disk"}, nil,
	)
	selectDisk = "SELECT @@total_read disk_reads_total, @@total_write disk_writes_total, @@total_errors disk_errors_total"
)

type diskCollector struct {
	diskRead  *prometheus.Desc
	diskWrite *prometheus.Desc
	diskError *prometheus.Desc
}

func NewDiskCollector() (Collector, error) {
	return &diskCollector{
		diskRead:  diskReadsTotal,
		diskWrite: diskWritesTotal,
		diskError: diskErrorsTotal,
	}, nil
}
func (c *diskCollector) Update(ch chan<- prometheus.Metric) error {
	if err := c.updateDisk(ch); err != nil {
		return err
	}
	return nil
}
func (c *diskCollector) updateDisk(ch chan<- prometheus.Metric) error {
	rows, err := DB.Query(selectDisk)
	if err != nil {
		return err
	}
	var reads, writes, errors float64
	for rows.Next() {
		rows.Scan(&reads, &writes, &errors)
	}
	ch <- prometheus.MustNewConstMetric(
		c.diskRead,
		prometheus.GaugeValue,
		reads,
		"disk_reads_total",
	)
	ch <- prometheus.MustNewConstMetric(
		c.diskWrite,
		prometheus.GaugeValue,
		writes,
		"disk_writes_total",
	)
	ch <- prometheus.MustNewConstMetric(
		c.diskError,
		prometheus.GaugeValue,
		errors,
		"disk_errors_total",
	)
	return nil
}
