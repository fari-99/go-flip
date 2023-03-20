package constants

import (
	"fmt"
	"time"
)

const RSAKeyBitTotal = 2048

// TimeFormatHeader using ISO8601 2020-07-10 15:00:00.000 is represents the 10th of July 2020 at 3 p.m.
const TimeFormatHeader = time.RFC3339
const TimeFormatBirthDate = "2006-01-02"
const TimeFormatExpiredDate = "2006-01-02 15:04"
const TimeFormatPagination = "2006-01-02"

const ApiUrlDev = "https://bigflip.id/big_sandbox_api"
const ApiUrlKycDev = "https://api.flip.id/kyc-sandbox/api/v1"
const ApiUrlProdV2 = "https://bigflip.id/api/v2"
const ApiUrlProdV3 = "https://bigflip.id/api/v3"
const ApiUrlKyc = "https://api.flip.id/kyc/api/v1"

const ApiDev = "dev"
const ApiProdV2 = "v2"
const ApiProdV3 = "v3"
const ApiKYC = "kyc"

const IdentityAddressFlagTrue = 1
const IdentityAddressFlagFalse = 0

const GenderMale = "male"
const GenderFemale = "female"

const SelfieFlagTrue = "1"  // Denotes that the uploaded file is selfie with identity image.
const SelfieFlagFalse = "0" // Denotes that the uploaded file is identity image.

const MaintenanceMessage = "flip currently in maintenance, please try again later [LOG: %s]"

const HeaderIdempotencyKey = "idempotency-key"
const HeaderAuthorization = "Authorization"
const HeaderContentType = "Content-Type"
const HeaderRequestID = "Request-ID"
const HeaderTimestamp = "X-TIMESTAMP"

const ContentTypeFormUrlEncoded = "application/x-www-form-urlencoded"
const ContentTypeMultipart = "multipart/form-data"

const SortTypeDESC = "sort_desc" // descending
const SortTypeASC = "sort_asc"   // ascending

// ---------------- BANK STATUS

// BankStatusOperational
// Banks are operational, disbursement will be processed as soon as possible
const BankStatusOperational = "OPERATIONAL"

// BankStatusDisturbed
// Banks are slow or have another problem.
// Disbursement will still be processed, but in slower pace and might be delayed
const BankStatusDisturbed = "DISTURBED"

// BankStatusHeavilyDisturbed
// Banks are having an error, offline, or another problem that result in a nearly unusable system.
// Disbursement to this bank can not be processed in a short time, and maybe won’t be processed in the same day.
// You can ask for a refund if this happens.
const BankStatusHeavilyDisturbed = "HEAVILY_DISTURBED"

// ---------------- DISBURSEMENT STATUS

// DisbursementStatusPending
// Disbursement is still in process
const DisbursementStatusPending = "PENDING"

// DisbursementStatusCancelled
// The transaction is cancelled and the amount of the transaction plus the transaction fee will be credited to your balance.
// This will happen if the transfer process are failed for reason such as inactive recipient account, wrong account number, etc
const DisbursementStatusCancelled = "CANCELLED"

// DisbursementStatusDone
// Disbursement process is finished and the money have been sent to the recipient
const DisbursementStatusDone = "DONE"

// DisbursementReasonInactiveAccount Inactive account / Nomor rekening tidak aktif
const DisbursementReasonInactiveAccount = "INACTIVE_ACCOUNT "

// DisbursementReasonNotRegisteredAccount Not registered account / Nomor rekening tidak terdaftar
const DisbursementReasonNotRegisteredAccount = "NOT_REGISTERED_ACCOUNT "

// DisbursementReasonCantReceiveTransfer Can’t receive transfer / Rekening tujuan tidak dapat menerima transfer
const DisbursementReasonCantReceiveTransfer = "CANT_RECEIVE_TRANSFER "

// DisbursementReasonIntermittentDisturbance Intermittent disturbance on destination bank / Bank tujuan sementara mengalami gangguan
const DisbursementReasonIntermittentDisturbance = "INTERMITTENT_DISTURBANCE_ON_BENEFICIARY_BANK"

// DisbursementReasonNotVerified Account is not verified / Rekening tujuan belum terverifikasi
const DisbursementReasonNotVerified = "BENEFICIARY_ACCOUNT_NOT_VERIFIED "

