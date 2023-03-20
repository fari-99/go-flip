package models

/*
CustomerBillings
| Attribute | Description                   |
|-----------|-------------------------------|
| name      | Name of the Customer.         |
| email     | Email of the Customer.        |
| phone     | Phone number of the Customer. |
| address   | Address of the Customer.      |

*/
type CustomerBillings struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

/*
BillPayments
| Attribute             | Description                                                                                                                                                                                                                                                                                                                        |
|-----------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| id                    | ID of the transaction.                                                                                                                                                                                                                                                                                                             |
| amount                | Amount of the transaction. Please note that in case of Bank Transfer payment option, the actual amount to be transferred by your customer must be equal to the amount + unique_code returned in the response.                                                                                                                      |
| unique_code           | Unique code of the transaction                                                                                                                                                                                                                                                                                                     |
| status                | Status of the transaction NOT_CONFIRMED: If the transaction was just created.PENDING: If the transaction has been confirmed by the userPROCESSED: If the transaction is being processed by the system.CANCELED: If the transaction is canceled.FAILED: If the transaction is failed.DONE: If the transaction is successfully done. |
| sender_bank           | Bank that is used for the payment.  See Bank Code Constants.                                                                                                                                                                                                                                                                       |
| sender_bank_type      | Type of the sender bank that is used for the payment. bank_account: if the type of the sender bank uses bank transfer.virtual_account: if the type of the sender bank uses a virtual account.wallet_account: if the type of the sender bank uses e-wallet.                                                                         |
| receiver_bank_account | Account of the receiver of the payment.  See Bank Account Object.                                                                                                                                                                                                                                                                  |
| user_address          | Address of the User.                                                                                                                                                                                                                                                                                                               |
| user_phone            | Phone number of the User.                                                                                                                                                                                                                                                                                                          |
| created_at            | Unix timestamps of the transaction when it is created.                                                                                                                                                                                                                                                                             |

*/
type BillPayments struct {
	Id             string `json:"id"`
	Amount         int    `json:"amount"`
	UniqueCode     int    `json:"unique_code"`
	Status         string `json:"status"`
	SenderBank     string `json:"sender_bank"`
	SenderBankType string `json:"sender_bank_type"`
	UserAddress    string `json:"user_address"`
	UserPhone      string `json:"user_phone"`
	CreatedAt      int    `json:"created_at"`

	ReceiverBankAccount BillBankAccount `json:"receiver_bank_account"`
}

/*
BillBankAccount
| Attribute      | Description                                                                        |
|----------------|------------------------------------------------------------------------------------|
| account_number | Account number of the receiver bank.                                               |
| account_type   | Type of the account, the value could be: virtual_accountbank_accountwallet_account |
| bank_code      | Bank code of the account.  See Bank Code Constants.                                |
| account_holder | Name of the holder of the account.                                                 |

*/
type BillBankAccount struct {
	AccountNumber *string `json:"account_number"`
	AccountType   string  `json:"account_type"`
	BankCode      string  `json:"bank_code"`
	AccountHolder string  `json:"account_holder"`
}

