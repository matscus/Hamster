package search

import (
	"../csv"
	"../structs"
)

func Search(searchParam map[string]string) []structs.User {
	keys := make([]string, 0, len(searchParam))
	searchResult := []structs.User{}
	for k := range searchParam {
		keys = append(keys, k)
	}
	counter := 0
	for i := 0; i < len(keys); i++ {
		if counter == 0 {
			k := keys[i]
			v := searchParam[k]
			searchResult = First(k, v)
			counter++
		} else {
			k := keys[i]
			v := searchParam[k]
			searchResult = RemovеUnsuitableCondition(k, v, searchResult)
		}

	}

	return searchResult
}

func First(k string, v string) []structs.User {
	firstSearchResult := []structs.User{}
	switch k {
	case "surname":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "name":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}

		}
	case "Patronymic":
		for i := 0; i < len(csv.Datapool); i++ {

			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "Gender":
		for i := 0; i < len(csv.Datapool); i++ {

			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "GenderRawSource":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "FullNameQC":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "FullNameAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if i == 0 {
				if v == csv.Datapool[i].Surname {
					firstSearchResult = append(firstSearchResult, csv.Datapool[i])
				}
			} else {

			}
		}
	case "FullNameRawSource":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "SurnameQc":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "FirstnameQc":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "PatronymicQc":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "GenderQc":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "NameCommonQc":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "PatronymicLackFlag":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "PatronymicLackFlagAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "ForeignSurname":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "ForeignSurnameAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "ForeignName":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "Birthdate":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "BirthdateAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "BirthdateQC":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "BirthdateRawSource":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "BirthPlace":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "BirthPlaceAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "BirthCountry":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "BirthCountryAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "Citizenship":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "CitizenshipAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "CesidentFlag":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "CesidentFlagAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "MaritalStatus":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "MaritalStatusAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "Inn":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "InnAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "InnQC":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "InnRawSource":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "Snils":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "SnilsAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "SnilsQC":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "SnilsRawSource":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}

		}
	case "EmployeeFlag":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "EmployeeFiredDate":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "EmployeeFlagAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "Branch":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "BranchAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "CrossId":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "CrossIdAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "VipStatus":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "PersonalManager":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "VipStatusAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "OpenAccountsFlag":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "OpenAccountsFlagAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "AdvertisingConsent":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "AdvertisingConsentAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "DignitaryFlag":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "DignitaryDescription":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "DignitaryFlagAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "DeathDate":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "DeathDateAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "BankruptcyFlag":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "BankruptcyFlagAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "BlackListFlag":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "BlackListReason":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "BlackListDate":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "BlackListAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "LaunderingRiskReason":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "LaunderingRiskReasonAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "Compliance":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "ComplianceAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "ComplianceDate":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "ComplianceCheckDate":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "ComplianceDescription":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "IbankFlag":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "IbankFlagAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "ActualityDateAuthor":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	case "ActualityDate":
		for i := 0; i < len(csv.Datapool); i++ {
			if v == csv.Datapool[i].Surname {
				firstSearchResult = append(firstSearchResult, csv.Datapool[i])
			}
		}
	}
	return firstSearchResult
}

