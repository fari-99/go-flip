package models

/*
CreateDisbursementRequest
| Attribute        | Description                                                                                                                                                                                                                                                                                                                                                                            |
|------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| AccountNumber    | string (required)  The account number of the recipient. Validation:Numeric                                                                                                                                                                                                                                                                                                             |
| BankCode         | string (required)  Bank code of the recipient bank.Validation:Accepted values are listed here.                                                                                                                                                                                                                                                                                         |
| Amount           | integer (required)  The amount of money to be disbursed Validation:Numeric                                                                                                                                                                                                                                                                                                             |
| Remark           | string (optional)  Remark to be included in the transfer made to the recipient. Usually will appear as berita transfer or remark in the transfer receipt. Only for disbursement with the bank code being bri, the remark will be prepended with the beneficiary name. Example: tes remark will be john doe test remark. Validation:1-18 characterAlphanumericSpaces, ., -, /, (, and ) |
| RecipientCity    | integer (optional)  City code of the recipient city. Validation:Accepted values are listed here.                                                                                                                                                                                                                                                                                       |
| BeneficiaryEmail | string (optional)  List of the recipient emails. Validation:Alphanumeric with email format (xxx@xxx.xxx)Separate each email with a comma (if any)                                                                                                                                                                                                                                      |

Example CreateDisbursementRequest:
{
  "account_number":"1122333300",
  "bank_code":"bni",
  "amount":"10000",
  "remark":"some remark",
  "recipient_city":"391,",
  "beneficiary_email":"test@mail.com,user@mail.com"
}
*/
type CreateDisbursementRequest struct {
	AccountNumber    string `json:"account_number"`
	BankCode         string `json:"bank_code"`
	Amount           string `json:"amount"`
	Remark           string `json:"remark"`
	RecipientCity    string `json:"recipient_city"`
	BeneficiaryEmail string `json:"beneficiary_email"`
}

type CreateDisbursementResponse DisbursementModel

/*
DisbursementModel
| Attribute         | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
|-------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| id                | Flip’s transaction id                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| user_id           | Your account user id in our system                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| amount            | The amount of money to be disbursed in IDR                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| status            | Transaction status. Possible values are:  PENDINGDisbursement is still in processCANCELLEDThe transaction is cancelled and the amount of the transaction plus the transaction fee will be credited to your balance. This will happen if the transfer process are failed for reason such as inactive recipient account, wrong account number, etcDONEDisbursement process is finished and the money have been sent to the recipient                                                                                                                                      |
| reason            | The reason value will be given if the transaction status is CANCELLED. Possible values are:  INACTIVE_ACCOUNT Inactive account / Nomor rekening tidak aktifNOT_REGISTERED_ACCOUNT Not registered account / Nomor rekening tidak terdaftarCANT_RECEIVE_TRANSFER Can’t receive transfer / Rekening tujuan tidak dapat menerima transferINTERMITTENT_DISTURBANCE_ON_BENEFICIARY_BANK Intermittent disturbance on destination bank / Bank tujuan sementara mengalami gangguanBENEFICIARY_ACCOUNT_NOT_VERIFIED Account is not verified / Rekening tujuan belum terverifikasi |
| timestamp         | The time when the disbursement request created. Time will be in GMT+7 with yyyy-mm-dd hh:mm:ss format                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| bank_code         | Bank code of the recipient bank                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| account_number    | The account number of the recipient                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| recipient_name    | The name of the recipient account holder. If the account number haven’t cached by Flip yet, this attribute will show - (dash) instead                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| sender_bank       | The default sender bank in your account                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| remark            | Remark to be included in the transfer made to the recipient                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| receipt           | Url of the transfer receipt. The receipt will be a screenshot taken from the internet banking interface of each bank. This attribute will only show the url when the status is DONE                                                                                                                                                                                                                                                                                                                                                                                     |
| time_served       | The time when the transaction is finished. Will only show valid value when the status is DONE                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| bundle_id         | The bundle id of the transaction made from Flip for Business Dashboard (csv upload or manual input). For the transaction created from API, the value will always be 0                                                                                                                                                                                                                                                                                                                                                                                                   |
| company_id        | Your Flip for Business account user id in our system                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| recipient_city    | City code of the recipient city                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| created_from      | The channel of which the transaction was created. Possible values are: APIDASHBOARD                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| direction         | The direction of the transaction. Possible values are: DOMESTIC_TRANSFERCommon Disbursement from Indonesia to Indonesian recipientDOMESTIC_SPECIAL_TRANSFERSpecial disbursement from the user of a Money Transfer Company residing in Indonesia to Indonesian recipientFOREIGN_INBOUND_SPECIAL_TRANSFERSpecial disbursement from the user of a Money Transfer Company residing in a foreign country to Indonesian recipient                                                                                                                                             |
| sender            | Possible values are null if the transaction is a common disbursement, and sender object if the transaction is a special disbursment.                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| fee               | The fee of the transaction                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| beneficiary_email | List of the recipient emails                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| idempotency_key   | The idempotency key used to create the disbursement                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |

Example DisbursementModel:
{
  "id": 10,
  "user_id": 20,
  "amount": 10000,
  "status": "PENDING",
  "reason": "",
  "timestamp": "2017-08-28 14:32:47",
  "bank_code": "bni",
  "account_number": "1122333300",
  "recipient_name": "John Doe",
  "sender_bank": "bri",
  "remark": "some remark",
  "receipt": "",
  "time_served": "0000-00-00 00:00:00",
  "bundle_id": 0,
  "company_id": 7,
  "recipient_city": 391,
  "created_from": "API",
  "direction": "DOMESTIC_TRANSFER",
  "sender": {
    "sender_name": "John Doe",
    "place_of_birth": 391,
    "date_of_birth": "1992-01-01",
    "address": "Some Address Street 123",
    "sender_identity_type": "nat_id",
    "sender_identity_number": "asdas213123",
    "sender_country": 100252,
    "job": "entrepreneur"
  },
  "fee": 1500,
  "beneficiary_email": "test@mail.com,user@mail.com",
  "idempotency_key": "idem-key-1"
}
*/
type DisbursementModel struct {
	Id               int64                    `json:"id"`
	UserId           int64                    `json:"user_id"`
	Amount           float64                  `json:"amount"`
	Status           string                   `json:"status"`
	Reason           string                   `json:"reason"`
	Timestamp        string                   `json:"timestamp"`
	BankCode         string                   `json:"bank_code"`
	AccountNumber    string                   `json:"account_number"`
	RecipientName    string                   `json:"recipient_name"`
	SenderBank       string                   `json:"sender_bank"`
	Remark           string                   `json:"remark"`
	Receipt          string                   `json:"receipt"`
	TimeServed       string                   `json:"time_served"`
	BundleId         int64                    `json:"bundle_id"`
	CompanyId        int64                    `json:"company_id"`
	RecipientCity    int                      `json:"recipient_city"`
	CreatedFrom      string                   `json:"created_from"`
	Direction        string                   `json:"direction"`
	Sender           *DisbursementSenderModel `json:"sender"`
	Fee              float64                  `json:"fee"`
	BeneficiaryEmail string                   `json:"beneficiary_email"`
	IdempotencyKey   string                   `json:"idempotency_key"`
}

