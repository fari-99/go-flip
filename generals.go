package flip

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/fari-99/go-flip/constants"
	"github.com/fari-99/go-flip/models"
	"github.com/fari-99/go-flip/requests"
)

type BaseFlip struct {
	idempotencyKey string
	error          error
	errorDetails   []constants.ErrorDetailModel
}

func NewBaseFlip() *BaseFlip {
	base := &BaseFlip{
		idempotencyKey: uuid.New().String(), // generate new uuid
	}

	return base
}

func (base *BaseFlip) SetIdempotencyKey(idemKey string) *BaseFlip {
	if idemKey == "" {
		panic("you want to set Idempotency Key, but you give empty string")
	}

	base.idempotencyKey = idemKey
	return base
}

func (base *BaseFlip) GetIdempotencyKey() string {
	return base.idempotencyKey
}

func (base *BaseFlip) GetErrorDetails() []constants.ErrorDetailModel {
	return base.errorDetails
}

func (base *BaseFlip) IsMaintenance() (isMaintenance *models.GetMaintenanceStatusResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductGeneral, "/general/maintenance")
	if err != nil {
		return nil, err
	}

	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeFormUrlEncoded,
	})

	err = baseRequest.GetRequest()
	if err != nil {
		return nil, err
	}

	generalError, specialError := baseRequest.GetResponseError()
	if generalError != nil {
		return nil, fmt.Errorf(generalError.Message)
	}

	if specialError != nil {
		base.errorDetails = specialError.Errors
		return nil, fmt.Errorf(specialError.CodeLabel)
	}

	var outputModel models.GetMaintenanceStatusResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) GetBalance() (balanceModel *models.GetBalanceResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductGeneral, "/general/balance")
	if err != nil {
		return nil, err
	}

	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeFormUrlEncoded,
	})

	err = baseRequest.GetRequest()
	if err != nil {
		return nil, err
	}

	generalError, specialError := baseRequest.GetResponseError()
	if generalError != nil {
		return nil, fmt.Errorf(generalError.Message)
	}

	if specialError != nil {
		base.errorDetails = specialError.Errors
		return nil, fmt.Errorf(specialError.CodeLabel)
	}

	var outputModel models.GetBalanceResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) GetBankInfo(params models.GetBankInfoRequest) (bankList *models.GetBankInfoResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductGeneral, "/general/banks")
	if err != nil {
		return nil, err
	}

	baseRequest.SetQueryParams(params)
	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeFormUrlEncoded,
	})

	err = baseRequest.GetRequest()
	if err != nil {
		return nil, err
	}

	generalError, specialError := baseRequest.GetResponseError()
	if generalError != nil {
		return nil, fmt.Errorf(generalError.Message)
	}

	if specialError != nil {
		base.errorDetails = specialError.Errors
		return nil, fmt.Errorf(specialError.CodeLabel)
	}

	var outputModel models.GetBankInfoResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) SendBankAccountInquiry(params models.SendBankAccountInquiryRequest) (inquiry *models.SendBankAccountInquiryResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductGeneral, "/general/banks")
	if err != nil {
		return nil, err
	}

	baseRequest.SetQueryParams(params)
	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeFormUrlEncoded,
	})

	err = baseRequest.GetRequest()
	if err != nil {
		return nil, err
	}

	generalError, specialError := baseRequest.GetResponseError()
	if generalError != nil {
		return nil, fmt.Errorf(generalError.Message)
	}

	if specialError != nil {
		base.errorDetails = specialError.Errors
		return nil, fmt.Errorf(specialError.CodeLabel)
	}

	var outputModel models.SendBankAccountInquiryResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}
