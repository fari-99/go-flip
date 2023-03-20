package models

import (
	"time"

	"github.com/fari-99/go-flip/requests"
)

/*
AgentIdentity

| Attribute               | Description                                                                                                                                                                                                                                        |
|-------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| id                      | Id of the created Agent                                                                                                                                                                                                                            |
| company_id              | Id of the company which the Agent belongs to                                                                                                                                                                                                       |
| name                    | Agent’s full name                                                                                                                                                                                                                                  |
| identity_type           | Type of identity. Possible values are:ktpNational Id Card or KTP in Indonesiapassport                                                                                                                                                              |
| identity_number         | Identity number based on the type chosen                                                                                                                                                                                                           |
| birth_place             | Agent’s birth place                                                                                                                                                                                                                                |
| birth_date              | Agent’s birth date on DD-MM-YYYY format                                                                                                                                                                                                            |
| gender                  | Agent’s gender. Possible values are:malefemale                                                                                                                                                                                                     |
| country_id              | Agent’s country code                                                                                                                                                                                                                               |
| country_name            | Agent’s country name based on the country code                                                                                                                                                                                                     |
| province_id             | Agent’s province code                                                                                                                                                                                                                              |
| provice_name            | Agent’s province name based on the province code                                                                                                                                                                                                   |
| city_id                 | Agent’s city ID in the identity type chosen                                                                                                                                                                                                        |
| city_name               | Agent’s city name based on the city code                                                                                                                                                                                                           |
| district_id             | Agent’s district ID in the identity type chosen                                                                                                                                                                                                    |
| district_name           | Agent’s district name based on the district code                                                                                                                                                                                                   |
| address                 | Agent’s address                                                                                                                                                                                                                                    |
| residence_country_id    | Agent’s current residential address country code                                                                                                                                                                                                   |
| residence_country_name  | Agent’s current residential address country name                                                                                                                                                                                                   |
| residence_province_id   | Agent’s current residential address province code                                                                                                                                                                                                  |
| residence_provice_name  | Agent’s current residential address province name                                                                                                                                                                                                  |
| residence_city_id       | Agent’s current residential address city code                                                                                                                                                                                                      |
| residence_city_name     | Agent’s current residential address city name                                                                                                                                                                                                      |
| residence_district_id   | Agent’s current residential address district code                                                                                                                                                                                                  |
| residence_district_name | Agent’s current residential address district name                                                                                                                                                                                                  |
| residence_address       | Agent’s current residential address                                                                                                                                                                                                                |
| occupation              | Agent’s job. Possible values are:housewifeentrepreneurprivate_employeegovernment_employeefoundation_boardPeople who work at foundation as an employeeindonesian_migrant_workerAlso known as TKIcompanyIf sender is a company, not individualothers |
| phone_number            | Agent’s phone number                                                                                                                                                                                                                               |
| email                   | Agent’s email                                                                                                                                                                                                                                      |
| kyc_status              | Agent’s KYC status. Values can be seen on the KYC status enum.                                                                                                                                                                                     |
*/
type AgentIdentity struct {
	Id                    int    `json:"id"`
	CompanyId             int    `json:"company_id"`
	Name                  string `json:"name"`
	IdentityType          string `json:"identity_type"`
	IdentityNumber        string `json:"identity_number"`
	BirthPlace            string `json:"birth_place"`
	BirthDate             string `json:"birth_date"`
	Gender                string `json:"gender"`
	CountryId             int    `json:"country_id"`
	CountryName           string `json:"country_name"`
	ProvinceId            int    `json:"province_id"`
	ProvinceName          string `json:"province_name"`
	CityId                int    `json:"city_id"`
	CityName              string `json:"city_name"`
	DistrictId            int    `json:"district_id"`
	DistrictName          string `json:"district_name"`
	Address               string `json:"address"`
	ResidenceCountryId    int    `json:"residence_country_id"`
	ResidenceCountryName  string `json:"residence_country_name"`
	ResidenceProvinceId   int    `json:"residence_province_id"`
	ResidenceProvinceName string `json:"residence_province_name"`
	ResidenceCityId       int    `json:"residence_city_id"`
	ResidenceCityName     string `json:"residence_city_name"`
	ResidenceDistrictId   int    `json:"residence_district_id"`
	ResidenceDistrictName string `json:"residence_district_name"`
	ResidenceAddress      string `json:"residence_address"`
	Occupation            string `json:"occupation"`
	PhoneNumber           string `json:"phone_number"`
	Email                 string `json:"email"`
	KycStatus             string `json:"kyc_status"`
}

