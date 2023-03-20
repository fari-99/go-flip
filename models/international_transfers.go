package models

import (
	"github.com/fari-99/flip/requests"
)

type ExchangeRateArrivalMessage struct {
	Id string `json:"id"`
	En string `json:"en"`
}

/*
ExchangeRates Model
| Attribute          | Description                                                                       |
|--------------------|-----------------------------------------------------------------------------------|
| currency_code      | Currency code of the country                                                      |
| country_code       | Country code using ISO 3166 (Alpha-2)                                             |
| country_name       | Name of the country                                                               |
| country_iso_code   | Country code using ISO 3166 (Alpha-3)                                             |
| flip_exchange_rate | The country exchange rates to IDR                                                 |
| flip_transfer_fee  | Transfer fee to the country                                                       |
| minimum_amount     | Minimum transfer amount                                                           |
| maximum_amount     | Maximum transfer amount                                                           |
| payment_speed      | Expected payment completion speed: Real TimeSame Day +1 day or “+N day”           |
| arrival_message    | Message related to payment speed, available in Bahasa (ID) and English (EN)       |
| notes              | Additional notes                                                                  |
| transaction_type   | Transaction type information                                                      |
| flip_cutoff_time   | Transfer cutoff time. Time will be in GMT+7 with yyyy-mm-dd hh:mm:ss format       |
| flip_arrival_time  | Estimation of arrival time. Time will be in GMT+7 with yyyy-mm-dd hh:mm:ss format |
*/
type ExchangeRates struct {
	CurrencyCode     string      `json:"currency_code"`
	CountryCode      string      `json:"country_code"`
	CountryName      string      `json:"country_name"`
	CountryIsoCode   string      `json:"country_iso_code"`
	FlipExchangeRate float64     `json:"flip_exchange_rate"`
	FlipTransferFee  int         `json:"flip_transfer_fee"`
	MinimumAmount    float64     `json:"minimum_amount"`
	MaximumAmount    float64     `json:"maximum_amount"`
	PaymentSpeed     string      `json:"payment_speed"`
	ArrivalMessage   interface{} `json:"arrival_message"`
	Notes            *string     `json:"notes"`
	TransactionType  string      `json:"transaction_type"`
	FlipCutoffTime   string      `json:"flip_cutoff_time"`
	FlipArrivalTime  string      `json:"flip_arrival_time"`
}

/*
InternationalTransfer
| Attribute                 | Description                                                                                                                                |
|---------------------------|--------------------------------------------------------------------------------------------------------------------------------------------|
| id                        | Id of international transfer                                                                                                               |
| user_id                   | Id creator of transaction                                                                                                                  |
| company_id                | Id of company                                                                                                                              |
| exchange_rate             | The country exchange rates to IDR                                                                                                          |
| fee                       | Transfer fee to the country                                                                                                                |
| amount                    | Transfer amount in IDR                                                                                                                     |
| source_country            | Source country code using ISO 3166 (Alpha-3)                                                                                               |
| destination_country       | Destination country code using ISO 3166 (Alpha-3)                                                                                          |
| beneficiary_amount        | Amount to be received by beneficiary                                                                                                       |
| beneficiary_currency_code | Currency code of the beneficiary                                                                                                           |
| status                    | Transaction status. Possible values are: PENDING,CANCELLED,DONE                                                                              |
| timestamp                 | The time when the disbursement request was created. Time will be in GMT+7 with yyyy-mm-dd hh:mm:ss format                                  |
| time_served               | The time when the disbursement is finished                                                                                                 |
| created_from              | The channel of which the transaction was created. Possible values are: APIDashboard                                                        |
| receipt                   | URL of the transfer receipt                                                                                                                |
| transaction_type          | Transaction type information. Possible values are: C2CC2BB2CB2B                                                                            |
| idempotency_key           | Idempotency key information                                                                                                                |
| beneficiary               | Beneficiary object                                                                                                                         |
| sender                    | Sender object                                                                                                                              |
| reason                    | The reason value will be given if the transaction status is CANCELLED. Possible values can be seen in Cancelled Transaction Reasons table. |
*/
type InternationalTransfer struct {
	Id                      int64   `json:"id"`
	UserId                  int64   `json:"user_id"`
	CompanyId               int64   `json:"company_id"`
	ExchangeRate            float64 `json:"exchange_rate"`
	Fee                     int64   `json:"fee"`
	Amount                  int64   `json:"amount"`
	SourceCountry           string  `json:"source_country"`
	DestinationCountry      string  `json:"destination_country"`
	BeneficiaryAmount       float64 `json:"beneficiary_amount"`
	BeneficiaryCurrencyCode string  `json:"beneficiary_currency_code"`
	Status                  string  `json:"status"`
	Timestamp               string  `json:"timestamp"`
	TimeServed              string  `json:"time_served"`
	CreatedFrom             string  `json:"created_from"`
	Receipt                 string  `json:"receipt"`
	TransactionType         string  `json:"transaction_type"`
	IdempotencyKey          string  `json:"idempotency_key"`
	Reason                  string  `json:"reason"`

	Beneficiary InterTransferBeneficiary `json:"beneficiary"`
	Sender      InterTransferSender      `json:"sender"`
}

