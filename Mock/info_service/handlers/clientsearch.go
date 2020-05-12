package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/matscus/Hamster/Mock/info_service/datapool"
)

type ClientSearchRQ struct {
	Meta struct {
		Channel string `json:"channel"`
	} `json:"meta"`
	Data struct {
		RequestFields []string `json:"requestFields"`
		Filter        struct {
			GUID  string `json:"guid"`
			Phone struct {
				FullNumber string `json: "fullNumber"`
			} `json:"phone"`
		} `json:"filter"`
	} `json:"data"`
}

type ClientSearchRS struct {
	Status          string `json:"status"`
	ActualTimestamp int64  `json:"actualTimestamp"`
	Data            struct {
		Clients []Client `json:"clients"`
	} `json:"data"`
}
type ClientSearchRSv2 struct {
	Status          string `json:"status"`
	ActualTimestamp int64  `json:"actualTimestamp"`
	Data            struct {
		Clientsv2 []Clientv2 `json:"clients"`
	} `json:"data"`
}

type Client struct {
	Addresses []Addresse `json:"addresses"`
	Base      struct {
		ActualDate       string        `json:"actualDate"`
		BirthPlace       string        `json:"birthPlace"`
		Birthdate        string        `json:"birthdate"`
		Categories       []Categorie   `json:"categories"`
		Citizenships     []Citizenship `json:"citizenships"`
		FullName         string        `json:"fullName"`
		Gender           string        `json:"gender"`
		GUID             string        `json:"guid"`
		Hid              string        `json:"hid"`
		IdentityType     string        `json:"identityType"`
		IsPatronymicLack bool          `json:"isPatronymicLack"`
		Name             string        `json:"name"`
		Patronymic       string        `json:"patronymic"`
		Residents        []Resident    `json:"residents"`
		Surname          string        `json:"surname"`
	} `json:"base"`
	Detail struct {
		Biometrics struct {
			IsAgreement bool `json:"isAgreement"`
		} `json:"biometrics"`
		LastFio []interface{} `json:"lastFio"`
	} `json:"detail"`
	Documents []Document `json:"documents"`
	Fatca     struct{}   `json:"fatca"`
	StopLists []StopList `json:"stopList"`
	Mails     []string   `json:"mails"`
	Phones    []Phone    `json:"phones"`
	Sources   []Source   `json:"sources"`
}

type Clientv2 struct {
	Base struct {
		ActualDate       string        `json:"actualDate"`
		BirthPlace       string        `json:"birthPlace"`
		Birthdate        string        `json:"birthdate"`
		Categories       []Categorie   `json:"categories"`
		Citizenships     []Citizenship `json:"citizenships"`
		FullName         string        `json:"fullName"`
		Gender           string        `json:"gender"`
		GUID             string        `json:"guid"`
		Hid              string        `json:"hid"`
		IdentityType     string        `json:"identityType"`
		IsPatronymicLack bool          `json:"isPatronymicLack"`
		Name             string        `json:"name"`
		Patronymic       string        `json:"patronymic"`
		Residents        []Resident    `json:"residents"`
		Surname          string        `json:"surname"`
	} `json:"base"`
	Phones  []Phone  `json:"phones"`
	Sources []Source `json:"sources"`
}

type StopList struct {
	TotalMatches string `json:"totalMatches"`
}

type Addresse struct {
	ActualDate      string `json:"actualDate"`
	Area            string `json:"area"`
	AreaType        string `json:"areaType"`
	City            string `json:"city"`
	CityType        string `json:"cityType"`
	CountryName     string `json:"countryName"`
	District        string `json:"district"`
	Flat            string `json:"flat"`
	FullAddress     string `json:"fullAddress"`
	Hid             string `json:"hid"`
	HouseNumber     string `json:"houseNumber"`
	IsForeign       bool   `json:"isForeign"`
	KladrCode       string `json:"kladrCode"`
	KladrPostalCode string `json:"kladrPostalCode"`
	OkatoCode       string `json:"okatoCode"`
	PostalCode      string `json:"postalCode"`
	Primary         bool   `json:"primary"`
	RegionName      string `json:"regionName"`
	RegionType      string `json:"regionType"`
	RegionCode      string `json:"regionCode"`
	Settlement      string `json:"settlement"`
	SettlementType  string `json:"settlementType"`
	Street          string `json:"street"`
	StreetType      string `json:"streetType"`
	Type            string `json:"type"`
}

type Categorie struct {
	Params []Param `json:"params"`
	Type   string  `json:"type"`
}

type Param struct {
	Key string `json:"key"`
}

type Citizenship struct {
	CountryName string `json:"countryName"`
}

type Resident struct {
	State struct {
		TerminalFlag bool `json:"terminalFlag"`
	} `json:"state"`
	Type string `json:"type"`
}

