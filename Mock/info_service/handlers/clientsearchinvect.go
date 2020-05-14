package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Mock/info_service/datapool"
)

type BaseInvest struct {
	GUID             string               `json:"guid"`
	Hid              string               `json:"hid"`
	IdentityType     string               `json:"identityType"`
	ActualDate       string               `json:"actualDate"`
	FullName         string               `json:"fullName"`
	Surname          string               `json:"surname"`
	Name             string               `json:"name"`
	Patronymic       string               `json:"patronymic"`
	IsPatronymicLack bool                 `json:"isPatronymicLack"`
	Gender           string               `json:"gender"`
	Birthdate        string               `json:"birthdate"`
	BirthPlace       string               `json:"birthPlace"`
	BirthCountry     string               `json:"birthCountry"`
	DeathDate        string               `json:"deathDate"`
	Bankruptcy       string               `json:"bankruptcy"`
	Categories       []CategorieInvest    `json:"categories"`
	Residents        []ResidentInsest     `json:"residents"`
	Citizenships     []CitizenshipsInvest `json:"citizenships"`
}
type ClientSearchInvestRS struct {
	Base      BaseInvest       `json:"base"`
	Addresses []AddressInvest  `json:"addresses"`
	Documents []DocumentInvest `json:"documents"`
	Phones    []PhoneInvest    `json:"phones"`
	Mails     []MailInvest     `json:"mails"`
	Sources   []SourceInvest   `json:"sources"`
}

type State struct {
	Code         string `json:"code"`
	TerminalFlag bool   `json:"terminalFlag"`
}
type CategorieInvest struct {
	Type string `json:"type"`
}

type ResidentInsest struct {
	Type  string `json:"type"`
	State State  `json:"state"`
}

type CitizenshipsInvest struct {
	CountryName string `json:"countryName"`
}

type AddressInvest struct {
	Hid             string `json:"hid"`
	Type            string `json:"type"`
	Primary         bool   `json:"primary"`
	ActualDate      string `json:"actualDate"`
	PostalCode      string `json:"postalCode"`
	KladrPostalCode string `json:"kladrPostalCode"`
	CountryName     string `json:"countryName"`
	District        string `json:"district"`
	RegionType      string `json:"regionType"`
	RegionName      string `json:"regionName"`
	CityType        string `json:"cityType"`
	City            string `json:"city"`
	StreetType      string `json:"streetType"`
	Street          string `json:"street"`
	HouseNumber     string `json:"houseNumber"`
	Flat            string `json:"flat"`
	OkatoCode       string `json:"okatoCode"`
	KladrCode       string `json:"kladrCode"`
	FullAddress     string `json:"fullAddress"`
	IsForeign       bool   `json:"isForeign"`
}

type DocumentInvest struct {
	Hid            string `json:"hid,omitempty"`
	Past           bool   `json:"past"`
	Type           string `json:"type"`
	Primary        bool   `json:"primary"`
	ActualDate     string `json:"actualDate,omitempty"`
	Series         string `json:"series,omitempty"`
	Number         string `json:"number,omitempty"`
	FullValue      string `json:"fullValue"`
	IssueDate      string `json:"issueDate,omitempty"`
	IssueAuthority string `json:"issueAuthority,omitempty"`
	DepartmentCode string `json:"departmentCode,omitempty"`
	State          State  `json:"state,omitempty"`
}

type PhoneInvest struct {
	Hid           string `json:"hid"`
	Type          string `json:"type"`
	Primary       bool   `json:"primary"`
	ActualDate    string `json:"actualDate"`
	CountryCode   string `json:"countryCode"`
	CityCode      string `json:"cityCode"`
	Number        string `json:"number"`
	FullNumber    string `json:"fullNumber"`
	Timezone      string `json:"timezone"`
	NumberProfile string `json:"numberProfile"`
	RawSource     string `json:"rawSource"`
	State         State  `json:"state"`
	IsForeign     bool   `json:"isForeign"`
}

type MailInvest struct {
	Hid        string      `json:"hid"`
	GUID       interface{} `json:"guid"`
	Type       string      `json:"type"`
	Primary    bool        `json:"primary"`
	ActualDate string      `json:"actualDate"`
	Value      string      `json:"value"`
	State      interface{} `json:"state"`
}
type SourceInvest struct {
	Hid        string           `json:"hid"`
	SystemInfo SystemInfoInvest `json:"systemInfo"`
}
type SystemInfoInvest struct {
	SystemID string `json:"systemId"`
	RawID    string `json:"rawId"`
}

