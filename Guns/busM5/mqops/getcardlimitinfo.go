package mqops

import (
	"encoding/json"

	"../errors"
)

func init() {
	GetCardLimitInfo()
}

type getCardLimitInfoJSON struct {
	Data struct {
		Pan string `json:"pan"`
	} `json:"data"`
}

//GetCardLimitInfo -  init script struct
func GetCardLimitInfo() {
	getCardLimitInfo := New()
	getCardLimitInfo.Name = "GetCardLimitInfo"
	getCardLimitInfo.PoolCh = NewCh(10)
	stringProperty := make(map[string]string)
	stringProperty["autorization"] = "Bearer"
	stringProperty["esfl_methodName"] = "getCardLimitInfo"
	stringProperty["src_systemID"] = "OMNI"
	stringProperty["src_channel"] = "test_channel"
	getCardLimitInfo.StringProperty = stringProperty
	go getCardLimitInfoJSONBody(getCardLimitInfo.PoolCh)
	MQOps = append(MQOps, getCardLimitInfo)
}

func getCardLimitInfoJSONBody(ch chan string) {
	var d getCardLimitInfoJSON
	for {
		data := <-PoolCh
		d.Data.Pan = data.Pan
		e, err := json.Marshal(d)
		errors.CheckError(err, "Error marshal boby")
		ch <- string(e)
	}
}
