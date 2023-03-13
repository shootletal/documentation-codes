package types

// BranchesDynamo struct info.
type BranchesDynamo struct {
	Currency                  string        `json:"currency,omitempty" dynamodbav:"currency,omitEmpty"`
	CountryTaxResidency       string        `json:"countryTaxResidency" dynamodbav:"countryTaxResidency"`
	Name                      string        `json:"name,omitempty" dynamodbav:"name,omitEmpty"`
	CompanyRegistrationNumber string        `json:"companyRegistrationNumber" dynamodbav:"companyRegistrationNumber"`
	TaxID                     string        `json:"taxId,omitempty" dynamodbav:"taxId,omitEmpty"`
	LegalName                 string        `json:"legalName,omitempty" dynamodbav:"legalName,omitEmpty"`
	LegalEntityType           string        `json:"legalEntityType" dynamodbav:"legalEntityType"`
	MCC                       int           `json:"mcc,omitempty" dynamodbav:"mcc,omitEmpty"`
	LegalRepresentative       []PersonLegal `json:"legalRepresentative,omitempty" dynamodbav:"legalRepresentative,omitempty"`
	Shareholder               PersonLegal   `json:"shareholder,omitempty" dynamodbav:"shareholder,omitempty"`
	AMLRiskCategory           string        `json:"AMLRiskCategory,omitempty" dynamodbav:"AMLRiskCategory,omitempty"`
	SettlementBankCountry     string        `json:"settlementBankCountry,omitempty" dynamodbav:"settlementBankCountry,omitempty"`
	SettlementBankName        string        `json:"settlementBankName,omitempty" dynamodbav:"settlementBankName,omitempty"`
	RegulatedBusiness         string        `json:"regulatedBusiness,omitempty" dynamodbav:"regulatedBusiness,omitempty"`
	ConcurrenceNeeded         string        `json:"concurrenceNeeded,omitempty" dynamodbav:"concurrenceNeeded,omitempty"`
	URL                       string        `json:"url,omitempty" dynamodbav:"url,omitempty"`
	ClientType                string        `json:"clientType,omitempty" dynamodbav:"clientType,omitEmpty"`
	ReportingDate             string        `json:"reportingDate,omitempty" dynamodbav:"reportingDate,omitEmpty"`
	PspTaxID                  string        `json:"pspTaxId,omitempty" dynamodbav:"pspTaxId,omitEmpty"`
	PspName                   string        `json:"pspName,omitempty" dynamodbav:"pspName,omitEmpty"`
	PspCustomerID             string        `json:"pspCustomerId,omitempty" dynamodbav:"pspCustomerId,omitEmpty"`
	PspNotificationEmail      string        `json:"pspNotificationEmail" dynamodbav:"pspNotificationEmail"`
	PspUserEmail              string        `json:"pspUserEmail" dynamodbav:"pspUserEmail"`
	KushkiStatus              string        `json:"kushkiStatus,omitempty" dynamodbav:"kushkiStatus,omitEmpty"`
	RiskMCC                   string        `json:"riskMCC,omitempty" dynamodbav:"riskMCC,omitEmpty"`
	ID                        string        `json:"id,omitempty" dynamodbav:"id,omitEmpty"`
	ListWC1                   []string      `json:"listWC1,omitempty" dynamodbav:"listWC1,omitEmpty"`
	StatusWC1                 string        `json:"statusWC1,omitempty" dynamodbav:"statusWC1,omitEmpty"`
	CaseSystemId              string        `json:"caseSystemId,omitempty" dynamodbav:"caseSystemId,omitEmpty"`
	ErrorWC1                  string        `json:"errorWC1,omitempty" dynamodbav:"errorWC1,omitEmpty"`
	CreatedAt                 int64         `json:"createdAt,omitempty" dynamodbav:"createdAt,omitEmpty"`
	RiskCountryGafiFatf       string        `json:"riskCountryGafiFatf,omitempty" dynamodbav:"riskCountryGafiFatf,omitEmpty"`
	RiskBankGafiFatf          string        `json:"riskBankGafiFatf,omitempty" dynamodbav:"riskBankGafiFatf,omitEmpty"`
	Coincidencia              string        `json:"coincidencia,omitempty" dynamodbav:"coincidencia,omitEmpty"`
	BranchId                  string        `json:"branchId,omitempty" dynamodbav:"branchId,omitEmpty"`
	Demo                      []PersonLegal `json:"demo,omitempty" dynamodbav:"demo,omitempty"`
}

// PersonLegal struct info.
type PersonLegal struct {
	FirstName        string      `json:"firstName,omitempty" dynamodbav:"firstName,omitEmpty"`
	LastName         string      `json:"lastName,omitempty" dynamodbav:"lastName,omitEmpty"`
	IdNumber         string      `json:"idNumber" dynamodbav:"idNumber"`
	Id               string      `json:"id,omitempty" dynamodbav:"id,omitempty"`
	BirthDate        string      `json:"birthDate" dynamodbav:"birthdate"`
	ResidencyCountry string      `json:"residencyCountry,omitempty" dynamodbav:"residencyCountry,omitempty"`
	ListWC1          []string    `json:"listWC1,omitempty" dynamodbav:"listWC1,omitEmpty"`
	StatusWC1        string      `json:"statusWC1,omitempty" dynamodbav:"statusWC1,omitEmpty"`
	CaseSystemId     string      `json:"caseSystemId,omitempty" dynamodbav:"caseSystemId,omitEmpty"`
	ErrorWC1         string      `json:"errorWC1,omitempty" dynamodbav:"errorWC1,omitEmpty"`
	SecondLvl        []AuxStruct `json:"secondLvl,omitempty" dynamodbav:"secondLvl,omitempty"`
}

// Shareholder struct info.
type Shareholder struct {
	FirstName        string   `json:"firstName" dynamodbav:"firstName"`
	LastName         string   `json:"lastName" dynamodbav:"lastName"`
	IdNumber         string   `json:"idNumber" dynamodbav:"idNumber"`
	Id               string   `json:"id,omitempty" dynamodbav:"id,omitempty"`
	BirthDate        string   `json:"birthDate" dynamodbav:"birthdate"`
	ResidencyCountry string   `json:"residencyCountry" dynamodbav:"residencyCountry"`
	ListWC1          []string `json:"listWC1,omitempty" dynamodbav:"listWC1,omitEmpty"`
	StatusWC1        string   `json:"statusWC1,omitempty" dynamodbav:"statusWC1,omitEmpty"`
	CaseSystemId     string   `json:"caseSystemId,omitempty" dynamodbav:"caseSystemId,omitEmpty"`
	ErrorWC1         string   `json:"errorWC1,omitempty" dynamodbav:"errorWC1,omitEmpty"`
}

type AuxStruct struct {
	FirstName string `json:"firstName" dynamodbav:"firstName"`
	LastName  string `json:"lastName" dynamodbav:"lastName"`
	BirthDate string `json:"birthDate" dynamodbav:"birthdate"`
}
