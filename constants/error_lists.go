package constants

import "fmt"

/*
ErrorResponse
Example ErrorResponse:
{
    "name": "Unauthorized",
    "message": "You are requesting with an invalid credential.",
    "status": 401,
}

{
    "name": "Not Found",
    "message": "Page not found.",
    "status": 404,
}
*/
type ErrorResponse struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

/*
ErrorClientResponse
Example:
{
  "code": [outer_code],
  "errors": [
    {
      "attribute": "[attribute]",
      "code": [inner_code],
      "message": "[message]"
    }
  ]
}
*/
type ErrorClientResponse struct {
	Code      string             `json:"code"` // outer_code
	CodeLabel string             `json:"code_label"`
	Errors    []ErrorDetailModel `json:"errors"`
}

type ErrorDetailModel struct {
	Attribute string `json:"attribute"`
	Code      int    `json:"code"` // inner_code
	Message   string `json:"message"`
	CodeLabel string `json:"code_label"`
}

type ErrorResponses struct {
	ErrorMessage string             `json:"error_message"`
	ErrorDetails []ErrorDetailModel `json:"error_details"`
}

const ProductGeneral = "general"
const ProductMoneyTransfer = "money_transfer"
const ProductSpecialMoneyTransfer = "special_money_transfer"
const ProductAgentMoneyTransfer = "agent_money_transfer"
const ProductAgentVerification = "agent_verification"
const ProductAcceptPayment = "accept_payment"
const ProductInternationalTransfer = "international_transfer"

const GeneralErrorBalanceInsufficient = "BALANCE_INSUFFICIENT"
const GeneralErrorValidationError = "VALIDATION_ERROR"

func GetAllGeneralError() map[string]string {
	generalError := map[string]string{
		GeneralErrorBalanceInsufficient: "Flip account balance is insufficient for the current transaction.", // (balance < (transfer amount + transfer fee))
		GeneralErrorValidationError:     "Error related to the validation of your payload data",
	}

	return generalError
}

func GetGeneralErrorLabel(generalError string) (string, error) {
	allGeneralError := GetAllGeneralError()
	if value, ok := allGeneralError[generalError]; !ok {
		return "", fmt.Errorf("general error [%s] not found", generalError)
	} else {
		return value, nil
	}
}

/*
// ------------------ TRANSFER ERROR

This TRANSFER ERROR can be used for
- Money Transfer,
- Special Money Transfer, and
- Agent Money Transfer
*/

const TransferErrorUndefined = "999"
const TransferErrorAttributeEmpty = "1001"
const TransferErrorAttributeNotCleanSymbols = "1002"
const TransferErrorAttributeNumber = "1020"
const TransferErrorAttributeMin = "1021"
const TransferErrorAttributeMax = "1022"
const TransferErrorMaxCharExceeded = "1024"
const TransferErrorAccountInvalid = "1025"
const TransferErrorAccountFround = "1026"
const TransferErrorAccountClosed = "1027"
const TransferErrorPagination = "1032"
const TransferErrorInvalidBankCode = "1033"
const TransferErrorInvalidCountry = "1034"
const TransferErrorInsufficientBalance = "1035"
const TransferErrorInvalidCity = "1038"
const TransferErrorInvalidDateFormat = "1039"
const TransferErrorInvalidDate = "1040"
const TransferErrorInvalidAttribute = "1041"
const TransferErrorMinOvo = "1042"
const TransferErrorBillTitleEmpty = "1043"
const TransferErrorMaxBeneficiaryEmail = "1070"
const TransferErrorInvalidBeneficiaryEmail = "1071"
const TransferErrorDisbursementIDNotFound = "1072"
const TransferErrorDisbursementIdemNotFound = "1073"
const TransferErrorDailyLimitExceeded = "1074"
const TransferErrorMaxLimitActive = "1080"
const TransferErrorMinDana = "1086"
const TransferErrorBeneficiaryBankDisturbance = "1088"
const TransferErrorFlipAccount = "1089"
const TransferErrorAgentStatusNotApproved = "1090"
const TransferErrorAgentNotActive = "1091"
const TransferErrorAgentNotAllowUpdate = "1092"
const TransferErrorCantProcess = "1093"
const TransferErrorStale = "1094"
const TransferErrorXTimestampInvalid = "1095"
const TransferErrorAttributeMinChar = "2001"
const TransferErrorAttributeDuplicated = "2002"
const TransferErrorAttributeNotClean = "2003"
const TransferErrorAttributeNonAlphaNum = "2004"

