package mqops

import (
	"encoding/json"

	"../errors"
)

func init() {
	GetCardCreditLimitHistory()
}

type getCardCreditLimitHistoryJSON struct {
	Data struct {
		Pan        string `json:"pan"`
		SystemCode string `json:"systemCode"`
	} `json:"data"`
}

//GetCardCreditLimitHistory -  init script struct
func GetCardCreditLimitHistory() {
	getCardCreditLimitHistory := New()
	getCardCreditLimitHistory.Name = "GetCardCreditLimitHistory"
	getCardCreditLimitHistory.PoolCh = NewCh(10)
	stringProperty := make(map[string]string)
	stringProperty["autorization"] = "Bearer"
	stringProperty["esfl_methodName"] = "getCardCreditLimitHistory"
	stringProperty["src_systemID"] = "OMNI"
	stringProperty["src_channel"] = "test_channel"
	getCardCreditLimitHistory.StringProperty = stringProperty
	go getCardCreditLimitHistoryJSONBody(getCardCreditLimitHistory.PoolCh)
	MQOps = append(MQOps, getCardCreditLimitHistory)
}

func getCardCreditLimitHistoryJSONBody(ch chan string) {
	var d getCardCreditLimitHistoryJSON
	for {
		data := <-PoolCh
		d.Data.Pan = data.Pan
		d.Data.SystemCode = data.SystemCode
		e, err := json.Marshal(d)
		errors.CheckError(err, "Error marshal boby")
		ch <- string(e)
	}
}
