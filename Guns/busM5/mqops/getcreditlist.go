package mqops

import (
	"encoding/json"

	"github.com/matscus/Hamster/Guns/busM5/errors"
)

func init() {
	GetCreditList()
}

type getCreditListJSON struct {
	Data struct {
		Hid string `json:"hid"`
	} `json:"data"`
}

//GetCreditList -  init script struct
func GetCreditList() {
	getCreditList := New()
	getCreditList.Name = "GetCreditList"
	getCreditList.PoolCh = NewCh(10)
	stringProperty := make(map[string]string)
	stringProperty["autorization"] = "Bearer"
	stringProperty["esfl_methodName"] = "getCreditList"
	stringProperty["src_systemID"] = "OMNI"
	stringProperty["src_channel"] = "test_channel"
	getCreditList.StringProperty = stringProperty
	go getCreditListJSONBody(getCreditList.PoolCh)
	MQOps = append(MQOps, getCreditList)
}

func getCreditListJSONBody(ch chan string) {
	var d getCreditListJSON
	for {
		data := <-PoolCh
		d.Data.Hid = data.Hid
		e, err := json.Marshal(d)
		errors.CheckError(err, "Error marshal boby")
		ch <- string(e)
	}
}
