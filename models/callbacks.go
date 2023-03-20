package models

/*
BankInquiryCallback
If our system has completed the bank account inquiry process, then we will hit the provided URL for the bank inquiry callback.

Example BankInquiryCallback:
{
  "bank_code": "bca",
  "account_number": "5465327020",
  "account_holder": "PT Fliptech Lentera IP",
  "status": "SUCCESS",
  "inquiry_key": "aVncCDdKW9dciRvH9qSH"
}
*/
type BankInquiryCallback struct {
	BankCode      string `json:"bank_code"`
	AccountNumber string `json:"account_number"`
	AccountHolder string `json:"account_holder"`
	Status        string `json:"status"`
	InquiryKey    string `json:"inquiry_key"`
}

/*
DisbursementCallback
Disbursement callback will be sent to you if you transaction status is changed to DONE or CANCELLED.
This callback is used for the V2 and V3 Disbursement.

Example DisbursementCallback:
{
  "id": 790,
  "user_id": 23,
  "amount": 10000,
  "status": "DONE",
  "reason": "",
  "timestamp": "2017-08-28 14:32:47",
  "bank_code": "bni",
  "account_number": "0437051936",
  "recipient_name": "- FLIPTECH LENTERA INSPIRASI P",
  "sender_bank": "bri",
  "remark": "testing",
  "receipt": "https://storage.biznetgiocloud.com/v1/AUTH_GIOOST443831/bukti_transfer/123993_2017-08-04%202017:07:26.jpg",
  "time_served": "2017-08-28 14:42:47",
  "bundle_id": 0,
  "company_id": 7,
  "recipient_city": 391,
  "created_from": "API",
  "direction": "DOMESTIC_TRANSFER",
  "sender": null,
  "fee": 1500
}
*/
type DisbursementCallback struct {
	Id            int         `json:"id"`
	UserId        int         `json:"user_id"`
	Amount        int         `json:"amount"`
	Status        string      `json:"status"`
	Reason        string      `json:"reason"`
	Timestamp     string      `json:"timestamp"`
	BankCode      string      `json:"bank_code"`
	AccountNumber string      `json:"account_number"`
	RecipientName string      `json:"recipient_name"`
	SenderBank    string      `json:"sender_bank"`
	Remark        string      `json:"remark"`
	Receipt       string      `json:"receipt"`
	TimeServed    string      `json:"time_served"`
	BundleId      int         `json:"bundle_id"`
	CompanyId     int         `json:"company_id"`
	RecipientCity int         `json:"recipient_city"`
	CreatedFrom   string      `json:"created_from"`
	Direction     string      `json:"direction"`
	Sender        interface{} `json:"sender"`
	Fee           int         `json:"fee"`
}

/*
SpecialDisbursementCallback
Special Disbursement callback will be sent to you if you transaction status is changed to DONE or CANCELLED.
This callback is used for the V2 and V3 Special Disbursement.

Example SpecialDisbursementCallback:
{
  "id": 812,
  "user_id": 23,
  "amount": "10000",
  "status": "DONE",
  "reason": "",
  "timestamp": "2017-09-06 14:29:55",
  "bank_code": "bni",
  "account_number": "0437051936",
  "recipient_name": "- FLIPTECH LENTERA INSPIRASI P",
  "sender_bank": "bri",
  "remark": "testing",
  "receipt": "https://storage.biznetgiocloud.com/v1/AUTH_GIOOST443831/bukti_transfer/123993_2017-08-04%202017:07:26.jpg",
  "time_served": "2017-09-06 14:39:55",
  "bundle_id": 0,
  "company_id": 7,
  "recipient_city": 391,
  "created_from": "API",
  "direction": "DOMESTIC_SPECIAL_TRANSFER",
  "sender": {
    "sender_name": "John Doe",
    "place_of_birth": 391,
    "date_of_birth": "1992-01-31",
    "address": "taman bakokekok di jalan bakokekok 15 no.2 - 230",
    "sender_identity_type": "nat_id",
    "sender_identity_number": "asdas213123",
    "sender_country": 100252,
    "job": "private_employee"
  },
  "fee": 1500
}
*/
type SpecialDisbursementCallback struct {
	Id            int    `json:"id"`
	UserId        int    `json:"user_id"`
	Amount        string `json:"amount"`
	Status        string `json:"status"`
	Reason        string `json:"reason"`
	Timestamp     string `json:"timestamp"`
	BankCode      string `json:"bank_code"`
	AccountNumber string `json:"account_number"`
	RecipientName string `json:"recipient_name"`
	SenderBank    string `json:"sender_bank"`
	Remark        string `json:"remark"`
	Receipt       string `json:"receipt"`
	TimeServed    string `json:"time_served"`
	BundleId      int    `json:"bundle_id"`
	CompanyId     int    `json:"company_id"`
	RecipientCity int    `json:"recipient_city"`
	CreatedFrom   string `json:"created_from"`
	Direction     string `json:"direction"`
	Sender        struct {
		SenderName           string `json:"sender_name"`
		PlaceOfBirth         int    `json:"place_of_birth"`
		DateOfBirth          string `json:"date_of_birth"`
		Address              string `json:"address"`
		SenderIdentityType   string `json:"sender_identity_type"`
		SenderIdentityNumber string `json:"sender_identity_number"`
		SenderCountry        int    `json:"sender_country"`
		Job                  string `json:"job"`
	} `json:"sender"`
	Fee int `json:"fee"`
}