/*
InterTransferBeneficiary
| Attribute                 | Description                                                                |
|---------------------------|----------------------------------------------------------------------------|
| id_number                 | Id number                                                                  |
| id_expiration_date        | ID expiration date using YYYY-MM-DD&nbsp;format                            |
| full_name                 | Full name of the beneficiary                                               |
| bank_account_number       | Account number of the beneficiary bank                                     |
| bank                      | Bank name of beneficiary bank                                              |
| email                     | Email of beneficiary                                                       |
| msisdn                    | A number used to identify a phone number internationally                   |
| nationality               | Country code using ISO 3166 (Alpha-3) of beneficiary nationality           |
| country                   | Country code using ISO 3166 (Alpha-3) of the beneficiary’s country         |
| province                  | Province or state of the beneficiary                                       |
| city                      | City of the beneficiary                                                    |
| address                   | Address of the beneficiary                                                 |
| postal_code               | Postal code of the beneficiary                                             |
| relationship              | Relationship of sender and beneficiary                                     |
| source_of_funds           | Source of funds                                                            |
| remittance_purpose        | Purpose of remittance                                                      |
| iban                      | Identifier of individual account involved in the international transaction |
| swift_bic_code            | Identifier of specific bank during an international transaction            |
| sort_code                 | Digits of code which is used by British and Irish banks                    |
| ifs_code                  | Indian financial system code                                               |
| bsb_number                | Identifier of banks and branches across Australia                          |
| branch_number             | Identifier of bank for Japan                                               |
| document_reference_number | Reference number of document related to the transaction                    |
| registration_number       | Registration number                                                        |
*/
type InterTransferBeneficiary struct {
	IdNumber                string  `json:"id_number"`
	IdExpirationDate        string  `json:"id_expiration_date"`
	AchCode                 string  `json:"ach_code"`
	FullName                string  `json:"full_name"`
	BankAccountNumber       string  `json:"bank_account_number"`
	Bank                    string  `json:"bank"`
	Email                   *string `json:"email"`
	Msisdn                  string  `json:"msisdn"`
	Nationality             string  `json:"nationality"`
	Country                 string  `json:"country"`
	Province                string  `json:"province"`
	City                    string  `json:"city"`
	Address                 string  `json:"address"`
	PostalCode              *string `json:"postal_code"`
	Relationship            string  `json:"relationship"`
	SourceOfFunds           string  `json:"source_of_funds"`
	RemittancePurpose       string  `json:"remittance_purpose"`
	Iban                    *string `json:"iban"`
	SwiftBicCode            *string `json:"swift_bic_code"`
	SortCode                *string `json:"sort_code"`
	IfsCode                 *string `json:"ifs_code"`
	BsbNumber               *string `json:"bsb_number"`
	BranchNumber            *string `json:"branch_number"`
	DocumentReferenceNumber string  `json:"document_reference_number"`
	RegistrationNumber      *string `json:"registration_number"`
}

