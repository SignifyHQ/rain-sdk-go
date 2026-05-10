// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainsdk

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"
	"time"

	"github.com/SignifyHQ/rain-sdk-go/internal/apiform"
	"github.com/SignifyHQ/rain-sdk-go/internal/apijson"
	"github.com/SignifyHQ/rain-sdk-go/internal/requestconfig"
	"github.com/SignifyHQ/rain-sdk-go/option"
	"github.com/SignifyHQ/rain-sdk-go/packages/param"
	"github.com/SignifyHQ/rain-sdk-go/packages/respjson"
)

// ApplicationUserService contains methods and other services that help with
// interacting with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApplicationUserService] method instead.
type ApplicationUserService struct {
	options []option.RequestOption
}

// NewApplicationUserService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewApplicationUserService(opts ...option.RequestOption) (r ApplicationUserService) {
	r = ApplicationUserService{}
	r.options = opts
	return
}

// Submits an application to create a consumer account for a user. The application
// can be submitted using a Sumsub share token, Persona share token, or directly
// via API. The user must provide details about their wallet, occupation, salary,
// and other account-related information.
func (r *ApplicationUserService) New(ctx context.Context, body ApplicationUserNewParams, opts ...option.RequestOption) (res *IssuingUser, err error) {
	opts = slices.Concat(r.options, opts)
	path := "applications/user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the current status and details of a user's consumer application. This
// includes the user's application progress and related information.
func (r *ApplicationUserService) Get(ctx context.Context, userID string, opts ...option.RequestOption) (res *ApplicationUserGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("applications/user/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the application information for a user, including personal details such
// as name, birth date, occupation, national ID, and account purpose.
func (r *ApplicationUserService) Update(ctx context.Context, userID string, body ApplicationUserUpdateParams, opts ...option.RequestOption) (res *IssuingUser, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("applications/user/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Submits an initial application for creating a consumer account. This request
// gathers basic personal details of the user, including their first and last name,
// email address, and optional wallet address if using a Rain-managed solution or
// hosted completion flow.
func (r *ApplicationUserService) Initiate(ctx context.Context, body ApplicationUserInitiateParams, opts ...option.RequestOption) (res *IssuingUser, err error) {
	opts = slices.Concat(r.options, opts)
	path := "applications/user/initiate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Allows a user to reapply or respond to a request for additional information
// regarding their consumer application. This is used when the initial application
// needs updating or more information is required.
//
// Deprecated: deprecated
func (r *ApplicationUserService) Reapply(ctx context.Context, userID string, body ApplicationUserReapplyParams, opts ...option.RequestOption) (res *IssuingUser, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("applications/user/%s/reapply", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Uploads a document for a user to support their consumer application. This is
// used to provide additional verification documents such as IDs, utility bills,
// and other required legal documents.
func (r *ApplicationUserService) UploadDocument(ctx context.Context, userID string, body ApplicationUserUploadDocumentParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return err
	}
	path := fmt.Sprintf("applications/user/%s/document", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

// The details of an issuing application.
type IssuingUser struct {
	// The user's unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// The user's email address
	Email string `json:"email" api:"required"`
	// The user's first name
	FirstName string `json:"firstName" api:"required"`
	// Indicates whether the user account is active
	IsActive bool `json:"isActive" api:"required"`
	// Indicates whether the user has accepted the terms of service
	IsTermsOfServiceAccepted bool `json:"isTermsOfServiceAccepted" api:"required"`
	// The user's last name
	LastName string `json:"lastName" api:"required"`
	// The user's address
	Address PhysicalAddress `json:"address"`
	// The identifier of the company the user belongs to, if applicable
	CompanyID string `json:"companyId" format:"uuid"`
	// The country code for the user's phone number
	PhoneCountryCode string `json:"phoneCountryCode"`
	// The user's phone number
	PhoneNumber string `json:"phoneNumber"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		Email                    respjson.Field
		FirstName                respjson.Field
		IsActive                 respjson.Field
		IsTermsOfServiceAccepted respjson.Field
		LastName                 respjson.Field
		Address                  respjson.Field
		CompanyID                respjson.Field
		PhoneCountryCode         respjson.Field
		PhoneNumber              respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
	IssuingApplication
}

// Returns the unmodified JSON received from the API
func (r IssuingUser) RawJSON() string { return r.JSON.raw }
func (r *IssuingUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The details of an issuing application.
type ApplicationUserGetResponse struct {
	// The identifier of the user's application
	ID string `json:"id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	IssuingApplication
}

// Returns the unmodified JSON received from the API
func (r ApplicationUserGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ApplicationUserGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationUserNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfUsingSumsubShareToken *ApplicationUserNewParamsBodyUsingSumsubShareToken `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfObject *ApplicationUserNewParamsBodyObject `json:",inline"`
	// This field is a request body variant, only one variant field can be set. The
	// user seeking to create an account. The user must have a wallet, and their wallet
	// will be linked as an owner on their Rain smart contract.
	OfUsingAPI *ApplicationUserNewParamsBodyUsingAPI `json:",inline"`

	paramObj
}

func (u ApplicationUserNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUsingSumsubShareToken, u.OfObject, u.OfUsingAPI)
}
func (r *ApplicationUserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccountPurpose, AnnualSalary, ExpectedMonthlyVolume, IPAddress,
// IsTermsOfServiceAccepted, Occupation, SumsubShareToken are required.
type ApplicationUserNewParamsBodyUsingSumsubShareToken struct {
	// The purpose of the user's account
	AccountPurpose string `json:"accountPurpose" api:"required"`
	// The user's annual salary
	AnnualSalary string `json:"annualSalary" api:"required"`
	// The estimated monthly spending amount for the user
	ExpectedMonthlyVolume string `json:"expectedMonthlyVolume" api:"required"`
	// This user's IP address
	IPAddress string `json:"ipAddress" api:"required"`
	// Indicates whether the user has accepted the terms of service
	//
	// Any of true.
	IsTermsOfServiceAccepted bool `json:"isTermsOfServiceAccepted,omitzero" api:"required"`
	// The user's occupation
	Occupation string `json:"occupation" api:"required"`
	// The Sumsub share token used for user verification
	SumsubShareToken string `json:"sumsubShareToken" api:"required"`
	// The chain ID of the user's external collateral contract, if applicable. Not
	// required when using Rain's collateral contracts.
	ChainID param.Opt[string] `json:"chainId,omitzero"`
	// The address of the user's external collateral contract, if applicable. Not
	// required when using Rain's collateral contracts.
	ContractAddress param.Opt[string] `json:"contractAddress,omitzero"`
	// Indicates whether the user will use existing documents for additional
	// verification
	HasExistingDocuments param.Opt[bool] `json:"hasExistingDocuments,omitzero"`
	// The user's Solana address. Either walletAddress or solanaAddress is required if
	// using a Rain-managed solution, but optional otherwise.
	SolanaAddress param.Opt[string] `json:"solanaAddress,omitzero"`
	// A unique identifier for the source of this user.
	SourceKey param.Opt[string] `json:"sourceKey,omitzero"`
	// The user's Ethereum Virtual Machine (EVM) address. Either walletAddress or
	// solanaAddress is required if using a Rain-managed solution, but optional
	// otherwise.
	WalletAddress param.Opt[string] `json:"walletAddress,omitzero"`
	paramObj
}

func (r ApplicationUserNewParamsBodyUsingSumsubShareToken) MarshalJSON() (data []byte, err error) {
	type shadow ApplicationUserNewParamsBodyUsingSumsubShareToken
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplicationUserNewParamsBodyUsingSumsubShareToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ApplicationUserNewParamsBodyUsingSumsubShareToken](
		"isTermsOfServiceAccepted", true,
	)
}

// The properties AccountPurpose, AnnualSalary, ExpectedMonthlyVolume, IPAddress,
// IsTermsOfServiceAccepted, Occupation, PersonaShareToken are required.
type ApplicationUserNewParamsBodyObject struct {
	// The purpose of the user's account
	AccountPurpose string `json:"accountPurpose" api:"required"`
	// The user's annual salary
	AnnualSalary string `json:"annualSalary" api:"required"`
	// The estimated monthly spending amount for the user
	ExpectedMonthlyVolume string `json:"expectedMonthlyVolume" api:"required"`
	// This user's IP address
	IPAddress string `json:"ipAddress" api:"required"`
	// Indicates whether the user has accepted the terms of service
	//
	// Any of true.
	IsTermsOfServiceAccepted bool `json:"isTermsOfServiceAccepted,omitzero" api:"required"`
	// The user's occupation
	Occupation string `json:"occupation" api:"required"`
	// The Persona inquiry ID
	PersonaShareToken string `json:"personaShareToken" api:"required"`
	// The chain ID of the user's external collateral contract, if applicable. Not
	// required when using Rain's collateral contracts.
	ChainID param.Opt[string] `json:"chainId,omitzero"`
	// The address of the user's external collateral contract, if applicable. Not
	// required when using Rain's collateral contracts.
	ContractAddress param.Opt[string] `json:"contractAddress,omitzero"`
	// Indicates whether the user will use existing documents for additional
	// verification
	HasExistingDocuments param.Opt[bool] `json:"hasExistingDocuments,omitzero"`
	// The user's Solana address. Either walletAddress or solanaAddress is required if
	// using a Rain-managed solution, but optional otherwise.
	SolanaAddress param.Opt[string] `json:"solanaAddress,omitzero"`
	// A unique identifier for the source of this user.
	SourceKey param.Opt[string] `json:"sourceKey,omitzero"`
	// The user's Ethereum Virtual Machine (EVM) address. Either walletAddress or
	// solanaAddress is required if using a Rain-managed solution, but optional
	// otherwise.
	WalletAddress param.Opt[string] `json:"walletAddress,omitzero"`
	paramObj
}

func (r ApplicationUserNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	type shadow ApplicationUserNewParamsBodyObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplicationUserNewParamsBodyObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ApplicationUserNewParamsBodyObject](
		"isTermsOfServiceAccepted", true,
	)
}

// The user seeking to create an account. The user must have a wallet, and their
// wallet will be linked as an owner on their Rain smart contract.
type ApplicationUserNewParamsBodyUsingAPI struct {
	// The purpose of the user's account
	AccountPurpose string `json:"accountPurpose" api:"required"`
	// The user's annual salary
	AnnualSalary string `json:"annualSalary" api:"required"`
	// The estimated monthly spending amount for the user
	ExpectedMonthlyVolume string `json:"expectedMonthlyVolume" api:"required"`
	// This user's IP address
	IPAddress string `json:"ipAddress" api:"required"`
	// Indicates whether the user has accepted the terms of service
	IsTermsOfServiceAccepted bool `json:"isTermsOfServiceAccepted,omitzero" api:"required"`
	// The user's occupation
	Occupation string `json:"occupation" api:"required"`
	// The chain ID of the user's external collateral contract, if applicable. Not
	// required when using Rain's collateral contracts.
	ChainID param.Opt[string] `json:"chainId,omitzero"`
	// The address of the user's external collateral contract, if applicable. Not
	// required when using Rain's collateral contracts.
	ContractAddress param.Opt[string] `json:"contractAddress,omitzero"`
	// Indicates whether the user will use existing documents for additional
	// verification
	HasExistingDocuments param.Opt[bool] `json:"hasExistingDocuments,omitzero"`
	// The user's Solana address. Either walletAddress or solanaAddress is required if
	// using a Rain-managed solution, but optional otherwise.
	SolanaAddress param.Opt[string] `json:"solanaAddress,omitzero"`
	// A unique identifier for the source of this user.
	SourceKey param.Opt[string] `json:"sourceKey,omitzero"`
	// The user's Ethereum Virtual Machine (EVM) address. Either walletAddress or
	// solanaAddress is required if using a Rain-managed solution, but optional
	// otherwise.
	WalletAddress param.Opt[string] `json:"walletAddress,omitzero"`
	IssuingApplicationPersonParam
}

func (r ApplicationUserNewParamsBodyUsingAPI) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ApplicationUserNewParamsBodyUsingAPI
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type ApplicationUserUpdateParams struct {
	// The purpose of the user's account
	AccountPurpose param.Opt[string] `json:"accountPurpose,omitzero"`
	// The user's annual salary
	AnnualSalary param.Opt[string] `json:"annualSalary,omitzero"`
	// The user's birth date
	BirthDate param.Opt[time.Time] `json:"birthDate,omitzero" format:"date"`
	// The 2-digit country code of the user's national ID issuer
	CountryOfIssue param.Opt[string] `json:"countryOfIssue,omitzero"`
	// The estimated monthly spending amount for the user
	ExpectedMonthlyVolume param.Opt[string] `json:"expectedMonthlyVolume,omitzero"`
	// The user's first name
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// Indicates whether the user will use existing documents for additional
	// verification
	HasExistingDocuments param.Opt[bool] `json:"hasExistingDocuments,omitzero"`
	// The user's IP address
	IPAddress param.Opt[string] `json:"ipAddress,omitzero"`
	// The user's last name
	LastName param.Opt[string] `json:"lastName,omitzero"`
	// The user's national ID number. For the US, this is a 9-digit SSN.
	NationalID param.Opt[string] `json:"nationalId,omitzero"`
	// The user's occupation
	Occupation param.Opt[string] `json:"occupation,omitzero"`
	// The user's address
	Address PhysicalAddressParam `json:"address,omitzero"`
	// Indicates whether the user has accepted the terms of service
	//
	// Any of true.
	IsTermsOfServiceAccepted bool `json:"isTermsOfServiceAccepted,omitzero"`
	paramObj
}

func (r ApplicationUserUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ApplicationUserUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplicationUserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationUserInitiateParams struct {
	// The user's email address
	Email param.Opt[string] `json:"email,omitzero"`
	// The user's first name
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The user's last name
	LastName param.Opt[string] `json:"lastName,omitzero"`
	// The user's wallet address. Required if using the hosted completion flow and a
	// Rain-managed solution. Not required otherwise.
	WalletAddress param.Opt[string] `json:"walletAddress,omitzero"`
	paramObj
}

func (r ApplicationUserInitiateParams) MarshalJSON() (data []byte, err error) {
	type shadow ApplicationUserInitiateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplicationUserInitiateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationUserReapplyParams struct {
	// The purpose of the user's account
	AccountPurpose string `json:"accountPurpose" api:"required"`
	// The user's address
	Address PhysicalAddressParam `json:"address,omitzero" api:"required"`
	// The user's annual salary
	AnnualSalary string `json:"annualSalary" api:"required"`
	// The user's birth date
	BirthDate time.Time `json:"birthDate" api:"required" format:"date"`
	// The 2-digit country code of the user's national ID issuer
	CountryOfIssue string `json:"countryOfIssue" api:"required"`
	// The estimated monthly spending amount for the user
	ExpectedMonthlyVolume string `json:"expectedMonthlyVolume" api:"required"`
	// The user's IP address
	IPAddress string `json:"ipAddress" api:"required"`
	// Indicates whether the user has accepted the terms of service
	//
	// Any of true.
	IsTermsOfServiceAccepted bool `json:"isTermsOfServiceAccepted,omitzero" api:"required"`
	// The user's national ID number. For the US, this is a 9-digit SSN
	NationalID string `json:"nationalId" api:"required"`
	// The user's occupation
	Occupation string `json:"occupation" api:"required"`
	// Indicates whether the user will use existing documents for additional
	// verification
	HasExistingDocuments param.Opt[bool] `json:"hasExistingDocuments,omitzero"`
	paramObj
}

func (r ApplicationUserReapplyParams) MarshalJSON() (data []byte, err error) {
	type shadow ApplicationUserReapplyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplicationUserReapplyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationUserUploadDocumentParams struct {
	// The actual document file to be uploaded. The document must be in binary format,
	// and the maximum allowed size is 20 MB.
	Document io.Reader `json:"document,omitzero" api:"required" format:"binary"`
	// The country where the document was issued
	Country param.Opt[string] `json:"country,omitzero"`
	// The name or title of the document being uploaded
	Name param.Opt[string] `json:"name,omitzero"`
	// The side of the document being uploaded
	//
	// Any of "front", "back".
	Side ApplicationUserUploadDocumentParamsSide `json:"side,omitzero"`
	// The type of the document being uploaded
	//
	// Any of "idCard", "passport", "drivers", "residencePermit", "utilityBill",
	// "selfie", "videoSelfie", "profileImage", "idDocPhoto", "agreement", "contract",
	// "driversTranslation", "investorDoc", "vehicleRegistrationCertificate",
	// "incomeSource", "paymentMethod", "bankCard", "covidVaccinationForm", "other".
	Type ApplicationUserUploadDocumentParamsType `json:"type,omitzero"`
	paramObj
}

func (r ApplicationUserUploadDocumentParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// The side of the document being uploaded
type ApplicationUserUploadDocumentParamsSide string

const (
	ApplicationUserUploadDocumentParamsSideFront ApplicationUserUploadDocumentParamsSide = "front"
	ApplicationUserUploadDocumentParamsSideBack  ApplicationUserUploadDocumentParamsSide = "back"
)

// The type of the document being uploaded
type ApplicationUserUploadDocumentParamsType string

const (
	ApplicationUserUploadDocumentParamsTypeIDCard                         ApplicationUserUploadDocumentParamsType = "idCard"
	ApplicationUserUploadDocumentParamsTypePassport                       ApplicationUserUploadDocumentParamsType = "passport"
	ApplicationUserUploadDocumentParamsTypeDrivers                        ApplicationUserUploadDocumentParamsType = "drivers"
	ApplicationUserUploadDocumentParamsTypeResidencePermit                ApplicationUserUploadDocumentParamsType = "residencePermit"
	ApplicationUserUploadDocumentParamsTypeUtilityBill                    ApplicationUserUploadDocumentParamsType = "utilityBill"
	ApplicationUserUploadDocumentParamsTypeSelfie                         ApplicationUserUploadDocumentParamsType = "selfie"
	ApplicationUserUploadDocumentParamsTypeVideoSelfie                    ApplicationUserUploadDocumentParamsType = "videoSelfie"
	ApplicationUserUploadDocumentParamsTypeProfileImage                   ApplicationUserUploadDocumentParamsType = "profileImage"
	ApplicationUserUploadDocumentParamsTypeIDDocPhoto                     ApplicationUserUploadDocumentParamsType = "idDocPhoto"
	ApplicationUserUploadDocumentParamsTypeAgreement                      ApplicationUserUploadDocumentParamsType = "agreement"
	ApplicationUserUploadDocumentParamsTypeContract                       ApplicationUserUploadDocumentParamsType = "contract"
	ApplicationUserUploadDocumentParamsTypeDriversTranslation             ApplicationUserUploadDocumentParamsType = "driversTranslation"
	ApplicationUserUploadDocumentParamsTypeInvestorDoc                    ApplicationUserUploadDocumentParamsType = "investorDoc"
	ApplicationUserUploadDocumentParamsTypeVehicleRegistrationCertificate ApplicationUserUploadDocumentParamsType = "vehicleRegistrationCertificate"
	ApplicationUserUploadDocumentParamsTypeIncomeSource                   ApplicationUserUploadDocumentParamsType = "incomeSource"
	ApplicationUserUploadDocumentParamsTypePaymentMethod                  ApplicationUserUploadDocumentParamsType = "paymentMethod"
	ApplicationUserUploadDocumentParamsTypeBankCard                       ApplicationUserUploadDocumentParamsType = "bankCard"
	ApplicationUserUploadDocumentParamsTypeCovidVaccinationForm           ApplicationUserUploadDocumentParamsType = "covidVaccinationForm"
	ApplicationUserUploadDocumentParamsTypeOther                          ApplicationUserUploadDocumentParamsType = "other"
)
