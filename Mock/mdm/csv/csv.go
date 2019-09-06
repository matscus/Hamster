package csv

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"

	"../structs"

	"../errors"
)

var Datapool []structs.User

func ReadCsv() {
	file, err := os.Open("/home/matscus/work/stend/plugs/mdm/csv/user.csv")
	errors.CheckError(err, "fail to open csv file")
	defer file.Close()
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		r, err := reader.Read()
		if err == io.EOF {
			break
		}
		errors.CheckError(err, "fail to read rows")
		Datapool = append(Datapool, structs.User{Surname: r[0], Name: r[1], Patronymic: r[2], Gender: r[3], GenderRawSource: r[4], FullNameQC: r[5], FullNameAuthor: r[6], FullNameRawSource: r[7], SurnameQc: r[8], FirstnameQc: r[9], PatronymicQc: r[10], GenderQc: r[11], NameCommonQc: r[12], PatronymicLackFlag: r[13], PatronymicLackFlagAuthor: r[14], ForeignSurname: r[15], ForeignSurnameAuthor: r[16], ForeignName: r[17], Birthdate: r[18], BirthdateAuthor: r[19], BirthdateQC: r[20], BirthdateRawSource: r[21], BirthPlace: r[22], BirthPlaceAuthor: r[23], BirthCountry: r[24], BirthCountryAuthor: r[25], Citizenship: r[26], CitizenshipAuthor: r[27], CesidentFlag: r[28], CesidentFlagAuthor: r[29], MaritalStatus: r[30], MaritalStatusAuthor: r[31], Inn: r[32], InnAuthor: r[33], InnQC: r[34], InnRawSource: r[35], Snils: r[36], SnilsAuthor: r[37], SnilsQC: r[38], SnilsRawSource: r[39], EmployeeFlag: r[40], EmployeeFiredDate: r[41], EmployeeFlagAuthor: r[42], Branch: r[43], BranchAuthor: r[44], CrossId: r[45], CrossIdAuthor: r[46], VipStatus: r[47], PersonalManager: r[48], VipStatusAuthor: r[49], OpenAccountsFlag: r[50], OpenAccountsFlagAuthor: r[51], AdvertisingConsent: r[52], AdvertisingConsentAuthor: r[53], DignitaryFlag: r[54], DignitaryDescription: r[55], DignitaryFlagAuthor: r[56], DeathDate: r[57], DeathDateAuthor: r[58], BankruptcyFlag: r[59], BankruptcyFlagAuthor: r[60], BlackListFlag: r[61], BlackListReason: r[62], BlackListDate: r[63], BlackListAuthor: r[64], LaunderingRiskReason: r[65], LaunderingRiskReasonAuthor: r[66], Compliance: r[67], ComplianceAuthor: r[68], ComplianceDate: r[69], ComplianceCheckDate: r[70], ComplianceDescription: r[71], IbankFlag: r[72], IbankFlagAuthor: r[73], ActualityDateAuthor: r[74], ActualityDate: r[75]})
	}
}