/*
Billings
| Attribute                | Description                                                                                                                                                                                                                                                                                                            |
|--------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| link_id                  | Bill link ID.                                                                                                                                                                                                                                                                                                          |
| link_url                 | Bill link URL for the user.                                                                                                                                                                                                                                                                                            |
| title                    | Title of the bill.                                                                                                                                                                                                                                                                                                     |
| type                     | Bill type: SINGLE: single useMULTIPLE: multiple use                                                                                                                                                                                                                                                                    |
| amount                   | Payment amount.                                                                                                                                                                                                                                                                                                        |
| redirect_url             | Redirect URL after payment is success.                                                                                                                                                                                                                                                                                 |
| expired_date             | Bill expiry date. The bill can’t be used any more beyond the expiry date.                                                                                                                                                                                                                                              |
| created_from             | Bill created from method.                                                                                                                                                                                                                                                                                              |
| status                   | Bill activation status: ACTIVEINACTIVE                                                                                                                                                                                                                                                                                 |
| is_address_required      | A flag if user needs to input their address when creating payment. 0: false1: true                                                                                                                                                                                                                                     |
| is_phone_number_required | A flag if user needs to input their phone number when creating payment. 0: false1: true                                                                                                                                                                                                                                |
| step                     | Which step the customers will be redirected to when opening the payment link. 1: input-data (default)2: payment-method3: payment-confirmation                                                                                                                                                                          |
| payment_url              | (Returned only if customer’s merchant chooses step 3)  URL of the confirmation link or instruction to do the payment action.URL confirmation link is for payment using bank transfer, VA payment method, and QRIS with internal urlURL instruction link is for payment using e-wallet payment method with external url |
| customer                 | (Returned only if customer’s merchant chooses step 2 or 3)  Customer Data that is being saved.  See Customer Object.                                                                                                                                                                                                   |
| bill_payment             | (Returned only if customer’s merchant chooses step 3)  Bill Payment data that has been created.  See Bill Payment Object.                                                                                                                                                                                              |

*/
type Billings struct {
	LinkId                int     `json:"link_id"`
	LinkUrl               string  `json:"link_url"`
	Title                 string  `json:"title"`
	Type                  string  `json:"type"`
	Amount                int     `json:"amount"`
	RedirectUrl           string  `json:"redirect_url"`
	ExpiredDate           *string `json:"expired_date"`
	CreatedFrom           string  `json:"created_from"`
	Status                string  `json:"status"`
	IsAddressRequired     int     `json:"is_address_required"`
	IsPhoneNumberRequired int     `json:"is_phone_number_required"`
	Step                  int     `json:"step"`
	PaymentUrl            string  `json:"payment_url"`

	Customer    *CustomerBillings `json:"customer,omitempty"`
	BillPayment *BillPayments     `json:"bill_payment,omitempty"`
}

/*
Payments
| Attribute         | Description                                                                                 |
|-------------------|---------------------------------------------------------------------------------------------|
| id                | Payment ID.                                                                                 |
| bill_link         | Bill link URL of the payment.                                                               |
| bill_title        | Title of the bill.                                                                          |
| sender_name       | Name of the user who did the payment.                                                       |
| sender_bank       | Bank code of the user’s bank.                                                               |
| amount            | Payment amount done by the user.                                                            |
| status            | Payment status. FAILEDSUCCESSFULPENDING                                                     |
| settlement_status | Merchant settlement status. CancelledSettledPending                                         |
| created_at        | The time when the payment is created. Time will be in GMT+7 with yyyy-mm-dd hh:mm:ss format |

Example Payments:
{
  "id": "PGPWF3453271107384587",
  "bill_link": "flip.id/$companyname/#coffeetable",
  "bill_title": "Coffee Table",
  "sender_name": "John Smith",
  "sender_bank": "qris",
  "amount": 900000,
  "status": "SUCCESSFUL",
  "settlement_status": "Pending",
  "created_at": "2022-07-01 14:57:44"
}
*/
type Payments struct {
	Id               string `json:"id"`
	BillLink         string `json:"bill_link"`
	BillTitle        string `json:"bill_title"`
	SenderName       string `json:"sender_name"`
	SenderBank       string `json:"sender_bank"`
	Amount           int    `json:"amount"`
	Status           string `json:"status"`
	SettlementStatus string `json:"settlement_status"`
	CreatedAt        string `json:"created_at"`
}

// --------------------------- RESPONSE REQUEST
// --------------------------- Create Bill

