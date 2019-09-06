package mqops

import (
	"encoding/json"

	"../errors"
)

func init() {
	GetCardList()
}

type getCardListJSON struct {
	Data struct {
		Hid string `json:"hid"`
	} `json:"data"`
}

//GetCardList -  init script struct
func GetCardList() {
	getCardList := New()
	getCardList.Name = "GetCardList"
	getCardList.PoolCh = NewCh(10)
	stringProperty := make(map[string]string)

	stringProperty["esfl_methodName"] = "getCardList"
	stringProperty["autorization"] = "Bearer"
	stringProperty["src_systemID"] = "OMNI"
	stringProperty["src_channel"] = "test_channel"
	getCardList.StringProperty = stringProperty
	go getCardListJSONBody(getCardList.PoolCh)
	MQOps = append(MQOps, getCardList)
}

func getCardListJSONBody(ch chan string) {
	var d getCardListJSON
	for {
		data := <-PoolCh
		d.Data.Hid = data.Hid
		e, err := json.Marshal(d)
		errors.CheckError(err, "Error marshal boby")
		ch <- string(e)
	}
}