func GetAllDisbursementStatusLabel() map[string]string {
	disbursementStatus := map[string]string{
		DisbursementStatusPending:                 "Disbursement is still in process",
		DisbursementStatusCancelled:               "The transaction is cancelled",
		DisbursementStatusDone:                    "Disbursement process is finished",
		DisbursementReasonInactiveAccount:         "Inactive account",
		DisbursementReasonNotRegisteredAccount:    "Not registered account",
		DisbursementReasonCantReceiveTransfer:     "Can’t receive transfer",
		DisbursementReasonIntermittentDisturbance: "Intermittent disturbance on destination bank",
		DisbursementReasonNotVerified:             "Account is not verified",
	}

	return disbursementStatus
}

func GetDisbursementStatusLabel(disbursementStatus string) (string, error) {
	disbursementStatuses := GetAllDisbursementStatusLabel()
	if value, ok := disbursementStatuses[disbursementStatus]; !ok {
		return "", fmt.Errorf("disbursement status [%s] not found", disbursementStatus)
	} else {
		return value, nil
	}
}

// ---------------- DISBURSEMENT DIRECTION

// DisbursementDirectionDomesticTransfer
// Common Disbursement from Indonesia to Indonesian recipient
const DisbursementDirectionDomesticTransfer = "DOMESTIC_TRANSFER"

// DisbursementDirectionDomesticSpecialTransfer
// Special disbursement from the user of a Money Transfer Company residing in Indonesia to Indonesian recipient
const DisbursementDirectionDomesticSpecialTransfer = "DOMESTIC_SPECIAL_TRANSFER"

// DisbursementDirectionForeignInboundSpecialTransfer
// Special disbursement from the user of a Money Transfer Company residing in a foreign country to Indonesian recipient
const DisbursementDirectionForeignInboundSpecialTransfer = "FOREIGN_INBOUND_SPECIAL_TRANSFER"

// ---------------- DISBURSEMENT SENDER IDENTITY TYPE

const SenderIdentityNatID = "nat_id"
const SenderIdentityDriverLicense = "drv_lic"
const SenderIdentityPassport = "passport"
const SenderIdentityBankAccount = "bank_acc" // Default

func GetAllSenderIdentity() map[string]string {
	labels := map[string]string{
		SenderIdentityNatID:         "National ID",
		SenderIdentityDriverLicense: "Driver License",
		SenderIdentityPassport:      "Passport",
		SenderIdentityBankAccount:   "Bank Account", // Default
	}

	return labels
}

func GetSenderIdentity(senderIdentity string) (string, error) {
	labels := GetAllSenderIdentity()
	if value, ok := labels[senderIdentity]; !ok {
		return "", fmt.Errorf("sender identity [%s] not found", senderIdentity)
	} else {
		return value, nil
	}
}

// ---------------- JOB TYPE

const JobTypeHouseWife = "housewife"
const JobTypeEntrepreneur = "entrepreneur"
const JobTypePrivateEmployee = "private_employee"
const JobTypeGovernmentEmployee = "government_employee"
const JobTypeFoundationBoard = "foundation_board"
const JobTypeIndonesianMigrantWorker = "indonesian_migrant_worker"
const JobTypeCompany = "company"
const JobTypeOthers = "others"

func GetAllJobTypeLabel() map[string]string {
	labels := map[string]string{
		JobTypeHouseWife:               "Housewife",
		JobTypeEntrepreneur:            "Entrepreneur",
		JobTypePrivateEmployee:         "Private Employee",
		JobTypeGovernmentEmployee:      "Government Employee",
		JobTypeFoundationBoard:         "Foundation Board",          // People who work at foundation as an employee
		JobTypeIndonesianMigrantWorker: "Indonesian Migrant Worker", // Also known as TKI
		JobTypeCompany:                 "Company",
		JobTypeOthers:                  "Others",
	}

	return labels
}

func GetJobTypeLabel(jobType string) (string, error) {
	labels := GetAllJobTypeLabel()
	if value, ok := labels[jobType]; !ok {
		return "", fmt.Errorf("job type [%s] not found", jobType)
	} else {
		return value, nil
	}
}

