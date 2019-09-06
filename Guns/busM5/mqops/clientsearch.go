package mqops

import (
	"encoding/json"

	"../errors"
)

type clientSearchJSON struct {
	Data struct {
		RequestFields []string `json:"requestFields"`
		Filter        struct {
			Surname    string `json:"surname"`
			Name       string `json:"name"`
			Patronymic string `json:"patronymic"`
			Documents  struct {
				Type      string `json:"type"`
				FullValue string `json:"fullValue"`
			} `json:"documents"`
			Birthdate string `json:"birthdate"`
			Phone     struct {
				Type   string `json:"type"`
				Number string `json:"number"`
			} `json:"phone"`
			Snils          string `json:"snils"`
			Inn            string `json:"inn"`
			SystemID       string `json:"systemId"`
			RawID          string `json:"rawId"`
			Hid            string `json:"hid"`
			ContractNumber string `json:"contractNumber"`
			AccountNumber  string `json:"accountNumber"`
			CardNumber     string `json:"cardNumber"`
			Mail           string `json:"mail"`
		} `json:"filter"`
	} `json:"data"`
}

func init() {
	ClientSearch()
}

//ClientSearch init script struct
func ClientSearch() {
	clientSearch := New()
	clientSearch.Name = "ClientSearch"
	clientSearch.PoolCh = NewCh(10)
	stringProperty := make(map[string]string)
	stringProperty["autorization"] = "Bearer"
	stringProperty["esfl_methodName"] = "clientSearch"
	stringProperty["src_systemID"] = "OMNI"
	stringProperty["src_channel"] = "test_channel"
	clientSearch.StringProperty = stringProperty
	go getClientSearchJSONBody(clientSearch.PoolCh)
	MQOps = append(MQOps, clientSearch)
}

func getClientSearchJSONBody(ch chan string) {
	var d clientSearchJSON
	for {
		data := <-PoolCh
		d.Data.RequestFields = []string{"base", "addresses", "documents", "phones", "mails", "sources", "pastdocuments"}
		d.Data.Filter.Surname = data.Surname
		d.Data.Filter.Name = data.Surname
		d.Data.Filter.Patronymic = data.Surname
		d.Data.Filter.Phone.Number = data.Phone
		e, err := json.Marshal(d)
		errors.CheckError(err, "Error marshal boby")
		ch <- string(e)
	}
}
