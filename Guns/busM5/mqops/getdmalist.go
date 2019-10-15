package mqops

import (
	"encoding/json"

	"github.com/matscus/Hamster/Guns/busM5/errors"
)

func init() {
	GetDMAList()
}

type getDMAListJSON struct {
	Data struct {
		Hid string `json:"hid"`
	} `json:"data"`
}

//GetDMAList -  init script struct
func GetDMAList() {
	getDMAList := New()
	getDMAList.Name = "GetDMAList"
	getDMAList.PoolCh = NewCh(10)
	stringProperty := make(map[string]string)
	stringProperty["autorization"] = "Bearer"
	stringProperty["esfl_methodName"] = "getDMAList"
	stringProperty["src_systemID"] = "OMNI"
	stringProperty["src_channel"] = "test_channel"
	getDMAList.StringProperty = stringProperty
	go getDMAListJSONBody(getDMAList.PoolCh)
	MQOps = append(MQOps, getDMAList)
}

func getDMAListJSONBody(ch chan string) {
	var d getDMAListJSON
	for {
		data := <-PoolCh
		d.Data.Hid = data.Hid
		e, err := json.Marshal(d)
		errors.CheckError(err, "Error marshal boby")
		ch <- string(e)
	}
}