func GetAllTransferError() map[string]string {
	transferError := map[string]string{
		TransferErrorUndefined:                  "Undefined error",
		TransferErrorAttributeEmpty:             "The related attribute should not be empty",
		TransferErrorAttributeNotCleanSymbols:   "Value for the related attribute is considered not clean. Things that considered as not clean is html tag and ?, #, $, ' (single quote), \" (double quote), and ; character",
		TransferErrorAttributeNumber:            "The related attribute can only contain number",
		TransferErrorAttributeMin:               "Transfer amount are less than the minimum amount (Rp10.000)",
		TransferErrorAttributeMax:               "Transfer amount are more than the maximum amount",
		TransferErrorMaxCharExceeded:            "Max char for the related attribute exceeded",
		TransferErrorAccountInvalid:             "Invalid bank account number or possibly a virtual account number",
		TransferErrorAccountFround:              "The recipient bank account is suspected with fraud. You can’t send money to this account",
		TransferErrorAccountClosed:              "The account number is closed",
		TransferErrorPagination:                 "Pagination should be a number more than 0",
		TransferErrorInvalidBankCode:            "Invalid bank code",
		TransferErrorInvalidCountry:             "Country Code is invalid.",
		TransferErrorInsufficientBalance:        "Your balance is insufficient for doing this transaction",
		TransferErrorInvalidCity:                "Country/city code is not valid.",
		TransferErrorInvalidDateFormat:          "Date format is invalid",
		TransferErrorInvalidDate:                "Date is invalid",
		TransferErrorInvalidAttribute:           "Attribute is invalid",
		TransferErrorMinOvo:                     "Minimum transfer ovo amount is Rp10.000",
		TransferErrorBillTitleEmpty:             "Bill title param is required when creating bill",
		TransferErrorMaxBeneficiaryEmail:        "Max Beneficiary Email 20",
		TransferErrorInvalidBeneficiaryEmail:    "Invalid Beneficiary Email value",
		TransferErrorDisbursementIDNotFound:     "Disbursement ID not found",
		TransferErrorDisbursementIdemNotFound:   "Disbursement Idempotency Key not found",
		TransferErrorDailyLimitExceeded:         "Daily limit transaction exceeded",
		TransferErrorMaxLimitActive:             "You have reached the maximum limit of active transactions. Please complete your previous transactions.",
		TransferErrorMinDana:                    "Minimum transfer dana amount is Rp20.500",
		TransferErrorBeneficiaryBankDisturbance: "The beneficiary bank is in severe disturbance",
		TransferErrorFlipAccount:                "The account number is Flip’s. You can’t send money to Flip account",
		TransferErrorAgentStatusNotApproved:     "Agent KYC status is not approved",
		TransferErrorAgentNotActive:             "Agent status is not active",
		TransferErrorAgentNotAllowUpdate:        "Agent is not allowed to update",
		TransferErrorCantProcess:                "Cannot process transaction because bank is in cut off time. Try again later",
		TransferErrorStale:                      "The sent request/transaction is marked as stale request (more than 1 minute than the server time)",
		TransferErrorXTimestampInvalid:          "X-TIMESTAMP format is invalid. It should follow the ISO8601 format",
		TransferErrorAttributeMinChar:           "The related attribute should contain [at least or at most] [number] of characters",
		TransferErrorAttributeDuplicated:        "The related attribute is duplicated",
		TransferErrorAttributeNotClean:          "The related attribute is considered not clean. Only letters (a-z), numbers (0-9), spaces ( ), periods (.), commas (,), apostrophes (‘), and dashes (-) are allowed",
		TransferErrorAttributeNonAlphaNum:       "The related attribute contains non-alphanumeric characters",
	}

	return transferError
}

func GetTransferErrorCodeLabel(errorCode string) (message string, err error) {
	transferError := GetAllTransferError()
	if value, ok := transferError[errorCode]; !ok {
		return "", fmt.Errorf("transfer money error code [%s] not found", errorCode)
	} else {
		return value, nil
	}
}