type Document struct {
	ActualDate     string `json:"actualDate"`
	DepartmentCode string `json:"departmentCode"`
	FullValue      string `json:"fullValue"`
	Hid            string `json:"hid"`
	IssueAuthority string `json:"issueAuthority"`
	IssueDate      string `json:"issueDate"`
	Number         string `json:"number"`
	Primary        bool   `json:"primary"`
	Series         string `json:"series"`
	State          struct {
		Code string `json:"code"`
	} `json:"state"`
	Type string `json:"type"`
}

type Phone struct {
	ActualDate    string `json:"actualDate"`
	CityCode      string `json:"cityCode"`
	CountryCode   string `json:"countryCode"`
	FullNumber    string `json:"fullNumber"`
	Hid           string `json:"hid"`
	IsForeign     bool   `json:"isForeign"`
	Number        string `json:"number"`
	NumberProfile string `json:"numberProfile"`
	Primary       bool   `json:"primary"`
	RawSource     string `json:"rawSource"`
	State         struct {
		Code string `json:"code"`
	} `json:"state"`
	Timezone string `json:"timezone"`
	Type     string `json:"type"`
}

type Source struct {
	Hid        string `json:"hid"`
	SystemInfo struct {
		RawID    string `json:"rawId"`
		SystemID string `json:"systemId"`
	} `json:"systemInfo"`
}

func ClientSearchCommon(w http.ResponseWriter, r *http.Request) {
	rq := ClientSearchRQ{}
	err := json.NewDecoder(r.Body).Decode(&rq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	client := datapool.GUIDPool[rq.Data.Filter.GUID]
	if client.GUID != "" {
		if len(rq.Data.RequestFields) == 8 {
			ClientSearchInvest(rq, w)
		} else {
			ClientSearchV1(rq, w)
		}
	}
	if client.GUID == "" {
		ClientSearchV2(rq, w)
	}

}

func ClientSearchV1(rq ClientSearchRQ, w http.ResponseWriter) {
	// rq := ClientSearchRQ{}
	// // log.Println("ClientSearchRQ = " + rq.Data.Filter.GUID)
	// err := json.NewDecoder(r.Body).Decode(&rq)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	// 	if errWrite != nil {
	// 		log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
	// 	}
	// 	return
	// }
	client := datapool.GUIDPool[rq.Data.Filter.GUID]
	rs := ClientSearchRS{}

	cli := Client{}
	cli.Base.GUID = client.GUID

	cli.Base.Hid = "162847723"
	cli.Base.IdentityType = "3"
	cli.Base.ActualDate = "2020-01-30"
	cli.Base.FullName = "ААААА КРИСТИНА ВИКТОРОВНА"
	cli.Base.Surname = "Ааааа"
	cli.Base.Name = "Кристина"
	cli.Base.Patronymic = "Викторовна"
	cli.Base.Gender = "FEMALE"
	cli.Base.BirthPlace = "РОССИЯ, ГОРОД МОСКВА"
	categorie := Categorie{}
	categorie.Type = "EMPLOYEE"
	param := Param{}
	param.Key = "employeeFireDate"
	categorie.Params = append(categorie.Params, param)
	categorie.Type = "REGULAR"
	cli.Base.Categories = append(cli.Base.Categories, categorie)
	resident := Resident{}
	resident.Type = "base"
	resident.State.TerminalFlag = false
	cli.Base.Residents = append(cli.Base.Residents, resident)
	citizenship := Citizenship{}
	citizenship.CountryName = "Российская федерация"
	cli.Base.Citizenships = append(cli.Base.Citizenships, citizenship)
	cli.Base.IsPatronymicLack = true
	cli.Base.Birthdate = "1990-07-10"
	addresse := Addresse{}
	addresse.Hid = "82728384"
	addresse.Type = "HOME"
	addresse.Primary = true
	addresse.ActualDate = "2020-01-24"
	addresse.PostalCode = "117461"
	addresse.KladrPostalCode = "117461"
	addresse.CountryName = "Россия"
	addresse.District = "Центральный"
	addresse.RegionType = "г"
	addresse.RegionName = "Москва"
	addresse.RegionCode = "099"
	addresse.CityType = "г"
	addresse.City = "Москва"
	addresse.StreetType = "ул"
	addresse.Street = "Херсонская"
	addresse.HouseNumber = "1"
	addresse.Flat = "123"
	addresse.OkatoCode = "45293562000"
	addresse.KladrCode = "7700000000030250061"
	addresse.FullAddress = "117461, Россия, г Москва, ул Херсонская, д. 1, кв. 123"
	addresse.IsForeign = false
	cli.Addresses = append(cli.Addresses, addresse)
	document := Document{}
	document.Hid = "96938652"
	document.Type = "PASSPORT_RU"
	document.Primary = true
	document.ActualDate = "2020-01-21"
	document.Series = "4515"
	document.Number = "111222"
	document.FullValue = "4515 111222"
	document.IssueDate = "2016-03-25"
	document.IssueAuthority = "ОТДЕЛОМ УФМС РОССИИ ПО ГОР. МОСКВЕ ПО РАЙОНУ ЗЮЗИНО"
	document.DepartmentCode = "770-116"
	document.State.Code = "ACTUAL"
	document.Type = "SNILS"
	document.FullValue = "001-ААА-ААА 17"
	cli.Documents = append(cli.Documents, document)

	phone := Phone{}
	phone.Hid = "67381378"
	phone.Type = "PC"
	phone.Primary = true
	phone.ActualDate = "2020-01-30"
	phone.CountryCode = "7"
	phone.CityCode = "985"
	phone.Number = client.Phone
	phone.FullNumber = client.Phone
	phone.Timezone = "UTC+3"
	phone.NumberProfile = "MOBILE"
	phone.RawSource = "+" + client.Phone
	phone.State.Code = "ACTUAL"
	phone.IsForeign = false
	cli.Phones = append(cli.Phones, phone)

	stopList := StopList{}
	stopList.TotalMatches = "0"
	cli.StopLists = append(cli.StopLists, stopList)

	source := Source{}
	source.Hid = "146798431"
	source.SystemInfo.SystemID = "DRTL"
	source.SystemInfo.RawID = client.UserID
	cli.Sources = append(cli.Sources, source)
	cli.Detail.Biometrics.IsAgreement = false

	rs.Status = "success"
	rs.ActualTimestamp = time.Now().Unix()
	rs.Data.Clients = append(rs.Data.Clients, cli)

	err := json.NewEncoder(w).Encode(rs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter  due: %s", errWrite.Error())
		}
	}
}

