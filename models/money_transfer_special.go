package models

/*
CreateSpecialDisbursementRequest
| Attribute              | Description                                                                                                                                                                                                                                                                                                                                                                            |
|------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| account_number         | string (required)  The account number of the recipient. Validation:Numeric                                                                                                                                                                                                                                                                                                             |
| bank_code              | string (required)  Bank code of the recipient bank. Validation:Accepted value are listed here                                                                                                                                                                                                                                                                                          |
| amount                 | integer (required)  The amount of money to be disbursed. Validation:Numeric                                                                                                                                                                                                                                                                                                            |
| remark                 | string (optional)  Remark to be included in the transfer made to the recipient. Usually will appear as berita transfer or remark in the transfer receipt. Only for disbursement with the bank code being bri, the remark will be prepended with the beneficiary name. Example: tes remark will be john doe test remark. Validation:1-18 characterAlphanumericSpaces, ., -, /, (, and ) |
| recipient_city         | integer (optional)  City code of the recipient city. Default value is 394. Validation:Accepted values are listed here.                                                                                                                                                                                                                                                                 |
| sender_country         | integer (required)  Country code of the sender’s residence. Validation:Accepted values are listed here.                                                                                                                                                                                                                                                                                |
| sender_place_of_birth  | integer (optional)  City/country code of the sender’s place of birth. Use city code if the sender’s place of birth is in Indonesia, and country code if outside Indonesia. Validation:Accepted values are listed here.                                                                                                                                                                 |
| sender_date_of_birth   | string/date (optional)  Sender’s birth date. Validation:YYYY-MM-DD format                                                                                                                                                                                                                                                                                                              |
| sender_identity_type   | string (optional)  Sender’s ID type. Validation:Accepted value are as follows: nat_idNational Id Card or KTP in Indonesiadrv_licDriving licensepassportbank_accBank Account (DEFAULT)                                                                                                                                                                                                  |
| sender_name            | string (required)  The name of the user of the Money Transfer Company that act as a sender. Validation:Alphanumeric1-50 characters                                                                                                                                                                                                                                                     |
| sender_address         | string (required)  Sender’s address. Validation:Alphanumeric255 characters                                                                                                                                                                                                                                                                                                             |
| sender_identity_number | string (optional)  Sender’s ID number. Default value is from attribute account_number. Validation:16 characters for KTP20 characters for Passport                                                                                                                                                                                                                                      |
| sender_job             | string (required) Sender’s job. Validation:Accepted value are as follows:housewifeentrepreneurprivate_employeegovernment_employeefoundation_boardPeople who work at foundation as an employeeindonesian_migrant_workerAlso known as TKIcompanyIf sender is a company, not individualothers                                                                                             |
| direction              | string (required) The direction of the transaction. Validation:Accepted values are as follows: DOMESTIC_SPECIAL_TRANSFERWhen the sender and the recipient are both residing in IndonesiaFOREIGN_INBOUND_SPECIAL_TRANSFERWhen the sender are in a foreign country, but the recipient are in Indonesia                                                                                   |
| beneficiary_email      | string (optional)  List of the recipient emails. Validation:Alphanumeric with email format (xxx@xxx.xxx)Separate each email with a comma (if any)                                                                                                                                                                                                                                      |

Example CreateSpecialDisbursementRequest:
{
  "account_number":"1122333301",
  "bank_code":"bni",
  "amount":"10000",
  "remark":"some remark",
  "recipient_city":"391",
  "sender_country":"100252",
  "sender_place_of_birth":"391",
  "sender_date_of_birth":"1992-01-01",
  "sender_identity_type":"nat_id",
  "sender_name":"John Doe",
  "sender_address":"Some Address Street 123",
  "sender_identity_number":"123456789",
  "sender_job":"entrepreneur",
  "direction":"DOMESTIC_SPECIAL_TRANSFER",
  "beneficiary_email":"test@mail.com,user@mail.com"
}
*/
type CreateSpecialDisbursementRequest struct {
	AccountNumber        string `json:"account_number"`
	BankCode             string `json:"bank_code"`
	Amount               string `json:"amount"`
	Remark               string `json:"remark"`
	RecipientCity        string `json:"recipient_city"`
	SenderCountry        string `json:"sender_country"`
	SenderPlaceOfBirth   string `json:"sender_place_of_birth"`
	SenderDateOfBirth    string `json:"sender_date_of_birth"`
	SenderIdentityType   string `json:"sender_identity_type"`
	SenderName           string `json:"sender_name"`
	SenderAddress        string `json:"sender_address"`
	SenderIdentityNumber string `json:"sender_identity_number"`
	SenderJob            string `json:"sender_job"`
	Direction            string `json:"direction"`
	BeneficiaryEmail     string `json:"beneficiary_email"`
}

type CreateSpecialDisbursementResponse DisbursementModel

type CityListResponse map[string]string        // key: city code, value: city name
type CountryListResponse map[string]string     // key: country code, value: country name
type CountryCityListResponse map[string]string // key: country/city code, value: country/city name
