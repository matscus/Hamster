package check

//CHDD - The structure of the building is data obtained from the use of hard disks, on all monitored hosts, in terms of mount points.
type CHDD struct {
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
