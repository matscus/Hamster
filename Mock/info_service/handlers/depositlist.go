package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/matscus/Hamster/Mock/info_service/datapool"
)

type DepositListRQ struct {
	Meta struct {
		Channel string `json:"channel"`
	} `json:"meta"`
	Data struct {
		GUID string `json:"guid"`
	} `json:"data"`
}
type DepositListRS struct {
	Status          string `json:"status"`
	ActualTimestamp int64  `json:"actualTimestamp"`
	Data            struct {
		Contracts []Contract `json:"contracts"`
	} `json:"data"`
}
type Contract struct {
	Base struct {
		ClientRequest struct {
			SystemID string `json:"systemId"`
			RawID    string `json:"rawId"`
		} `json:"clientRequest"`
		SystemInfo struct {
			SystemID string `json:"systemId"`
			RawID    string `json:"rawId"`
		} `json:"systemInfo"`
		Owner struct {
			SystemID string `json:"systemId"`
			RawID    string `json:"rawId"`
			FullName string `json:"fullName"`
		} `json:"owner"`
		Number  string `json:"number"`
		Product struct {
			Code string `json:"code"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"product"`
		CreationDate  string `json:"creationDate"`
		StartDate     string `json:"startDate"`
		CloseDatePlan string `json:"closeDatePlan"`
		CloseDateFact string `json:"closeDateFact"`
		State         struct {
			Code string `json:"code"`
			Name string `json:"name"`
			Date string `json:"date"`
		} `json:"state"`
		Bank struct {
			Bic         string       `json:"bic"`
			FullName    string       `json:"fullName"`
			Departments []Department `json:"departments"`
		} `json:"bank"`
		EmployeeFullName string `json:"employeeFullName"`
		IsWorking        bool   `json:"isWorking"`
	} `json:"base"`
	Deposits []Deposit `json:"deposits"`
}
type Deposit struct {
	Base struct {
		SystemInfo struct {
			SystemID string `json:"systemId"`
			RawID    string `json:"rawId"`
		} `json:"systemInfo"`
		Term struct {
			Unit  string `json:"unit"`
			Value int    `json:"value"`
		} `json:"term"`
		MainAccount struct {
			Base struct {
				SystemInfo struct {
					SystemID string `json:"systemId"`
					RawID    string `json:"rawId"`
				} `json:"systemInfo"`
				Number string `json:"number"`
				Bank   struct {
					Bic string `json:"bic"`
				} `json:"bank"`
				Currency struct {
					NumCode  string `json:"numCode"`
					CharCode string `json:"charCode"`
				} `json:"currency"`
				Type  string `json:"type"`
				State struct {
					Code string `json:"code"`
					Name string `json:"name"`
				} `json:"state"`
				Balance []Balance `json:"balance"`
				Topup   struct {
					IsPosibility bool `json:"isPosibility"`
				} `json:"topup"`
				Withdrawal struct {
					IsPosibility bool `json:"isPosibility"`
				} `json:"withdrawal"`
				IsArrest bool `json:"isArrest"`
			} `json:"base"`
		} `json:"mainAccount"`
		Rate struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"rate"`
		IsMultiCurrency bool `json:"isMultiCurrency"`
	} `json:"base"`
}

func DepositList(w http.ResponseWriter, r *http.Request) {
	rq := DepositListRQ{}
	// log.Println("DepositListRQ = " + rq.Data.GUID)
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
	rs := DepositListRS{}

	contract := Contract{}
	contract.Base.ClientRequest.SystemID = "DRTL"
	contract.Base.ClientRequest.RawID = client.UserID
	contract.Base.SystemInfo.SystemID = "DRTL"
	contract.Base.SystemInfo.RawID = client.DealID
	contract.Base.Owner.SystemID = "DRTL"
	contract.Base.Owner.RawID = client.UserID
	contract.Base.Owner.FullName = "ААААА КРИСТИНА ВИКТОРОВНА"
	contract.Base.Number = client.ContractNum
	contract.Base.Product.Code = "ДЕПОЗ_ВУП"
	contract.Base.Product.Name = "Ваш успех"
	contract.Base.Product.Type = "2.0"
	contract.Base.CreationDate = "2020-01-13"
	contract.Base.StartDate = "2020-03-13"
	contract.Base.CloseDatePlan = "2020-09-10"
	contract.Base.CloseDateFact = "1900-01-01"
	contract.Base.State.Code = "Введен"
	contract.Base.State.Name = "Введен"
	contract.Base.State.Date = "2020-03-13"
	contract.Base.Bank.Bic = "044525823"
	contract.Base.Bank.FullName = "\"Газпромбанк\" (Акционерное общество)"
	department := Department{}
	department.FullName = "028/1003"
	department.Address.FullAddress = "РОССИЙСКАЯ ФЕДЕРАЦИЯ, Московская обл, Балашихинский р-н, Балашиха г, Ленина пр-кт,  д. ФФФ"
	contract.Base.Bank.Departments = append(contract.Base.Bank.Departments, department)
	contract.Base.EmployeeFullName = "LOADER5NT"
	contract.Base.IsWorking = true

	deposit := Deposit{}
	deposit.Base.SystemInfo.RawID = client.DealID
	deposit.Base.SystemInfo.SystemID = "DRTL"
	deposit.Base.Term.Unit = "day"
	deposit.Base.Term.Value = 181
	deposit.Base.MainAccount.Base.SystemInfo.RawID = client.DealID
	deposit.Base.MainAccount.Base.SystemInfo.SystemID = "DRTL"
	deposit.Base.MainAccount.Base.Number = client.AccNum
	deposit.Base.MainAccount.Base.Bank.Bic = "044525823"
	deposit.Base.MainAccount.Base.Currency.NumCode = "810"
	deposit.Base.MainAccount.Base.Currency.CharCode = "RUR"
	deposit.Base.MainAccount.Base.Type = "depositAccount"
	deposit.Base.MainAccount.Base.State.Code = "open"
	deposit.Base.MainAccount.Base.State.Name = "Открыт"
	balance := Balance{}
	balance.Type = "mainAmount"
	balance.Value.Amount = "1.9955E8"
	balance.Value.Currency.NumCode = "810"
	balance.Value.Currency.CharCode = "RUR"
	deposit.Base.MainAccount.Base.Balance = append(deposit.Base.MainAccount.Base.Balance, balance)
	deposit.Base.MainAccount.Base.Topup.IsPosibility = true
	deposit.Base.MainAccount.Base.Withdrawal.IsPosibility = false
	deposit.Base.MainAccount.Base.IsArrest = false
	deposit.Base.Rate.Type = "percentsRate"
	deposit.Base.Rate.Value = 0
	deposit.Base.IsMultiCurrency = false

	contract.Deposits = append(contract.Deposits, deposit)

	rs.Status = "success"
	rs.ActualTimestamp = time.Now().Unix()
	rs.Data.Contracts = append(rs.Data.Contracts, contract)

	err = json.NewEncoder(w).Encode(rs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter  due: %s", errWrite.Error())
		}
	}
}
