package check

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Services/service"
)

//InitGetResponseAllData - function to obtain information about all services from the database. all services are append to slice Services
func InitGetResponseAllData(project string) (responsealldata *[]service.Service, err error) {
	pgclient := client.PGClient{}.New()
	services, err := pgclient.GetProjectServices(project)
	if err != nil {
		return nil, err
	}
	l := len(*services)
	getResponceAllData := make([]service.Service, l, l)
	for i := 0; i < l; i++ {
		var s service.Service
		t := (*services)[i]
		s.ID = t.ID
		s.Name = t.Name
		s.Host = t.Host
		s.URI = t.URI
		s.Type = t.Type
		s.Projects = t.Projects
		getResponceAllData[i] = s
	}
	return &getResponceAllData, nil
}

//CheckStend - function to check stehd, checks the status of monitoring agents, memory utilization, hard disks and processors.
func CheckStend(getResponceAllData *[]service.Service) (res Result, err error) {
	prometheusstate := true
	temp := make(map[string]Host)
	l := len(*getResponceAllData)
	checkhdd := CheckHDD{}
	checkcpu := CheckCPU{}
	checkmem := CheckMemory{}
	client, err := client.NewHTTPSClient()
	for i := 0; i < l; i++ {
		id := (*getResponceAllData)[i].ID
		u := (*getResponceAllData)[i].URI
		_, err = client.Get(u)
		if err != nil {
			res.ServiceID = append(res.ServiceID, id)
			if (*getResponceAllData)[i].Name == "prometheus" {
				prometheusstate = false
			}
		}
	}
	if prometheusstate {
		responsefs, err := http.Get(os.Getenv("PROMETHEUSURI") + "?query=node_filesystem_avail_bytes/node_filesystem_size_bytes*100")
		if err != nil {
			err = errors.New("error Get responsefs: %s" + err.Error())
		}
		responsecpu, err := http.Get(os.Getenv("PROMETHEUSURI") + "?query=avg%20by(instance)(max_over_time(node_cpu_seconds_total{mode!=\"idle\"}[5m])-(min_over_time(node_cpu_seconds_total{mode!=\"idle\"}[5m])))")
		if err != nil {
			err = errors.New("error Get responsecpu: %s" + err.Error())
		}
		responsemem, err := http.Get(os.Getenv("PROMETHEUSURI") + "?query=node_memory_MemAvailable_bytes/node_memory_MemTotal_bytes*100")
		if err != nil {
			err = errors.New("error Get responsemem: %s" + err.Error())
		}
		defer responsefs.Body.Close()
		defer responsecpu.Body.Close()
		defer responsemem.Body.Close()
		contentsfs, _ := ioutil.ReadAll(responsefs.Body)
		contentscpu, _ := ioutil.ReadAll(responsecpu.Body)
		contentsmem, _ := ioutil.ReadAll(responsemem.Body)
		err = json.Unmarshal(contentsfs, &checkhdd)
		if err != nil {
			err = errors.New("error unmarshal contentsfs: %s" + err.Error())
		}
		err = json.Unmarshal(contentscpu, &checkcpu)
		if err != nil {
			err = errors.New("error unmarshal contentscpu: %s" + err.Error())
		}
		err = json.Unmarshal(contentsmem, &checkmem)
		if err != nil {
			err = errors.New("error unmarshal contentsmem: %s" + err.Error())
		}
		for i := 0; i < len(checkcpu.Data.Result); i++ {
			var h Host
			tt := fmt.Sprint(checkcpu.Data.Result[i].Value[1])
			v, _ := strconv.ParseFloat(tt, 64)
			if v >= 70 {
				h.Host = checkcpu.Data.Result[i].Metric.Instance
				h.CPU = "cpu is used over 70%\n"
				temp[checkcpu.Data.Result[i].Metric.Instance] = h
			}
		}
		for i := 0; i < len(checkhdd.Data.Result); i++ {
			var h Host
			tt := fmt.Sprint(checkhdd.Data.Result[i].Value[1])
			v, _ := strconv.ParseFloat(tt, 64)
			if v <= 10 {
				if v, ok := temp[checkhdd.Data.Result[i].Metric.Instance]; ok {
					h.Host = v.Host
					h.CPU = v.CPU
					h.HDD = v.HDD + "free space in mountpoint " + checkhdd.Data.Result[i].Metric.Mountpoint + " <10%\n"
					temp[checkhdd.Data.Result[i].Metric.Instance] = h
				} else {
					h.Host = checkhdd.Data.Result[i].Metric.Instance
					h.HDD = "free space in mountpoint " + checkhdd.Data.Result[i].Metric.Mountpoint + " <10%\n"
					temp[checkhdd.Data.Result[i].Metric.Instance] = h
				}
			}
		}
		for i := 0; i < len(checkmem.Data.Result); i++ {
			var h Host
			tt := fmt.Sprint(checkmem.Data.Result[i].Value[1])
			v, _ := strconv.ParseFloat(tt, 64)
			if v <= 80 {
				if v, ok := temp[checkmem.Data.Result[i].Metric.Instance]; ok {
					h.Host = v.Host
					h.CPU = v.CPU
					h.HDD = v.HDD
					h.Memory = v.Memory + "memory is used over 20%"
					temp[checkmem.Data.Result[i].Metric.Instance] = h
				} else {
					h.Host = checkmem.Data.Result[i].Metric.Instance
					h.Memory = "memory is used over 20%"
					temp[checkmem.Data.Result[i].Metric.Instance] = h
				}
			}
		}
		for _, v := range temp {
			res.Host = append(res.Host, v)
		}
	} else {
		host := Host{}
		res.Host = append(res.Host, host)
		err = errors.New("Prometheus is not avalible")
	}
	return res, err
}
