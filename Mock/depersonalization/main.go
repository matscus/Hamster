package main

import (
	"encoding/json"
	"log"
	"os"
)

type DepositListRS struct {
	Data struct {
		GetDepositList []struct {
			Contracts []struct {
				SystemInfoContractID         string `json:"SystemInfoContractId"`
				SystemInfoSystemID           string `json:"SystemInfoSystemId"`
				ContractStateName            string `json:"ContractStateName"`
				OwnerRawID                   string `json:"OwnerRawId"`
				ContractNumber               string `json:"ContractNumber"`
				ContractBankDepartmentObject string `json:"ContractBankDepartmentObject"`
				ContractCloseDateFact        string `json:"ContractCloseDateFact"`
				PercentRates                 []struct {
					PretermPercentAmount  string `json:"PretermPercentAmount"`
					RankedPercentAmount   string `json:"RankedPercentAmount"`
					SystemInfoContractID  string `json:"SystemInfoContractId"`
					ClosePercentAmount    string `json:"ClosePercentAmount"`
					PercentCharCode       string `json:"PercentCharCode"`
					PercentStartDate      string `json:"PercentStartDate"`
					RateType              string `json:"RateType"`
					ClientRequestRawID    string `json:"ClientRequestRawId"`
					RateValue             string `json:"RateValue"`
					PercentNumCode        string `json:"PercentNumCode"`
					ClientRequestSystemID string `json:"ClientRequestSystemId"`
					PercentAmount         string `json:"PercentAmount"`
					AddedPercentAmount    string `json:"AddedPercentAmount"`
					RateValueBon          string `json:"RateValueBon"`
					PercentEndDate        string `json:"PercentEndDate"`
				} `json:"PercentRates"`
				DepositContracts []struct {
					SystemInfoContractID           string      `json:"SystemInfoContractId"`
					MainAccountType                string      `json:"MainAccountType"`
					MainAccountWithdrawalStartDate interface{} `json:"MainAccountWithdrawalStartDate"`
					MainAccountNumber              string      `json:"MainAccountNumber"`
					MainAccountStateName           string      `json:"MainAccountStateName"`
					MainAccountStateCode           string      `json:"MainAccountStateCode"`
					CardMaskNum                    interface{} `json:"CardMaskNum"`
					MainAccountRawID               string      `json:"MainAccountRawId"`
					MainAccountWithdrawalEndDate   interface{} `json:"MainAccountWithdrawalEndDate"`
					CardVirtualNum                 interface{} `json:"CardVirtualNum"`
					TermValue                      int         `json:"TermValue"`
					IsMultiCurrency                bool        `json:"IsMultiCurrency"`
					MainAccountBankBic             string      `json:"MainAccountBankBic"`
					ClientRequestRawID             string      `json:"ClientRequestRawId"`
					MainAccountCreationDate        string      `json:"MainAccountCreationDate"`
					IsCapitalisation               bool        `json:"IsCapitalisation"`
					ProlongCount                   int         `json:"ProlongCount"`
					PayPercentAccountObject        interface{} `json:"PayPercentAccountObject"`
					DvAccountObject                struct {
						TopupMinAmount          interface{} `json:"TopupMinAmount"`
						SystemInfoContractID    string      `json:"SystemInfoContractId"`
						WithdrawalEndDate       interface{} `json:"WithdrawalEndDate"`
						StateDate               string      `json:"StateDate"`
						TopupMaxAmount          interface{} `json:"TopupMaxAmount"`
						Number                  string      `json:"Number"`
						CreationDate            string      `json:"CreationDate"`
						CloseDateFact           interface{} `json:"CloseDateFact"`
						WithdrawalStartDate     interface{} `json:"WithdrawalStartDate"`
						BalanceValueNumCode     string      `json:"BalanceValueNumCode"`
						WithdrawalMinAmount     interface{} `json:"WithdrawalMinAmount"`
						IsArrest                bool        `json:"IsArrest"`
						WithdrawalIsPossibility bool        `json:"WithdrawalIsPossibility"`
						StateCode               string      `json:"StateCode"`
						TopupStartDate          interface{} `json:"TopupStartDate"`
						TopupIsPossibility      bool        `json:"TopupIsPossibility"`
						TopupEndDate            interface{} `json:"TopupEndDate"`
						Type                    string      `json:"Type"`
						CurrencyCharCode        string      `json:"CurrencyCharCode"`
						WithdrawalMaxAmount     interface{} `json:"WithdrawalMaxAmount"`
						BalanceValueAmount      string      `json:"BalanceValueAmount"`
						ClientRequestRawID      string      `json:"ClientRequestRawId"`
						BalanceValueAmountRur   string      `json:"BalanceValueAmountRur"`
						SystemInfoRawID         string      `json:"SystemInfoRawId"`
						ClientRequestSystemID   string      `json:"ClientRequestSystemId"`
						CurrencyNumCode         string      `json:"CurrencyNumCode"`
						StateName               string      `json:"StateName"`
						BalanceValueCharCode    string      `json:"BalanceValueCharCode"`
						BankBic                 string      `json:"BankBic"`
					} `json:"DvAccountObject"`
					InitialAmount                  string      `json:"InitialAmount"`
					GiverFullName                  string      `json:"GiverFullName"`
					GiverRawID                     string      `json:"GiverRawId"`
					PayPercentAccount              interface{} `json:"PayPercentAccount"`
					MainAccountBalanceCharCode     string      `json:"MainAccountBalanceCharCode"`
					GiverSystemID                  string      `json:"GiverSystemId"`
					IsEndTermPayment               bool        `json:"IsEndTermPayment"`
					MainAccountStateDate           string      `json:"MainAccountStateDate"`
					MinBalance                     interface{} `json:"MinBalance"`
					IsProlongation                 bool        `json:"IsProlongation"`
					TermUnit                       string      `json:"TermUnit"`
					MaxDepositAmount               interface{} `json:"MaxDepositAmount"`
					MinDepositAmount               interface{} `json:"MinDepositAmount"`
					PeriodKind                     string      `json:"PeriodKind"`
					IsPercentToDeposit             bool        `json:"IsPercentToDeposit"`
					DvAccount                      string      `json:"DvAccount"`
					MainAccountWithdrawalMinAmount interface{} `json:"MainAccountWithdrawalMinAmount"`
					MainAccountNumCode             string      `json:"MainAccountNumCode"`
					MainAccountWithdrawalMaxAmount interface{} `json:"MainAccountWithdrawalMaxAmount"`
					MainAccountBalanceAmount       string      `json:"MainAccountBalanceAmount"`
					MainAccountTopupMaxAmount      interface{} `json:"MainAccountTopupMaxAmount"`
					MainAccountTopupEndDate        interface{} `json:"MainAccountTopupEndDate"`
					MainAccountBalanceNumCode      string      `json:"MainAccountBalanceNumCode"`
					ProductType                    int         `json:"ProductType"`
					MainAccountSystemID            string      `json:"MainAccountSystemId"`
					MainAccountTopupStartDate      interface{} `json:"MainAccountTopupStartDate"`
					MainAccountCharCode            string      `json:"MainAccountCharCode"`
					ClientRequestSystemID          string      `json:"ClientRequestSystemId"`
					MainAccountCloseDateFact       interface{} `json:"MainAccountCloseDateFact"`
					MainAccountIsArrest            bool        `json:"MainAccountIsArrest"`
					MainAccountTopupIsPossibility  bool        `json:"MainAccountTopupIsPossibility"`
					MainAccountBalanceAmountRur    string      `json:"MainAccountBalanceAmountRur"`
				} `json:"DepositContracts"`
				ProductType              int    `json:"ProductType"`
				ContractCreationDate     string `json:"ContractCreationDate"`
				ContractIsWorking        bool   `json:"ContractIsWorking"`
				ProductCode              string `json:"ProductCode"`
				ContractStateDate        string `json:"ContractStateDate"`
				ProductName              string `json:"ProductName"`
				ClientFullName           string `json:"ClientFullName"`
				ClientRequestRawID       string `json:"ClientRequestRawId"`
				ContractStateCode        string `json:"ContractStateCode"`
				OwnerFullName            string `json:"OwnerFullName"`
				ContractStartDate        string `json:"ContractStartDate"`
				OwnerSystemID            string `json:"OwnerSystemId"`
				ClientRequestSystemID    string `json:"ClientRequestSystemId"`
				ContractEmployeeFullName string `json:"ContractEmployeeFullName"`
				ContractCloseDatePlan    string `json:"ContractCloseDatePlan"`
			} `json:"Contracts"`
			Lag struct {
				OriginLag interface{} `json:"originLag"`
				Lag       int         `json:"Lag"`
				Timestamp interface{} `json:"timestamp"`
				Origin    interface{} `json:"origin"`
				NodeName  interface{} `json:"node_name"`
			} `json:"lag"`
		} `json:"getDepositList"`
	} `json:"data"`
}