// ---------------- AGENT IDENTITY TYPE

const AgentIdentityKTP = "ktp"
const AgentIdentityPassport = "passport"
const AgentIdentityKTPEnum = "1"
const AgentIdentityPassportEnum = "2"

func GetAllIdentityTypeLabel() map[string]string {
	labels := map[string]string{
		AgentIdentityKTPEnum:      AgentIdentityKTP,
		AgentIdentityPassportEnum: AgentIdentityPassport,
	}

	return labels
}

func GetIdentityTypeLabel(identityType string) (string, error) {
	labels := GetAllIdentityTypeLabel()
	if value, ok := labels[identityType]; !ok {
		return "", fmt.Errorf("identity type [%s] for agent not found", identityType)
	} else {
		return value, nil
	}
}

// ---------------- AGENT STATUS

const AgentStatusDeleted = "0"
const AgentStatusActive = "10"
const AgentStatusBlacklisted = "20"
const AgentStatusBlocked = "30"

func GetAllAgentStatusLabel() map[string]string {
	agentStatus := map[string]string{
		AgentStatusDeleted:     "Deleted",
		AgentStatusActive:      "Active",
		AgentStatusBlacklisted: "Blacklisted",
		AgentStatusBlocked:     "Blocked",
	}

	return agentStatus
}

func GetAgentStatusLabel(agentStatus string) (string, error) {
	agentStatuses := GetAllAgentStatusLabel()
	if value, ok := agentStatuses[agentStatus]; !ok {
		return "", fmt.Errorf("agent status [%s] not found", agentStatus)
	} else {
		return value, nil
	}
}

// ---------------- USER TYPE

const UserTypeAgent = "1"

// ---------------- AGENT KYC STATUS

const KycStatusBasicData = "BASIC_DATA"
const KycStatusUploadIdentityFailed = "UPLOAD_IDENTITY_FAILED"
const KycStatusUploadIdentitySuccess = "UPLOAD_IDENTITY_SUCCESS"
const KycStatusUploadSelfieProgress = "UPLOAD_IDENTITY_SELFIE_PROGRESS"
const KycStatusUploadSelfieFailed = "UPLOAD_IDENTITY_SELFIE_FAILED"
const KycStatusUploadSelfieSuccess = "UPLOAD_IDENTITY_SELFIE_SUCCESS"
const KycStatusPending = "PENDING"
const KycStatusSuspicious = "SUSPICIOUS"
const KycStatusNeedCheck = "NEED_TO_CHECK"
const KycStatusHelp = "HELP"
const KycStatusApproved = "APPROVED"
const KycStatusRejected = "REJECTED"

func GetAllKycStatusLabel() map[string]string {
	labels := map[string]string{
		KycStatusBasicData:             "Agent has been created successfully",
		KycStatusUploadIdentityFailed:  "Identity image is failed to be uploaded",
		KycStatusUploadIdentitySuccess: "Identity image is succeed to be uploaded",
		KycStatusUploadSelfieProgress:  "Selfie identity image has been uploaded",
		KycStatusUploadSelfieFailed:    "Selfie identity image is failed to be uploaded",
		KycStatusUploadSelfieSuccess:   "Selfie identity image is succeed to be uploaded",
		KycStatusPending:               "Agent is waiting to be checked by Flip ops",
		KycStatusSuspicious:            "Agent is flagged by Flip’s admin, and needs to be checked.",
		KycStatusNeedCheck:             "Agent is flagged by Flip’s admin, and needs to be checked.",
		KycStatusHelp:                  "Need to upload supporting documents, usually after Agent is rejected",
		KycStatusApproved:              "KYC Approved",
		KycStatusRejected:              "KYC Rejected",
	}

	return labels
}

func GetKycStatusLabel(kycStatus string) (string, error) {
	kycStatuses := GetAllKycStatusLabel()
	if value, ok := kycStatuses[kycStatus]; !ok {
		return "", fmt.Errorf("agent kyc status [%s] not found", kycStatus)
	} else {
		return value, nil
	}
}

// ---------------- AGENT KYC REJECTION REASONS