/*
InterTransferSender
| Attribute       | Description                                                                                                                                |
|-----------------|--------------------------------------------------------------------------------------------------------------------------------------------|
| name            | The name of the user of the Money Transfer Company that act as a sender                                                                    |
| place_of_birth  | City/country code of the Sender’s place of birth                                                                                           |
| date_of_birth   | Sender’s date of birth                                                                                                                     |
| address         | Sender’s address                                                                                                                           |
| identity_type   | Sender’s identity type. Possible values are: nat_idpassport                                                                                |
| identity_number | Sender’s identity number                                                                                                                   |
| country         | Country code of the Sender’s country                                                                                                       |
| job             | Sender’s job. Possible values are: housewifeentrepreneurprivate_employeegovernment_employeefoundation_boardindonesian_migrant_workerothers |
| email           | Sender’s email                                                                                                                             |
| city            | Sender’s city                                                                                                                              |
| phone_number    | Sender’s phone number                                                                                                                      |
*/
type InterTransferSender struct {
	Name           string `json:"name"`
	PlaceOfBirth   int64  `json:"place_of_birth"`
	DateOfBirth    string `json:"date_of_birth"`
	Address        string `json:"address"`
	IdentityType   string `json:"identity_type"`
	IdentityNumber string `json:"identity_number"`
	Country        int64  `json:"country"`
	Job            string `json:"job"`
	Email          string `json:"email"`
	City           string `json:"city"`
	PhoneNumber    string `json:"phone_number"`
}

// ---------------------------- REQUEST/RESPONSE
//----------------------------- Get Exchange Rates

/*
GetExchangeRatesRequest
| Param            | Value                                                                                                                                           |
|------------------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| country_iso_code | string (optional)  Country code using ISO 3166 (Alpha-3). You can choose multiple countries separate by comma. See supported Country ISO Codes. |
| transaction_type | string (required)  Transaction type information. Possible values are: C2CC2BB2CB2B                                                              |

Example GetExchangeRatesRequest:
[URL]?country_iso_code=country_iso_code&transaction_type=transaction_type
*/
type GetExchangeRatesRequest struct {
	CountryIsoCode  string `json:"country_iso_code,omitempty"`
	TransactionType string `json:"transaction_type"`
}

/*
GetExchangeRateResponse
Detail params on ExchangeRates
*/
type GetExchangeRateResponse ExchangeRates

//----------------------------- Get Form Data

/*
GetFormDataRequest
| Param            | Value                                                                                                                                           |
|------------------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| country_iso_code | string (optional)  Country code using ISO 3166 (Alpha-3). You can choose multiple countries separate by comma. See supported Country ISO Codes. |
| transaction_type | string (required)  Transaction type information. Possible values are: C2CC2BB2CB2B                                                              |

Example GetFormDataRequest:
[URL]?country_iso_code=country_iso_code&transaction_type=transaction_type
*/
type GetFormDataRequest struct {
	CountryIsoCode  string `json:"country_iso_code,omitempty"`
	TransactionType string `json:"transaction_type"`
}

/*
GetFormDataResponse
| Attribute                 | Description                              |
|---------------------------|------------------------------------------|
| country_iso_code          | Country code using ISO 3166 (Alpha-3)    |
| currency_code             | Currency code of the country             |
| beneficiary_relationships | Object of relationship of beneficiary    |
| source_of_funds           | Object of source of funds                |
| remittance_purposes       | Object of purpose of remittance          |
| banks                     | Object of bank in specific country       |
| special_identifiers       | Required identifier for specific country |
| regions                   | Array of region object                   |
| nationality_countries     | Array of country object                  |
*/
type GetFormDataResponse struct {
	CountryIsoCode           string            `json:"country_iso_code"`
	CurrencyCode             string            `json:"currency_code"`
	BeneficiaryRelationships map[string]string `json:"beneficiary_relationships"`
	SourceOfFunds            map[string]string `json:"source_of_funds"`
	RemittancePurposes       map[string]string `json:"remittance_purposes"`
	Banks                    map[string]string `json:"banks"`
	SpecialIdentifiers       []string          `json:"special_identifiers"`
	Regions                  []interface{}     `json:"regions"`
	NationalityCountries     []interface{}     `json:"nationality_countries"`
}

//----------------------------- Create International Transfer C2C/C2B