/*
DisbursementSenderModel
| Attribute              | Description                                                                                                                                     |
|------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| sender_name            | The name of the user of the Money Transfer Company that act as a sender                                                                         |
| place_of_birth         | City/country code of the Sender’s place of birth                                                                                                |
| date_of_birth          | Sender’s date of birth                                                                                                                          |
| address                | Sender’s address                                                                                                                                |
| sender_identity_type   | Sender’s identity type. Possible value are: nat_id,drv_lic,passport,bank_acc                                                                    |
| sender_identity_number | Sender’s identity number                                                                                                                        |
| sender_country         | Country code of the Sender’s country                                                                                                            |
| job                    | Sender’s job. Possible values are:housewife,entrepreneur,private_employee,government_employee,foundation_board,indonesian_migrant_worker,others |

Example DisbursementSenderModel:
{
    "sender_name": "John Doe",
    "place_of_birth": 391,
    "date_of_birth": "1992-01-01",
    "address": "Some Address Street 123",
    "sender_identity_type": "nat_id",
    "sender_identity_number": "asdas213123",
    "sender_country": 100252,
    "job": "entrepreneur"
}
*/
type DisbursementSenderModel struct {
	SenderName           string `json:"sender_name"`
	PlaceOfBirth         int    `json:"place_of_birth"`
	DateOfBirth          string `json:"date_of_birth"`
	Address              string `json:"address"`
	SenderIdentityType   string `json:"sender_identity_type"`
	SenderIdentityNumber string `json:"sender_identity_number"`
	SenderCountry        int    `json:"sender_country"`
	Job                  string `json:"job"`
}

/*
GetAllDisbursementRequest
| param      | value                                                                                                                                                                                                                                   |
|------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| pagination | integer (optional)  The pagination of the result. Default value is 20. Validation:Numeric                                                                                                                                               |
| page       | integer (optional)  The page number of the result to be viewed. Validation:Numeric                                                                                                                                                      |
| sort       | string (optional)  Sort the result by the attribute. Use the attribute name (e.g sort=id) to sort in ascending order or dash+attribute name (e.g sort=-id) to sort in descending order. Validation:Accepted values are as listed below. |

You can also filter the result by changing attribute with the attribute to be filtered and value with the filter value.
You can filter more than one attribute by appending another attribute filter to the url.

Attribute that can be filtered are:

    id - exact comparison
    amount - exact comparison
    status - exact comparison
    timestamp - “like” comparison
    bank_code - “like” comparison
    account_number - “like” comparison
    recipient_name - “like” comparison
    remark - “like” comparison
    time_served - “like” comparison
    created_from - “like” comparison
    direction - exact comparison

Example GetAllDisbursementRequest:
pagination:10
page:1
sort:id
attribute:11223333
*/
type GetAllDisbursementRequest struct {
	Pagination string  `json:"pagination"`
	Page       string  `json:"page"`
	Sort       string  `json:"sort"`
	Attribute  *string `json:"attribute,omitempty"`
}

type GetAllDisbursementResponse struct {
	TotalData   int                 `json:"total_data"`
	DataPerPage int                 `json:"data_per_page"`
	TotalPage   int                 `json:"total_page"`
	Page        int                 `json:"page"`
	Reason      string              `json:"reason"`
	Data        []DisbursementModel `json:"data"`
}
