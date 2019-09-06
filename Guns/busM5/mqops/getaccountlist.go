package mqops

import (
	"encoding/json"

	"../errors"
)

func init() {
	GetAccountList()
}

type getAccountListJSON struct {
	Data struct {
		Hid string `json:"hid"`
	} `json:"data"`
}

//GetAccountList -  init script struct
func GetAccountList() {
	getAccountList := New()
	getAccountList.Name = "GetAccountList"
	getAccountList.PoolCh = NewCh(10)
	stringProperty := make(map[string]string)
	stringProperty["autorization"] = "Bearer"
	stringProperty["esfl_methodName"] = "getAccountList"
	stringProperty["src_systemID"] = "OMNI"
	stringProperty["src_channel"] = "test_channel"
	getAccountList.StringProperty = stringProperty
	go getAccountListJSONBody(getAccountList.PoolCh)
	MQOps = append(MQOps, getAccountList)
}

func getAccountListJSONBody(ch chan string) {
	var d getAccountListJSON
	for {
		data := <-PoolCh
		d.Data.Hid = data.Hid
		e, err := json.Marshal(d)
		errors.CheckError(err, "Error marshal boby")
		ch <- string(e)
	}
}