func RemovеUnsuitableCondition(k string, v string, users []structs.User) []structs.User {
	switch k {
	case "surname":
		for i := 0; i < len(users); i++ {
			if users[i].Surname != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "name":
		for i := 0; i < len(users); i++ {
			if users[i].Name != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "Patronymic":
		for i := 0; i < len(users); i++ {
			if users[i].Patronymic != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "Gender":
		for i := 0; i < len(users); i++ {
			if users[i].Gender != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "GenderRawSource":
		for i := 0; i < len(users); i++ {
			if users[i].GenderRawSource != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "FullNameQC":
		for i := 0; i < len(users); i++ {
			if users[i].FullNameQC != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "FullNameAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].FullNameAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "FullNameRawSource":
		for i := 0; i < len(users); i++ {
			if users[i].FullNameRawSource != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "SurnameQc":
		for i := 0; i < len(users); i++ {
			if users[i].SurnameQc != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "FirstnameQc":
		for i := 0; i < len(users); i++ {
			if users[i].FirstnameQc != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "PatronymicQc":
		for i := 0; i < len(users); i++ {
			if users[i].PatronymicQc != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "GenderQc":
		for i := 0; i < len(users); i++ {
			if users[i].GenderQc != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "NameCommonQc":
		for i := 0; i < len(users); i++ {
			if users[i].NameCommonQc != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "PatronymicLackFlag":
		for i := 0; i < len(users); i++ {
			if users[i].PatronymicLackFlag != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "PatronymicLackFlagAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].PatronymicLackFlagAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "ForeignSurname":
		for i := 0; i < len(users); i++ {
			if users[i].ForeignSurname != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "ForeignSurnameAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].ForeignSurnameAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "ForeignName":
		for i := 0; i < len(users); i++ {
			if users[i].ForeignName != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "Birthdate":
		for i := 0; i < len(users); i++ {
			if users[i].Birthdate != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "BirthdateAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].BirthdateAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "BirthdateQC":
		for i := 0; i < len(users); i++ {
			if users[i].BirthdateQC != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "BirthdateRawSource":
		for i := 0; i < len(users); i++ {
			if users[i].BirthdateRawSource != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "BirthPlace":
		for i := 0; i < len(users); i++ {
			if users[i].BirthPlace != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "BirthPlaceAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].BirthPlaceAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "BirthCountry":
		for i := 0; i < len(users); i++ {
			if users[i].BirthCountryAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "BirthCountryAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].BirthCountryAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "Citizenship":
		for i := 0; i < len(users); i++ {
			if users[i].Citizenship != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "CitizenshipAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].CitizenshipAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "CesidentFlag":
		for i := 0; i < len(users); i++ {
			if users[i].CesidentFlag != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "CesidentFlagAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].CesidentFlagAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "MaritalStatus":
		for i := 0; i < len(users); i++ {
			if users[i].MaritalStatus != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "MaritalStatusAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].MaritalStatusAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "Inn":
		for i := 0; i < len(users); i++ {
			if users[i].Inn != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "InnAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].InnAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "InnQC":
		for i := 0; i < len(users); i++ {
			if users[i].InnQC != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "InnRawSource":
		for i := 0; i < len(users); i++ {
			if users[i].InnRawSource != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "Snils":
		for i := 0; i < len(users); i++ {
			if users[i].Snils != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "SnilsAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].SnilsAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "SnilsQC":
		for i := 0; i < len(users); i++ {
			if users[i].SnilsQC != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "SnilsRawSource":
		for i := 0; i < len(users); i++ {
			if users[i].SnilsRawSource != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "EmployeeFlag":
		for i := 0; i < len(users); i++ {
			if users[i].EmployeeFlag != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "EmployeeFiredDate":
		for i := 0; i < len(users); i++ {
			if users[i].EmployeeFiredDate != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "EmployeeFlagAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].EmployeeFlagAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "Branch":
		for i := 0; i < len(users); i++ {
			if users[i].Branch != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "BranchAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].BranchAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "CrossId":
		for i := 0; i < len(users); i++ {
			if users[i].CrossId != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "CrossIdAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].CrossIdAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "VipStatus":
		for i := 0; i < len(users); i++ {
			if users[i].VipStatus != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "PersonalManager":
		for i := 0; i < len(users); i++ {
			if users[i].PersonalManager != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "VipStatusAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].VipStatusAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "OpenAccountsFlag":
		for i := 0; i < len(users); i++ {
			if users[i].OpenAccountsFlag != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "OpenAccountsFlagAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].OpenAccountsFlagAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "AdvertisingConsent":
		for i := 0; i < len(users); i++ {
			if users[i].AdvertisingConsent != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "AdvertisingConsentAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].AdvertisingConsentAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "DignitaryFlag":
		for i := 0; i < len(users); i++ {
			if users[i].DignitaryFlag != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "DignitaryDescription":
		for i := 0; i < len(users); i++ {
			if users[i].DignitaryDescription != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "DignitaryFlagAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].DignitaryFlagAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "DeathDate":
		for i := 0; i < len(users); i++ {
			if users[i].DeathDate != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "DeathDateAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].DeathDateAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "BankruptcyFlag":
		for i := 0; i < len(users); i++ {
			if users[i].BankruptcyFlag != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "BankruptcyFlagAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].BankruptcyFlagAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "BlackListFlag":
		for i := 0; i < len(users); i++ {
			if users[i].BlackListReason != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "BlackListReason":
		for i := 0; i < len(users); i++ {
			if users[i].BlackListReason != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "BlackListDate":
		for i := 0; i < len(users); i++ {
			if users[i].BlackListDate != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "BlackListAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].BlackListAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "LaunderingRiskReason":
		for i := 0; i < len(users); i++ {
			if users[i].LaunderingRiskReason != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "LaunderingRiskReasonAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].LaunderingRiskReasonAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "Compliance":
		for i := 0; i < len(users); i++ {
			if users[i].Compliance != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "ComplianceAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].ComplianceAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "ComplianceDate":
		for i := 0; i < len(users); i++ {
			if users[i].ComplianceDate != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "ComplianceCheckDate":
		for i := 0; i < len(users); i++ {
			if users[i].ComplianceCheckDate != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "ComplianceDescription":
		for i := 0; i < len(users); i++ {
			if users[i].ComplianceDescription != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "IbankFlag":
		for i := 0; i < len(users); i++ {
			if users[i].IbankFlag != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "IbankFlagAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].IbankFlagAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "ActualityDateAuthor":
		for i := 0; i < len(users); i++ {
			if users[i].ActualityDateAuthor != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	case "ActualityDate":
		for i := 0; i < len(users); i++ {
			if users[i].ActualityDate != v {
				users = append(users[:i], users[i+1:]...)
			}
		}
	}
	return users
}