/*
SupportDocuments
| Attribute  | Description                                   |
|------------|-----------------------------------------------|
| id         | ID of the uploaded image                      |
| type       | User type: 1 (Agent)                          |
| selfie     | Flag if the uploaded file is the selfie image |
| path       | Uploaded file path URL                        |
| created_at | Timestamp of creation of the uploaded file    |
| updated_at | Timestamp of update of the uploaded file      |

Example SupportDocuments:
{
    "id": 1,
    "type": 1,
    "selfie": false,
    "path": "someurl.png",
    "created_at": "2022-01-02T15:38:49.317601182+07:00",
    "updated_at": "2022-01-02T15:38:49.317601182+07:00"
}
*/
type SupportDocuments struct {
	Id        int       `json:"id"`
	Type      int       `json:"type"`
	Selfie    bool      `json:"selfie"`
	Path      string    `json:"path"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

/*
Locations
| Attribute | Description                           |
|-----------|---------------------------------------|
| id        | ID Country/Province/Cities/District   |
| name      | Country/Province/Cities/District Name |

Example Locations:
{
  "countries": [
    {
      "id": 1,
      "name": "Indonesia"
    }
  ],
  "provinces": [
    {
      "id": 1,
      "name": "West Java"
    }
  ],
  "cities": [
    {
      "id": 1,
      "name": "Depok"
    }
  ],
  "districts": [
    {
      "id": 1,
      "name": "Sukmajaya"
    }
  ]
}
*/
type Locations struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// --------------------------- RESPONSE REQUEST
// --------------------------- Create Agent Identity

/*
CreateAgentIdentityRequest
| Attribute             | Description                                                                                                                                                                                                                                                                              |
|-----------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| name                  | string (required)Agent’s full name. Validation:Alphanumeric1-50 characters                                                                                                                                                                                                               |
| identity_type         | string (required)Type of identity. Validation:Accepted values are as follows:ktp (National Id Card or KTP in Indonesia)passport (Passport)                                                                                                                                               |
| identity_number       | string (required)Identity number based on the type chosen. Validation:16 characters for KTP20 characters for Passport                                                                                                                                                                    |
| birth_date            | string (required)Agent’s birth date. Validation:DD-MM-YYYY format                                                                                                                                                                                                                        |
| birth_place           | string (required)Agent’s birth place. Validation:Alphanumeric255 characters                                                                                                                                                                                                              |
| country_id            | integer (required)Agent’s country code based on the identity chosen. Validation:Available values can be retrieved from the country list API here.                                                                                                                                        |
| province_id           | integer (required)Agent’s province code based on the identity chosen. Validation:Available values can be retrieved from the province list API here.                                                                                                                                      |
| city_id               | integer (required)Agent’s city code based on the identity chosen. Validation:Available values can be retrieved from the city list API here.                                                                                                                                              |
| district_id           | integer (required)Agent’s district code based on the identity chosen. Validation:Available values can be retrieved from the district list API here.                                                                                                                                      |
| address               | string (required)Agent’s address based on the identity chosen. Validation:Alphanumeric10-300 charactersSpaces, ., -, /, (, and )                                                                                                                                                         |
| gender                | string (required)Agent’s gender. Validation:Accepted value are as follows:malefemale                                                                                                                                                                                                     |
| occupation            | string (required)Agent’s job. Validation:Accepted value are as follows:housewifeentrepreneurprivate_employeegovernment_employeefoundation_boardPeople who work at foundation as an employeeindonesian_migrant_workerAlso known as TKIcompanyIf sender is a company, not individualothers |
| phone_number          | string (required)Agent’s phone number. It must be the one that has been verified by the company. Validation:If it starts with “0”, length is 8-24 charactersIf it starts with “+”, length is 10-26 characters                                                                            |
| use_identity_address  | integer (optional)A flag to set if the current residential address of the Agent is the same as the identity address. Validation:Accepted value are as follows:0: false1: true                                                                                                            |
| residence_country_id  | integer (optional)Agent’s current residential address country code. Validation:Available values can be retrieved from the country list API here.                                                                                                                                         |
| residence_province_id | integer (optional)Agent’s current residential address province code. Validation:Available values can be retrieved from the province list API here.                                                                                                                                       |
| residence_city_id     | integer (optional)Agent’s current residential address city code. Validation:Available values can be retrieved from the city list API here.                                                                                                                                               |
| residence_district_id | integer (optional)Agent’s current residential address district code. Validation:Available values can be retrieved from the district list API here.                                                                                                                                       |
| residence_address     | string (optional)Agent’s current residential address. Validation:Alphanumeric10-300 charactersSpaces, ., -, /, (, and )                                                                                                                                                                  |
| email                 | string (optional)Agent’s email. Validation:Alphanumeric with email format (xxx@xxx.xxx)                                                                                                                                                                                                  |

Example CreateAgentIdentityRequest:
{
    "name": "John Doe",
    "identity_type": "ktp",
    "identity_number": "1234567890123456",
    "birth_date": "31-10-1990",
    "birth_place": "Kota Depok",
    "country_id": "1",
    "province_id": "12",
    "city_id": "184",
    "district_id": "1234",
    "address": "Jalan Margonda Raya",
    "gender": "male",
    "occupation": "entrepreneur",
    "phone_number": "0812345678901",
    "use_identity_address": "1",
    "residence_country_id": 1,
    "residence_province_id": 1,
    "residence_city_id": 1,
    "residence_district_id": 1,
    "residence_address": "Jalan Kematian",
    "email": "example@mail.com"
}
*/
type CreateAgentIdentityRequest struct {
	Id                  int64  `json:"id,omitempty"`
	Name                string `json:"name"`
	IdentityType        string `json:"identity_type"`
	IdentityNumber      string `json:"identity_number"`
	BirthDate           string `json:"birth_date"`
	BirthPlace          string `json:"birth_place"`
	CountryId           string `json:"country_id"`
	ProvinceId          string `json:"province_id"`
	CityId              string `json:"city_id"`
	DistrictId          string `json:"district_id"`
	Address             string `json:"address"`
	Gender              string `json:"gender"`
	Occupation          string `json:"occupation"`
	PhoneNumber         string `json:"phone_number"`
	UseIdentityAddress  string `json:"use_identity_address"`
	ResidenceCountryId  int    `json:"residence_country_id"`
	ResidenceProvinceId int    `json:"residence_province_id"`
	ResidenceCityId     int    `json:"residence_city_id"`
	ResidenceDistrictId int    `json:"residence_district_id"`
	ResidenceAddress    string `json:"residence_address"`
	Email               string `json:"email"`
}

/*
CreateAgentIdentityResponse
The response the same as AgentIdentity

Example CreateAgentIdentityResponse:
{
   "id": 1,
   "company_id": 1,
   "name": "John Doe",
   "identity_type": "ktp",
   "identity_number": "1234567890123456",
   "birth_place": "Kota Depok",
   "birth_date": "31-10-1990",
   "gender": "male",
   "country_id": 1,
   "country_name": "Indonesia",
   "province_id": 12,
   "province_name": "Jawa Barat",
   "city_id": 184,
   "city_name": "Kota Depok",
   "district_id": 1234,
   "district_name": "Sukmajaya",
   "address": "Jalan Margonda Raya",
   "residence_country_id": 1,
   "residence_country_name": "Indonesia",
   "residence_province_id": 12,
   "residence_province_name": "Jawa Barat",
   "residence_city_id": 128,
   "residence_city_name": "Kota Depok",
   "residence_district_id": 1234,
   "residence_district_name": "Sukmajaya",
   "residence_address": "Jalan Margonda Raya",
   "occupation": "entrepreneur",
   "phone_number": "+628123456789",
   "email": "example@mail.com",
   "kyc_status":"BASIC_DATA"
}
*/
type CreateAgentIdentityResponse AgentIdentity

// --------------------------- Update Agent Identity

// UpdateAgentIdentityRequest
//// the params the same as CreateAgentIdentityRequest
type UpdateAgentIdentityRequest CreateAgentIdentityRequest

// UpdateAgentIdentityResponse
// the response the same as AgentIdentity
type UpdateAgentIdentityResponse AgentIdentity

// --------------------------- Get Agent Identity

/*
GetAgentIdentityResponse
| Attribute            | Description                                                                                                                                             |
|------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------|
| id                   | Id of the created Agent                                                                                                                                 |
| company_id           | Id of the company which the Agent belongs to                                                                                                            |
| name                 | Agent’s full name                                                                                                                                       |
| identity_type        | Type of identity. Possible values are:ktpNational Id Card or KTP in Indonesiapassport                                                                   |
| identity_number      | Identity number based on the type chosen                                                                                                                |
| phone_number         | Agent’s phone number                                                                                                                                    |
| email                | Agent’s email                                                                                                                                           |
| kyc_status           | Agent’s KYC status. Values can be seen on the KYC status enum.                                                                                          |
| rejected_reason_code | Rejected reason code for Agent’s KYC. This value will be:null - if there’s no valueinteger from Agent’s KYC Rejection Reasons Code - if there’s a value |
| rejected_reason      | Rejected reason for Agent’s KYC. Full list of rejected reasons can be seen in the Agent’s KYC Rejection Reasons list.                                   |
| status               | Enum of agent status:0 - Deleted10 - Active20 - Blacklisted30 - Blocked                                                                                 |
| created_at           | Created timestamp                                                                                                                                       |

Example GetAgentIdentityResponse:
*/
type GetAgentIdentityResponse struct {
	Id                 int         `json:"id"`
	CompanyId          int         `json:"company_id"`
	Name               string      `json:"name"`
	IdentityType       string      `json:"identity_type"`
	IdentityNumber     string      `json:"identity_number"`
	PhoneNumber        string      `json:"phone_number"`
	Email              string      `json:"email"`
	KycStatus          string      `json:"kyc_status"`
	RejectedReasonCode interface{} `json:"rejected_reason_code"`
	RejectedReason     string      `json:"rejected_reason"`
	Status             string      `json:"status"`
	CreatedAt          string      `json:"created_at"`
}

// --------------------------- Upload Agent Identity Image

/*
UploadAgentIdentityRequest
| Attribute     | Description                                                                                                                                                                                                                                                                               |
|---------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| user_type     | string (required)Type of user. Validation:User type Agent: 1                                                                                                                                                                                                                              |
| image         | file (required)Image file that wants to be uploaded. Validation:Size less than 2MBDimension between 256 (W) x 256 (H) and 4096 (W) x 4096 (H)File type should be png, jpg, or jpegImage should be uploaded right from Camera app (not Gallery or File or Document app)                    |
| identity_type | string (required)Type of identity. Validation:Accepted value are as follows:1: National Id Card or KTP in Indonesia2: Passport                                                                                                                                                            |
| selfie        | string (optional)Flag denotes the uploaded file is the identity image or selfie with identity image. Validation:Accepted value are as follows:0: False (default). Denotes that the uploaded file is identity image.1: True. Denotes that the uploaded file is selfie with identity image. |

Example :
{
  "user_type":"1",
  "image":(binary),
  "identity_type":"1",
  "selfie":"1"
}
*/
type UploadAgentIdentityRequest struct {
	UserType     string `json:"user_type"`
	IdentityType string `json:"identity_type"`
	Selfie       string `json:"selfie,omitempty"`
}

type UploadAgentIdentityFileRequest struct {
	Image requests.UploadFile `json:"image"`
}

/*
UploadAgentIdentityResponse
| Attribute | Description               |
|-----------|---------------------------|
| image_url | URL of the uploaded image |

Example UploadAgentIdentityResponse:
*/
type UploadAgentIdentityResponse struct {
	ImageUrl string `json:"image_url"`
}

// --------------------------- Upload Supporting Documents

/*
UploadSupportingDocumentsRequest
| Attribute               | Description                                                  |
|-------------------------|--------------------------------------------------------------|
| user_id                 | string (required)Agent’s ID. Validation:Numeric              |
| user_type               | string (required)Type of user. Validation:User type Agent: 1 |
| student_card            | file (optional)Student card image file.                      |
| student_card_selfie     | file (optional)Selfie with student card image file.          |
| employee_card           | file (optional)Employee card image file.                     |
| employee_card_selfie    | file (optional)Selfie with employee card image file.         |
| last_certificate        | file (optional)Certificate image file.                       |
| last_certificate_selfie | file (optional)Selfie with certificate image file.           |
| passport                | file (optional)Passport image file.                          |
| passport_selfie         | file (optional)Selfie with passport image file.              |
| family_card             | file (optional)Family card image file.                       |
| family_card_selfie      | file (optional)Selfie with family card image file.           |
| driving_license         | file (optional)Driving license image file.                   |
| driving_license_selfie  | file (optional)Selfie with driving license image file.       |
| married_card            | file (optional)Married card image file.                      |
| married_card_selfie     | file (optional)Selfie with married card image file.          |
| npwp                    | file (optional)NPWP image file.                              |
| npwp_selfie             | file (optional)Selfie with NPWP image file.                  |
| bpjs_kesehatan          | file (optional)BPJS Kesehatan image file.                    |
| bpjs_kesehatan_selfie   | file (optional)Selfie with BPJS Kesehatan image file.        |

Example UploadSupportingDocumentsRequest:
{
  "user_id":"1",
  "user_type":"1",
  "passport":"(binary)"
}
*/
type UploadSupportingDocumentsRequest struct {
	UserId   int64  `json:"user_id"`
	UserType string `json:"user_type"`
}

type UploadSupportingDocumentFilesRequest map[string]requests.UploadFile

// UploadSupportingDocumentResponse
// the response the same as SupportDocuments
type UploadSupportingDocumentResponse []SupportDocuments

// --------------------------- KYC SUBMISSION

type KycSubmissionRequest struct {
	UserType string `json:"user_type"`
}

type KycSubmissionResponse struct {
	Message string `json:"message"`
}

// --------------------------- REPAIR DATA

/*
RepairDataRequest
| Attribute             | Description                                                                                                                                                                                                                                                                              |
|-----------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| user_type             | string (required)Type of user. Validation:User type Agent: 1                                                                                                                                                                                                                             |
| name                  | string (optional)Agent’s full name. Validation:Alphanumeric1-50 characters                                                                                                                                                                                                               |
| identity_type         | string (optional)Type of identity. Validation:Accepted value are as follows:1: National Id Card or KTP in Indonesia2: Passport                                                                                                                                                           |
| identity_number       | string (optional)Identity number based on the type chosen. Validation:16 characters for KTP20 characters for Passport                                                                                                                                                                    |
| birth_date            | string (optional)Agent’s birth date. Validation:DD-MM-YYYY format                                                                                                                                                                                                                        |
| birth_place           | string (optional)Agent’s birth place.Validation:Alphanumeric255 characters                                                                                                                                                                                                               |
| gender                | string (optional)Agent’s gender. Validation:Accepted value are as follows:1 for male2 for female                                                                                                                                                                                         |
| occupation            | string (optional)Agent’s job. Validation:Accepted value are as follows:housewifeentrepreneurprivate_employeegovernment_employeefoundation_boardPeople who work at foundation as an employeeindonesian_migrant_workerAlso known as TKIcompanyIf sender is a company, not individualothers |
| country_id            | string (optional)Agent’s country code based on the identity chosen. Validation:Available values can be retrieved from the country list API here.                                                                                                                                         |
| province_id           | string (optional)Agent’s province code based on the identity chosen. Validation:Available values can be retrieved from the province list API here.                                                                                                                                       |
| city_id               | string (optional)Agent’s city code based on the identity chosen. Validation:Available values can be retrieved from the city list API here.                                                                                                                                               |
| district_id           | string (optional)Agent’s district code based on the identity chosen. Validation:Available values can be retrieved from the district list API here.                                                                                                                                       |
| address               | string (optional)Agent’s address based on the identity chosen. Validation:Alphanumeric10-300 charactersSpaces, ., -, /, (, and )                                                                                                                                                         |
| residence_country_id  | string (optional)Agent’s current residential address country code. Validation:Available values can be retrieved from the country list API here.                                                                                                                                          |
| residence_province_id | string (optional)Agent’s current residential address province code. Validation:Available values can be retrieved from the province list API here.                                                                                                                                        |
| residence_city_id     | string (optional)Agent’s current residential address city code. Validation:Available values can be retrieved from the city list API here.                                                                                                                                                |
| residence_district_id | string (optional)Agent’s current residential address district code. Validation:Available values can be retrieved from the district list API here.                                                                                                                                        |
| residence_address     | string (optional)Agent’s current residential address. Validation:Alphanumeric10-300 charactersSpaces, ., -, /, (, and )                                                                                                                                                                  |

Example RepairDataRequest:
{
  "user_type":"1",
  "name":"John Smith",
  "identity_type":"1",
  "identity_number":"1234567890123456",
  "birth_date":"31-10-1990",
  "birth_place":"Kota Depok",
  "gender":"1",
  "occupation":"entrepreneur",
  "country_id":"1",
  "province_id":"12",
  "city_id":"184",
  "district_id":"1234",
  "address":"Jalan Margonda Raya",
  "residence_country_id":"1",
  "residence_province_id":"12",
  "residence_city_id":"184",
  "residence_district_id":"1234",
  "residence_address":"Jalan Margonda Raya"
}
*/
type RepairDataRequest struct {
	UserType            string `json:"user_type"`
	Name                string `json:"name"`
	IdentityType        string `json:"identity_type"`
	IdentityNumber      string `json:"identity_number"`
	BirthDate           string `json:"birth_date"`
	BirthPlace          string `json:"birth_place"`
	Gender              string `json:"gender"`
	Occupation          string `json:"occupation"`
	CountryId           string `json:"country_id"`
	ProvinceId          string `json:"province_id"`
	CityId              string `json:"city_id"`
	DistrictId          string `json:"district_id"`
	Address             string `json:"address"`
	ResidenceCountryId  string `json:"residence_country_id"`
	ResidenceProvinceId string `json:"residence_province_id"`
	ResidenceCityId     string `json:"residence_city_id"`
	ResidenceDistrictId string `json:"residence_district_id"`
	ResidenceAddress    string `json:"residence_address"`
}

/*
RepairDataResponse
| Attribute             | Description                                                                                                                                                                                                                                        |
|-------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| user_id               | Id of the Agent                                                                                                                                                                                                                                    |
| user_type             | User type Agent: 1                                                                                                                                                                                                                                 |
| name                  | Agent’s full name                                                                                                                                                                                                                                  |
| birth_date            | Agent’s birth date                                                                                                                                                                                                                                 |
| birth_place           | Agent’s birth place                                                                                                                                                                                                                                |
| gender                | Agent’s gender. Possible values are:1 for male2 for female                                                                                                                                                                                         |
| status                | For this endpoint, the returned status will always be 20. It means PENDING (See KYC status enum)                                                                                                                                                   |
| basic_data_filled     | Flag if the Agent’s basic data has been filled                                                                                                                                                                                                     |
| identity_type         | Type of identity. Possible values:1: National Id Card or KTP in Indonesia2: Passport                                                                                                                                                               |
| nik                   | Identity number value if the identity_type is KTP (1)                                                                                                                                                                                              |
| passport_number       | Identity number value if the identity_type is Passport (2)                                                                                                                                                                                         |
| occupation            | Agent’s job. Possible values are:housewifeentrepreneurprivate_employeegovernment_employeefoundation_boardPeople who work at foundation as an employeeindonesian_migrant_workerAlso known as TKIcompanyIf sender is a company, not individualothers |
| country_id            | Agent’s country code                                                                                                                                                                                                                               |
| province_id           | Agent’s province code                                                                                                                                                                                                                              |
| city_id               | Agent’s city ID in the identity type chosen                                                                                                                                                                                                        |
| district_id           | Agent’s district ID in the identity type chosen                                                                                                                                                                                                    |
| address               | Agent’s address                                                                                                                                                                                                                                    |
| residence_country_id  | Agent’s current residential address country code                                                                                                                                                                                                   |
| residence_province_id | Agent’s current residential address province code                                                                                                                                                                                                  |
| residence_city_id     | Agent’s current residential address city code                                                                                                                                                                                                      |
| residence_district_id | Agent’s current residential address district code                                                                                                                                                                                                  |
| residence_address     | Agent’s current residential address                                                                                                                                                                                                                |
| created_at            | Agent’s identity data created at timestamp                                                                                                                                                                                                         |
| updated_at            | Agent’s identity data updated at timestamp                                                                                                                                                                                                         |

Example RepairDataResponse:
*/
type RepairDataResponse struct {
	Attribute           string `json:"Attribute"`
	UserId              string `json:"user_id"`
	UserType            string `json:"user_type"`
	Name                string `json:"name"`
	BirthDate           string `json:"birth_date"`
	BirthPlace          string `json:"birth_place"`
	Gender              string `json:"gender"`
	Status              string `json:"status"`
	BasicDataFilled     string `json:"basic_data_filled"`
	IdentityType        string `json:"identity_type"`
	Nik                 string `json:"nik"`
	PassportNumber      string `json:"passport_number"`
	Occupation          string `json:"occupation"`
	CountryId           string `json:"country_id"`
	ProvinceId          string `json:"province_id"`
	CityId              string `json:"city_id"`
	DistrictId          string `json:"district_id"`
	Address             string `json:"address"`
	ResidenceCountryId  string `json:"residence_country_id"`
	ResidenceProvinceId string `json:"residence_province_id"`
	ResidenceCityId     string `json:"residence_city_id"`
	ResidenceDistrictId string `json:"residence_district_id"`
	ResidenceAddress    string `json:"residence_address"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
}