/*
AcceptPaymentCallback
Accept payment callback will be sent to you if there is a payment status that changed to SUCCESSFUL, CANCELED, or FAILED.

Example AcceptPaymentCallback:
{
  "id": "FT1",
  "bill_link": "flip.id/$<company_code>/#<product_code>",
  "bill_link_id": 4740,
  "bill_title": "Cimol Goreng",
  "sender_name": "Jon Doe",
  "sender_bank": "bni",
  "amount": 10000,
  "status": "SUCCESSFUL",
  "sender_bank_type": "bank_account",
  "created_at": "2021-11-29 10:10:10"
}
*/
type AcceptPaymentCallback struct {
	Id             string `json:"id"`
	BillLink       string `json:"bill_link"`
	BillLinkId     int    `json:"bill_link_id"`
	BillTitle      string `json:"bill_title"`
	SenderName     string `json:"sender_name"`
	SenderBank     string `json:"sender_bank"`
	Amount         int    `json:"amount"`
	Status         string `json:"status"`
	SenderBankType string `json:"sender_bank_type"`
	CreatedAt      string `json:"created_at"`
}

/*
InternationalTransferCallback
International transfer callback will be sent to you if you transaction status is changed.

Example InternationalTransferCallback:
{
  "id": 10,
  "user_id": 12345,
  "company_id": 123456,
  "exchange_rate": 19000.55,
  "amount": 1000000,
  "beneficiary_amount": 52.63,
  "fee": 88888,
  "source_country": "IDN",
  "destination_country": "GBR",
  "beneficiary_currency_code": "GBP",
  "status": "DONE",
  "timestamp": "2021-12-02 15:27:24",
  "time_served": "(not set)",
  "created_from": "API",
  "receipt": "someurl.png",
  "transaction_type": "C2C",
  "idempotency_key": "idem-1",
  "beneficiary": {
    "full_name": "John Smith",
    "bank_account_number": "1122333300",
    "bank": "Allica Bank",
    "email": "emial@email.com",
    "msisdn": "09584008222",
    "nationality": "IDN",
    "country": "GBR",
    "province": "United Kingdom",
    "city": "Manchester",
    "address": "Mosley St Manchester",
    "postal_code": "M2",
    "relationship": "SON",
    "source_of_funds": "BUSINESS",
    "remittance_purpose": "EDUCATION",
    "iban": null,
    "swift_bic_code": null,
    "sort_code": "506967",
    "ifs_code": null,
    "bsb_number": null,
    "branch_number": null,
    "document_reference_number": null,
    "registration_number": null
  },
  "sender": {
    "name": "John Doe",
    "place_of_birth": 100230,
    "date_of_birth": "1963-12-01",
    "address": "Some Address Street 123",
    "identity_type": "nat_id",
    "identity_number": "1234789012347890",
    "country": 100252,
    "job": "entrepreneur",
    "email": "sender@mail.com",
    "city": "Sender City",
    "phone_number": "+628123456789"
  }
}
*/
type InternationalTransferCallback struct {
	Id                      int     `json:"id"`
	UserId                  int     `json:"user_id"`
	CompanyId               int     `json:"company_id"`
	ExchangeRate            float64 `json:"exchange_rate"`
	Amount                  int     `json:"amount"`
	BeneficiaryAmount       float64 `json:"beneficiary_amount"`
	Fee                     int     `json:"fee"`
	SourceCountry           string  `json:"source_country"`
	DestinationCountry      string  `json:"destination_country"`
	BeneficiaryCurrencyCode string  `json:"beneficiary_currency_code"`
	Status                  string  `json:"status"`
	Timestamp               string  `json:"timestamp"`
	TimeServed              string  `json:"time_served"`
	CreatedFrom             string  `json:"created_from"`
	Receipt                 string  `json:"receipt"`
	TransactionType         string  `json:"transaction_type"`
	IdempotencyKey          string  `json:"idempotency_key"`
	Beneficiary             struct {
		FullName                string      `json:"full_name"`
		BankAccountNumber       string      `json:"bank_account_number"`
		Bank                    string      `json:"bank"`
		Email                   string      `json:"email"`
		Msisdn                  string      `json:"msisdn"`
		Nationality             string      `json:"nationality"`
		Country                 string      `json:"country"`
		Province                string      `json:"province"`
		City                    string      `json:"city"`
		Address                 string      `json:"address"`
		PostalCode              string      `json:"postal_code"`
		Relationship            string      `json:"relationship"`
		SourceOfFunds           string      `json:"source_of_funds"`
		RemittancePurpose       string      `json:"remittance_purpose"`
		Iban                    interface{} `json:"iban"`
		SwiftBicCode            interface{} `json:"swift_bic_code"`
		SortCode                string      `json:"sort_code"`
		IfsCode                 interface{} `json:"ifs_code"`
		BsbNumber               interface{} `json:"bsb_number"`
		BranchNumber            interface{} `json:"branch_number"`
		DocumentReferenceNumber interface{} `json:"document_reference_number"`
		RegistrationNumber      interface{} `json:"registration_number"`
	} `json:"beneficiary"`
	Sender struct {
		Name           string `json:"name"`
		PlaceOfBirth   int    `json:"place_of_birth"`
		DateOfBirth    string `json:"date_of_birth"`
		Address        string `json:"address"`
		IdentityType   string `json:"identity_type"`
		IdentityNumber string `json:"identity_number"`
		Country        int    `json:"country"`
		Job            string `json:"job"`
		Email          string `json:"email"`
		City           string `json:"city"`
		PhoneNumber    string `json:"phone_number"`
	} `json:"sender"`
}

/*
AgentKYCStatusCallback
A callback notification will be sent to your Agent KYC callback URL if there is a status update related to the agent’s KYC process.

Example :
{
  "agent_id": 1,
  "agent_name": "John Doe",
  "kyc_status": "UPLOAD_IDENTITY_SELFIE_FAILED",
  "rejected_reason_code": 10,
  "rejected_reason": "The data that is filled in does not match with identity card",
  "created_at": "2022-02-18 05:03:32",
  "updated_at": "2022-03-17 09:48:58",
  "submitted_at": "2022-03-17 09:48:58",
  "verified_at": "2022-01-03 09:11:49"
}
*/
type AgentKYCStatusCallback struct {
	AgentId            int    `json:"agent_id"`
	AgentName          string `json:"agent_name"`
	KycStatus          string `json:"kyc_status"`
	RejectedReasonCode int    `json:"rejected_reason_code"`
	RejectedReason     string `json:"rejected_reason"`
	CreatedAt          string `json:"created_at"`
	UpdatedAt          string `json:"updated_at"`
	SubmittedAt        string `json:"submitted_at"`
	VerifiedAt         string `json:"verified_at"`
}
