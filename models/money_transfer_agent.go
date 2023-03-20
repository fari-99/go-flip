package models

/*
CreateDisbursementAgentRequest
Example CreateDisbursementAgentRequest:
{
  "agent_id":"1",
  "account_number":"1122333301",
  "amount":"10000",
  "bank_code":"bni",
  "direction":"DOMESTIC_SPECIAL_TRANSFER",
  "remark":"some remark",
  "beneficiary_email":"example@mail.com"
}
*/
type CreateDisbursementAgentRequest struct {
	AgentId          string `json:"agent_id"`
	AccountNumber    string `json:"account_number"`
	Amount           string `json:"amount"`
	BankCode         string `json:"bank_code"`
	Direction        string `json:"direction"`
	Remark           string `json:"remark"`
	BeneficiaryEmail string `json:"beneficiary_email"`
}

type CreateDisbursementAgentResponse DisbursementModel

/*
GetDisbursementAgentListRequest
| Param      | Value                                                                                                                                                                                                                                  |
|------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| agent_id   | integer (optional)  The agent ID that wants to be retrieved. If not provided, will return all agent disbursement for the company instead. Validation:Numeric                                                                           |
| pagination | integer (optional)  The pagination of the result. Default value is 20. Validation:Numeric                                                                                                                                              |
| page       | integer (optional)  The page number of the result to be viewed. Validation:Numeric                                                                                                                                                     |
| sort       | string (optional)  Sort the result by the attribute. Use the attribute name (e.g sort=id) to sort in ascending order or dash+attribute name (e.g sort=-id) to sort in descending order. Validation:Possible values are: idamountstatus |

Example GetDisbursementAgentListRequest:
agent_id:1
pagination:10
page:1
sort:id
*/
type GetDisbursementAgentListRequest struct {
	AgentId    string `json:"agent_id,omitempty"`
	Pagination string `json:"pagination,omitempty"`
	Page       string `json:"page,omitempty"`
	Sort       string `json:"sort,omitempty"`
}

type GetDisbursementAgentListResponse struct {
	TotalData   int                 `json:"total_data"`
	DataPerPage int                 `json:"data_per_page"`
	TotalPage   int                 `json:"total_page"`
	Page        int                 `json:"page"`
	Data        []DisbursementModel `json:"data"`
}