// ------------------ AGENT VERIFICATION ERROR

const AgentVerErrorInvalidUserID = "400"
const AgentVerErrorInvalidToken = "401"
const AgentVerErrorForbidden = "403"
const AgentVerErrorEndpointNotFound = "404"
const AgentVerErrorInternalServerError = "500"
const AgentVerErrorInternalServerError503 = "503"
const AgentVerErrorAttributeEmpty = "1001"
const AgentVerErrorAttributeInvalid = "1002"
const AgentVerErrorAttributeCharInvalid = "2001"
const AgentVerErrorAttributeDuplicate = "2002"
const AgentVerErrorAttributeNonSterile = "2003"
const AgentVerErrorAttributeNonAlphaNum = "2004"
const AgentVerErrorAgentIDNotFound = "1011"
const AgentVerErrorInvalidInput = "4001"
const AgentVerErrorInvalidUserType = "4002"
const AgentVerErrorInvalidUserID4003 = "4003"
const AgentVerErrorInvalidIdentityType = "4004"
const AgentVerErrorRejectedReasonNotMatch = "4006"
const AgentVerErrorImageSizeMax = "4007"
const AgentVerErrorImageNotValid = "4008"
const AgentVerErrorFailReadImage = "4009"
const AgentVerErrorAgentStatusNotValid = "4010"
const AgentVerErrorRetryMaxAchieved = "4011"
const AgentVerErrorIdentityNumExists = "4012"
const AgentVerErrorInvalidDocType = "4013"
const AgentVerErrorImageNotSimilar = "4014"
const AgentVerErrorCantCreate = "4015"

func GetAllAgentVerificationErrorLabel() map[string]string {
	agentVerificationLabel := map[string]string{
		AgentVerErrorInvalidUserID:          "Invalid user ID",                                      // "Recheck the userID in URL params"
		AgentVerErrorInvalidToken:           "Invalid token",                                        // "Header Request-ID must be filledMissing user_typeIP must be whitelisted"
		AgentVerErrorForbidden:              "Forbidden",                                            // "Missing user_type"
		AgentVerErrorEndpointNotFound:       "Endpoint not found",                                   // "Recheck API URL"
		AgentVerErrorInternalServerError:    "Internal server error",                                // "Reach out to FLIP KYC team"
		AgentVerErrorInternalServerError503: "Internal server error",                                // "Reach out to FLIP KYC team"
		AgentVerErrorAttributeEmpty:         "Attribute payload is empty",                           // "Fill the attribute payload"
		AgentVerErrorAttributeInvalid:       "Param attribute is invalid",                           // "Ensure attribute’s validation"
		AgentVerErrorAttributeCharInvalid:   "Param attribute characters is invalid",                // "Ensure attribute’s validation"
		AgentVerErrorAttributeDuplicate:     "Param attribute is duplicate",                         // "Use different identity number"
		AgentVerErrorAttributeNonSterile:    "Param attribute contains non-sterile characters",      // "Ensure attribute’s validation"
		AgentVerErrorAttributeNonAlphaNum:   "Param attribute contains non-alphanumeric characters", // "Ensure attribute’s validation"
		AgentVerErrorAgentIDNotFound:        "Agent ID is not found",                                // "Agent ID not registered yet"
		AgentVerErrorInvalidInput:           "Invalid input parameter",                              // "Ensure attribute’s validation"
		AgentVerErrorInvalidUserType:        "Invalid user type",                                    // "Ensure attribute’s validation"
		AgentVerErrorInvalidUserID4003:      "Invalid user ID",                                      // "Recheck the userID in URL params"
		AgentVerErrorInvalidIdentityType:    "Invalid identity type",                                // "Ensure attribute’s validation"
		AgentVerErrorRejectedReasonNotMatch: "Rejected reason is not match",                         // "Reheck rejected_reason_code. It should match with the API Repair type."
		AgentVerErrorImageSizeMax:           "Image file size is too large",                         // "Image size should be less than 2MB"
		AgentVerErrorImageNotValid:          "Image is not valid",                                   // "Please ensure image validations"
		AgentVerErrorFailReadImage:          "Failed reading image data",                            // "Ensure uploaded image is not corrupted"
		AgentVerErrorAgentStatusNotValid:    "Invalid agent status",                                 // "Depends on which API this error displayed: KYC Submission: Ensure agent’s KYC status is 19 (UPLOAD_IDENTITY_SELFIE_SUCCESS)Upload Supporting Documents: Ensure agent’s KYC status is 50 (REJECTED) or 15 (UPLOAD_IDENTITY_SELFIE_FAILED)"
		AgentVerErrorRetryMaxAchieved:       "Maximum retry has been achieved",                      // "Upload supporting documents"
		AgentVerErrorIdentityNumExists:      "Identity number has been used/registered",             // "Use different identity number"
		AgentVerErrorInvalidDocType:         "Invalid document type",                                // "Ensure attribute’s validation"
		AgentVerErrorImageNotSimilar:        "Image not similar",                                    // "Reach out to FLIP KYC team"
		AgentVerErrorCantCreate:             "Cannot create/update agent data",                      // "Reach out to FLIP KYC team"
	}

	return agentVerificationLabel
}

