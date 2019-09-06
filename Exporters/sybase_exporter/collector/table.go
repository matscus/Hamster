package collector

import (
	"log"
	"os"

	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	registerCollector("tablesize", true, NewTableCollector)

}

const (
	tableCollectorSubsystem = "tablesize"
)

type tableCollector struct {
	rowTotal *prometheus.Desc
}

var (
	tableRowTotal = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, tableCollectorSubsystem, "rowtotal"),
		"table rows total",
		[]string{"table"}, nil,
	)
	selectTable = `SELECT o.name AS table_name,
	SUM(row_count(?,?)) AS rowtotal`
	dbname      = os.Getenv("DBNAME")
	tablesNames = []string{"tContractCredit", "tCard", "tContract", "tInstitution", "tNode", "tResource "}
	tablesID    = make([]int, 0, 6)
)

func NewTableCollector() (Collector, error) {
	return &tableCollector{
		rowTotal: tableRowTotal,
	}, nil
}
func (c *tableCollector) Update(ch chan<- prometheus.Metric) error {
	if err := c.updateTable(ch); err != nil {
		return err
	}
	return nil
}
func (c *tableCollector) updateTable(ch chan<- prometheus.Metric) error {
	var tableName string
	var rowTotal, heapSize, indexSize, unusedSize float64
	rows, err := DB.Query(selectTable)
	if err != nil {
		return err
	}
	for rows.Next() {
		rows.Scan(&tableName, &heapSize, &indexSize, &unusedSize)
	}
	ch <- prometheus.MustNewConstMetric(
		c.rowTotal,
		prometheus.GaugeValue,
		rowTotal,
		tableName,
	)
	return nil
}

func getTableID() {
	l := len(tablesNames)
	for i := 0; i < l; i++ {
		rows, err := DB.Query("select data_pages(?,?)", dbname, tablesNames[i])
		if err != nil {
			log.Printf("error get table id for table %s :%s", tablesNames[i], err)
		}
		var id int
		for rows.Next() {
			rows.Scan(&id)
		}
		tablesID = append(tablesID, id)
	}
}