func ClientSearchInvest(rq ClientSearchRQ, w http.ResponseWriter) {
	client := datapool.GUIDPool[rq.Data.Filter.GUID]
	rs := ClientSearchInvestRS{}
	rs.Base = BaseInvest{
		GUID:             client.GUID,
		Hid:              "162847723",
		IdentityType:     "3",
		ActualDate:       "2020-01-30",
		FullName:         "ААААА КРИСТИНА ВИКТОРОВНА",
		Surname:          "Ааааа",
		Name:             "Кристина",
		Patronymic:       "Викторовна",
		Gender:           "FEMALE",
		BirthPlace:       "РОССИЯ, ГОРОД МОСКВА",
		IsPatronymicLack: false,
		Birthdate:        "1990-04-16",
		BirthCountry:     "СССР",
		DeathDate:        "3120-04-16",
		Bankruptcy:       "false",
		Categories:       append(rs.Base.Categories, CategorieInvest{Type: "REGULAR"}),
	}

	rs.Base.Residents = append(rs.Base.Residents, ResidentInsest{
		State: State{TerminalFlag: true},
		Type:  "base",
	})
	rs.Base.Residents = append(rs.Base.Residents, ResidentInsest{
		State: State{TerminalFlag: true},
		Type:  "tax",
	})
	rs.Base.Citizenships = append(rs.Base.Citizenships, CitizenshipsInvest{CountryName: "РОССИЯ"})

	rs.Addresses = append(rs.Addresses, AddressInvest{
		Hid:             "47501197",
		Type:            "CONSTANT_REGISTRATION",
		Primary:         false,
		ActualDate:      "2018-03-01",
		PostalCode:      "109029",
		KladrPostalCode: "109029",
		CountryName:     "Россия",
		District:        "Центральный",
		RegionName:      "Москва",
		CityType:        "г",
		City:            "Москва",
		StreetType:      "ул",
		Street:          "Калитниковская М.",
		HouseNumber:     "22",
		Flat:            "66",
		OkatoCode:       "45286580000",
		KladrCode:       "7700000000014190013",
		FullAddress:     "109029, Россия, г Москва, ул Калитниковская М., д. 22, кв. 66",
		IsForeign:       false,
	})
	rs.Addresses = append(rs.Addresses, AddressInvest{
		Hid:             "23343433",
		Type:            "HOME",
		Primary:         false,
		ActualDate:      "2018-03-01",
		PostalCode:      "109029",
		KladrPostalCode: "109029",
		CountryName:     "Россия",
		District:        "Центральный",
		RegionName:      "Москва",
		CityType:        "г",
		City:            "Москва",
		StreetType:      "ул",
		Street:          "Калитниковская М.",
		HouseNumber:     "22",
		Flat:            "66",
		OkatoCode:       "45286580000",
		KladrCode:       "7700000000014190013",
		FullAddress:     "109029, Россия, г Москва, ул Коровий вал., д. 5, кв. 66",
		IsForeign:       false,
	})

	rs.Documents = append(rs.Documents, DocumentInvest{
		Hid:            "462626332",
		Past:           false,
		Type:           "PASSPORT_RU",
		Primary:        true,
		ActualDate:     "2020-04-16",
		Series:         "89 89",
		Number:         "999777",
		FullValue:      "89 89999777",
		IssueDate:      "2020-04-16",
		IssueAuthority: "ОТДЕЛОМ КРИВОРУКИХ УБЛЮДКОВ ИЗ ГОРОДА МОСКВЫ",
		DepartmentCode: "Отлел 777",
		State:          State{Code: "ACTUAL", TerminalFlag: false},
	})
	rs.Documents = append(rs.Documents, DocumentInvest{
		Past:      false,
		Type:      "INN",
		Primary:   false,
		FullValue: "770973997100",
		State:     State{TerminalFlag: false},
	})
	rs.Phones = append(rs.Phones, PhoneInvest{
		Hid:           "87575",
		Type:          "MOBILE",
		Primary:       true,
		ActualDate:    "2020-04-16",
		CountryCode:   "1",
		CityCode:      "222",
		Number:        "3334455",
		FullNumber:    client.Phone,
		Timezone:      "UTC+3",
		NumberProfile: "MOBILE",
		RawSource:     "+" + client.Phone,
		State:         State{Code: "ACTUAL", TerminalFlag: false},
		IsForeign:     false,
	})
	rs.Phones = append(rs.Phones, PhoneInvest{
		Hid:           "39465318",
		Type:          "PC",
		Primary:       true,
		ActualDate:    "2020-04-16",
		CountryCode:   "1",
		CityCode:      "0111",
		Number:        "222333444",
		FullNumber:    "01112223344",
		Timezone:      "UTC+3",
		NumberProfile: "MOBILE",
		RawSource:     "+01112223344",
		State:         State{Code: "ACTUAL", TerminalFlag: false},
		IsForeign:     false,
	})
	rs.Mails = append(rs.Mails, MailInvest{
		Hid:        "3109903",
		GUID:       nil,
		Type:       "HOME",
		Primary:    true,
		ActualDate: "2019-07-17",
		Value:      "mu@mumu.ru",
		State:      nil,
	})
	rs.Mails = append(rs.Mails, MailInvest{
		Hid:        "3109904",
		GUID:       nil,
		Type:       "PC",
		Primary:    true,
		ActualDate: "2019-07-17",
		Value:      "mu@mu.ru",
		State:      nil,
	})
	rs.Mails = append(rs.Mails, MailInvest{
		Hid:        "3109905",
		GUID:       nil,
		Type:       "PC",
		Primary:    true,
		ActualDate: "2019-07-17",
		Value:      "ru@ru.ru",
		State:      nil,
	})
	rs.Sources = append(rs.Sources, SourceInvest{
		Hid:        "72943076",
		SystemInfo: SystemInfoInvest{SystemID: "system one", RawID: "49524"},
	})
	rs.Sources = append(rs.Sources, SourceInvest{
		Hid:        "74875458",
		SystemInfo: SystemInfoInvest{SystemID: "system two", RawID: "10013786933"},
	})
	jsonStr, _ := json.Marshal(rs)
	response := "\"ClientSearchResponseData{clients=\"" + string(jsonStr) + "}"
	log.Println(response)
	_, err := w.Write([]byte(response))
	//err := json.NewEncoder(w).Encode(rs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter  due: %s", errWrite.Error())
		}
	}

}