func GetAgentVerificationErrorLabel(agentVerStatus string) (string, error) {
	labels := GetAllAgentVerificationErrorLabel()
	if value, ok := labels[agentVerStatus]; !ok {
		return "", fmt.Errorf("agent verification status [%s] not found", agentVerStatus)
	} else {
		return value, nil
	}
}

// ------------------ ACCEPT PAYMENT ERROR LIST

const AcpPaymentErrorUndefined = "999"
const AcpPaymentErrorBillTypeRequired = "1044"
const AcpPaymentErrorBillTypeInvalid = "1045"
const AcpPaymentErrorBillMaxExceeded = "1046"
const AcpPaymentErrorBillExpInvalid = "1047"
const AcpPaymentErrorBillAlreadyExp = "1048"
const AcpPaymentErrorBillAlreadyUsed = "1049"
const AcpPaymentErrorBillIDNotFound = "1050"
const AcpPaymentErrorProductCodeRequired = "1051"
const AcpPaymentErrorPageRequired = "1052"
const AcpPaymentErrorPaginationRequired = "1053"
const AcpPaymentErrorPaginationNumber = "1054"
const AcpPaymentErrorProductBillIDNotFound = "1055"
const AcpPaymentErrorStartDateInvalid = "1056"
const AcpPaymentErrorEndDateInvalid = "1057"
const AcpPaymentErrorInvalidDateRange = "1058"
const AcpPaymentErrorTransferMin = "1059"
const AcpPaymentErrorRedirectUrlInvalid = "1060"
const AcpPaymentErrorStatusInvalid = "1061"
const AcpPaymentErrorTitleInvalid = "1062"
const AcpPaymentErrorSortByInvalid = "1063"
const AcpPaymentErrorSortTypeInvalid = "1064"
const AcpPaymentErrorAmountExceeded = "1067"

func GetAllAcceptPaymentErrorLabel() map[string]string {
	labels := map[string]string{
		AcpPaymentErrorUndefined:             "Undefined error",
		AcpPaymentErrorBillTypeRequired:      "Bill type param is required when creating bill",
		AcpPaymentErrorBillTypeInvalid:       "Bill type param is invalid when creating bill",
		AcpPaymentErrorBillMaxExceeded:       "Maximum amount for the bill is Rp10.000.000",
		AcpPaymentErrorBillExpInvalid:        "Bill expired date param is invalid when creating bill",
		AcpPaymentErrorBillAlreadyExp:        "Bill link has already expired",
		AcpPaymentErrorBillAlreadyUsed:       "Bill link has already used",
		AcpPaymentErrorBillIDNotFound:        "Bill link ID is not Found",
		AcpPaymentErrorProductCodeRequired:   "Param product_code is required",
		AcpPaymentErrorPageRequired:          "Param page should be a number and more than 0",
		AcpPaymentErrorPaginationRequired:    "Param pagination should be a number and more than 0",
		AcpPaymentErrorPaginationNumber:      "Product bill link ID is not found",
		AcpPaymentErrorProductBillIDNotFound: "Param start_date is invalid",
		AcpPaymentErrorStartDateInvalid:      "Param end_date is invalid",
		AcpPaymentErrorEndDateInvalid:        "Invalid date range. Param end_date must be greater than start_date",
		AcpPaymentErrorInvalidDateRange:      "Minimum transfer amount is Rp10.000",
		AcpPaymentErrorTransferMin:           "Param redirect_url is invalid",
		AcpPaymentErrorRedirectUrlInvalid:    "Param status is invalid",
		AcpPaymentErrorStatusInvalid:         "Param title is invalid",
		AcpPaymentErrorTitleInvalid:          "Param sort_by is invalid",
		AcpPaymentErrorSortByInvalid:         "Param sort_type is invalid",
		AcpPaymentErrorSortTypeInvalid:       "Param amount must be a number",
		AcpPaymentErrorAmountExceeded:        "Param amount is exceeding company’s PWF maximum limit amount",
	}

	return labels
}

