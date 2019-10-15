package mqops

import (
	"encoding/json"

	"github.com/matscus/Hamster/Guns/busM5/errors"
)

func init() {
	GetCreditSchedule()
}

type getCredidScheduleJSON struct {
	Data struct {
		SystemCode string `json:"systemCode"`
		ContractID string `json:"contractId"`
	} `json:"data"`
}

//GetCreditSchedule -  init script struct
func GetCreditSchedule() {
	getCredidSchedule := New()
	getCredidSchedule.Name = "GetCreditSchedule"
	getCredidSchedule.PoolCh = NewCh(10)
	stringProperty := make(map[string]string)
	stringProperty["autorization"] = "Bearer"
	stringProperty["esfl_methodName"] = "getCreditSchedule"
	stringProperty["src_systemID"] = "OMNI"
	stringProperty["src_channel"] = "test_channel"
	getCredidSchedule.StringProperty = stringProperty
	go getCredidScheduleJSONBody(getCredidSchedule.PoolCh)
	MQOps = append(MQOps, getCredidSchedule)
}

func getCredidScheduleJSONBody(ch chan string) {
	var d getCredidScheduleJSON
	for {
		data := <-PoolCh
		d.Data.ContractID = data.ContractID
		d.Data.SystemCode = data.SystemCode
		e, err := json.Marshal(d)
		errors.CheckError(err, "Error marshal boby")
		ch <- string(e)
	}
}