const KycRejectionDataNotMatch = "10"
const KycRejectionNIKFormat = "11"
const KycRejectionNIKAlreadyRegistered = "12"
const KycRejectionNIKBolded = "13"
const KycRejectionCantRead = "14"
const KycRejectionIdentityTypeNotValid = "15"
const KycRejectionIdentityNotCaptured = "16"
const KycRejectionIdentityPartCantRead = "17"
const KycRejectionIdentityBrokenUp = "18"
const KycRejectionIdentityExpired = "19"
const KycRejectionNotOriginal = "20"
const KycRejectionBackground = "21"
const KycRejectionNotOriginalPhoto = "22"
const KycRejectionCombinationNotVisible = "23"
const KycRejectionCombinationFaceImage = "24"
const KycRejectionCombinationNotMatch = "25"
const KycRejectionCombinationNotHold = "26"
const KycRejectionPhotoTakenIndirectly = "27"
const KycRejectionUsingVeil = "28"
const KycRejectionViolateTOC = "29"
const KycRejectionNeedConfirmation = "30"

func GetAllKycRejectionReasonLabel() map[string]string {
	labels := map[string]string{
		KycRejectionDataNotMatch:          "The data that is filled in does not match with identity card [Repair ID data]",
		KycRejectionNIKFormat:             "NIK does not match the format [Upload supporting document (Kartu Keluarga/NPWP/BPJS Kesehatan)]",
		KycRejectionNIKAlreadyRegistered:  "NIK already registered in another account [Chat with CX team]",
		KycRejectionNIKBolded:             "NIK data is bolded manually [Upload supporting document (Kartu Keluarga/NPWP/BPJS Kesehatan)]",
		KycRejectionCantRead:              "Identity card blurry or got spotlight or cannot be read [Upload ID again]",
		KycRejectionIdentityTypeNotValid:  "The Identity card uploaded is not (KTP/Passport/KITAS) [Upload ID again]",
		KycRejectionIdentityNotCaptured:   "Identity card is not fully captured [Upload ID again]",
		KycRejectionIdentityPartCantRead:  "Part of the data on the photo ID card is missing, so it can’t be read in its entirety [Upload ID again OR Upload supporting documents (Kartu Keluarga/NPWP/BPJS Kesehatan)]",
		KycRejectionIdentityBrokenUp:      "If the identity card is broken up to the data section on the identity card [Upload ID again]",
		KycRejectionIdentityExpired:       "ID card has expired [Upload ID again]",
		KycRejectionNotOriginal:           "Photo ID card is not original [Upload ID again]",
		KycRejectionBackground:            "The background of the user’s pass photo on the ID card is not blue/red [Upload ID again OR Upload supporting documents (Kartu Keluarga/NPWP/BPJS Kesehatan)]",
		KycRejectionNotOriginalPhoto:      "Pass photo in identity card is not original but only patch photo [Chat with CX team]",
		KycRejectionCombinationNotVisible: "Combination photos are not clearly visible or partially covered [Upload selfie with ID again]",
		KycRejectionCombinationFaceImage:  "Combination photos only contain face images [Upload selfie with ID again]",
		KycRejectionCombinationNotMatch:   "Combination photo of the user holding another identity card other than (KTP/Passport/KITAS) [Upload selfie with ID again]",
		KycRejectionCombinationNotHold:    "Combination photos do not directly hold an identity card [Upload selfie with ID again]",
		KycRejectionPhotoTakenIndirectly:  "Photos taken indirectly [Upload selfie with ID again]",
		KycRejectionUsingVeil:             "User wearing veil [Chat with CX team]",
		KycRejectionViolateTOC:            "Indicated in violation of the Terms and Conditions of the application [Fraud user - Can’t do anything]",
		KycRejectionNeedConfirmation:      "We need confirm your KYC with some data [Chat with CX team]",
	}

	return labels
}

func GetKycRejectionReasonLabel(kycRejectionReason string) (string, error) {
	kycRejectionReasons := GetAllKycRejectionReasonLabel()
	if value, ok := kycRejectionReasons[kycRejectionReason]; !ok {
		return "", fmt.Errorf("agent kyc rejection reasons [%s] not found", kycRejectionReason)
	} else {
		return value, nil
	}
}

// ---------------- AGENT KYC STATUS CALLBACK

