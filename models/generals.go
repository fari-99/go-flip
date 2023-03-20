package models

type GetBalanceResponse struct {
	Balance float64 `json:"balance"`
}

/*
GetBankInfoRequest
Example GetBankInfoRequest:
{
  "code":"mandiri"
}
*/
type GetBankInfoRequest struct {
	BankCode string `json:"code,omitempty"`
}

/*
GetBankInfoResponse
| Attribute | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
|-----------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| bank_code | Flip’s bank code. cimb is the code for both CIMB Niaga and CIMB Niaga Syariah                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| name      | The name of the bank as we usually say it in Indonesian                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| fee       | The fee that you’ll be charged if you send money to this bank                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| queue     | Current queue for related bank. The longer/higher the queue number, the longer the transaction will be finished.                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| status    | The status of the disbursement process in related bank. Possible values are: OPERATIONALBanks are operational, disbursement will be processed as soon as possibleDISTURBEDBanks are slow or have another problem. Disbursement will still be processed, but in slower pace and might be delayedHEAVILY_DISTURBEDBanks are having an error, offline, or another problem that result in a nearly unusable system. Disbursement to this bank can not be processed in a short time, and maybe won’t be processed in the same day. You can ask for a refund if this happen. |

Example GetBankInfoResponse:
[
    {
        "bank_code": "mandiri",
        "name": "Mandiri",
        "fee": 5000,
        "queue": 8,
        "status": "DISTURBED"
    }
]
*/
type GetBankInfoResponse []struct {
	BankCode string `json:"bank_code"`
	Name     string `json:"name"`
	Fee      int    `json:"fee"`
	Queue    int    `json:"queue"`
	Status   string `json:"status"`
}

/*
GetMaintenanceStatusResponse
| Attribute   | Description             |
|-------------|-------------------------|
| maintenance | Flip maintenance status |

Example GetMaintenanceStatusResponse:
{
    "maintenance": false
}
*/
type GetMaintenanceStatusResponse struct {
	Maintenance bool `json:"maintenance"`
}

/*
SendBankAccountInquiryRequest
| Attribute      | Description                                                                         |
|----------------|-------------------------------------------------------------------------------------|
| account_number | string (required)  The account number of the bank account                           |
| bank_code      | string (required)  Bank code of the account. Accepted value are listed above        |
| inquiry_key    | string (optional)  Inquiry key for handling multiple request with multiple callback |

Example SendBankAccountInquiryRequest:
{
  "account_number": "5465327020",
  "bank_code": "bca",
  "inquiry_key": "aVncCDdKW9dciRvH9qSH"
}
*/
type SendBankAccountInquiryRequest struct {
	AccountNumber string `json:"account_number,omitempty"`
	BankCode      string `json:"bank_code,omitempty"`
	InquiryKey    string `json:"inquiry_key,omitempty"`
}

/*
SendBankAccountInquiryResponse
| Attribute      | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| bank_code      | Bank code of the account                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| inquiry_key    | Inquiry key for handling multiple request with multiple callback                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| account_number | Account number of the bank account                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| account_holder | Name of the bank account holder                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| status         | Possible values are  PENDINGInquiry still in processSUCCESSInquiry process is complete and bank account number is validINVALID_ACCOUNT_NUMBERInquiry process is complete but the account number is invalid or maybe a virtual account numberSUSPECTED_ACCOUNTBank account have been suspected on doing fraud. You still can do a disbursement to this account.BLACK_LISTEDBank account have been confirmed on doing a fraud and therefore is blacklisted. You can’t do a disbursment to this account.FAILEDThe inquiry process is failed before we get the final status of the inquiry, e.g due to timeout or any other errors from the bank. If you get this response, please retry the inquiry to trigger reverification of the account.CLOSEDThe inquiry process is complete and the account is valid, but it is closed/inactive so that it cannot receive money. You cannot do a disbursement to this account. |

Example SendBankAccountInquiryResponse:
{
    "bank_code": "bca",
    "account_number": "5465327020",
    "account_holder": "PT Fliptech Lentera IP",
    "status": "SUCCESS",
    "inquiry_key": "aVncCDdKW9dciRvH9qSH",
}
*/
type SendBankAccountInquiryResponse []struct {
	BankCode      string `json:"bank_code"`
	AccountNumber string `json:"account_number"`
	AccountHolder string `json:"account_holder"`
	Status        string `json:"status"`
	InquiryKey    string `json:"inquiry_key"`
}
