package mqops

import (
	"encoding/json"
	"os"

	"../errors"
	"github.com/matscus/Hamster/Package/Datapool/pools"
	"github.com/matscus/Hamster/Package/Datapool/structs"
)

var (
	//MQConf global config struct
	MQConf  config
	poollen = 1000
	//PoolCh - chan for all operations
	PoolCh = make(chan structs.DefaultDatapool, poollen)
)

type config struct {
	Data struct {
		Defaultproduser   int `json:"defaultproduser"`
		Defaultconsumer   int `json:"defaultconsumer"`
		Defaultreceiver   int `json:"defaultreceiver"`
		Defaultstep       int `json:"defaultstep"`
		Defaultrumpup     int `json:"defaultrumpup"`
		Defaultthroughput int `json:"defaultthroughput"`
		Operation         []struct {
			Data struct {
				Name          string `json:"name"`
				Defaultparams bool   `json:"defaultparams"`
				Produser      int    `json:"produser"`
				Consumer      int    `json:"consumer"`
				Receiver      int    `json:"receiver"`
				Step          int    `json:"step"`
				Rumpup        int    `json:"rumpup"`
				Throughput    int    `json:"throughput"`
				Host          string `json:"host"`
				Port          int    `json:"port"`
				Manager       string `json:"manager"`
				Cannel        string `json:"cannel"`
				UserName      string `json:"userName"`
				Password      string `json:"password"`
				Queuein       string `json:"queuein"`
				Queueout      string `json:"queueout"`
			} `json:"data"`
		} `json:"Operation"`
	} `json:"Data"`
}

//InitConfigs -  init MQConf
func InitConfigs(configpath string) {
	go func() {
		err := pools.Datapool{}.New().GetDefaultDatapool(poollen, PoolCh)
		if err != nil {
			errors.CheckError(err, "error init pool: ")
		}
	}()
	fl, err := os.Open(configpath)
	errors.CheckError(err, "error open config: ")
	defer fl.Close()
	err = json.NewDecoder(fl).Decode(&MQConf)
	errors.CheckError(err, "error decode config: ")
}