func ClientSearchV2(rq ClientSearchRQ, w http.ResponseWriter) {
	// rq := ClientSearchRQ{}
	// // log.Println("ClientSearchRQ = " + rq.Data.Filter.Phone.FullNumber)
	// err := json.NewDecoder(r.Body).Decode(&rq)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	// 	if errWrite != nil {
	// 		log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
	// 	}
	// 	return
	// }
	client := datapool.PhonePool[rq.Data.Filter.Phone.FullNumber]
	rs := ClientSearchRSv2{}

	cli := Clientv2{}
	cli.Base.GUID = client.GUID
	cli.Base.Hid = "162847723"
	cli.Base.IdentityType = "3"
	cli.Base.ActualDate = "2020-01-30"
	cli.Base.FullName = "ААААА КРИСТИНА ВИКТОРОВНА"
	cli.Base.Surname = "Ааааа"
	cli.Base.Name = "Кристина"
	cli.Base.Patronymic = "Викторовна"
	cli.Base.Gender = "FEMALE"
	cli.Base.BirthPlace = "РОССИЯ, ГОРОД МОСКВА"
	categorie := Categorie{}
	categorie.Type = "EMPLOYEE"
	param := Param{}
	param.Key = "employeeFireDate"
	categorie.Params = append(categorie.Params, param)
	categorie.Type = "REGULAR"
	cli.Base.Categories = append(cli.Base.Categories, categorie)
	resident := Resident{}
	resident.Type = "base"
	resident.State.TerminalFlag = false
	cli.Base.Residents = append(cli.Base.Residents, resident)
	citizenship := Citizenship{}
	citizenship.CountryName = "Российская федерация"
	cli.Base.Citizenships = append(cli.Base.Citizenships, citizenship)
	cli.Base.IsPatronymicLack = true
	cli.Base.Birthdate = "1990-07-10"

	phone := Phone{}
	phone.Hid = "67381378"
	phone.Type = "PC"
	phone.Primary = true
	phone.ActualDate = "2020-01-30"
	phone.CountryCode = "7"
	phone.CityCode = "985"
	phone.Number = client.Phone
	phone.FullNumber = client.Phone
	phone.Timezone = "UTC+3"
	phone.NumberProfile = "MOBILE"
	phone.RawSource = "+" + client.Phone
	phone.State.Code = "ACTUAL"
	phone.IsForeign = false
	cli.Phones = append(cli.Phones, phone)

	source := Source{}
	source.Hid = "146798431"
	source.SystemInfo.SystemID = "DRTL"
	source.SystemInfo.RawID = client.UserID
	cli.Sources = append(cli.Sources, source)

	rs.Status = "success"
	rs.ActualTimestamp = time.Now().Unix()
	rs.Data.Clientsv2 = append(rs.Data.Clientsv2, cli)

	err := json.NewEncoder(w).Encode(rs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter  due: %s", errWrite.Error())
		}
	}
}