/*
CreateInternationalTransferC2CC2BRequest
| Attribute                             | Description                                                                                                                                                                                                                                                                     |
|---------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| id_number                             | string (required, if the destination country need it) Id number.                                                                                                                                                                                                                |
| id_expiration_date                    | string (required, if the destination country need it) ID expiration date using YYYY-MM-DD&nbsp;format.                                                                                                                                                                          |
| amount                                | string (required) Amount of money that wants to be disbursed to the beneficiary (in beneficiary currency). Maximum 2 decimal places, for Japan there can be no decimal. Must be within the minimum and maximum amount based on transaction currency.                            |
| source_country                        | string (required) Source country code using ISO 3166 (Alpha-3). Currently possible value is IDN.                                                                                                                                                                                |
| destination_country                   | string (required) Destination country code using ISO 3166 (Alpha-3). See supported Destination Countries.                                                                                                                                                                       |
| transaction_type                      | string (required) Transaction type information. Possible values are: C2CC2B                                                                                                                                                                                                     |
| beneficiary_full_name                 | string (required) Full name of the beneficiary                                                                                                                                                                                                                                  |
| beneficiary_account_number            | string (required) Account number of the beneficiary bank                                                                                                                                                                                                                        |
| beneficiary_bank_id                   | string (required) ID of beneficiary bank                                                                                                                                                                                                                                        |
| beneficiary_bank_name                 | string (optional) Name of beneficiary bank. Can be used for destination countries AUS and GBR                                                                                                                                                                                   |
| beneficiary_email                     | string (optional) Email of beneficiary                                                                                                                                                                                                                                          |
| beneficiary_msisdn                    | string (required, if the destination country need it) A number used to identify a phone number internationally                                                                                                                                                                  |
| beneficiary_nationality               | string (required) Country code using ISO 3166 (Alpha-3) of beneficiary nationality                                                                                                                                                                                              |
| beneficiary_province                  | string (required) Province or state of the beneficiary                                                                                                                                                                                                                          |
| beneficiary_city                      | string (required) City of the beneficiary                                                                                                                                                                                                                                       |
| beneficiary_address                   | string (required) Address of the beneficiary                                                                                                                                                                                                                                    |
| beneficiary_postal_code               | string (optional) Postal code of the beneficiary                                                                                                                                                                                                                                |
| beneficiary_relationship              | string (required) Relationship of beneficiary                                                                                                                                                                                                                                   |
| beneficiary_source_of_funds           | string (required) Source of funds                                                                                                                                                                                                                                               |
| beneficiary_remittance_purposes       | string (required) Purpose of remittance                                                                                                                                                                                                                                         |
| beneficiary_iban                      | string (required, if the destination country need it) Identifier of individual account involved in the international transaction                                                                                                                                                |
| beneficiary_swift_bic_code            | string (required, if the destination country need it) Identifier of specific bank during an international transaction                                                                                                                                                           |
| beneficiary_sort_code                 | string (required, if the destination country need it) Digits of code which is used by British and Irish banks                                                                                                                                                                   |
| beneficiary_ifs_code                  | string (required, if the destination country need it) Indian financial system code                                                                                                                                                                                              |
| beneficiary_bsb_number                | string (required, if the destination country need it) Identifier of banks and branches across Australia                                                                                                                                                                         |
| beneficiary_branch_number             | string (required, if the destination country need it) Identifier of banks and branches across Japan                                                                                                                                                                             |
| beneficiary_document_reference_number | string (required, if the destination country need it) Reference number of document related to the transaction                                                                                                                                                                   |
| beneficiary_registration_number       | string (required, if the destination country need it) Registration number                                                                                                                                                                                                       |
| beneficiary_region                    | string (required, if the destination country need it) Required for destination country China code.                                                                                                                                                                              |
| sender_name                           | string (required) The name of the user of the Money Transfer Company that act as a sender                                                                                                                                                                                       |
| sender_country                        | integer (required) Country code of the sender’s residence. Available value can be retrieved from country list.                                                                                                                                                                  |
| sender_place_of_birth                 | integer (required) City/country code of the sender’s place of birth. Use city code if the sender’s place of birth is in Indonesia, and country code if outside Indonesia. Available value can be retrieved from city/country list.                                              |
| sender_date_of_birth                  | string (required) Sender’s date of birth with YYYY-MM-DD format                                                                                                                                                                                                                 |
| sender_address                        | string (required) Sender’s address                                                                                                                                                                                                                                              |
| sender_identity_type                  | string (required) Sender’s ID type. Accepted value are: nat_id (National Id Card or KTP in Indonesia)passport (Passport)                                                                                                                                                        |
| sender_identity_number                | string (required) Sender’s ID number                                                                                                                                                                                                                                            |
| sender_job                            | string (required) Sender’s job. Accepted values are: housewifeentrepreneurprivate_employeegovernment_employeefoundation_board (People who work at foundation as an employee)indonesian_migrant_worker (Also known as TKI)company (If sender is a company, not individual)others |
| sender_email                          | string (required) Sender’s email. Only accept one email.                                                                                                                                                                                                                        |
| sender_city                           | string (required) Sender’s city.                                                                                                                                                                                                                                                |
| sender_phone_number                   | string (required) Sender’s phone number. Start with country phone code.                                                                                                                                                                                                         |


Example CreateInternationalTransferC2CC2BRequest:
{
  "id_number":"44943733088",
  "id_expiration_date":"2022-11-29",
  "amount":"52.63",
  "source_country":"IDN",
  "destination_country":"GBR",
  "transaction_type":"C2C",
  "beneficiary_full_name":"John Smith",
  "beneficiary_account_number":"1122333300",
  "beneficiary_bank_id":"1807",
  "beneficiary_bank_name":"Allica Bank",
  "beneficiary_email":"emial@email.com",
  "beneficiary_msisdn":"09584008222",
  "beneficiary_nationality":"IDN",
  "beneficiary_province":"United Kingdom",
  "beneficiary_city":"Manchester",
  "beneficiary_address":"Mosley St Manchester",
  "beneficiary_postal_code":"M2",
  "beneficiary_relationship":"SON",
  "beneficiary_source_of_funds":"BUSINESS",
  "beneficiary_remittance_purposes":"EDUCATION",
  "beneficiary_sort_code":"506967",
  "sender_name":"John Doe",
  "sender_country":100252,
  "sender_place_of_birth":100230,
  "sender_date_of_birth":"1963-12-01",
  "sender_address":"Some Address Street 123",
  "sender_identity_type":"nat_id",
  "sender_identity_number":"1234789012347890",
  "sender_job":"entrepreneur",
  "sender_email":"sender@mail.com",
  "sender_city":"Sender City",
  "sender_phone_number":" 628123456789",
  "beneficiary_region":"10274"
}
*/
type CreateInternationalTransferC2CC2BRequest struct {
	IdNumber                           string `json:"id_number"`
	IdExpirationDate                   string `json:"id_expiration_date"`
	Amount                             string `json:"amount"`
	SourceCountry                      string `json:"source_country"`
	DestinationCountry                 string `json:"destination_country"`
	TransactionType                    string `json:"transaction_type"`
	BeneficiaryFullName                string `json:"beneficiary_full_name"`
	BeneficiaryAccountNumber           string `json:"beneficiary_account_number"`
	BeneficiaryBankId                  string `json:"beneficiary_bank_id"`
	BeneficiaryBankName                string `json:"beneficiary_bank_name"`
	BeneficiaryEmail                   string `json:"beneficiary_email"`
	BeneficiaryMsisdn                  string `json:"beneficiary_msisdn"`
	BeneficiaryNationality             string `json:"beneficiary_nationality"`
	BeneficiaryProvince                string `json:"beneficiary_province"`
	BeneficiaryCity                    string `json:"beneficiary_city"`
	BeneficiaryAddress                 string `json:"beneficiary_address"`
	BeneficiaryPostalCode              string `json:"beneficiary_postal_code"`
	BeneficiaryRelationship            string `json:"beneficiary_relationship"`
	BeneficiaryBeneficiaryRelationship string `json:"beneficiary_beneficiary_relationship"`
	BeneficiarySourceOfFunds           string `json:"beneficiary_source_of_funds"`
	BeneficiaryRemittancePurposes      string `json:"beneficiary_remittance_purposes"`
	BeneficiaryIban                    string `json:"beneficiary_iban"`
	BeneficiarySwiftBicCode            string `json:"beneficiary_swift_bic_code"`
	BeneficiarySortCode                string `json:"beneficiary_sort_code"`
	BeneficiaryIfsCode                 string `json:"beneficiary_ifs_code"`
	BeneficiaryBsbNumber               string `json:"beneficiary_bsb_number"`
	BeneficiaryBranchNumber            string `json:"beneficiary_branch_number"`
	BeneficiaryDocumentReferenceNumber string `json:"beneficiary_document_reference_number"`
	BeneficiaryRegistrationNumber      string `json:"beneficiary_registration_number"`
	BeneficiaryRegion                  string `json:"beneficiary_region"`
	SenderName                         string `json:"sender_name"`
	SenderCountry                      string `json:"sender_country"`
	SenderPlaceOfBirth                 string `json:"sender_place_of_birth"`
	SenderDateOfBirth                  string `json:"sender_date_of_birth"`
	SenderAddress                      string `json:"sender_address"`
	SenderIdentityType                 string `json:"sender_identity_type"`
	SenderIdentityNumber               string `json:"sender_identity_number"`
	SenderJob                          string `json:"sender_job"`
	SenderEmail                        string `json:"sender_email"`
	SenderCity                         string `json:"sender_city"`
	SenderPhoneNumber                  string `json:"sender_phone_number"`
}

