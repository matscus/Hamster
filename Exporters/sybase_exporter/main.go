package main

import (
	"fmt"

	"net/http"
	"os"
	"sort"

	"./collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	listenport  string
	metricsPath = "/metrics"
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
		// No filters, use the prepared unfiltered handler.
		h.unfilteredHandler.ServeHTTP(w, r)
		return
	}
	// To serve filtered metrics, we create a filtering handler on the fly.
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
	r.MustRegister(version.NewCollector("sybase_exporter"))
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
		).Default(":9300").String()
		dbname = kingpin.Flag(
			"web.db-name",
			"Address on which to expose metrics and web interface.",
		).Default("master").String()
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
	os.Setenv("DBNAME", *dbname)
	log.AddFlags(kingpin.CommandLine)
	kingpin.Version(version.Print("sybase_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	log.Infoln("Starting node_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())
	http.Handle(*metricsPath, newHandler(!*disableExporterMetrics, *maxRequests))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Sybase Exporter</title></head>
			<body>
			<h1>Sybase Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})

	log.Infoln("Listening on", *listenAddress)
	if err := http.ListenAndServe(*listenAddress, nil); err != nil {
		log.Fatal(err)
	}
}
