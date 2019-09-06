package mqops

import (
	"encoding/json"

	"../errors"
)

type getDepositInfoJSON struct {
	Data struct {
		SystemCode string `json:"systemCode"`
		ContractID string `json:"contractId"`
	} `json:"data"`
}

func init() {
	GetDepositInfo()
}

//GetDepositInfo -  init script struct
func GetDepositInfo() {
	getDepositInfo := New()
	getDepositInfo.Name = "GetDepositInfo"
	getDepositInfo.PoolCh = NewCh(10)
	stringProperty := make(map[string]string)
	stringProperty["autorization"] = "Bearer"
	stringProperty["esfl_methodName"] = "getDepositInfo"
	stringProperty["src_systemID"] = "OMNI"
	stringProperty["src_channel"] = "test_channel"
	getDepositInfo.StringProperty = stringProperty
	go getDepositInfoJSONBody(getDepositInfo.PoolCh)
	MQOps = append(MQOps, getDepositInfo)
}

func getDepositInfoJSONBody(ch chan string) {
	var d getDepositInfoJSON
	for {
		data := <-PoolCh
		d.Data.ContractID = data.ContractID
		d.Data.SystemCode = data.SystemCode
		e, err := json.Marshal(d)
		errors.CheckError(err, "Error marshal boby")
		ch <- string(e)
	}
}