func main() {
	file, err := os.Open("deposit.json")
	if err != nil {
		log.Println(err)
	}
	depositRS := DepositListRS{}

	json.NewDecoder(file).Decode(&depositRS)
	fileOut, err := os.Create("depositRS.json")
	if err != nil {
		log.Println(err)
	}
	for i := 0; i < len(depositRS.Data.GetDepositList); i++ {

		for ii := 0; ii < len(depositRS.Data.GetDepositList[i].Contracts); ii++ {
			depositRS.Data.GetDepositList[i].Contracts[ii].SystemInfoContractID = "00000000000"
			depositRS.Data.GetDepositList[i].Contracts[ii].SystemInfoSystemID = "dr00"
			depositRS.Data.GetDepositList[i].Contracts[ii].OwnerRawID = "00000000000"
			depositRS.Data.GetDepositList[i].Contracts[ii].ContractNumber = "LLL -0000/00-00000"
			depositRS.Data.GetDepositList[i].Contracts[ii].ProductCode = "ФФФФФ_ФффФ"
			depositRS.Data.GetDepositList[i].Contracts[ii].ProductName = "БанкБанк - Мега Банк"
			depositRS.Data.GetDepositList[i].Contracts[ii].ClientRequestSystemID = "ИВАНОВ ИВАН ИВАНОВИЧ"
			depositRS.Data.GetDepositList[i].Contracts[ii].ClientRequestRawID = "00000000000"
			depositRS.Data.GetDepositList[i].Contracts[ii].OwnerFullName = "ИВАНОВ ИВАН ИВАНОВИЧ"
			depositRS.Data.GetDepositList[i].Contracts[ii].OwnerSystemID = "DR00"
			depositRS.Data.GetDepositList[i].Contracts[ii].ClientRequestSystemID = "DR00"
			depositRS.Data.GetDepositList[i].Contracts[ii].ContractEmployeeFullName = "ПЕТРОВ ПЕТР ПЕТРОВИЧ"
			for iii := 0; iii < len(depositRS.Data.GetDepositList[i].Contracts[ii].PercentRates); iii++ {
				depositRS.Data.GetDepositList[i].Contracts[ii].PercentRates[iii].SystemInfoContractID = "00000000000"
				depositRS.Data.GetDepositList[i].Contracts[ii].PercentRates[iii].ClientRequestRawID = "00000000000"
				depositRS.Data.GetDepositList[i].Contracts[ii].PercentRates[iii].ClientRequestSystemID = "DR00"
			}
			for iii := 0; iii < len(depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts); iii++ {
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].SystemInfoContractID = "00000000000"
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].MainAccountNumber = "00000000800000000000"
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].MainAccountRawID = "00000000000"
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].MainAccountBankBic = "000000000"
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].ClientRequestRawID = "00000000000"
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].DvAccountObject.ClientRequestRawID = "00000000000"
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].DvAccountObject.SystemInfoContractID = "00000000000"
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].DvAccountObject.Number = "00000000000000000000"
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].DvAccountObject.SystemInfoRawID = "00000000000"
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].DvAccountObject.ClientRequestSystemID = "DR00"
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].DvAccountObject.BankBic = "000000000"
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].GiverFullName = "ИВАНОВ ИВАН ИВАНОВИЧ"
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].GiverRawID = "00000000000"
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].DvAccount = "0000000000"
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].GiverSystemID = "DR00"
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].MainAccountSystemID = "DR00"
				depositRS.Data.GetDepositList[i].Contracts[ii].DepositContracts[iii].ClientRequestSystemID = "DR00"
			}
		}
	}

	depositBytes, err := json.Marshal(depositRS)
	if err != nil {
		log.Println(err)
	}
	fileOut.Write([]byte(depositBytes))
}

