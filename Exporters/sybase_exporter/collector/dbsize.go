package collector

import (
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	registerCollector("dbsize", true, NewDBSizeCollector)
}

const (
	dbSizeCollectorSubsystem = "dbsize"
)

var (
	dbSize = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, dbSizeCollectorSubsystem, "db_size"),
		"db_size",
		[]string{"db_name"}, nil,
	)
	dbFree = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, dbSizeCollectorSubsystem, "db_free"),
		"db_sfree",
		[]string{"db_name"}, nil,
	)
	selectDBSize = `sp_helpdb ` + os.Getenv("DBNAME")
)

type dbSizeCollector struct {
	dbSize *prometheus.Desc
	dbFree *prometheus.Desc
}

func NewDBSizeCollector() (Collector, error) {
	return &dbSizeCollector{
		dbSize: dbSize,
		dbFree: dbFree,
	}, nil
}
func (c *dbSizeCollector) Update(ch chan<- prometheus.Metric) error {
	if err := c.updateDB(ch); err != nil {
		return err
	}
	return nil
}
func (c *dbSizeCollector) updateDB(ch chan<- prometheus.Metric) error {
	rows, err := DB.Query(selectDBSize)
	if err != nil {
		return err
	}
	var name, device, log string
	var mb, free float64
	for rows.Next() {
		rows.Scan(&name, &device, &mb, &log, &free)
	}
	ch <- prometheus.MustNewConstMetric(
		c.dbSize,
		prometheus.GaugeValue,
		mb,
		name,
	)
	ch <- prometheus.MustNewConstMetric(
		c.dbFree,
		prometheus.GaugeValue,
		free,
		name,
	)
	return nil
}