/*
CreateInternationalTransferC2CC2BResponse
Detail params can see on InternationalTransfer
*/
type CreateInternationalTransferC2CC2BResponse InternationalTransfer

//----------------------------- Create International Transfer B2C/B2B

/*
CreateInternationalTransferB2XRequest
| Attribute                             | Description                                                                                                                                                                                                                                          |
|---------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| destination_country                   | string (required) Destination country code using ISO 3166 (Alpha-3). See supported Destination Countries.                                                                                                                                            |
| source_country                        | string (required) Source country code using ISO 3166 (Alpha-3). Currently possible value is IDN.                                                                                                                                                     |
| transaction_type                      | string (required) Transaction type information. Possible values are: B2CB2B                                                                                                                                                                          |
| amount                                | string (required) Amount of money that wants to be disbursed to the beneficiary (in beneficiary currency). Maximum 2 decimal places, for Japan there can be no decimal. Must be within the minimum and maximum amount based on transaction currency. |
| attachment_data                       | file (required, if the destination country need it) File of invoice for Malaysia B2B Accepted extensions: txt, pdf, doc, docx, jpg, jpeg, png, bmp, rtf, xls, xlsx                                                                                   |
| attachment_type                       | string (required, if the destination country need it)Currently possible value is invoice for MYS                                                                                                                                                     |
| beneficiary_account_number            | string (required) Account number of the beneficiary bank                                                                                                                                                                                             |
| beneficiary_ach_code                  | string (required, if the destination country need it)Ach code of the beneficiary                                                                                                                                                                     |
| beneficiary_address                   | string (required, if the destination country need it) Address of the beneficiary                                                                                                                                                                     |
| beneficiary_bank_id                   | string (required, if the destination country need it) ID of beneficiary bank                                                                                                                                                                         |
| beneficiary_bank_name                 | string (required, if the destination country need it) Name of beneficiary bank. Can be used for destination countries AUS and GBR                                                                                                                    |
| beneficiary_branch_number             | string (required, if the destination country need it) Identifier of banks and branches across Japan                                                                                                                                                  |
| beneficiary_bsb_number                | string (required, if the destination country need it) Identifier of banks and branches across Australia                                                                                                                                              |
| beneficiary_city                      | string (required, if the destination country need it) City of the beneficiary                                                                                                                                                                        |
| beneficiary_document_reference_number | string (required, if the destination country need it) Reference number of document related to the transaction                                                                                                                                        |
| beneficiary_email                     | string (optional) Email of beneficiary                                                                                                                                                                                                               |
| beneficiary_full_name                 | string (required, if the destination country need it) Full name of the beneficiary                                                                                                                                                                   |
| beneficiary_iban                      | string (required, if the destination country need it) Identifier of individual account involved in the international transaction                                                                                                                     |
| beneficiary_id_expiration_date        | string (required, if the destination country need it)ID expiration date using YYYY-MM-DD format                                                                                                                                                      |
| beneficiary_ifs_code                  | string (required, if the destination country need it)Indian financial system code                                                                                                                                                                    |
| beneficiary_id_number                 | string (required, if the destination country need it)ID Number                                                                                                                                                                                       |
| beneficiary_msisdn                    | string (required, if the destination country need it) A number used to identify a phone number internationally                                                                                                                                       |
| beneficiary_nationality               | string (required, if the destination country need it) Country code using ISO 3166 (Alpha-3) of beneficiary nationality                                                                                                                               |
| beneficiary_postal_code               | string (required, if the destination country need it) Postal code of the beneficiary                                                                                                                                                                 |
| beneficiary_province                  | string (required, if the destination country need it) Province or state of the beneficiary                                                                                                                                                           |
| beneficiary_relationship              | string (required, if the destination country need it) Relationship of beneficiary                                                                                                                                                                    |
| beneficiary_remittance_purposes       | string (required, if the destination country need it) Purpose of remittance                                                                                                                                                                          |
| beneficiary_sort_code                 | string (required, if the destination country need it) Digits of code which is used by British and Irish banks                                                                                                                                        |
| beneficiary_source_of_funds           | string (required, if the destination country need it) Source of funds                                                                                                                                                                                |

Example CreateInternationalTransferB2XRequest:
{
  "destination_country":"MYS",
  "source_country":"IDN",
  "transaction_type":"B2B",
  "amount":"300.6",
  "attachment_data":(binary),
  "attachment_type":"invoice",
  "beneficiary_account_number":"1187710200",
  "beneficiary_bank_id":"2122",
  "beneficiary_document_reference_number":"doc_reference_123456",
  "beneficiary_full_name":"Jane Doe",
  "beneficiary_remittance_purposes":"TRAVEL"
}
*/
type CreateInternationalTransferB2XRequest struct {
	DestinationCountry                 string `json:"destination_country"`
	SourceCountry                      string `json:"source_country"`
	TransactionType                    string `json:"transaction_type"`
	Amount                             string `json:"amount"`
	AttachmentType                     string `json:"attachment_type"`
	BeneficiaryAccountNumber           string `json:"beneficiary_account_number"`
	BeneficiaryAchCode                 string `json:"beneficiary_ach_code"`
	BeneficiaryAddress                 string `json:"beneficiary_address"`
	BeneficiaryBankId                  string `json:"beneficiary_bank_id"`
	BeneficiaryBankName                string `json:"beneficiary_bank_name"`
	BeneficiaryBranchNumber            string `json:"beneficiary_branch_number"`
	BeneficiaryBsbNumber               string `json:"beneficiary_bsb_number"`
	BeneficiaryCity                    string `json:"beneficiary_city"`
	BeneficiaryDocumentReferenceNumber string `json:"beneficiary_document_reference_number"`
	BeneficiaryEmail                   string `json:"beneficiary_email"`
	BeneficiaryFullName                string `json:"beneficiary_full_name"`
	BeneficiaryIban                    string `json:"beneficiary_iban"`
	BeneficiaryIdExpirationDate        string `json:"beneficiary_id_expiration_date"`
	BeneficiaryIfsCode                 string `json:"beneficiary_ifs_code"`
	BeneficiaryIdNumber                string `json:"beneficiary_id_number"`
	BeneficiaryMsisdn                  string `json:"beneficiary_msisdn"`
	BeneficiaryNationality             string `json:"beneficiary_nationality"`
	BeneficiaryPostalCode              string `json:"beneficiary_postal_code"`
	BeneficiaryProvince                string `json:"beneficiary_province"`
	BeneficiaryRelationship            string `json:"beneficiary_relationship"`
	BeneficiaryRemittancePurposes      string `json:"beneficiary_remittance_purposes"`
	BeneficiarySortCode                string `json:"beneficiary_sort_code"`
	BeneficiarySourceOfFunds           string `json:"beneficiary_source_of_funds"`
}