const KyCStatusCallbackUploadSelfieFailed = "15"
const KyCStatusCallbackUploadSelfieSuccess = "19"
const KyCStatusCallbackApproved = "40"
const KyCStatusCallbackRejected = "50"

func GetAllAgentKycStatusCallbackLabel() map[string]string {
	labels := map[string]string{
		KyCStatusCallbackUploadSelfieFailed:  "Selfie identity image is failed to be uploaded",
		KyCStatusCallbackUploadSelfieSuccess: "Selfie identity image is succeed to be uploaded",
		KyCStatusCallbackApproved:            "KYC Approved",
		KyCStatusCallbackRejected:            "KYC Rejected",
	}

	return labels
}

func GetAgentKycStatusCallbackLabel(statusCode string) (string, error) {
	labels := GetAllAgentKycStatusCallbackLabel()
	if value, ok := labels[statusCode]; !ok {
		return "", fmt.Errorf("agent kyc status callback [%s] not found", statusCode)
	} else {
		return value, nil
	}
}

// ---------------- BILLING

const BillTypeSingle = "SINGLE"
const BillTypeMultiple = "MULTIPLE"

const BillStepOne = "1"
const BillStepTwo = "2"
const BillStepThree = "3"

const BillStatusActive = "ACTIVE"
const BillStatusInActive = "INACTIVE"

const BillPaymentStatusNotConfirmed = "NOT_CONFIRMED"
const BillPaymentStatusPending = "PENDING"
const BillPaymentStatusProcessed = "PROCESSED"
const BillPaymentStatusCancelled = "CANCELED"
const BillPaymentStatusFailed = "FAILED"
const BillPaymentStatusDone = "DONE"

func GetAllBillPaymentStatusLabel() map[string]string {
	labels := map[string]string{
		BillPaymentStatusNotConfirmed: "The transaction was just created.",
		BillPaymentStatusPending:      "The transaction has been confirmed by the user.",
		BillPaymentStatusProcessed:    "The transaction is being processed by the system.",
		BillPaymentStatusCancelled:    "The transaction is canceled.",
		BillPaymentStatusFailed:       "The transaction is failed.",
		BillPaymentStatusDone:         "The transaction is successfully done.",
	}

	return labels
}

func GetBillPaymentStatusLabel(billPaymentStatus string) (string, error) {
	labels := GetAllBillPaymentStatusLabel()
	if value, ok := labels[billPaymentStatus]; !ok {
		return "", fmt.Errorf("bill payment status [%s] not found", billPaymentStatus)
	} else {
		return value, nil
	}
}

const PaymentSortByID = "id"
const PaymentSortByBillLink = "bill_link"
const PaymentSortByBillTitle = "bill_title"
const PaymentSortBySenderBank = "sender_bank"
const PaymentSortByAmount = "amount"
const PaymentSortByCreatedAt = "created_at"

func GetAllPaymentSortByLabel() map[string]string {
	labels := map[string]string{
		PaymentSortByID:         "ID",
		PaymentSortByBillLink:   "Bill Link",
		PaymentSortByBillTitle:  "Bill Title",
		PaymentSortBySenderBank: "Sender Bank",
		PaymentSortByAmount:     "Amount",
		PaymentSortByCreatedAt:  "Created At",
	}

	return labels
}

func GetPaymentSortByLabels(sortBy string) (string, error) {
	labels := GetAllPaymentSortByLabel()
	if value, ok := labels[sortBy]; !ok {
		return "", fmt.Errorf("payment sort by code [%s] not found", sortBy)
	} else {
		return value, nil
	}
}

// ---------------- SENDER BANK TYPE

const BankTypeBankAccount = "bank_account"
const BankTypeVirtualAccount = "virtual_account"
const BankTypeWalletAccount = "wallet_account"

// ---------------- TRANSACTION TYPE

const TransactionTypeC2C = "C2C"
const TransactionTypeC2B = "C2B"
const TransactionTypeB2C = "B2C"
const TransactionTypeB2B = "B2B"

// ---------------- INTERNATIONAL CANCELLED TRANSACTION REASONS

