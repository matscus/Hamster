package collector

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/common/log"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"

	"../prometheus/client_golang/prometheus"
)

const (
	socketSubsystem = "socketCounter"
)

type socketCollector struct {
	socketLocalCounter  *prometheus.Desc
	socketRemoteCounter *prometheus.Desc
	socketStatCounter   *prometheus.Desc
	socketJavaCounter   *prometheus.Desc
}

func init() {
	registerCollector("matscus_socket_counter", defaultEnabled, NewSocketCountCollector)
}

func NewSocketCountCollector() (Collector, error) {
	return &socketCollector{
		socketLocalCounter: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, socketSubsystem, "socketLocalCounter"),
			"tested.",
			[]string{"ip", "port"}, nil,
		),
		socketRemoteCounter: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, socketSubsystem, "socketRemoteCounter"),
			"tested.",
			[]string{"ip", "port"}, nil,
		),
		socketStatCounter: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, socketSubsystem, "socketStatCounter"),
			"tested.",
			[]string{"collector", "state"}, nil,
		),
		socketJavaCounter: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, socketSubsystem, "socketJavaCounter"),
			"tested.",
			[]string{"process_name", "values"}, nil,
		),
	}, nil
}
func (c *socketCollector) Update(ch chan<- prometheus.Metric) error {
	if err := c.updateLocalCounter(ch); err != nil {
		return err
	}
	if err := c.updateRemoteCounter(ch); err != nil {
		return err
	}
	if err := c.updateStatCounter(ch); err != nil {
		return err
	}
	if err := c.updateJavaCounter(ch); err != nil {
		return err
	}
	return nil
}
func (c *socketCollector) updateJavaCounter(ch chan<- prometheus.Metric) error {
	proc, err := process.Processes()
	res := make(map[string]map[string]float64)
	metric := make(map[string]float64)
	CheckError("error func process.Processes()", err)
	for i := 0; i < len(proc); i++ {
		process, err := proc[i].Name()
		CheckError("error read procces for name", err)
		switch process {
		case "java":
			p := proc[i]
			pid, err := p.Tgid()
			CheckError("error get proccer Tgid", err)
			con, err := net.ConnectionsPid("tcp", pid)
			CheckError("error get connection pid", err)
			var lPorts, rPorts []uint32
			for i := 0; i < len(con); i++ {
				lPorts = append(lPorts, con[i].Laddr.Port)
				rPorts = append(rPorts, con[i].Raddr.Port)
			}
			jps, err := exec.Command("sh", "-c", "jps  | grep "+strconv.Itoa(int(pid))+"| awk '{print $2}'").Output()
			CheckError("error get process name", err)
			name := strings.Replace(string(jps), "\n", "", -1)
			metric["localPorts"] = float64(len(lPorts))
			metric["remotePorts"] = float64(len(rPorts))
			upTime, err := p.CreateTime()
			CheckError("error get process create time", err)
			str := strconv.Itoa(int(upTime / 1000))
			i, err := strconv.ParseInt(str, 10, 64)
			CheckError("error parse process time create", err)
			tm := time.Unix(i, 0)
			t := time.Now()
			duration := t.Sub(tm)
			min := duration.Minutes()
			pUsed, err := p.CPUPercent()
			CheckError("error get process cpu percent", err)
			mUsed, err := p.MemoryPercent()
			CheckError("error get process memory percent", err)
			treads, err := p.NumThreads()
			CheckError("error get process treads count", err)
			//test:=p.NumFDs()
			metric["treads"] = float64(treads)
			metric["upTime"] = min
			metric["useCpuPercent"] = pUsed
			metric["useMemoryPercent"] = float64(mUsed)
			res[name] = metric
		}
	}
	for physicalPackageID, coreMap := range res {
		for processName, values := range coreMap {
			ch <- prometheus.MustNewConstMetric(c.socketJavaCounter,
				prometheus.CounterValue,
				values,
				physicalPackageID,
				processName)
		}
	}
	return nil
}

func (c *socketCollector) updateLocalCounter(ch chan<- prometheus.Metric) error {
	proc, err := process.Processes()
	var allCon []net.ConnectionStat
	CheckError("error func process.Processes()", err)
	for i := 0; i < len(proc); i++ {
		p := proc[i]
		pid, err := p.Tgid()
		CheckError("error get proccer Tgid", err)
		con, err := net.ConnectionsPid("tcp", pid)
		CheckError("error get connection pid", err)
		for i := 0; i < len(con); i++ {
			allCon = append(allCon, con[i])
		}
	}
	localConnection := GetUniqueLocalConenctionStat(allCon)

	for physicalPackageID, coreMap := range localConnection {
		for ip, socketCounter := range coreMap {
			ch <- prometheus.MustNewConstMetric(c.socketLocalCounter,
				prometheus.CounterValue,
				socketCounter,
				physicalPackageID,
				ip)
		}
	}
	return nil
}
func (c *socketCollector) updateRemoteCounter(ch chan<- prometheus.Metric) error {
	proc, err := process.Processes()
	var allCon []net.ConnectionStat
	CheckError("error func process.Processes()", err)
	for i := 0; i < len(proc); i++ {
		p := proc[i]
		pid, err := p.Tgid()
		CheckError("error get proccer Tgid", err)
		con, err := net.ConnectionsPid("tcp", pid)
		CheckError("error get connection pid", err)
		for i := 0; i < len(con); i++ {
			allCon = append(allCon, con[i])
		}
	}
	remoteConnection := GetUniqueRemoteConenctionStat(allCon)

	for physicalPackageID, coreMap := range remoteConnection {
		for ip, socketCounter := range coreMap {
			ch <- prometheus.MustNewConstMetric(c.socketRemoteCounter,
				prometheus.CounterValue,
				socketCounter,
				physicalPackageID,
				ip)
		}
	}
	return nil
}

