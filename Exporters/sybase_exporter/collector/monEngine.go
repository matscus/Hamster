package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	registerCollector("monEngine", true, NewmonEngineCollector)
}

type monEngine struct {
	EngineNumber          float64
	CurrentKPID           float64
	PreviousKPID          float64
	CPUTime               float64
	SystemCPUTime         float64
	UserCPUTime           float64
	IOCPUTime             float64
	IdleCPUTime           float64
	Yields                float64
	Connections           float64
	DiskIOChecks          float64
	DiskIOPolled          float64
	DiskIOCompleted       float64
	MaxOutstandingIOs     float64
	ProcessesAffinitied   float64
	ContextSwitches       float64
	HkgcMaxQSize          float64
	HkgcPendingItems      float64
	HkgcHWMItems          float64
	HkgcOverflows         float64
	HkgcPendingItemsDcomp float64
	HkgcOverflowsDcomp    float64
	AffinitiedToCPU       float64
	OSPID                 float64
}

const (
	monEngineCollectorSubsystem = "monEngine"
)

var (
	monEngines = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, monEngineCollectorSubsystem, "monEngine"),
		"monEngine",
		[]string{"monEngine"}, nil,
	)

	//selectMonEngine = "SELECT * from monEngine"
	selectMonEngine = `select EngineNumber,CurrentKPID,PreviousKPID,CPUTime,SystemCPUTime,UserCPUTime,IOCPUTime,
	IdleCPUTime,Yields,Connections,DiskIOChecks,DiskIOPolled,DiskIOCompleted,MaxOutstandingIOs,
	ProcessesAffinitied,ContextSwitches,HkgcMaxQSize,HkgcPendingItems,HkgcHWMItems,HkgcOverflows,
	HkgcPendingItemsDcomp,HkgcOverflowsDcomp,AffinitiedToCPU,OSPID from monEngine`
)

type monEngineCollector struct {
	monengine *prometheus.Desc
}

func NewmonEngineCollector() (Collector, error) {
	return &monEngineCollector{
		monengine: monEngines,
	}, nil
}
func (c *monEngineCollector) Update(ch chan<- prometheus.Metric) error {
	if err := c.updateMonEngine(ch); err != nil {
		return err
	}
	return nil
}
func (c *monEngineCollector) updateMonEngine(ch chan<- prometheus.Metric) error {
	rows, err := DB.Query(selectMonEngine)
	if err != nil {
		return err
	}
	e := monEngine{}
	for rows.Next() {
		rows.Scan(&e.EngineNumber, &e.CurrentKPID, &e.PreviousKPID, &e.CPUTime, &e.SystemCPUTime, &e.UserCPUTime, &e.IOCPUTime, &e.IdleCPUTime, &e.Yields, &e.Connections, &e.DiskIOChecks, &e.DiskIOPolled, &e.DiskIOCompleted, &e.MaxOutstandingIOs, &e.ProcessesAffinitied, &e.ContextSwitches, &e.HkgcMaxQSize, &e.HkgcPendingItems, &e.HkgcHWMItems, &e.HkgcOverflows, &e.HkgcPendingItemsDcomp, &e.HkgcOverflowsDcomp, &e.AffinitiedToCPU, &e.OSPID)
	}
	res := getMap(&e)
	for key, value := range res {
		ch <- prometheus.MustNewConstMetric(
			c.monengine,
			prometheus.CounterValue,
			value,
			key,
		)
	}
	return nil
}

func getMap(e *monEngine) map[string]float64 {
	res := make(map[string]float64)
	res["EngineNumber"] = e.EngineNumber
	res["CurrentKPID"] = e.CurrentKPID
	res["PreviousKPID"] = e.PreviousKPID
	res["CPUTime"] = e.CPUTime
	res["SystemCPUTime"] = e.SystemCPUTime
	res["UserCPUTime"] = e.UserCPUTime
	res["IOCPUTime"] = e.IOCPUTime
	res["IdleCPUTime"] = e.IdleCPUTime
	res["Yields"] = e.Yields
	res["Connections"] = e.Connections
	res["DiskIOChecks"] = e.DiskIOChecks
	res["DiskIOPolled"] = e.DiskIOPolled
	res["DiskIOCompleted"] = e.DiskIOCompleted
	res["MaxOutstandingIOs"] = e.MaxOutstandingIOs
	res["ProcessesAffinitied"] = e.ProcessesAffinitied
	res["ContextSwitches"] = e.ContextSwitches
	res["HkgcMaxQSize"] = e.HkgcMaxQSize
	res["HkgcPendingItems"] = e.HkgcPendingItems
	res["HkgcHWMItems"] = e.HkgcHWMItems
	res["HkgcOverflows"] = e.HkgcOverflows
	res["HkgcPendingItemsDcomp"] = e.HkgcPendingItemsDcomp
	res["HkgcOverflowsDcomp"] = e.HkgcOverflowsDcomp
	res["AffinitiedToCPU"] = e.AffinitiedToCPU
	res["OSPID"] = e.OSPID
	return res
}