const IntCancelReasonDeclined = "DECLINED"
const IntCancelReasonCancelled = "CANCELLED"
const IntCancelReasonRefused = "REFUSED"
const IntCancelReasonInvalidBeneficiary = "INVALID_BENEFICIARY"
const IntCancelReasonInvalidBeneficiaryDetails = "INVALID_BENEFICIARY_DETAILS"
const IntCancelReasonLimitTransValue = "LIMITATIONS_ON_TRANSACTION_VALUE"
const IntCancelReasonSLSSender = "SLS_SENDER"
const IntCancelReasonSLSBeneficiary = "SLS_BENEFICIARY"
const IntCancelReasonBarredBeneficiary = "BARRED_BENEFICIARY"
const IntCancelReasonBarredSender = "BARRED_SENDER"
const IntCancelReasonCompliance = "COMPLIANCE_REASON"
const IntCancelReasonUnsupportedBeneficiary = "UNSUPPORTED_BENEFICIARY"
const IntCancelReasonDuplicate = "DUPLICATED_TRANSACTION"
const IntCancelReasonLimitSenderQty = "LIMITATIONS_ON_SENDER_QUANTITY"
const IntCancelReasonLimitBeneficiaryQty = "LIMITATIONS_ON_BENEFICIARY_QUANTITY"
const IntCancelReasonLimitAccQty = "LIMITATIONS_ON_ACCOUNT_QUANTITY"
const IntCancelReasonLimitAccValue = "LIMITATIONS_ON_ACCOUNT_VALUE"

func GetAllInternationalCancelledReason() map[string]string {
	labels := map[string]string{
		IntCancelReasonDeclined:                  "Transaction cannot be processed.",
		IntCancelReasonCancelled:                 "Transaction was cancelled by system.",
		IntCancelReasonRefused:                   "Transaction was not complete because beneficiary declined the fund.",
		IntCancelReasonInvalidBeneficiary:        "Transaction was failed due to invalid beneficiary account, e.g. bank account number.",
		IntCancelReasonInvalidBeneficiaryDetails: "Transaction was failed due to invalid beneficiary details, e.g. name, email, phone number.",
		IntCancelReasonLimitTransValue:           "Transaction was failed due to value of transaction exceeding the limit.",
		IntCancelReasonSLSSender:                 "Transaction was rejected due to sender failing the sanction list screening.",
		IntCancelReasonSLSBeneficiary:            "Transaction was not complete due to beneficiary failing the sanction list screening.",
		IntCancelReasonBarredBeneficiary:         "Transaction was not complete due to beneficiary being blacklisted.",
		IntCancelReasonBarredSender:              "Transaction was rejected due to sender being blacklisted.",
		IntCancelReasonCompliance:                "Transaction was rejected due to compliance reason.",
		IntCancelReasonUnsupportedBeneficiary:    "Transaction was not complete due to beneficiary’s inability to receive the fund.",
		IntCancelReasonDuplicate:                 "Transaction failed due it was a duplicate.",
		IntCancelReasonLimitSenderQty:            "Transaction was rejected due to sender had exceeded the transaction count limit.",
		IntCancelReasonLimitBeneficiaryQty:       "Transaction was not complete due to beneficiary had exceeded the transaction count limit.",
		IntCancelReasonLimitAccQty:               "Transaction was not complete due to beneficiary had exceeded the account transaction count limit.",
		IntCancelReasonLimitAccValue:             "Transaction was not complete due to beneficiary had reached the account transaction value limit.",
	}

	return labels
}

func GetInternationalCancelledReason(cancelReason string) (string, error) {
	labels := GetAllInternationalCancelledReason()
	if value, ok := labels[cancelReason]; !ok {
		return "", fmt.Errorf("international transaction cancel reason [%s] not found", cancelReason)
	} else {
		return value, nil
	}
}

// ---------------- BENEFICIARY RELATIONSHIP

