package mqops

import (
	"encoding/json"

	"github.com/matscus/Hamster/Guns/busM5/errors"
)

func init() {
	GetDepositList()
}

type getDepositListJSON struct {
	Data struct {
		Hid string `json:"hid"`
	} `json:"data"`
}

//GetDepositList -  init script struct
func GetDepositList() {
	getDepositList := New()
	getDepositList.Name = "GetDepositList"
	getDepositList.PoolCh = NewCh(10)
	stringProperty := make(map[string]string)
	stringProperty["autorization"] = "Bearer"
	stringProperty["esfl_methodName"] = "getDepositList"
	stringProperty["src_systemID"] = "test_systemID"
	stringProperty["src_channel"] = "test_channel"
	getDepositList.StringProperty = stringProperty
	go getDepositListJSONBody(getDepositList.PoolCh)
	MQOps = append(MQOps, getDepositList)
}

func getDepositListJSONBody(ch chan string) {
	var d getDepositListJSON
	for {
		data := <-PoolCh
		d.Data.Hid = data.Hid
		e, err := json.Marshal(d)
		errors.CheckError(err, "Error marshal boby")
		ch <- string(e)
	}
}
