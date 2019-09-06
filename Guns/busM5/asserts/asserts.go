package asserts

import (
	"regexp"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	Asserts        sync.Map
	assertsmetrics = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "mqops_operation_status",
			Help: "Response status",
		},
		[]string{"status", "operation", "code"},
	)
	success, fail int
)

func init() {
	prometheus.MustRegister(assertsmetrics)
	success = 0
	fail = 0
}

func CheckAssert(body string, operation string) {
	var code string
	json := parseStatusMessage(body)
	switch json {
	case "success":
		assertsmetrics.With(prometheus.Labels{"status": "success", "operation": operation, "code": "200"}).Inc()
		success++
	case "warning":
		code = parseCodeMessage(body)
		assertsmetrics.With(prometheus.Labels{"status": "warning", "operation": operation, "code": code}).Inc()
		fail++
	case "error":
		code = parseCodeMessage(body)
		assertsmetrics.With(prometheus.Labels{"status": "error", "operation": operation, "code": code}).Inc()
		fail++
	case "code_not_found":
		assertsmetrics.With(prometheus.Labels{"status": "code_not_found", "operation": operation, "code": "404"}).Inc()
		fail++
	}
}
func parseStatusMessage(str string) (res string) {
	var re = regexp.MustCompile(`(?m)"status":\s?"(\w+)"`)
	r := re.FindStringSubmatch(str)
	if len(r) > 0 {
		return r[1]
	}
	return "code_not_found"
}
func parseCodeMessage(str string) (res string) {
	var re = regexp.MustCompile(`(?m)"code":\s?"(\w+)"`)
	return re.FindStringSubmatch(str)[1]
}
func CheckTestResult() (res bool) {
	if success <= fail {
		res = false
	} else {
		r := (fail * 100) / (success + fail)
		if r > 5 {
			res = false
		} else {
			res = true
		}
	}
	return res
}
