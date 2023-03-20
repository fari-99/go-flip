package flip

import (
	"fmt"

	"github.com/fari-99/flip/constants"
	"github.com/fari-99/flip/models"
	"github.com/fari-99/flip/requests"
)

func (base *BaseFlip) CreateBill(formData models.CreateBillRequest) (bill *models.CreateBillResponse, idempotencyKey string, err error) {
	idempotencyKey = base.GetIdempotencyKey()
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductAcceptPayment, "/pwf/bill")
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

	var outputModel models.CreateBillResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, idempotencyKey, err
	}

	return &outputModel, idempotencyKey, nil
}

func (base *BaseFlip) EditBill(billID int64, formData models.EditBillingRequest) (bill *models.EditBillingResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductAcceptPayment, fmt.Sprintf("/pwf/%d/bill", billID))
	if err != nil {
		return nil, err
	}

	baseRequest.SetFormData(formData)
	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeFormUrlEncoded,
	})

	err = baseRequest.PutRequest()
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

	var outputModel models.EditBillingResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) GetBill(billID int64) (bill *models.EditBillingResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductAcceptPayment, fmt.Sprintf("/pwf/%d/bill", billID))
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

	var outputModel models.EditBillingResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) GetAllBill() (bills *models.GetAllBillingResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductAcceptPayment, "/pwf/bill")
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

	var outputModel models.GetAllBillingResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) GetPayment(billID int64, params models.GetPaymentRequest) (payments *models.GetPaymentResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductAcceptPayment, fmt.Sprintf("/pwf/%d/payment", billID))
	if err != nil {
		return nil, err
	}

	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeFormUrlEncoded,
	})

	baseRequest.SetQueryParams(params)
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

	var outputModel models.GetPaymentResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) GetAllPayment(params models.GetAllPaymentRequest) (payments *models.GetAllPaymentResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductAcceptPayment, "/pwf/payment")
	if err != nil {
		return nil, err
	}

	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeFormUrlEncoded,
	})

	baseRequest.SetQueryParams(params)
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

	var outputModel models.GetAllPaymentResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) ConfirmBillPayment(transactionID int64) (payment *models.ConfirmBillPaymentResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductAcceptPayment, fmt.Sprintf("/pwf/bill-payment/%d/confirm", transactionID))
	if err != nil {
		return nil, err
	}

	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeFormUrlEncoded,
	})

	err = baseRequest.PutRequest()
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

	var outputModel models.ConfirmBillPaymentResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}
