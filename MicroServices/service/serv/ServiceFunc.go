package serv

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Services/service"
)

var (
	//GetResponceAllData - slice services
	GetResponseAllData []service.Service
	//HostsAndUsers - sync map users from remote host
	HostsAndUsers sync.Map
)

func init() {
	pgclient := client.PGClient{}.New()
	hostsAndUsers, err := pgclient.GetUsersAndHosts()
	if err != nil {
		log.Println("[ERR] Error get users and hosts: ", err)
	}
	for k, v := range hostsAndUsers {
		HostsAndUsers.Store(k, v)
	}
}

//InitGetResponseAllData - function to obtain information about all services from the database. all services are append to slice Services
func InitGetResponseAllData() error {
	pgclient := client.PGClient{}.New()
	for {
		services, err := pgclient.GetAllServices()
		if err != nil {
			return err
		}
		l := len(*services)
		GetResponseAllData = make([]service.Service, l, l)
		for i := 0; i < l; i++ {
			var s service.Service
			t := (*services)[i]
			s.ID = t.ID
			s.Name = t.Name
			s.Host = t.Host
			s.URI = t.URI
			s.Type = t.Type
			s.Projects = t.Projects
			GetResponseAllData[i] = s
		}
		time.Sleep(1 * time.Minute)
	}
}

//CheckService - The function to check the status of web services.
//Performs a GET request at the of each service. if available, sets the service in the status Run.
func CheckService() {
	var err error
	l := len(GetResponseAllData)
	for i := 0; i < l; i++ {
		_, err = http.Get(GetResponseAllData[i].URI)
		if err != nil {
			GetResponseAllData[i].Status = "stop"
		} else {
			GetResponseAllData[i].Status = "run"
		}
	}
}

//UpdateState - update service info
func UpdateState(s *service.Service) {
	l := len(GetResponseAllData)
	for i := 0; i < l; i++ {
		if s.Name == GetResponseAllData[i].Name && s.Host == GetResponseAllData[i].Host {
			GetResponseAllData[i].Status = s.Status
			break
		}
	}
}

//GetSTR - returt string for service run
func GetSTR(name string) (str string) {
	switch name {
	case "prometheus":
		str = "nohup " + os.Getenv("PROMETHEUSDIR") + "./prometheus --web.listen-address=" + os.Getenv("PROMETHEUSPORT") + " --config.file=" + os.Getenv("PROMETHEUSDIR") + "prometheus.yml &> /dev/null"
	case "grafana":
		str = "nohup " + os.Getenv("GRAFANADIR") + "./grafana-server -config=" + os.Getenv("GRAFANADIR") + "/conf/default.ini &> /dev/null"
	case "influxd":
		str = "nohup " + os.Getenv("INFLUXDIR") + "./influxd &> /dev/null"
	case "postgres_exporter":
		str = "nohup " + os.Getenv("POSTGRESEXPORTERDIR") + "./postgres_exporter --web.listen-address=" + os.Getenv("POSTGRESEXPORTERPORT") + " -extend.query-path=" + os.Getenv("POSTGRESEXPORTERCONF") + " &> /dev/null"
	case "node_exporter":
		str = "nohup " + os.Getenv("NODEEXPORTERDIR") + "./node_exporter --web.listen-address=" + os.Getenv("NODEEXPORTERPORT") + " &> /dev/null"
	case "rabbit_exporter":
		str = "nohup " + os.Getenv("RABBITEXPORTERDIR") + "./rabbit_exporter &> /dev/null"
	case "5NT_mock":
		str = "nohup " + os.Getenv("5NTMOCKRDIR") + "./.5NT &> /dev/null"
	case "mdm_mock":
		str = "nohup " + os.Getenv("MDMMOCKRDIR") + "./.mdm &> /dev/null"
	case "zapk_mock":
		str = "nohup " + os.Getenv("ZAPKMOCKRDIR") + "./.zapk &> /dev/null"
	case "graphite_exporter":
		str = "nohup " + os.Getenv("GRAPHITEEXPORTERDIR") + "./.graphite_exporter &> /dev/null"
	default:
		str = ""
	}
	return str
}
