package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Mock/info_service/datapool"
)

type AccountListRQ struct {
	Meta struct {
		Channel string `json:"channel"`
	} `json:"meta"`
	Data struct {
		GUID string `json:"guid"`
	} `json:"data"`
}

type AccountListRS struct {
	Contracts []ContractAccount `json:"contracts"`
}

type Department struct {
	Address struct {
		FullAddress string `json:"fullAddress"`
	} `json:"address"`
	FullName string `json:"fullName"`
}

type ContractAccount struct {
	Account struct {
		Base struct {
			Balance []Balance `json:"balance"`
			Bank    struct {
				Bic string `json:"bic"`
			} `json:"bank"`
			Currency struct {
				CharCode string `json:"charCode"`
				NumCode  string `json:"numCode"`
			} `json:"currency"`
			IsArrest bool   `json:"isArrest"`
			Number   string `json:"number"`
			State    struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"state"`
			SystemInfo struct {
				RawID    string `json:"rawId"`
				SystemID string `json:"systemId"`
			} `json:"systemInfo"`
			Topup struct {
				IsPosibility bool   `json:"isPosibility"`
				MinAmount    string `json:"minAmount"`
			} `json:"topup"`
			Type       string `json:"type"`
			Withdrawal struct {
				IsPosibility bool `json:"isPosibility"`
			} `json:"withdrawal"`
		} `json:"base"`
	} `json:"account"`
	Base struct {
		Bank struct {
			Bic         string       `json:"bic"`
			Departments []Department `json:"departments"`
			FullName    string       `json:"fullName"`
		} `json:"bank"`
		ClientRequest struct {
			RawID    string `json:"rawId"`
			SystemID string `json:"systemId"`
		} `json:"clientRequest"`
		CloseDateFact    string `json:"closeDateFact"`
		CloseDatePlan    string `json:"closeDatePlan"`
		CreationDate     string `json:"creationDate"`
		EmployeeFullName string `json:"employeeFullName"`
		IsWorking        bool   `json:"isWorking"`
		Number           string `json:"number"`
		Owner            struct {
			FullName string `json:"fullName"`
			RawID    string `json:"rawId"`
			SystemID string `json:"systemId"`
		} `json:"owner"`
		Product struct {
			Code string `json:"code"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"product"`
		StartDate string `json:"startDate"`
		State     struct {
			Code string `json:"code"`
			Date string `json:"date"`
			Name string `json:"name"`
		} `json:"state"`
		SystemInfo struct {
			RawID    string `json:"rawId"`
			SystemID string `json:"systemId"`
		} `json:"systemInfo"`
	} `json:"base"`
}

func AccountList(w http.ResponseWriter, r *http.Request) {
	rq := AccountListRQ{}
	err := json.NewDecoder(r.Body).Decode(&rq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	client := datapool.GUIDPool[rq.Data.GUID]
	rs := AccountListRS{}

	contractAccount := ContractAccount{}
	contractAccount.Base.ClientRequest.SystemID = "DR36"
	contractAccount.Base.ClientRequest.RawID = client.UserID
	contractAccount.Base.SystemInfo.SystemID = "DR36"
	contractAccount.Base.SystemInfo.RawID = "10057524431"
	contractAccount.Base.Owner.RawID = client.UserID
	contractAccount.Base.Owner.SystemID = "DR36"
	contractAccount.Base.Owner.FullName = "ПППП ФФФФ ББББ"
	contractAccount.Base.Number = client.ContractNum
	contractAccount.Base.Product.Code = "ГПБТекущ"
	contractAccount.Base.Product.Name = "Банковский (текущий) счет"
	contractAccount.Base.Product.Type = "3.0"
	contractAccount.Base.CreationDate = "2019-10-23"
	contractAccount.Base.StartDate = "2019-10-23"
	contractAccount.Base.CloseDatePlan = "2019-10-23"
	contractAccount.Base.CloseDateFact = ""
	contractAccount.Base.State.Code = "Оформлен"
	contractAccount.Base.State.Name = "Оформлен"
	contractAccount.Base.State.Date = "2019-10-23"
	contractAccount.Base.IsWorking = true
	contractAccount.Base.Bank.Bic = "044525823"
	contractAccount.Base.Bank.FullName = "\"Газпромбанк\" (Акционерное общество)"
	department := Department{}
	department.FullName = "015/0001"
	department.Address.FullAddress = "РОССИЙСКАЯ ФЕДЕРАЦИЯ, 191124, Санкт-Петербург г, 1231231231 ул,  д. 333, лит. А"
	contractAccount.Base.Bank.Departments = append(contractAccount.Base.Bank.Departments, department)
	contractAccount.Base.EmployeeFullName = "Авввввв Авввввв Авввввв"
	contractAccount.Account.Base.SystemInfo.SystemID = "DR36"
	contractAccount.Account.Base.SystemInfo.RawID = "10095094495"

	contractAccount.Account.Base.Number = client.AccNum
	contractAccount.Account.Base.Bank.Bic = "444555666"
	contractAccount.Account.Base.Currency.NumCode = "810"
	contractAccount.Account.Base.Currency.CharCode = "RUR"

	contractAccount.Account.Base.Type = "depositAccount"
	contractAccount.Account.Base.State.Code = "open"
	contractAccount.Account.Base.State.Name = "Открыт"
	balance := Balance{}
	balance.Type = "mainAmount"
	balance.Value.Amount = "1.9955E8"
	balance.Value.Currency.NumCode = "810"
	balance.Value.Currency.CharCode = "RUR"
	contractAccount.Account.Base.Balance = append(contractAccount.Account.Base.Balance, balance)
	contractAccount.Account.Base.IsArrest = false
	contractAccount.Account.Base.Topup.IsPosibility = true
	contractAccount.Account.Base.Topup.MinAmount = "0.0"
	contractAccount.Account.Base.Withdrawal.IsPosibility = true

	rs.Contracts = append(rs.Contracts, contractAccount)

	err = json.NewEncoder(w).Encode(rs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter  due: %s", errWrite.Error())
		}
	}
}