const BeneficiaryRelSelf = "SELF"
const BeneficiaryRelBrother = "BROTHER"
const BeneficiaryRelSister = "SISTER"
const BeneficiaryRelSon = "SON"
const BeneficiaryRelDaughter = "DAUGHTER"
const BeneficiaryRelNephew = "NEPHEW"
const BeneficiaryRelNiece = "NIECE"
const BeneficiaryRelFather = "FATHER"
const BeneficiaryRelMother = "MOTHER"
const BeneficiaryRelUncle = "UNCLE"
const BeneficiaryRelAunt = "AUNT"
const BeneficiaryRelCousin = "COUSIN"
const BeneficiaryRelFIL = "FATHER_IN_LAW"
const BeneficiaryRelMIL = "MOTHER_IN_LAW"
const BeneficiaryRelBIL = "BROTHER_IN_LAW"
const BeneficiaryRelSIL = "SISTER_IN_LAW"
const BeneficiaryRelGrandFather = "GRAND_FATHER"
const BeneficiaryRelGrandMother = "GRAND_MOTHER"
const BeneficiaryRelHusband = "HUSBAND"
const BeneficiaryRelWife = "WIFE"
const BeneficiaryRelFriend = "FRIEND"

func GetAllBeneficiaryRelationshipLabel() map[string]string {
	labels := map[string]string{
		BeneficiaryRelSelf:        "Diri sendiri",
		BeneficiaryRelBrother:     "Saudara laki-laki",
		BeneficiaryRelSister:      "Saudara perempuan",
		BeneficiaryRelSon:         "Anak laki-laki",
		BeneficiaryRelDaughter:    "Anak perempuan",
		BeneficiaryRelNephew:      "Keponakan laki-laki",
		BeneficiaryRelNiece:       "Keponakan perempuan",
		BeneficiaryRelFather:      "Ayah",
		BeneficiaryRelMother:      "Ibu",
		BeneficiaryRelUncle:       "Paman",
		BeneficiaryRelAunt:        "Bibi",
		BeneficiaryRelCousin:      "Sepupu",
		BeneficiaryRelFIL:         "Bapak mertua",
		BeneficiaryRelMIL:         "Ibu mertua",
		BeneficiaryRelBIL:         "Ipar laki-laki",
		BeneficiaryRelSIL:         "Ipar perempuan",
		BeneficiaryRelGrandFather: "Kakek",
		BeneficiaryRelGrandMother: "Nenek",
		BeneficiaryRelHusband:     "Suami",
		BeneficiaryRelWife:        "Istri",
		BeneficiaryRelFriend:      "Teman",
	}

	return labels
}

func GetBeneficiaryRelationshipLabel(relationshipCode string) (string, error) {
	labels := GetAllBeneficiaryRelationshipLabel()
	if value, ok := labels[relationshipCode]; !ok {
		return "", fmt.Errorf("beneficiary relationship code [%s] not found", relationshipCode)
	} else {
		return value, nil
	}
}

// ---------------- SOURCE OF FUNDS

const SourceFundBusiness = "BUSINESS"
const SourceFundSalary = "SALARY"
const SourceFundSavings = "SAVINGS"
const SourceFundGift = "GIFT"

func GetAllSourceFundLabel() map[string]string {
	labels := map[string]string{
		SourceFundBusiness: "Bisnis",
		SourceFundSalary:   "Gaji",
		SourceFundSavings:  "Tabungan",
		SourceFundGift:     "Hadiah",
	}

	return labels
}

func GetSourceFundLabel(sourceFundCode string) (string, error) {
	labels := GetAllSourceFundLabel()
	if value, ok := labels[sourceFundCode]; !ok {
		return "", fmt.Errorf("source of funds code [%s] not found", sourceFundCode)
	} else {
		return value, nil
	}
}

// ---------------- REMITTANCE PURPOSES

