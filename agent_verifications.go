package flip

import (
	"fmt"

	"github.com/fari-99/flip/constants"
	"github.com/fari-99/flip/models"
	"github.com/fari-99/flip/requests"
)

func (base *BaseFlip) CreateAgentIdentity(formData models.CreateAgentIdentityRequest) (agentIdentity *models.CreateAgentIdentityResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductAgentVerification, "/agents")
	if err != nil {
		return nil, err
	}

	baseRequest.SetFormData(formData)
	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeFormUrlEncoded,
	})

	err = baseRequest.PostRequest()
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

	var outputModel models.CreateAgentIdentityResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) UpdateAgentIdentity(agentID int64, formData models.UpdateAgentIdentityRequest) (identityModel *models.UpdateAgentIdentityResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductAgentVerification, fmt.Sprintf("/agents/%d", agentID))
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

	var outputModel models.UpdateAgentIdentityResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) GetAgentIdentity(agentID int64) (agentIdentity *models.GetAgentIdentityResponse, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiProdV2, constants.ProductAgentVerification, fmt.Sprintf("/agents/%d", agentID))
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

	var outputModel models.GetAgentIdentityResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, err
	}

	return &outputModel, nil
}

func (base *BaseFlip) UploadAgentIdentityImage(agentID int64, uuid string, multipartData models.UploadAgentIdentityRequest, multipartFile models.UploadAgentIdentityFileRequest) (imageFlip *models.UploadAgentIdentityResponse, requestID string, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiKYC, constants.ProductAgentVerification, fmt.Sprintf("/users/%d/identities", agentID))
	if err != nil {
		return nil, requestID, err
	}

	baseRequest.SetRequestID(uuid)
	baseRequest.SetMultipartData(multipartData)
	baseRequest.SetMultipartFile(multipartFile.Image, "image")
	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeMultipart,
	})

	requestID = baseRequest.GetRequestID()
	err = baseRequest.PutRequest()
	if err != nil {
		return nil, requestID, err
	}

	generalError, specialError := baseRequest.GetResponseError()
	if generalError != nil {
		return nil, requestID, fmt.Errorf(generalError.Message)
	}

	if specialError != nil {
		base.errorDetails = specialError.Errors
		return nil, requestID, fmt.Errorf(specialError.CodeLabel)
	}

	var outputModel models.UploadAgentIdentityResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, requestID, err
	}

	return &outputModel, requestID, nil
}

func (base *BaseFlip) UploadSupportingDocuments(uuid string, multipartData models.UploadSupportingDocumentsRequest, multipartFile models.UploadSupportingDocumentFilesRequest) (documents *models.UploadSupportingDocumentResponse, requestID string, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiKYC, constants.ProductAgentVerification, "/documents")
	if err != nil {
		return nil, requestID, err
	}

	baseRequest.SetRequestID(uuid)
	baseRequest.SetMultipartData(multipartData)

	for paramName, fileData := range multipartFile {
		baseRequest.SetMultipartFile(fileData, paramName)
	}

	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeMultipart,
	})

	requestID = baseRequest.GetRequestID()
	err = baseRequest.PostRequest()
	if err != nil {
		return nil, requestID, err
	}

	generalError, specialError := baseRequest.GetResponseError()
	if generalError != nil {
		return nil, requestID, fmt.Errorf(generalError.Message)
	}

	if specialError != nil {
		base.errorDetails = specialError.Errors
		return nil, requestID, fmt.Errorf(specialError.CodeLabel)
	}

	var outputModel models.UploadSupportingDocumentResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, requestID, err
	}

	return &outputModel, requestID, nil
}

