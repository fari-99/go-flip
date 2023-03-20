package flip

import (
	"fmt"

	"github.com/fari-99/go-flip/constants"
	"github.com/fari-99/go-flip/models"
	"github.com/fari-99/go-flip/requests"
)

func (base *BaseFlip) GetExchangeRate(params models.GetExchangeRatesRequest) (exchangeRates []models.GetExchangeRateResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductInternationalTransfer, "/international-disbursement/exchange-rates")
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

	var outputModel []models.GetExchangeRateResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return outputModel, nil
}

func (base *BaseFlip) GetFormData(params models.GetFormDataRequest) (formData *models.GetFormDataResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductInternationalTransfer, "/international-disbursement/form-data")
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

	var outputModel models.GetFormDataResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) CreateInternationalTransferC2CC2B(formData models.CreateInternationalTransferC2CC2BRequest) (internationalTransfer *models.CreateInternationalTransferC2CC2BResponse, idempotencyKey string, err error) {
	idempotencyKey = base.GetIdempotencyKey()
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductInternationalTransfer, "/international-disbursement")
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

	var outputModel models.CreateInternationalTransferC2CC2BResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, idempotencyKey, err
	}

	return &outputModel, idempotencyKey, nil
}

func (base *BaseFlip) CreateInternationalTransferB2CB2B(multipartData models.CreateInternationalTransferB2XRequest, multipartFile models.CreateInternationalTransferB2XFileRequest) (internationalTransfer *models.CreateInternationalTransferB2CB2BResponse, idempotencyKey string, err error) {
	idempotencyKey = base.GetIdempotencyKey()
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductInternationalTransfer, "/international-disbursement/create-with-attachment")
	if err != nil {
		return nil, idempotencyKey, err
	}

	baseRequest.SetIdempotencyKey(idempotencyKey)
	baseRequest.SetMultipartData(multipartData)
	baseRequest.SetMultipartFile(multipartFile.AttachmentData, "attachment_data")
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

	var outputModel models.CreateInternationalTransferB2CB2BResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, idempotencyKey, err
	}

	return &outputModel, idempotencyKey, nil
}

func (base *BaseFlip) GetInternationalTransfer(transactionID int64) (internationalTransfer *models.GetInternationalTransferResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductInternationalTransfer, fmt.Sprintf("/international-disbursement/%d", transactionID))
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

	var outputModel models.GetInternationalTransferResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) GetAllInternationalTransfer(params models.GetAllInternationalTransferRequest) (internationalTransfers *models.GetAllInternationalTransferResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductInternationalTransfer, "/international-disbursement")
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

	var outputModel models.GetAllInternationalTransferResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}