func (depositRS *DepositListRS) Depersonalization() {
	for _, getDepositList := range depositRS.Data.GetDepositList {
		for _, contracts := range getDepositList.Contracts {
			contracts.SystemInfoContractID = "00000000000"
			contracts.SystemInfoSystemID = "dr00"
			contracts.OwnerRawID = "00000000000"
			contracts.ContractNumber = "LLL -0000/00-00000"
			for _, percentRates := range contracts.PercentRates {
				percentRates.SystemInfoContractID = "00000000000"
				percentRates.ClientRequestRawID = "00000000000"
				percentRates.ClientRequestSystemID = "DR00"
			}
			for _, depositContracts := range contracts.DepositContracts {
				depositContracts.SystemInfoContractID = "00000000000"
				depositContracts.MainAccountNumber = "00000000800000000000"
				depositContracts.MainAccountRawID = "00000000000"
				depositContracts.MainAccountBankBic = "000000000"
				depositContracts.ClientRequestRawID = "00000000000"
				depositContracts.DvAccountObject.ClientRequestRawID = "00000000000"
				depositContracts.DvAccountObject.SystemInfoContractID = "00000000000"
				depositContracts.DvAccountObject.Number = "00000000000000000000"
				depositContracts.DvAccountObject.SystemInfoRawID = "00000000000"
				depositContracts.DvAccountObject.ClientRequestSystemID = "DR00"
				depositContracts.DvAccountObject.BankBic = "000000000"
				depositContracts.GiverFullName = "ИВАНОВ ИВАН ИВАНОВИЧ"
				depositContracts.GiverRawID = "00000000000"
				depositContracts.DvAccount = "0000000000"
				depositContracts.GiverSystemID = "DR00"
				depositContracts.MainAccountSystemID = "DR00"
				depositContracts.ClientRequestSystemID = "DR00"
			}
			contracts.ProductCode = "ФФФФФ_ФффФ"
			contracts.ProductName = "БанкБанк - Мега Банк"
			contracts.ClientRequestSystemID = "ИВАНОВ ИВАН ИВАНОВИЧ"
			contracts.ClientRequestRawID = "00000000000"
			contracts.OwnerFullName = "ИВАНОВ ИВАН ИВАНОВИЧ"
			contracts.OwnerSystemID = "DR00"
			contracts.ClientRequestSystemID = "DR00"
			contracts.ContractEmployeeFullName = "ПЕТРОВ ПЕТР ПЕТРОВИЧ"

		}
	}

	log.Println(depositRS)
	log.Println("depositRS")
}
