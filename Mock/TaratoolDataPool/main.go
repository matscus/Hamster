package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	uri     string
	token   string
	poolLen int
)

type Req struct {
	Query string `json:"query"`
}
type Res struct {
	Data struct {
		CardContractsBase []struct {
			ClientRequestRawID   string `json:"ClientRequestRawId"`
			PAN                  string `json:"PAN"`
			SystemInfoContractID string `json:"SystemInfoContractId"`
			SystemInfoSystemID   string `json:"SystemInfoSystemId"`
		} `json:"cardContractsBase"`
	} `json:"data"`
}

func main() {
	flag.StringVar(&uri, "uri", "https://10.178.49.176:433/graphql", "uri")
	flag.StringVar(&token, "token", "", "tarantool token")
	flag.IntVar(&poolLen, "poollen", 1000, "len tarantool pool")
	fileCardlist, err := os.Create("cardlist.csv")
	if err != nil {
		log.Println(err)
	}
	fileCardlist.WriteString("system_id,")
	fileCardlist.WriteString("raw_id\n")
	filePans, err := os.Create("pans.csv")
	filePans.WriteString("pan\n")
	if err != nil {
		log.Println(err)
	}
	filePans.WriteString("pan")
	resSlice := make([]Res, 0, poolLen)
	dr := []string{"DR01", "DR02", "DR04", "DR25", "DR29", "DR36", "DR37", "DR38", "DR24"}
	lenPool := poolLen / len(dr)
	client := &http.Client{}
	for _, v := range dr {
		req, err := http.NewRequest("POST", uri, bytes.NewBuffer([]byte(`
		{"query": "query {cardContractsBase(first: `+strconv.Itoa(lenPool)+`, SystemIdContractId_le: [\"`+v+`\",\"20030000000\"]) {
			SystemInfoContractId
			SystemInfoSystemId
			ClientRequestRawId
			PAN
		  }
		  }"
		  }`)))
		if err != nil {
			log.Println(err)
		}
		defer req.Body.Close()
		req.Header.Set("auth-token", token)
		resp, err := client.Do(req)
		res := Res{}
		err = json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			log.Println(err)
		}
		resSlice = append(resSlice, res)
	}
	for _, v := range resSlice {
		for _, vv := range v.Data.CardContractsBase {
			fileCardlist.WriteString(vv.SystemInfoSystemID + "," + vv.ClientRequestRawID + "\n")
			filePans.WriteString(vv.PAN + "\n")
		}
	}
}
