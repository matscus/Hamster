package check

//Result - result structure, contains the stend status, slice monitoring services/host slice and their status.
type Result struct {
	Status    bool    `json:"Status"`
	ServiceID []int64 `json:"serviceid"`
	//Monitoring []Monitoring `json:"Monitoring"`
	Host []Host `json:"Host"`
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
