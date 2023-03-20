package requests

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/spf13/cast"

	"github.com/fari-99/flip/constants"
)

type RequestFlip struct {
	isDev       bool
	url         string
	productName string

	defaultError *constants.ErrorResponse
	specialError *constants.ErrorClientResponse

	response      *resty.Response
	body          interface{}
	params        interface{}
	formData      interface{}
	multipartData interface{}
	multipartFile map[string]UploadFile
	headers       Headers

	idempotencyKey string
	requestID      string

	Error error
}

type UploadFile struct {
	ContentType string
	FileName    string
	File        io.Reader
	FileByte    []byte
}

type Headers map[string]string

func NewRequestFlip(appVersion, productName, path string) (base *RequestFlip, err error) {
	newRequest := &RequestFlip{
		productName: productName,
		isDev:       checkDev(),
		headers: Headers{
			constants.HeaderAuthorization: base.getAuthentication(),
		},
	}

	err = newRequest.getUrl(appVersion, path)
	if err != nil {
		return nil, err
	}

	return newRequest, nil
}

func (base *RequestFlip) SetIdempotencyKey(idempotencyKey string) *RequestFlip {
	base.headers[constants.HeaderIdempotencyKey] = idempotencyKey
	base.idempotencyKey = idempotencyKey
	return base
}

func (base *RequestFlip) SetRequestID(uuidString string) *RequestFlip {
	requestID, err := base.generateRequestID(uuidString)
	if err != nil {
		base.Error = err
		return base
	}

	base.headers[constants.HeaderRequestID] = requestID
	base.requestID = requestID
	return base
}

func (base *RequestFlip) GetRequestID() string {
	return base.requestID
}

func (base *RequestFlip) GetResponseError() (defaultError *constants.ErrorResponse, specialError *constants.ErrorClientResponse) {
	return base.defaultError, base.specialError
}

func (base *RequestFlip) GetResponse(output any) error {
	responseBody := base.response.Body()
	//log.Printf(string(responseBody))

	err := json.Unmarshal(responseBody, output)
	if err != nil {
		return err
	}

	return nil
}

func (base *RequestFlip) SetBodyData(bodyData any) *RequestFlip {
	base.body = bodyData
	return base
}

func (base *RequestFlip) SetQueryParams(params any) *RequestFlip {
	base.params = params
	return base
}

func (base *RequestFlip) SetFormData(formData any) *RequestFlip {
	base.formData = formData
	return base
}

func (base *RequestFlip) SetMultipartData(multipartData any) *RequestFlip {
	base.multipartData = multipartData
	return base
}

func (base *RequestFlip) SetMultipartFile(updateFileData UploadFile, paramName string) *RequestFlip {
	multipartFile := base.multipartFile
	if multipartFile == nil {
		multipartFile = make(map[string]UploadFile)
	}

	multipartFile[paramName] = updateFileData
	base.multipartFile = multipartFile
	return base
}

func (base *RequestFlip) SetHeaders(headers Headers) *RequestFlip {
	baseHeaders := base.headers
	for headerName, headerValue := range headers {
		baseHeaders[headerName] = headerValue
	}

	base.headers = baseHeaders
	return base
}

func (base *RequestFlip) GetRequest() error {
	headerData := base.headers

	client := resty.New()
	if base.isDev {
		client.SetDebug(true)
		client.EnableTrace()
	}

	clientRequest := client.R().SetHeaders(headerData)

	if base.params != nil {
		params, err := base.getQueryParams(base.params)
		if err != nil {
			return err
		}
		clientRequest.SetQueryParams(params)
	}

	if base.formData != nil {
		formData, err := base.getQueryParams(base.formData)
		if err != nil {
			return err
		}

		clientRequest.SetFormData(formData)
	}

	if base.multipartData != nil {
		multipartData, err := base.getQueryParams(base.multipartData)
		if err != nil {
			return err
		}

		clientRequest.SetMultipartFormData(multipartData)
	}

	if base.multipartFile != nil {
		for paramName, fileData := range base.multipartFile {
			clientRequest.SetFileReader(paramName, fileData.FileName, fileData.File)
		}
	}

	if base.body != nil {
		clientRequest.SetBody(base.body)
	}

	response, err := clientRequest.Get(base.url)
	if err != nil {
		return err
	}

	if !response.IsSuccess() {
		base.getErrorResponse(response, base.productName)
	}

	base.response = response
	return nil
}

func (base *RequestFlip) PostRequest() error {
	headerData := base.headers
	headerData[constants.HeaderTimestamp] = base.getTimestampHeader(time.Now(), "")

	client := resty.New()
	if base.isDev {
		client.SetDebug(true)
		client.EnableTrace()
	}

	clientRequest := client.R().SetHeaders(headerData)

	if base.params != nil {
		params, err := base.getQueryParams(base.params)
		if err != nil {
			return err
		}
		clientRequest.SetQueryParams(params)
	}

	if base.formData != nil {
		formData, err := base.getQueryParams(base.formData)
		if err != nil {
			return err
		}

		clientRequest.SetFormData(formData)
	}

	if base.multipartData != nil {
		multipartData, err := base.getQueryParams(base.multipartData)
		if err != nil {
			return err
		}

		clientRequest.SetMultipartFormData(multipartData)
	}

	if base.multipartFile != nil {
		for paramName, fileData := range base.multipartFile {
			clientRequest.SetMultipartField(paramName, fileData.FileName, fileData.ContentType, fileData.File)
		}
	}

	if base.body != nil {
		clientRequest.SetBody(base.body)
	}

	response, err := clientRequest.Post(base.url)
	if err != nil {
		return err
	}

	if !response.IsSuccess() {
		base.getErrorResponse(response, base.productName)
	}

	base.response = response
	return nil
}

