package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	registerCollector("cpu", true, NewCPUCollector)
}

const (
	cpuCollectorSubsystem = "cpu"
)

type cpuCollector struct {
	cpuBusy *prometheus.Desc
	cpuIdle *prometheus.Desc
}

var (
	cpuBusySecondsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, cpuCollectorSubsystem, "cpu_busy_seconds_total"),
		"cpu activity",
		[]string{"cpu"}, nil,
	)
	cpuIdleSecondsTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, cpuCollectorSubsystem, "cpu_idle_seconds_total"),
		"cpu idle",
		[]string{"cpu"}, nil,
	)
	selectCPU = "SELECT CONVERT(REAL, @@cpu_busy) * (CONVERT(REAL, @@timeticks) / 1000000) cpu_busy_seconds_total,CONVERT(REAL, @@idle) * (CONVERT(REAL, @@timeticks) / 1000000) cpu_idle_seconds_total"
)

func NewCPUCollector() (Collector, error) {
	return &cpuCollector{
		cpuBusy: cpuBusySecondsTotal,
		cpuIdle: cpuIdleSecondsTotal,
	}, nil
}
func (c *cpuCollector) Update(ch chan<- prometheus.Metric) error {
	if err := c.updateCPU(ch); err != nil {
		return err
	}
	return nil
}
func (c *cpuCollector) updateCPU(ch chan<- prometheus.Metric) error {
	rows, err := DB.Query(selectCPU)
	if err != nil {
		return err
	}
	var busy, idle float64
	for rows.Next() {
		rows.Scan(&busy, &idle)
	}
	ch <- prometheus.MustNewConstMetric(
		c.cpuBusy,
		prometheus.GaugeValue,
		busy,
		"busy_seconds_total",
	)
	ch <- prometheus.MustNewConstMetric(
		c.cpuIdle,
		prometheus.GaugeValue,
		idle,
		"idle_seconds_total",
	)
	return nil
}