func (base *BaseFlip) KYCSubmissions(agentID int64, uuid string, formData models.KycSubmissionRequest) (message *models.KycSubmissionResponse, requestID string, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiKYC, constants.ProductAgentVerification, fmt.Sprintf("/users/%d/submit", agentID))
	if err != nil {
		return nil, requestID, err
	}

	baseRequest.SetRequestID(uuid)
	baseRequest.SetFormData(formData)
	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeMultipart,
	})

	requestID = baseRequest.GetRequestID()
	err = baseRequest.PutRequest()
	if err != nil {
		return nil, requestID, err
	}

	generalError, specialError := baseRequest.GetResponseError()
	if generalError != nil {
		return nil, requestID, fmt.Errorf(generalError.Message)
	}

	if specialError != nil {
		base.errorDetails = specialError.Errors
		return nil, requestID, fmt.Errorf(specialError.CodeLabel)
	}

	var outputModel models.KycSubmissionResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, requestID, err
	}

	return &outputModel, requestID, nil
}

func (base *BaseFlip) RepairData(agentID int64, uuid string, formData models.RepairDataRequest) (repairedData *models.RepairDataResponse, requestID string, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiKYC, constants.ProductAgentVerification, fmt.Sprintf("/users/%d/repair", agentID))
	if err != nil {
		return nil, requestID, err
	}

	baseRequest.SetRequestID(uuid)
	baseRequest.SetFormData(formData)
	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeMultipart,
	})

	requestID = baseRequest.GetRequestID()
	err = baseRequest.PutRequest()
	if err != nil {
		return nil, requestID, err
	}

	generalError, specialError := baseRequest.GetResponseError()
	if generalError != nil {
		return nil, requestID, fmt.Errorf(generalError.Message)
	}

	if specialError != nil {
		base.errorDetails = specialError.Errors
		return nil, requestID, fmt.Errorf(specialError.CodeLabel)
	}

	var outputModel models.RepairDataResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, requestID, err
	}

	return &outputModel, requestID, nil
}

func (base *BaseFlip) RepairIdentityImage(agentID int64, uuid string, multipartData models.RepairIdentityImageRequest, multipartFile models.RepairIdentityImageFileRequest) (repairedData *models.RepairDataResponse, requestID string, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiKYC, constants.ProductAgentVerification, fmt.Sprintf("/users/%d/repairPhoto", agentID))
	if err != nil {
		return nil, requestID, err
	}

	baseRequest.SetRequestID(uuid)
	baseRequest.SetMultipartData(multipartData)
	baseRequest.SetMultipartFile(multipartFile.Image, "image")
	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeMultipart,
	})

	requestID = baseRequest.GetRequestID()
	err = baseRequest.PutRequest()
	if err != nil {
		return nil, requestID, err
	}

	generalError, specialError := baseRequest.GetResponseError()
	if generalError != nil {
		return nil, requestID, fmt.Errorf(generalError.Message)
	}

	if specialError != nil {
		base.errorDetails = specialError.Errors
		return nil, requestID, fmt.Errorf(specialError.CodeLabel)
	}

	var outputModel models.RepairDataResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, requestID, err
	}

	return &outputModel, requestID, nil
}

func (base *BaseFlip) RepairIdentitySelfieImage(agentID int64, uuid string, multipartData models.RepairIdentityImageRequest, multipartFile models.RepairIdentityImageFileRequest) (repairedData *models.RepairDataResponse, requestID string, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiKYC, constants.ProductAgentVerification, fmt.Sprintf("/users/%d/repairSelfie", agentID))
	if err != nil {
		return nil, requestID, err
	}

	baseRequest.SetRequestID(uuid)
	baseRequest.SetMultipartData(multipartData)
	baseRequest.SetMultipartFile(multipartFile.Image, "image")
	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeMultipart,
	})

	requestID = baseRequest.GetRequestID()
	err = baseRequest.PutRequest()
	if err != nil {
		return nil, requestID, err
	}

	generalError, specialError := baseRequest.GetResponseError()
	if generalError != nil {
		return nil, requestID, fmt.Errorf(generalError.Message)
	}

	if specialError != nil {
		base.errorDetails = specialError.Errors
		return nil, requestID, fmt.Errorf(specialError.CodeLabel)
	}

	var outputModel models.RepairDataResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, requestID, err
	}

	return &outputModel, requestID, nil
}