func (base *RequestFlip) PutRequest() error {
	headerData := base.headers
	headerData[constants.HeaderTimestamp] = base.getTimestampHeader(time.Now(), "")

	client := resty.New()
	if base.isDev {
		client.SetDebug(true)
		client.EnableTrace()
	}

	clientRequest := client.R().SetHeaders(headerData)

	if base.params != nil {
		params, err := base.getQueryParams(base.params)
		if err != nil {
			return err
		}
		clientRequest.SetQueryParams(params)
	}

	if base.formData != nil {
		formData, err := base.getQueryParams(base.formData)
		if err != nil {
			return err
		}

		clientRequest.SetFormData(formData)
	}

	if base.multipartData != nil {
		multipartData, err := base.getQueryParams(base.multipartData)
		if err != nil {
			return err
		}

		clientRequest.SetMultipartFormData(multipartData)
	}

	if base.multipartFile != nil {
		for paramName, fileData := range base.multipartFile {
			clientRequest.SetMultipartField(paramName, fileData.FileName, fileData.ContentType, fileData.File)
		}
	}

	if base.body != nil {
		clientRequest.SetBody(base.body)
	}

	response, err := clientRequest.Put(base.url)
	if err != nil {
		return err
	}

	if !response.IsSuccess() {
		base.getErrorResponse(response, base.productName)
	}

	base.response = response
	return nil
}

func (base *RequestFlip) getAuthentication() string {
	flipSecretKey := os.Getenv("FLIP_SECRET_TOKEN")
	if flipSecretKey == "" {
		panic("please set flip secret key [FLIP_SECRET_TOKEN] to ENV")
	}

	flipSecretKeyExtend := flipSecretKey + ":"
	encodeSecretKey := base64.StdEncoding.EncodeToString([]byte(flipSecretKeyExtend))

	return fmt.Sprintf("Basic " + encodeSecretKey)
}

func (base *RequestFlip) generateRequestID(uuidString string) (string, error) {
	if uuidString == "" {
		uuidString = uuid.New().String()
	} else {
		_, err := uuid.Parse(uuidString)
		if err != nil {
			return "", err
		}
	}

	requestID := "bigflip-" + uuidString
	return requestID, nil
}

func (base *RequestFlip) getErrorResponse(response *resty.Response, productsName string) *RequestFlip {
	statusCode := response.StatusCode()

	if statusCode != 422 {
		var errorModel constants.ErrorResponse
		_ = json.Unmarshal(response.Body(), &errorModel)

		if errorModel.Message != "" {
			base.defaultError = &errorModel
			return base
		}
	}

	var errorClientModel constants.ErrorClientResponse
	_ = json.Unmarshal(response.Body(), &errorClientModel)

	if errorClientModel.Code != "" {
		message, _ := constants.GetGeneralErrorLabel(errorClientModel.Code)
		errorClientModel.CodeLabel = message
	}

	var errorCodeLabel map[string]string
	switch productsName {
	case constants.ProductMoneyTransfer, constants.ProductSpecialMoneyTransfer, constants.ProductAgentMoneyTransfer:
		errorCodeLabel = constants.GetAllTransferError()
	case constants.ProductAgentVerification:
		errorCodeLabel = constants.GetAllAgentVerificationErrorLabel()
	case constants.ProductAcceptPayment:
		errorCodeLabel = constants.GetAllAcceptPaymentErrorLabel()
	case constants.ProductInternationalTransfer:
		errorCodeLabel = constants.GetAllInternationalTransferErrorLabel()
	}

	errorDetails := errorClientModel.Errors
	for idxError, errorDetail := range errorDetails {
		errorClientModel.Errors[idxError].CodeLabel = errorCodeLabel[cast.ToString(errorDetail.Code)]
	}

	base.specialError = &errorClientModel
	return base
}

func checkDev() bool {
	flipEnvironment := os.Getenv("FLIP_ENVIRONMENT")
	return flipEnvironment == "" || flipEnvironment == constants.ApiDev
}

func (base *RequestFlip) getUrl(appVersion string, path string) error {
	isDev := base.isDev

	var urlPath string
	switch appVersion {
	case constants.ApiProdV2:
		if !isDev {
			urlPath = constants.ApiUrlProdV2
		} else {
			urlPath = constants.ApiUrlDev + "/v2"
		}
	case constants.ApiProdV3:
		if !isDev {
			urlPath = constants.ApiUrlProdV3
		} else {
			urlPath = constants.ApiUrlDev + "/v2"
		}
	case constants.ApiKYC:
		if !isDev {
			urlPath = constants.ApiUrlKyc
		} else {
			urlPath = constants.ApiUrlKycDev
		}
	default:
		return fmt.Errorf("api version [%s] is not set, please write issue to add here", appVersion)
	}

	base.url = urlPath + path
	return nil
}

func (base *RequestFlip) getQueryParams(params any) (queryParams map[string]string, err error) {
	paramMarshal, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	var paramModel map[string]any
	err = json.Unmarshal(paramMarshal, &paramModel)
	if err != nil {
		return nil, err
	}

	queryParams = make(map[string]string)
	for paramIdx, paramData := range paramModel {
		queryParams[paramIdx] = cast.ToString(paramData)
	}

	return queryParams, nil
}

func (base *RequestFlip) getTimestampHeader(currentTime time.Time, targetLocation string) (formatDate string) {
	var targetLocationData *time.Location
	if targetLocation == "" {
		targetLocationData = time.UTC
	} else {
		targetLoc, err := time.LoadLocation(targetLocation)
		if err != nil {
			panic(err)
		}

		targetLocationData = targetLoc
	}

	serverTime := currentTime.In(targetLocationData)
	return serverTime.Format(constants.TimeFormatHeader)
}
