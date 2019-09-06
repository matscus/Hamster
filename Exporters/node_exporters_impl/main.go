package main

import (
	//"github.com/shirou/gopsutil/net"

	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"./collector"
	"./prometheus/client_golang/prometheus"
	"./prometheus/client_golang/prometheus/promhttp"
	"./prometheus/common/log"
	"./prometheus/common/version"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type handler struct {
	unfilteredHandler       http.Handler
	exporterMetricsRegistry *prometheus.Registry
	includeExporterMetrics  bool
	maxRequests             int
}

func newHandler(includeExporterMetrics bool, maxRequests int) *handler {
	h := &handler{
		exporterMetricsRegistry: prometheus.NewRegistry(),
		includeExporterMetrics:  includeExporterMetrics,
		maxRequests:             maxRequests,
	}
	if h.includeExporterMetrics {
		h.exporterMetricsRegistry.MustRegister(
			prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
			prometheus.NewGoCollector(),
		)
	}
	if innerHandler, err := h.innerHandler(); err != nil {
		log.Fatalf("Couldn't create metrics handler: %s", err)
	} else {
		h.unfilteredHandler = innerHandler
	}
	return h
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filters := r.URL.Query()["collect[]"]
	log.Debugln("collect query:", filters)

	if len(filters) == 0 {

		h.unfilteredHandler.ServeHTTP(w, r)
		return
	}

	filteredHandler, err := h.innerHandler(filters...)
	if err != nil {
		log.Warnln("Couldn't create filtered metrics handler:", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Couldn't create filtered metrics handler: %s", err)))
		return
	}
	filteredHandler.ServeHTTP(w, r)
}
func (h *handler) innerHandler(filters ...string) (http.Handler, error) {
	nc, err := collector.NewNodeCollector(filters...)
	if err != nil {
		return nil, fmt.Errorf("couldn't create collector: %s", err)
	}

	if len(filters) == 0 {
		log.Infof("Enabled collectors:")
		collectors := []string{}
		for n := range nc.Collectors {
			collectors = append(collectors, n)
		}
		sort.Strings(collectors)
		for _, n := range collectors {
			log.Infof(" - %s", n)
		}
	}

	r := prometheus.NewRegistry()
	r.MustRegister(version.NewCollector("matscus_exporter"))
	if err := r.Register(nc); err != nil {
		return nil, fmt.Errorf("couldn't register node collector: %s", err)
	}
	handler := promhttp.HandlerFor(
		prometheus.Gatherers{h.exporterMetricsRegistry, r},
		promhttp.HandlerOpts{
			ErrorLog:            log.NewErrorLogger(),
			ErrorHandling:       promhttp.ContinueOnError,
			MaxRequestsInFlight: h.maxRequests,
		},
	)
	if h.includeExporterMetrics {
		handler = promhttp.InstrumentMetricHandler(
			h.exporterMetricsRegistry, handler,
		)
	}
	return handler, nil
}

func main() {
	var (
		listenAddress = kingpin.Flag(
			"web.listen-address",
			"Address on which to expose metrics and web interface.",
		).Default(":9100").String()
		metricsPath = kingpin.Flag(
			"web.telemetry-path",
			"Path under which to expose metrics.",
		).Default("/metrics").String()
		disableExporterMetrics = kingpin.Flag(
			"web.disable-exporter-metrics",
			"Exclude metrics about the exporter itself (promhttp_*, process_*, go_*).",
		).Bool()
		maxRequests = kingpin.Flag(
			"web.max-requests",
			"Maximum number of parallel scrape requests. Use 0 to disable.",
		).Default("40").Int()
	)

	log.AddFlags(kingpin.CommandLine)
	kingpin.Version(version.Print("node_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	log.Infoln("Starting node_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	http.Handle(*metricsPath, newHandler(!*disableExporterMetrics, *maxRequests))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Node Exporter</title></head>
			<body>
			<h1>Node Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})

	log.Infoln("Listening on", *listenAddress)
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Fatal(err)
	}

}

func GetStat() {
	for {
		proc, err := process.Processes()
		CheckError(err)
		for i := 0; i < len(proc); i++ {
			process, err := proc[i].Name()
			CheckError(err)
			switch process {
			case "mdm":
				p := proc[i]
				pid, err := p.Tgid()
				CheckError(err)
				con, _ := net.ConnectionsPid("tcp", pid)
				status, _ := proc[i].Status()
				parent, _ := proc[i].Parent()
				numCtxSwitches, _ := proc[i].NumCtxSwitches()
				uids, _ := proc[i].Uids()
				gids, _ := proc[i].Gids()
				numThreads, _ := proc[i].NumThreads()
				meminfo, _ := p.MemoryInfoEx()
				lastCPUTimes, _ := p.Times()
				fmt.Println("--------------//////////////////-----------------")
				fmt.Println("process: " + process)
				fmt.Println("status: " + status)
				fmt.Print("parent: ")
				fmt.Println(parent)
				fmt.Print("numCtxSwitches: ")
				fmt.Println(numCtxSwitches)
				fmt.Print("uids: ")
				fmt.Println(uids)
				fmt.Print("guids: ")
				fmt.Println(gids)
				fmt.Print("numThreads: ")
				fmt.Println(numThreads)
				fmt.Print("memInfo:")
				fmt.Println(meminfo)
				fmt.Print("lastCPUTimes: ")
				fmt.Println(lastCPUTimes)
				fmt.Print("net: ")
				fmt.Println(con)
				break
			}
			time.Sleep(10 * time.Second)
		}
	}
}

func Float64FromBytes(b []byte) float64 {
	s := string(b)
	str := strings.Replace(s, "\n", "", 1)
	f, err := strconv.Atoi(str)
	CheckError(err)
	return float64(f)
}

func Float64FromString(s string) float64 {
	str := strings.Replace(s, "\n", "", 1)
	f, err := strconv.Atoi(str)
	CheckError(err)
	return float64(f)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