const RemittanceAdsExpenses = "ADVERTISING_EXPENSES"
const RemittanceAdvisoryFees = "ADVISORY_FEES"
const RemittanceBusinessInsurance = "BUSINESS_INSURANCE"
const RemittanceComputerServices = "COMPUTER_SERVICES"
const RemittanceConstructionExpenses = "CONSTRUCTION_EXPENSES"
const RemittanceDeliveryFees = "DELIVERY_FEES"
const RemittanceEducation = "EDUCATION"
const RemittanceExportedGoods = "EXPORTED_GOODS"
const RemittanceFamilySupport = "FAMILY_SUPPORT"
const RemittanceFundInvestment = "FUND_INVESTMENT"
const RemittanceGiftDonation = "GIFT_AND_DONATION"
const RemittanceHotelAccommodation = "HOTEL_ACCOMMODATION"
const RemittanceInfluencerPayment = "INFLUENCER_PAYMENT"
const RemittanceInsuranceClaims = "INSURANCE_CLAIMS"
const RemittanceLiberalized = "LIBERALIZED_REMITTANCE"
const RemittanceLoanPayment = "LOAN_PAYMENT"
const RemittanceMaintenanceExpenses = "MAINTENANCE_EXPENSES"
const RemittanceMedicalTreatment = "MEDICAL_TREATMENT"
const RemittanceOfficeExpenses = "OFFICE_EXPENSES"
const RemittanceOtherFees = "OTHER_FEES"
const RemittancePersonalTransfer = "PERSONAL_TRANSFER"
const RemittancePropertyPurchase = "PROPERTY_PURCHASE"
const RemittancePropertyRental = "PROPERTY_RENTAL"
const RemittanceRewardPayment = "REWARD_PAYMENT"
const RemittanceRoyaltyFees = "ROYALTY_FEES"
const RemittanceSalaryPayment = "SALARY_PAYMENT"
const RemittanceServiceCharges = "SERVICE_CHARGES"
const RemittanceSharesInvestment = "SHARES_INVESTMENT"
const RemittanceSmallValue = "SMALL_VALUE_REMITTANCE"
const RemittanceTaxPayment = "TAX_PAYMENT"
const RemittanceTransportationFees = "TRANSPORTATION_FEES"
const RemittanceTravel = "TRAVEL"
const RemittanceUtilityBills = "UTILITY_BILLS"

func GetAllRemittanceLabel() map[string]string {
	labels := map[string]string{
		RemittanceAdsExpenses:          "Beban Iklan",
		RemittanceAdvisoryFees:         "Biaya Konsultan",
		RemittanceBusinessInsurance:    "Asuransi Bisnis",
		RemittanceComputerServices:     "Layanan Komputer",
		RemittanceConstructionExpenses: "Beban Konstruksi",
		RemittanceDeliveryFees:         "Biaya Pengiriman",
		RemittanceEducation:            "Pendidikan",
		RemittanceExportedGoods:        "Pembayaran Barang Ekspor",
		RemittanceFamilySupport:        "Bantuan Keluarga",
		RemittanceFundInvestment:       "Investasi Dana",
		RemittanceGiftDonation:         "Hadiah Dan Donasi",
		RemittanceHotelAccommodation:   "Akomodasi Penginapan",
		RemittanceInfluencerPayment:    "Pembayaran Influencer",
		RemittanceInsuranceClaims:      "Klaim Asuransi",
		RemittanceLiberalized:          "Remitansi Liberal",
		RemittanceLoanPayment:          "Pembayaran Pinjaman",
		RemittanceMaintenanceExpenses:  "Beban Pemeliharaan",
		RemittanceMedicalTreatment:     "Perawatan Medis",
		RemittanceOfficeExpenses:       "Beban Kantor",
		RemittanceOtherFees:            "Biaya Komitmen Dan Garansi",
		RemittancePersonalTransfer:     "Pengiriman Pribadi",
		RemittancePropertyPurchase:     "Pembelian Properti",
		RemittancePropertyRental:       "Pembayaran Sewa Properti",
		RemittanceRewardPayment:        "Pembayaran Hadiah",
		RemittanceRoyaltyFees:          "Biaya Royalti Dan Hak Cipta",
		RemittanceSalaryPayment:        "Pembayaran Gaji",
		RemittanceServiceCharges:       "Biaya Layanan",
		RemittanceSharesInvestment:     "Investasi Saham",
		RemittanceSmallValue:           "Remitansi Nilai Kecil",
		RemittanceTaxPayment:           "Pembayaran Pajak",
		RemittanceTransportationFees:   "Biaya Transportasi",
		RemittanceTravel:               "Perjalanan",
		RemittanceUtilityBills:         "Tagihan Utilitas",
	}

	return labels
}

func GetRemittanceLabel(remittancePurposes string) (string, error) {
	labels := GetAllRemittanceLabel()
	if value, ok := labels[remittancePurposes]; !ok {
		return "", fmt.Errorf("remittance purposes code [%s] not found", remittancePurposes)
	} else {
		return value, nil
	}
}
