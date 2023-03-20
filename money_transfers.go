package flip

import (
	"fmt"

	"github.com/spf13/cast"

	"github.com/fari-99/flip/constants"
	"github.com/fari-99/flip/models"
	"github.com/fari-99/flip/requests"
)

func (base *BaseFlip) CreateDisbursement(formData models.CreateDisbursementRequest) (responseModel *models.CreateDisbursementResponse, idempotencyKey string, err error) {
	idempotencyKey = base.GetIdempotencyKey()
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV3, constants.ProductMoneyTransfer, "/disbursement")
	if err != nil {
		return nil, idempotencyKey, err
	}

	baseRequest.SetIdempotencyKey(idempotencyKey)
	baseRequest.SetFormData(formData)
	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeFormUrlEncoded,
	})

	err = baseRequest.PostRequest()
	if err != nil {
		return nil, idempotencyKey, err
	}

	generalError, specialError := baseRequest.GetResponseError()
	if generalError != nil {
		return nil, idempotencyKey, fmt.Errorf(generalError.Message)
	}

	if specialError != nil {
		base.errorDetails = specialError.Errors
		return nil, idempotencyKey, fmt.Errorf(specialError.CodeLabel)
	}

	var outputModel models.CreateDisbursementResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, idempotencyKey, err
	}

	return &outputModel, idempotencyKey, nil
}

func (base *BaseFlip) GetAllDisbursement(params models.GetAllDisbursementRequest) (disbursements *models.GetAllDisbursementResponse, idempotencyKey string, err error) {
	idempotencyKey = base.GetIdempotencyKey()
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV3, constants.ProductMoneyTransfer, "/disbursement")
	if err != nil {
		return nil, idempotencyKey, err
	}

	baseRequest.SetIdempotencyKey(idempotencyKey)

	baseRequest.SetQueryParams(params)
	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeFormUrlEncoded,
	})

	err = baseRequest.GetRequest()
	if err != nil {
		return nil, idempotencyKey, err
	}

	generalError, specialError := baseRequest.GetResponseError()
	if generalError != nil {
		return nil, idempotencyKey, fmt.Errorf(generalError.Message)
	}

	if specialError != nil {
		base.errorDetails = specialError.Errors
		return nil, idempotencyKey, fmt.Errorf(specialError.CodeLabel)
	}

	var outputModel models.GetAllDisbursementResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, idempotencyKey, err
	}

	return &outputModel, idempotencyKey, nil
}

func (base *BaseFlip) GetDisbursementByIdemKey(idempotencyKey string) (disbursementModel *models.DisbursementModel, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV3, constants.ProductMoneyTransfer, "/get-disbursement")
	if err != nil {
		return nil, err
	}

	baseRequest.SetQueryParams(map[string]string{
		"idempotency-key": idempotencyKey,
	})

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

	var outputModel models.DisbursementModel
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) GetDisbursementById(disbursementID int64) (disbursementModel *models.DisbursementModel, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV3, constants.ProductMoneyTransfer, "/get-disbursement")
	if err != nil {
		return nil, err
	}

	baseRequest.SetQueryParams(map[string]string{
		"id": cast.ToString(disbursementID),
	})

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

	var outputModel models.DisbursementModel
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}
