package mqops

import (
	"encoding/json"

	"../errors"
)

func init() {
	GetCardStateChangeInfo()
}

type getCardStateChangeInfoJSON struct {
	Data struct {
		Pan        string `json:"pan"`
		SystemCode string `json:"systemCode"`
	} `json:"data"`
}

//GetCardStateChangeInfo -  init script struct
func GetCardStateChangeInfo() {
	getCardStateChangeInfo := New()
	getCardStateChangeInfo.Name = "GetCardStateChangeInfo"
	getCardStateChangeInfo.PoolCh = NewCh(10)
	stringProperty := make(map[string]string)
	stringProperty["autorization"] = "Bearer"
	stringProperty["esfl_methodName"] = "getCardStateChangeInfo"
	stringProperty["src_systemID"] = "OMNI"
	stringProperty["src_channel"] = "test_channel"
	getCardStateChangeInfo.StringProperty = stringProperty
	go getCardStateChangeInfoJSONBody(getCardStateChangeInfo.PoolCh)
	MQOps = append(MQOps, getCardStateChangeInfo)
}

func getCardStateChangeInfoJSONBody(ch chan string) {
	var d getCardStateChangeInfoJSON
	for {
		data := <-PoolCh
		d.Data.Pan = data.Pan
		d.Data.SystemCode = data.SystemCode
		e, err := json.Marshal(d)
		errors.CheckError(err, "Error marshal boby")
		ch <- string(e)
	}
}
