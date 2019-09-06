package main

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"./search"

	"./csv"

	"./structs"
	"github.com/gorilla/mux"
)

func init() {
	csv.ReadCsv()
}

var (
	pemPath         string
	keyPath         string
	proto           string
	searchParamPool = []string{"surname", "name", "patronymic", "gender", "genderRawSource", "fullNameQC", "fullNameAuthor", "fullNameRawSource", "surnameQc", "firstnameQc", "patronymicQc", "genderQc", "nameCommonQc", "patronymicLackFlag", "patronymicLackFlagAuthor", "foreignSurname", "foreignSurnameAuthor", "foreignName", "birthdate", "birthdateAuthor", "birthdateQC", "birthdateRawSource", "birthPlace", "birthPlaceAuthor", "birthCountry", "birthCountryAuthor", "citizenship", "citizenshipAuthor", "residentFlag", "residentFlagAuthor", "maritalStatus", "maritalStatusAuthor", "inn", "innAuthor", "innQC", "innRawSource", "snils", "snilsAuthor", "snilsQC", "snilsRawSource", "employeeFlag", "employeeFiredDate", "employeeFlagAuthor", "branch", "branchAuthor", "crossId", "crossIdAuthor", "vipStatus", "personalManager", "vipStatusAuthor", "openAccountsFlag", "openAccountsFlagAuthor", "advertisingConsent", "advertisingConsentAuthor", "dignitaryFlag", "dignitaryDescription", "dignitaryFlagAuthor", "deathDate", "deathDateAuthor", "bankruptcyFlag", "bankruptcyFlagAuthor", "blackListFlag", "blackListReason", "blackListDate", "blackListAuthor", "launderingRiskReason", "launderingRiskReasonAuthor", "compliance", "complianceAuthor", "complianceDate", "complianceCheckDate", "complianceDescription", "ibankFlag", "ibankFlagAuthor", "actualityDateAuthor", "actualityDate"}
)

func main() {
	flag.StringVar(&pemPath, "pem", "./ssl/server.pem", "path to pem file")
	flag.StringVar(&keyPath, "key", "./ssl/server.key", "path to key file")
	flag.StringVar(&proto, "proto", "https", "Proxy protocol (http or https)")
	flag.Parse()
	fmt.Println("start serve")
	router := mux.NewRouter()
	router.HandleFunc("/cdi/soap/services/2_13/PartyWS", PartyWS)
	http.Handle("/", router)
	fmt.Println("Listen to port")
	//http.ListenAndServe(":9999", context.ClearHandler(http.DefaultServeMux))
	server := &http.Server{
		Addr: ":9999",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			PartyWS(w, r)
		}),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	if proto == "http" {
		log.Fatal(server.ListenAndServe())
	} else {
		log.Fatal(server.ListenAndServeTLS(pemPath, keyPath))
	}
}

