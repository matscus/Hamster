package structs

import "encoding/xml"

var (
	IdParamSurname                    = 0
	IdParamName                       = 1
	IdParamPatronymic                 = 2
	IdParamGender                     = 3
	IdParamGenderRawSource            = 4
	IdParamFullNameQC                 = 5
	IdParamFullNameAuthor             = 6
	IdParamFullNameRawSource          = 7
	IdParamSurnameQc                  = 8
	IdParamFirstnameQc                = 9
	IdParamPatronymicQc               = 10
	IdParamGenderQc                   = 11
	IdParamNameCommonQc               = 12
	IdParamPatronymicLackFlag         = 13
	IdParamPatronymicLackFlagAuthor   = 14
	IdParamForeignSurname             = 15
	IdParamForeignSurnameAuthor       = 16
	IdParamForeignName                = 17
	IdParamBirthdate                  = 18
	IdParamBirthdateAuthor            = 19
	IdParamBirthdateQC                = 20
	IdParamBirthdateRawSource         = 21
	IdParamBirthPlace                 = 22
	IdParamBirthPlaceAuthor           = 23
	IdParamBirthCountry               = 24
	IdParamBirthCountryAuthor         = 25
	IdParamCitizenship                = 26
	IdParamCitizenshipAuthor          = 27
	IdParamCesidentFlag               = 28
	IdParamCesidentFlagAuthor         = 29
	IdParamMaritalStatus              = 30
	IdParamMaritalStatusAuthor        = 31
	IdParamInn                        = 32
	IdParamInnAuthor                  = 33
	IdParamInnQC                      = 34
	IdParamInnRawSource               = 35
	IdParamSnils                      = 36
	IdParamSnilsAuthor                = 37
	IdParamSnilsQC                    = 38
	IdParamSnilsRawSource             = 39
	IdParamEmployeeFlag               = 40
	IdParamEmployeeFiredDate          = 41
	IdParamEmployeeFlagAuthor         = 42
	IdParamBranch                     = 43
	IdParamBranchAuthor               = 44
	IdParamCrossId                    = 45
	IdParamCrossIdAuthor              = 46
	IdParamVipStatus                  = 47
	IdParamPersonalManager            = 48
	IdParamVipStatusAuthor            = 49
	IdParamOpenAccountsFlag           = 50
	IdParamOpenAccountsFlagAuthor     = 51
	IdParamAdvertisingConsent         = 52
	IdParamAdvertisingConsentAuthor   = 53
	IdParamDignitaryFlag              = 54
	IdParamDignitaryDescription       = 55
	IdParamDignitaryFlagAuthor        = 56
	IdParamDeathDate                  = 57
	IdParamDeathDateAuthor            = 58
	IdParamBankruptcyFlag             = 59
	IdParamBankruptcyFlagAuthor       = 60
	IdParamBlackListFlag              = 61
	IdParamBlackListReason            = 62
	IdParamBlackListDate              = 63
	IdParamBlackListAuthor            = 64
	IdParamLaunderingRiskReason       = 65
	IdParamLaunderingRiskReasonAuthor = 66
	IdParamCompliance                 = 67
	IdParamComplianceAuthor           = 68
	IdParamComplianceDate             = 69
	IdParamComplianceCheckDate        = 70
	IdParamComplianceDescription      = 71
	IdParamIbankFlag                  = 72
	IdParamIbankFlagAuthor            = 73
	IdParamActualityDateAuthor        = 74
	IdParamActualityDate              = 75
)

var SoapHead = []byte(
	`<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">` +
		`<soap:Body>` +
		`<searchResponse xmlns="http://hflabs.ru/cdi/soap/2_13">` +
		`<party type="PHYSICAL" hid="10159862" sourceSystem="DR25" rawId="10012047703">`)
var SoapClose = []byte(
	`</party>` +
		`</searchResponse>` +
		`</soap:Body>` +
		`</soap:Envelope>`)

type Datapool struct {
	Name    string
	Surname string
	Mafia   string
}
type Request struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Soapenv string   `xml:"soapenv,attr"`
	_13     string   `xml:"_13,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text   string `xml:",chardata"`
		Search struct {
			Text      string `xml:",chardata"`
			Query     string `xml:"query"`
			PartyType string `xml:"partyType"`
		} `xml:"search"`
	} `xml:"Body"`
}

type User struct {
	Surname                    string
	Name                       string
	Patronymic                 string
	Gender                     string
	GenderRawSource            string
	FullNameQC                 string
	FullNameAuthor             string
	FullNameRawSource          string
	SurnameQc                  string
	FirstnameQc                string
	PatronymicQc               string
	GenderQc                   string
	NameCommonQc               string
	PatronymicLackFlag         string
	PatronymicLackFlagAuthor   string
	ForeignSurname             string
	ForeignSurnameAuthor       string
	ForeignName                string
	Birthdate                  string
	BirthdateAuthor            string
	BirthdateQC                string
	BirthdateRawSource         string
	BirthPlace                 string
	BirthPlaceAuthor           string
	BirthCountry               string
	BirthCountryAuthor         string
	Citizenship                string
	CitizenshipAuthor          string
	CesidentFlag               string
	CesidentFlagAuthor         string
	MaritalStatus              string
	MaritalStatusAuthor        string
	Inn                        string
	InnAuthor                  string
	InnQC                      string
	InnRawSource               string
	Snils                      string
	SnilsAuthor                string
	SnilsQC                    string
	SnilsRawSource             string
	EmployeeFlag               string
	EmployeeFiredDate          string
	EmployeeFlagAuthor         string
	Branch                     string
	BranchAuthor               string
	CrossId                    string
	CrossIdAuthor              string
	VipStatus                  string
	PersonalManager            string
	VipStatusAuthor            string
	OpenAccountsFlag           string
	OpenAccountsFlagAuthor     string
	AdvertisingConsent         string
	AdvertisingConsentAuthor   string
	DignitaryFlag              string
	DignitaryDescription       string
	DignitaryFlagAuthor        string
	DeathDate                  string
	DeathDateAuthor            string
	BankruptcyFlag             string
	BankruptcyFlagAuthor       string
	BlackListFlag              string
	BlackListReason            string
	BlackListDate              string
	BlackListAuthor            string
	LaunderingRiskReason       string
	LaunderingRiskReasonAuthor string
	Compliance                 string
	ComplianceAuthor           string
	ComplianceDate             string
	ComplianceCheckDate        string
	ComplianceDescription      string
	IbankFlag                  string
	IbankFlagAuthor            string
	ActualityDateAuthor        string
	ActualityDate              string
}