// --------------------------- REPAIR IDENTITY IMAGE

/*
RepairIdentityImageRequest
| Attribute | Description                                                                                                                                                                                                                                                   |
|-----------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| user_type | string (required)Type of user. Validation:User type Agent: 1                                                                                                                                                                                                  |
| image     | file (required)Agent’s identity image file. Validation:Size less than 2MBDimension between 256 (W) x 256 (H) and 4096 (W) x 4096 (H)File type should be png, jpg, or jpegImage should be uploaded right from Camera app (not Gallery or File or Document app) |

Example RepairIdentityImageRequest:
{
  "user_type":"1",
  "image":(binary)
}
*/
type RepairIdentityImageRequest struct {
	UserType string `json:"user_type"`
}

type RepairIdentityImageFileRequest struct {
	Image requests.UploadFile `json:"image"`
}

/*
RepairIdentityImageResponse
| Attribute | Description               |
|-----------|---------------------------|
| image_url | URL of the uploaded image |

Example RepairIdentityImageResponse:
*/
type RepairIdentityImageResponse struct {
	ImageUrl string `json:"image_url"`
}

// --------------------------- LOCATION KYC PARAMS

/*
LocationKycRequest
| Param       | Value                                                        |
|-------------|--------------------------------------------------------------|
| user_type   | string (required)Type of user. Validation:User type Agent: 1 |
| country_id  | integer (optional) Country ID for filtering the provinces    |
| province_id | integer (optional) City ID for filtering the province_id     |
| city_id     | integer (optional) City ID for filtering the districts       |

*/
type LocationKycRequest struct {
	UserType   string `json:"user_type"`
	CountryID  int64  `json:"country_id,omitempty"`
	ProvinceID int64  `json:"province_id,omitempty"`
	CityID     int64  `json:"city_id,omitempty"`
}

type LocationKycResponse struct {
	Countries []Locations `json:"countries,omitempty"`
	Provinces []Locations `json:"provinces,omitempty"`
	Cities    []Locations `json:"cities,omitempty"`
	Districts []Locations `json:"districts,omitempty"`
}