func GetAcceptPaymentErrorLabel(acceptPaymentError string) (string, error) {
	labels := GetAllAcceptPaymentErrorLabel()
	if value, ok := labels[acceptPaymentError]; !ok {
		return "", fmt.Errorf("accept payment error code [%s] not found", acceptPaymentError)
	} else {
		return value, nil
	}
}

// --------- INTERNATIONAL TRANSFER ERROR

const IntTransferErrorUndefined = "999"
const IntTransferErrorAttributeNumber = "1020"
const IntTransferErrorAttributeMaxChar = "1024"
const IntTransferErrorInvalidCountry = "1034"
const IntTransferErrorInvalidCity = "1038"
const IntTransferErrorInvalidDateFormat = "1039"
const IntTransferErrorInvalidDate = "1040"
const IntTransferErrorInvalidAttribute = "1041"
const IntTransferErrorInvalidCountryCode = "1068"
const IntTransferErrorInvalidTransType = "1069"
const IntTransferErrorAmountMaxExceeded = "1081"
const IntTransferErrorAmountMinNotReach = "1082"
const IntTransferErrorAmountDecimal = "1083"
const IntTransferErrorAttributeAlphaNum = "1084"
const IntTransferErrorWordMin = "1085"
const IntTransferErrorAttributeMinChar = "1087"
const IntTransferErrorRequestStale = "1094"
const IntTransferErrorTimestampInvalid = "1095"

func GetAllInternationalTransferErrorLabel() map[string]string {
	labels := map[string]string{
		IntTransferErrorUndefined:          "Undefined error",
		IntTransferErrorAttributeNumber:    "The related attribute can only contain number",
		IntTransferErrorAttributeMaxChar:   "Max char for the related attribute exceeded",
		IntTransferErrorInvalidCountry:     "Country is invalid. Please see available country code",
		IntTransferErrorInvalidCity:        "Country/city code is not valid", // The difference with 1037 code is 1037 will occur if the attribute only allows country code while this code will occur if the attribute allows country or city code
		IntTransferErrorInvalidDateFormat:  "Date format is invalid",
		IntTransferErrorInvalidDate:        "Date is invalid",
		IntTransferErrorInvalidAttribute:   "Attribute is invalid",
		IntTransferErrorInvalidCountryCode: "Param country_iso_code is invalid",
		IntTransferErrorInvalidTransType:   "Param transaction_type is invalid",
		IntTransferErrorAmountMaxExceeded:  "Maximum amount is exceeded",
		IntTransferErrorAmountMinNotReach:  "Minimum amount has not been reached",
		IntTransferErrorAmountDecimal:      "Cannot use decimal amount",
		IntTransferErrorAttributeAlphaNum:  "The related attribute can only contain alphanumeric character",
		IntTransferErrorWordMin:            "Minimum words required",
		IntTransferErrorAttributeMinChar:   "Attribute must consist of X character",
		IntTransferErrorRequestStale:       "The sent request/transaction is marked as stale request (more than 1 minute than the server time)",
		IntTransferErrorTimestampInvalid:   "X-TIMESTAMP format is invalid. It should follow the ISO8601 format",
	}

	return labels
}

func GetInternationalTransferErrorLabel(intTransferError string) (string, error) {
	labels := GetAllInternationalTransferErrorLabel()
	if value, ok := labels[intTransferError]; !ok {
		return "", fmt.Errorf("international transfer error code [%s] not found", intTransferError)
	} else {
		return value, nil
	}
}