func PartyWS(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}
	var res structs.Request
	err = xml.Unmarshal(body, &res)
	if err != nil {
		fmt.Println(err)
	}
	resParse := Parse(res.Body.Search.Query)
	searchResult := search.Search(resParse)
	var buffer bytes.Buffer
	buffer.Write(structs.SoapHead)
	for i := 0; i < len(searchResult); i++ {
		var soapBody = []byte(
			`<field name="surname">` + searchResult[i].Surname + `</field>` +
				`<field name="name">` + searchResult[i].Name + `</field>` +
				`<field name="patronymic">` + searchResult[i].Patronymic + `</field>` +
				`<field name="gender">` + searchResult[i].Gender + `</field>\` +
				`<field name="genderRawSource">` + searchResult[i].GenderRawSource + `</field>` +
				`<field name="fullNameQC">` + searchResult[i].FullNameQC + `</field>` +
				`<field name="fullNameAuthor">` + searchResult[i].FullNameAuthor + `</field>` +
				`<field name="fullNameRawSource">` + searchResult[i].Surname + `•` + searchResult[i].Name + `•` + searchResult[i].Patronymic + `</field>` +
				`<field name="surnameQc">Ok not changed</field>` +
				`<field name="firstnameQc">Ok not changed</field>` +
				`<field name="patronymicQc">Ok not changed</field>` +
				`<field name="genderQc">Ok not changed</field>` +
				`<field name="nameCommonQc">Ok not changed</field>` +
				`<field name="patronymicLackFlag">` + searchResult[i].PatronymicLackFlag + `</field>` +
				`<field name="patronymicLackFlagAuthor"/>` +
				`<field name="foreignSurname"/>` +
				`<field name="foreignSurnameAuthor"/>` +
				`<field name="foreignName"/>` +
				`<field name="birthdate">` + searchResult[i].Birthdate + `</field>` +
				`<field name="birthdateAuthor">` + searchResult[i].BirthdateAuthor + `</field>` +
				`<field name="birthdateQC">` + searchResult[i].BirthdateQC + `</field>` +
				`<field name="birthdateRawSource">` + searchResult[i].BirthdateRawSource + `</field>` +
				`<field name="birthPlace">` + searchResult[i].BirthPlace + `</field>` +
				`<field name="birthPlaceAuthor">` + searchResult[i].BirthPlaceAuthor + `</field>` +
				`<field name="birthCountry">` + searchResult[i].BirthCountry + `</field>` +
				`<field name="birthCountryAuthor">` + searchResult[i].BirthCountryAuthor + `</field>` +
				`<field name="citizenship">` + searchResult[i].Citizenship + `</field>` +
				`<field name="citizenshipAuthor">` + searchResult[i].CitizenshipAuthor + `</field>` +
				`<field name="residentFlag">TRUE</field>` +
				`<field name="residentFlagAuthor">DR34:10016865568</field>` +
				`<field name="maritalStatus">` + searchResult[i].MaritalStatus + `</field>` +
				`<field name="maritalStatusAuthor">` + searchResult[i].MaritalStatusAuthor + `</field>` +
				`<field name="inn"/>` +
				`<field name="innAuthor">` + searchResult[i].InnAuthor + `</field>` +
				`<field name="innQC">EMPTY</field>` +
				`<field name="innRawSource"/>` +
				`<field name="snils"/>` +
				`<field name="snilsAuthor">` + searchResult[i].SnilsAuthor + `</field>` +
				`<field name="snilsQC">EMPTY</field>` +
				`<field name="snilsRawSource"/>` +
				`<field name="employeeFlag">UNKNOWN</field>` +
				`<field name="employeeFiredDate"/>` +
				`<field name="employeeFlagAuthor"/>` +
				`<field name="branch">FIL_KRYAR</field>` +
				`<field name="branchAuthor">DR34:10016865568</field>` +
				`<field name="crossId">4646707</field>` +
				`<field name="crossIdAuthor">DR34:10016865568</field>` +
				`<field name="vipStatus">REGULAR</field>` +
				`<field name="personalManager"/>` +
				`<field name="vipStatusAuthor"/>` +
				`<field name="openAccountsFlag">TRUE</field>` +
				`<field name="openAccountsFlagAuthor">DR34:10016865568</field>` +
				`<field name="advertisingConsent">TRUE</field>` +
				`<field name="advertisingConsentAuthor">DR34:10016865568</field>` +
				`<field name="dignitaryFlag">FALSE</field>` +
				`<field name="dignitaryDescription"/>` +
				`<field name="dignitaryFlagAuthor">DR34:10016865568</field>` +
				`<field name="deathDate"/>` +
				`<field name="deathDateAuthor"/>` +
				`<field name="bankruptcyFlag">FALSE</field>` +
				`<field name="bankruptcyFlagAuthor">DR34:10016865568</field>` +
				`<field name="blackListFlag">FALSE</field>` +
				`<field name="blackListReason"/>` +
				`<field name="blackListDate"/>` +
				`<field name="blackListAuthor">DR34:10016865568</field>` +
				`<field name="launderingRiskReason"/>` +
				`<field name="launderingRiskReasonAuthor"/>` +
				`<field name="compliance">UNKNOWN</field>` +
				`<field name="complianceAuthor"/>` +
				`<field name="complianceDate"/>` +
				`<field name="complianceCheckDate"/>` +
				`<field name="complianceDescription"/>` +
				`<field name="ibankFlag">FALSE</field>` +
				`<field name="ibankFlagAuthor">` + searchResult[i].IbankFlagAuthor + `</field>` +
				`<field name="actualityDateAuthor">` + searchResult[i].ActualityDateAuthor + `</field>` +
				`<field name="actualityDate">` + searchResult[i].ActualityDate + `</field>`)
		buffer.Write(soapBody)
	}
	buffer.Write(structs.SoapClose)
	w.Header().Set("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
	w.Header().Set("X-Powered-By", "Undertow/1")
	w.Header().Set("Server", "WildFly/10")
	w.Header().Set("X-Powered-By", "Undertow/1")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Date", "no-cache")
	w.Header().Set("Connection", "close")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Set-Cookie", "SERVERNAME=10.96.67.125|XByIO; path=//*//")
	w.Write(buffer.Bytes())
}

func Parse(str string) map[string]string {
	result := make(map[string]string)
	for i := 0; i < len(searchParamPool); i++ {
		var re = regexp.MustCompile(`\.` + searchParamPool[i] + `=(.\S+)`)
		if len(re.FindStringIndex(str)) > 0 {
			// fmt.Println(re.FindString(str), "index: ", re.FindStringIndex(str)[0])
			r := strings.Replace(re.FindString(str), "."+searchParamPool[i]+"=", "", 1)
			result[searchParamPool[i]] = r
		}
	}
	return result
}
