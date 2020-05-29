package check

//Result - result structure, contains the stend status, slice monitoring services/host slice and their status.
type Result struct {
	Status    bool       `json:"status"`
	ServiceRS []ServerRS `json:"servicers"`
	//Monitoring []Monitoring `json:"Monitoring"`
	Hosts struct {
		PrometheusState bool   `json:"prometheusstate"`
		Нost            []Host `json:"hosts,omitempty"`
	}
}

//ServerRS - struct from response server status ad id
type ServerRS struct {
	ID     int64 `json:"id"`
	Status bool  `json:"status"`
}

//Host - struct for host info, contains host(IP addr),CPU value,HDD value,Memory value
type Host struct {
	Host   string `json:"host"`
	CPU    string `json:"CPU"`
	HDD    string `json:"HDD"`
	Memory string `json:"Memory"`
}

//Monitoring - struct for Monitoring info, contains name,host(IP addr) and launch status
type Monitoring struct {
	ID     int64  `json:"id"`
	Name   string `json:"name,omitempty"`
	Host   string `json:"host,omitempty"`
	Status bool   `json:"status,omitempty"`
}

//CheckCPU - data structure obtained from prometheus on CPU usage on hosts
type CheckCPU struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resulttype"`
		Result     []struct {
			Metric struct {
				Instance string `json:"instance"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

//CheckHDD - The structure of the building is data obtained from the use of hard disks, on all monitored hosts, in terms of mount points.
type CheckHDD struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Device     string `json:"device"`
				Fstype     string `json:"fstype"`
				Instance   string `json:"instance"`
				Job        string `json:"job"`
				Mountpoint string `json:"mountpoint"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

//CheckMemory - data structure obtained from prometheus on memory usage on hosts
type CheckMemory struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Instance string `json:"instance"`
				Job      string `json:"job"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}
