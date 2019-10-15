package mqops

import (
	"encoding/json"

	"github.com/matscus/Hamster/Guns/busM5/errors"
)

func init() {
	GetCardGeoRestrictions()
}

type getCardGeoRestrictionsJSON struct {
	Data struct {
		Pan string `json:"pan"`
	} `json:"data"`
}

//GetCardGeoRestrictions -  init script struct
func GetCardGeoRestrictions() {
	getCardGeoRestrictions := New()
	getCardGeoRestrictions.Name = "GetCardGeoRestrictions"
	getCardGeoRestrictions.PoolCh = NewCh(10)
	stringProperty := make(map[string]string)
	stringProperty["autorization"] = "Bearer"
	stringProperty["esfl_methodName"] = "getCardGeoRestrictions"
	stringProperty["src_systemID"] = "OMNI"
	stringProperty["src_channel"] = "test_channel"
	getCardGeoRestrictions.StringProperty = stringProperty
	go getCardGeoRestrictionsInfoJSONBody(getCardGeoRestrictions.PoolCh)
	MQOps = append(MQOps, getCardGeoRestrictions)
}

func getCardGeoRestrictionsInfoJSONBody(ch chan string) {
	var d getCardGeoRestrictionsJSON
	for {
		data := <-PoolCh
		d.Data.Pan = data.Pan
		e, err := json.Marshal(d)
		errors.CheckError(err, "Error marshal boby")
		ch <- string(e)
	}
}