func (c *socketCollector) updateStatCounter(ch chan<- prometheus.Metric) error {
	statConnection := GetSocketStatCounter()
	for physicalPackageID, coreMap := range statConnection {
		for state, socketCounter := range coreMap {
			ch <- prometheus.MustNewConstMetric(c.socketStatCounter,
				prometheus.CounterValue,
				socketCounter,
				physicalPackageID,
				state)
		}
	}
	return nil
}

func GetUniqueLocalConenctionStat(con []net.ConnectionStat) map[string]map[string]float64 {
	dumpLemote := make(map[string][]uint32)
	res := make(map[string]map[string]float64)
	var ip string
	for i := 0; i < len(con); i++ {
		ip = con[i].Laddr.IP
		// d := portRemoteSlice
		if v, ok := dumpLemote[ip]; ok {
			dump := append(v, con[i].Laddr.Port)
			dumpLemote[ip] = dump
		} else {
			var dump = []uint32{con[i].Laddr.Port}
			dumpLemote[ip] = dump
		}

	}
	for k, v := range dumpLemote {
		countConnect := getUniquePorts(v)
		res[k] = countConnect

	}
	return res
}
func GetUniqueRemoteConenctionStat(con []net.ConnectionStat) map[string]map[string]float64 {
	dumpRemote := make(map[string][]uint32)
	res := make(map[string]map[string]float64)
	var ip string
	for i := 0; i < len(con); i++ {
		ip = con[i].Raddr.IP
		if v, ok := dumpRemote[ip]; ok {
			dump := append(v, con[i].Raddr.Port)
			dumpRemote[ip] = dump
		} else {
			var dump = []uint32{con[i].Raddr.Port}
			dumpRemote[ip] = dump
		}

	}
	for k, v := range dumpRemote {
		countConnect := getUniquePorts(v)
		res[k] = countConnect

	}
	return res
}

func GetSocketStatCounter() map[string]map[string]float64 {
	var netStatSocketCounter = map[string]map[string]float64{}
	countListen, _ := exec.Command("sh", "-c", "ss -ant | grep LISTEN -c").Output()
	//CheckError("error get LISTEN connection count", err)
	countTimeWait, _ := exec.Command("sh", "-c", "ss -ant | grep TIME-WAIT -c").Output()
	//CheckError("error get TIME-WAIT connection count", err)
	countCloseWait, _ := exec.Command("sh", "-c", "ss -ant | grep CLOSE-WAIT -c").Output()
	//CheckError("error get CLOSE-WAIT connection count", err)
	countEstab, _ := exec.Command("sh", "-c", "ss -ant | grep ESTAB -c").Output()
	//CheckError("error get ESTAB connection count", err)
	countSynSent, _ := exec.Command("sh", "-c", "ss -ant | grep SYN-SENT -c").Output()
	// CheckError("error convert count SYN-SENT  to float", err)
	var connectionStatCount = map[string]float64{}
	connectionStatCount["LISTEN"] = float64(Float64FromBytes(countListen))
	connectionStatCount["ESTAB"] = float64(Float64FromBytes(countEstab))
	connectionStatCount["TIME_WAIT"] = float64(Float64FromBytes(countTimeWait))
	connectionStatCount["CLOSE_WAIT"] = float64(Float64FromBytes(countCloseWait))
	connectionStatCount["SYN_SENT"] = float64(Float64FromBytes(countSynSent))
	netStatSocketCounter["ConnectionCount"] = connectionStatCount

	return netStatSocketCounter
}
func getCoutConnectToPort(p []uint32) map[string]float64 {
	res := make(map[string]float64)
	for i := 0; i < len(p); i++ {
		port := fmt.Sprint(p[i])
		if _, ok := res[port]; ok {

			value := 1
			value++
			res[port] = float64(value)
		} else {
			res[port] = float64(1)
		}
	}
	return res
}

func portsToString(ports []uint32) []string {
	res := make([]string, len(ports), len(ports))
	for i := 0; i < len(ports); i++ {
		res[i] = fmt.Sprint(ports[i])
	}
	return res
}
func CheckError(str string, err error) {
	if err != nil {
		log.Infoln(str, err)
	}
}
func Float64FromBytes(b []byte) float64 {
	s := string(b)
	str := strings.Replace(s, "\n", "", 1)
	f, err := strconv.Atoi(str)
	CheckError("error []byte to float", err)
	return float64(f)
}

func Float64FromString(s string) float64 {
	str := strings.Replace(s, "\n", "", 1)
	f, err := strconv.Atoi(str)
	CheckError("error string to float", err)
	return float64(f)
}
func getUniquePorts(p []uint32) map[string]float64 {
	res := make(map[string]float64)
	ports := portsToInt(p)
	for _, port := range ports {
		str := fmt.Sprint(port)
		res[str] = float64(res[str] + 1)
	}
	return res
}

func portsToInt(ports []uint32) []int {
	res := make([]int, len(ports), len(ports))
	for i := 0; i < len(ports); i++ {
		res[i] = int(ports[i])
	}
	return res
}
