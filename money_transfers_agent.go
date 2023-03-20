package flip

import (
	"fmt"

	"github.com/fari-99/go-flip/constants"
	"github.com/fari-99/go-flip/models"
	"github.com/fari-99/go-flip/requests"
)

func (base *BaseFlip) CreateDisbursementAgent(formData models.CreateDisbursementAgentRequest) (disbursement *models.CreateDisbursementAgentResponse, idempotencyKey string, err error) {
	idempotencyKey = base.GetIdempotencyKey()
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductAgentMoneyTransfer, "/agent-disbursements")
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

	var outputModel models.CreateDisbursementAgentResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, idempotencyKey, err
	}

	return &outputModel, idempotencyKey, nil
}

func (base *BaseFlip) GetDisbursementAgentByID(transactionID int64) (disbursementModel *models.DisbursementModel, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductMoneyTransfer, fmt.Sprintf("/agent-disbursements/%d", transactionID))
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

	var outputModel models.DisbursementModel
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) GetDisbursementAgentList(params models.GetDisbursementAgentListRequest) (disbursementModel *models.GetDisbursementAgentListResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductMoneyTransfer, "/agent-disbursements")
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

	var outputModel models.GetDisbursementAgentListResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}