type CreateInternationalTransferB2XFileRequest struct {
	AttachmentData requests.UploadFile `json:"attachment_data"`
}

/*
CreateInternationalTransferB2CB2BResponse
Documentation on InternationalTransfer
*/
type CreateInternationalTransferB2CB2BResponse InternationalTransfer

//----------------------------- Get International Transfer

/*
GetInternationalTransferResponse
Documentation on InternationalTransfer
*/
type GetInternationalTransferResponse InternationalTransfer

//----------------------------- Get All International Transfer

/*
GetAllInternationalTransferRequest
| Param      | Value                                                                                                                                                                                                                                                                 |
|------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| pagination | integer (optional)  The pagination of the result. Default value is 20.                                                                                                                                                                                                |
| page       | integer (optional)  The page number of the result to be viewed. Default value is 1.                                                                                                                                                                                   |
| sort_by    | string (optional) Sort the result by the attribute. Use the attribute name (e.g sort=id) to sort in ascending order or dash+attribute name (e.g sort=-id) to sort in descending order. Possible values are: id (default)amountsource_countrydestination_countrystatus |

More Documentation on Pagination
*/
type GetAllInternationalTransferRequest Pagination

/*
GetAllInternationalTransferResponse
| Attribute     | Description                            |
|---------------|----------------------------------------|
| total_data    | Total data returned in all page        |
| data_per_page | Total data returned in current page    |
| total_page    | Total/max page available               |
| page          | Current page                           |
| data          | Array of InternationalTransfer         |
*/
type GetAllInternationalTransferResponse struct {
	TotalData   int                     `json:"total_data"`
	DataPerPage int                     `json:"data_per_page"`
	TotalPage   int                     `json:"total_page"`
	Page        int                     `json:"page"`
	Data        []InternationalTransfer `json:"data"`
}