/*
CreateBillRequest
| Attribute                | Description                                                                                                                                                                                                             |
|--------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| title                    | string (required)  The title of the bill.                                                                                                                                                                               |
| type                     | string (required)  Bill type: SINGLE: single useMULTIPLE: multiple use For bill created with step 3 only SINGLE is permissible.                                                                                         |
| amount                   | integer (optional)  Payment amount, minimum Rp10.000 and maximum Rp10.000.000. Leave blank if want to set a flexible amount.  Optional for bill created with Step 1 and mandatory for bill created with Step 2/ Step 3. |
| expired_date             | string (optional)  Bill expiry date. The bill can’t be used any more beyond the expiry date. Format: YYYY-MM-DD HH:mm.                                                                                                  |
| redirect_url             | string (optional)  Redirect URL after payment is success.                                                                                                                                                               |
| is_address_required      | integer (optional)  A flag if user needs to input their address when creating payment. 0: false (default)1: true                                                                                                        |
| is_phone_number_required | integer (optional)  A flag if user needs to input their phone number when creating payment. 0: false (default)1: true                                                                                                   |
| step                     | integer (required)  Which step the customers will be redirected to when opening the payment link. 1: input-data (default)2: payment-method3: payment-confirmation                                                       |
| sender_name              | string (required for step 2-payment method or step 3-payment-confirmation)  Name of the Customer.                                                                                                                       |
| sender_email             | string (required for step 2-payment method or step 3-payment-confirmation)  Email of the Customer.                                                                                                                      |
| sender_phone_number      | string (applicable for step 2-payment method or step 3-payment-confirmation)  Phone number of the Customer.                                                                                                             |
| sender_address           | string (applicable for step 2-payment method or step 3-payment-confirmation)  Address of the Customer.                                                                                                                  |
| sender_bank              | string (required for step 3-payment-confirmation)  Bank that is used for the payment.  Value will be bank_code of Flip. Please instruct your users to follow these steps for making BSI VA payments                     |
| sender_bank_type         | string (required for step 3-payment-confirmation) For a virtual account the value must be set as virtual_accountFor e-wallet the value must be set as wallet_account.See Bank Account Constants.                        |

Example CreateBillRequest:
{
    "title": "Coffee Table",
    "type": "SINGLE",
    "amount": "900000",
    "expired_date": "2024-12-30 15:50",
    "redirect_url": "https://someurl.com",
    "is_address_required": "0",
    "is_phone_number_required": "0",
    "step": "1"
}
*/
type CreateBillRequest struct {
	Attribute             string `json:"Attribute"`
	Title                 string `json:"title"`
	Type                  string `json:"type"`
	Amount                string `json:"amount"`
	ExpiredDate           string `json:"expired_date"`
	RedirectUrl           string `json:"redirect_url"`
	IsAddressRequired     string `json:"is_address_required"`
	IsPhoneNumberRequired string `json:"is_phone_number_required"`
	Step                  string `json:"step"`
	SenderName            string `json:"sender_name"`
	SenderEmail           string `json:"sender_email"`
	SenderPhoneNumber     string `json:"sender_phone_number"`
	SenderAddress         string `json:"sender_address"`
	SenderBank            string `json:"sender_bank"`
	SenderBankType        string `json:"sender_bank_type"`
}

type CreateBillResponse Billings

// --------------------------- Edit Billing

