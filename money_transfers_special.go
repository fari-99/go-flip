package flip

import (
	"fmt"

	"github.com/fari-99/go-flip/constants"
	"github.com/fari-99/go-flip/models"
	"github.com/fari-99/go-flip/requests"
)

func (base *BaseFlip) CreateSpecialDisbursement(formData models.CreateSpecialDisbursementRequest) (responseModel *models.CreateSpecialDisbursementResponse, idempotencyKey string, err error) {
	idempotencyKey = base.GetIdempotencyKey()
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV3, constants.ProductSpecialMoneyTransfer, "/special-disbursement")
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

	var outputModel models.CreateSpecialDisbursementResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, idempotencyKey, err
	}

	return &outputModel, idempotencyKey, nil
}

func (base *BaseFlip) GetDisbursementCityList() (cityList *models.CityListResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductGeneral, "/disbursement/city-list")
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

	var outputModel models.CityListResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) GetDisbursementCountryList() (country *models.CountryListResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductGeneral, "/disbursement/country-list")
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

	var outputModel models.CountryListResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) GetDisbursementCountryCityList() (country *models.CountryCityListResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductGeneral, "/disbursement/city-country-list")
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

	var outputModel models.CountryCityListResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}