func (base *BaseFlip) GetCountryList(uuid string, params models.LocationKycRequest) (locationData *models.LocationKycResponse, requestID string, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiKYC, constants.ProductAgentVerification, "/countries")
	if err != nil {
		return nil, "", err
	}

	baseRequest.SetRequestID(uuid)
	baseRequest.SetQueryParams(params)
	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeFormUrlEncoded,
	})

	requestID = baseRequest.GetRequestID()
	err = baseRequest.GetRequest()
	if err != nil {
		return nil, requestID, err
	}

	generalError, specialError := baseRequest.GetResponseError()
	if generalError != nil {
		return nil, requestID, fmt.Errorf(generalError.Message)
	}

	if specialError != nil {
		base.errorDetails = specialError.Errors
		return nil, requestID, fmt.Errorf(specialError.CodeLabel)
	}

	var outputModel models.LocationKycResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, requestID, err
	}

	return &outputModel, requestID, nil
}

func (base *BaseFlip) GetProvinceList(uuid string, params models.LocationKycRequest) (locationData *models.LocationKycResponse, requestID string, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiKYC, constants.ProductAgentVerification, "/provinces")
	if err != nil {
		return nil, "", err
	}

	baseRequest.SetRequestID(uuid)
	baseRequest.SetQueryParams(params)
	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeFormUrlEncoded,
	})

	requestID = baseRequest.GetRequestID()
	err = baseRequest.GetRequest()
	if err != nil {
		return nil, requestID, err
	}

	generalError, specialError := baseRequest.GetResponseError()
	if generalError != nil {
		return nil, requestID, fmt.Errorf(generalError.Message)
	}

	if specialError != nil {
		base.errorDetails = specialError.Errors
		return nil, requestID, fmt.Errorf(specialError.CodeLabel)
	}

	var outputModel models.LocationKycResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, requestID, err
	}

	return &outputModel, requestID, nil
}

func (base *BaseFlip) GetCityList(uuid string, params models.LocationKycRequest) (locationData *models.LocationKycResponse, requestID string, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiKYC, constants.ProductAgentVerification, "/cities")
	if err != nil {
		return nil, "", err
	}

	baseRequest.SetRequestID(uuid)
	baseRequest.SetQueryParams(params)
	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeFormUrlEncoded,
	})

	requestID = baseRequest.GetRequestID()
	err = baseRequest.GetRequest()
	if err != nil {
		return nil, requestID, err
	}

	generalError, specialError := baseRequest.GetResponseError()
	if generalError != nil {
		return nil, requestID, fmt.Errorf(generalError.Message)
	}

	if specialError != nil {
		base.errorDetails = specialError.Errors
		return nil, requestID, fmt.Errorf(specialError.CodeLabel)
	}

	var outputModel models.LocationKycResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, requestID, err
	}

	return &outputModel, requestID, nil
}

func (base *BaseFlip) GetDistrictList(uuid string, params models.LocationKycRequest) (locationData *models.LocationKycResponse, requestID string, err error) {
	baseRequest, err := requests.NewRequestFlip(constants.ApiKYC, constants.ProductAgentVerification, "/districts")
	if err != nil {
		return nil, "", err
	}

	baseRequest.SetRequestID(uuid)
	baseRequest.SetQueryParams(params)
	baseRequest.SetHeaders(requests.Headers{
		constants.HeaderContentType: constants.ContentTypeFormUrlEncoded,
	})

	requestID = baseRequest.GetRequestID()
	err = baseRequest.GetRequest()
	if err != nil {
		return nil, requestID, err
	}

	generalError, specialError := baseRequest.GetResponseError()
	if generalError != nil {
		return nil, requestID, fmt.Errorf(generalError.Message)
	}

	if specialError != nil {
		base.errorDetails = specialError.Errors
		return nil, requestID, fmt.Errorf(specialError.CodeLabel)
	}

	var outputModel models.LocationKycResponse
	err = baseRequest.GetResponse(&outputModel)
	if err != nil {
		return nil, requestID, err
	}

	return &outputModel, requestID, nil
}