/*
EditBillingRequest
| Attribute                | Description                                                                                                                  |
|--------------------------|------------------------------------------------------------------------------------------------------------------------------|
| title                    | string (optional)  The title of the bill.                                                                                    |
| type                     | string (optional)  Bill type: SINGLE: single useMULTIPLE: multiple use for bill with step 3, type can only be “SINGLE”       |
| amount                   | integer (optional)  Payment amount, minimum Rp10.000 and maximum Rp10.000.000. Leave blank if want to set a flexible amount. |
| expired_date             | string (optional)  Bill expiry date. The bill can’t be used any more beyond the expiry date. Format: YYYY-MM-DD HH:mm.       |
| redirect_url             | string (optional)  Redirect URL after payment is success.                                                                    |
| status                   | string (optional)  Bill activation status: ACTIVEINACTIVE                                                                    |
| is_address_required      | integer (optional)  A flag if user needs to input their address when creating payment. 0: false (default)1: true             |
| is_phone_number_required | integer (optional)  A flag if user needs to input their phone number when creating payment. 0: false (default)1: true        |

Example EditBillingRequest:
{
  "title":"Coffee Table",
  "amount":"900000",
  "type":"SINGLE",
  "expired_date":"2022-12-30 15:50",
  "redirect_url":"https://someurl.com",
  "status":"INACTIVE",
  "is_address_required":"0",
  "is_phone_number_required":"0"
}
*/
type EditBillingRequest struct {
	Title                 string `json:"title"`
	Type                  string `json:"type"`
	Amount                string `json:"amount"`
	ExpiredDate           string `json:"expired_date"`
	RedirectUrl           string `json:"redirect_url"`
	Status                string `json:"status"`
	IsAddressRequired     string `json:"is_address_required"`
	IsPhoneNumberRequired string `json:"is_phone_number_required"`
}

type EditBillingResponse Billings

// --------------------------- Get Billing

type GetBillingResponse Billings

// --------------------------- Get All Billing

type GetAllBillingResponse []Billings

// --------------------------- Get Payments

/*
GetPaymentRequest
See documentation on Pagination
*/
type GetPaymentRequest Pagination

/*
GetPaymentResponse
| Attribute     | Description                         |
|---------------|-------------------------------------|
| total_data    | Total data returned in all page     |
| data_per_page | Total data returned in current page |
| total_page    | Total/max page available            |
| page          | Current page                        |
| data          | Array of Payments object.           |

Example GetPaymentResponse:
{
  "link_id": 1065,
  "total_data": 2,
  "data_per_page": 2,
  "total_page": 1,
  "page": 1,
  "data": [
    {
      "id": "PGPWF3453271107384587",
      "bill_link": "flip.id/$companyname/#coffeetable",
      "bill_title": "Coffee Table",
      "sender_name": "John Smith",
      "sender_bank": "qris",
      "amount": 900000,
      "status": "SUCCESSFUL",
      "settlement_status": "Pending",
      "created_at": "2022-07-01 14:57:44"
    }
  ]
}
*/
type GetPaymentResponse struct {
	LinkId      int        `json:"link_id"`
	TotalData   int        `json:"total_data"`
	DataPerPage int        `json:"data_per_page"`
	TotalPage   int        `json:"total_page"`
	Page        int        `json:"page"`
	Data        []Payments `json:"data"`
}

// --------------------------- Get All Payments

/*
GetAllPaymentRequest
See documentation on Pagination
*/
type GetAllPaymentRequest Pagination

/*
GetAllPaymentResponse
| Attribute     | Description                         |
|---------------|-------------------------------------|
| total_data    | Total data returned in all page     |
| data_per_page | Total data returned in current page |
| total_page    | Total/max page available            |
| page          | Current page                        |
| data          | Array of Payments object.           |

Example GetAllPaymentResponse:
{
  "total_data": 3,
  "data_per_page": 3,
  "total_page": 1,
  "page": 1,
  "data": [
    {
      "id": "FT12345678",
      "bill_link": "flip.id/$companyname/#coffeetable",
      "bill_title": "Coffee Table",
      "sender_name": "John Smith",
      "sender_bank": "mandiri",
      "amount": 900000,
      "status": "SUCCESSFUL",
      "settlement_status": "Pending",
      "created_at": "2021-02-01 14:57:44"
    }
  ]
}

*/
type GetAllPaymentResponse struct {
	TotalData   int        `json:"total_data"`
	DataPerPage int        `json:"data_per_page"`
	TotalPage   int        `json:"total_page"`
	Page        int        `json:"page"`
	Data        []Payments `json:"data"`
}

// --------------------------- Confirm Bill Payment

type ConfirmBillPaymentResponse struct {
	ConfirmedAt string `json:"confirmed_at"`
	PaymentUrl  string `json:"payment_url"`
}
