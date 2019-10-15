package mqops

import (
	"encoding/json"

	"github.com/matscus/Hamster/Guns/busM5/errors"
)

func init() {
	GetCardInfo()
}

type getCardInfoJSON struct {
	Data struct {
		Pan        string `json:"pan"`
		SystemCode string `json:"systemCode"`
	} `json:"data"`
}

//GetCardInfo -  init script struct
func GetCardInfo() {
	getCardInfo := New()
	getCardInfo.Name = "GetCardInfo"
	getCardInfo.PoolCh = NewCh(10)
	stringProperty := make(map[string]string)
	stringProperty["autorization"] = "Bearer"
	stringProperty["esfl_methodName"] = "getCardInfo"
	stringProperty["src_systemID"] = "OMNI"
	stringProperty["src_channel"] = "test_channel"
	getCardInfo.StringProperty = stringProperty
	go getCardInfoJSONBody(getCardInfo.PoolCh)
	MQOps = append(MQOps, getCardInfo)
}

func getCardInfoJSONBody(ch chan string) {
	var d getCardInfoJSON
	for {
		data := <-PoolCh
		d.Data.Pan = data.Pan
		d.Data.SystemCode = data.SystemCode
		e, err := json.Marshal(d)
		errors.CheckError(err, "Error marshal boby")
		ch <- string(e)
	}
}
