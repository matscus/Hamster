package mqops

import (
	"encoding/json"

	"github.com/matscus/Hamster/Guns/busM5/errors"
)

func init() {
	GetCardMobilePayInfo()
}

type getCardMobilePayInfoJSON struct {
	Data struct {
		Pan string `json:"pan"`
	} `json:"data"`
}

//GetCardMobilePayInfo -  init script struct
func GetCardMobilePayInfo() {
	getCardMobilePayInfo := New()
	getCardMobilePayInfo.Name = "GetCardMobilePayInfo"
	getCardMobilePayInfo.PoolCh = NewCh(10)
	stringProperty := make(map[string]string)
	stringProperty["autorization"] = "Bearer"
	stringProperty["esfl_methodName"] = "getCardMobilePayInfo"
	stringProperty["src_systemID"] = "OMNI"
	stringProperty["src_channel"] = "test_channel"
	getCardMobilePayInfo.StringProperty = stringProperty
	go getCardMobilePayInfoJSONBody(getCardMobilePayInfo.PoolCh)
	MQOps = append(MQOps, getCardMobilePayInfo)
}

func getCardMobilePayInfoJSONBody(ch chan string) {
	var d getCardMobilePayInfoJSON
	for {
		data := <-PoolCh
		d.Data.Pan = data.Pan
		e, err := json.Marshal(d)
		errors.CheckError(err, "Error marshal boby")
		ch <- string(e)
	}
}
