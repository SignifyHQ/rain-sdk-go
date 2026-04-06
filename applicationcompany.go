// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainhelloworld

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/rain-hello-world-go/internal/apiform"
	"github.com/stainless-sdks/rain-hello-world-go/internal/apijson"
	"github.com/stainless-sdks/rain-hello-world-go/internal/requestconfig"
	"github.com/stainless-sdks/rain-hello-world-go/option"
	"github.com/stainless-sdks/rain-hello-world-go/packages/param"
	"github.com/stainless-sdks/rain-hello-world-go/packages/respjson"
)

// ApplicationCompanyService contains methods and other services that help with
// interacting with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApplicationCompanyService] method instead.
type ApplicationCompanyService struct {
	options []option.RequestOption
	Ubo     ApplicationCompanyUboService
}

// NewApplicationCompanyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewApplicationCompanyService(opts ...option.RequestOption) (r ApplicationCompanyService) {
	r = ApplicationCompanyService{}
	r.options = opts
	r.Ubo = NewApplicationCompanyUboService(opts...)
	return
}

// Submits an application to create a corporate account. The application requires
// details about the company, its legal entity, representatives, and beneficial
// owners. The initial user must provide a wallet address.
func (r *ApplicationCompanyService) New(ctx context.Context, body ApplicationCompanyNewParams, opts ...option.RequestOption) (res *IssuingCompany, err error) {
	opts = slices.Concat(r.options, opts)
	path := "applications/company"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves the current status and details of a company's corporate application,
// including the company's ultimate beneficial owners and application progress.
func (r *ApplicationCompanyService) Get(ctx context.Context, companyID string, opts ...option.RequestOption) (res *ApplicationCompanyGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if companyID == "" {
		err = errors.New("missing required companyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("applications/company/%s", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the information for an existing corporate account application. The
// company's details, including name, address, and legal entity information, can be
// modified through this endpoint.
func (r *ApplicationCompanyService) Update(ctx context.Context, companyID string, body ApplicationCompanyUpdateParams, opts ...option.RequestOption) (res *IssuingCompany, err error) {
	opts = slices.Concat(r.options, opts)
	if companyID == "" {
		err = errors.New("missing required companyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("applications/company/%s", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Allows a company to reapply or respond to a request for information after
// submitting their corporate application. This endpoint is typically used when
// additional information or corrections are needed.
//
// Deprecated: deprecated
func (r *ApplicationCompanyService) Reapply(ctx context.Context, companyID string, body ApplicationCompanyReapplyParams, opts ...option.RequestOption) (res *IssuingCompany, err error) {
	opts = slices.Concat(r.options, opts)
	if companyID == "" {
		err = errors.New("missing required companyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("applications/company/%s/reapply", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Uploads a document that supports a company's corporate application. This is
// typically used to provide additional documentation, such as proof of address,
// incorporation certificates, or other required legal documents.
func (r *ApplicationCompanyService) UploadDocument(ctx context.Context, companyID string, body ApplicationCompanyUploadDocumentParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if companyID == "" {
		err = errors.New("missing required companyId parameter")
		return err
	}
	path := fmt.Sprintf("applications/company/%s/document", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

// The details of an issuing application.
type IssuingApplication struct {
	// Represents the possible statuses of an application.
	//
	// Any of "approved", "pending", "needsInformation", "needsVerification",
	// "manualReview", "denied", "locked", "canceled".
	ApplicationStatus IssuingApplicationApplicationStatus `json:"applicationStatus" api:"required"`
	// The link to the application completion page
	ApplicationCompletionLink IssuingApplicationApplicationCompletionLink `json:"applicationCompletionLink"`
	// The link to the external verification page for the application
	//
	// Deprecated: deprecated
	ApplicationExternalVerificationLink IssuingApplicationApplicationExternalVerificationLink `json:"applicationExternalVerificationLink"`
	// The reason behind the current application status
	ApplicationReason string `json:"applicationReason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicationStatus                   respjson.Field
		ApplicationCompletionLink           respjson.Field
		ApplicationExternalVerificationLink respjson.Field
		ApplicationReason                   respjson.Field
		ExtraFields                         map[string]respjson.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingApplication) RawJSON() string { return r.JSON.raw }
func (r *IssuingApplication) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents the possible statuses of an application.
type IssuingApplicationApplicationStatus string

const (
	IssuingApplicationApplicationStatusApproved          IssuingApplicationApplicationStatus = "approved"
	IssuingApplicationApplicationStatusPending           IssuingApplicationApplicationStatus = "pending"
	IssuingApplicationApplicationStatusNeedsInformation  IssuingApplicationApplicationStatus = "needsInformation"
	IssuingApplicationApplicationStatusNeedsVerification IssuingApplicationApplicationStatus = "needsVerification"
	IssuingApplicationApplicationStatusManualReview      IssuingApplicationApplicationStatus = "manualReview"
	IssuingApplicationApplicationStatusDenied            IssuingApplicationApplicationStatus = "denied"
	IssuingApplicationApplicationStatusLocked            IssuingApplicationApplicationStatus = "locked"
	IssuingApplicationApplicationStatusCanceled          IssuingApplicationApplicationStatus = "canceled"
)

// The link to the application completion page
type IssuingApplicationApplicationCompletionLink struct {
	// The URL for the completion page
	URL    string                                            `json:"url" api:"required" format:"uri"`
	Params IssuingApplicationApplicationCompletionLinkParams `json:"params"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Params      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingApplicationApplicationCompletionLink) RawJSON() string { return r.JSON.raw }
func (r *IssuingApplicationApplicationCompletionLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IssuingApplicationApplicationCompletionLinkParams struct {
	// The user's unique identifier
	UserID string `json:"userId" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingApplicationApplicationCompletionLinkParams) RawJSON() string { return r.JSON.raw }
func (r *IssuingApplicationApplicationCompletionLinkParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The link to the external verification page for the application
//
// Deprecated: deprecated
type IssuingApplicationApplicationExternalVerificationLink struct {
	// The URL for the external verification page
	URL    string                                                      `json:"url" api:"required" format:"uri"`
	Params IssuingApplicationApplicationExternalVerificationLinkParams `json:"params"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Params      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingApplicationApplicationExternalVerificationLink) RawJSON() string { return r.JSON.raw }
func (r *IssuingApplicationApplicationExternalVerificationLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IssuingApplicationApplicationExternalVerificationLinkParams struct {
	// The user's unique identifier
	UserID string `json:"userId" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingApplicationApplicationExternalVerificationLinkParams) RawJSON() string {
	return r.JSON.raw
}
func (r *IssuingApplicationApplicationExternalVerificationLinkParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Address, BirthDate, CountryOfIssue, Email, FirstName, LastName,
// NationalID are required.
type IssuingApplicationPersonParam struct {
	// The person's address
	Address PhysicalAddressParam `json:"address,omitzero" api:"required"`
	// The person's birth date
	BirthDate time.Time `json:"birthDate" api:"required" format:"date"`
	// The 2-digit country code of the person's national ID issuer
	CountryOfIssue string `json:"countryOfIssue" api:"required"`
	// The user's email address
	Email string `json:"email" api:"required"`
	// The person's first name
	FirstName string `json:"firstName" api:"required"`
	// The person's last name
	LastName string `json:"lastName" api:"required"`
	// The person's national ID number. For the US, this is a 9-digit SSN
	NationalID string `json:"nationalId" api:"required"`
	// The person's unique identifier
	ID param.Opt[string] `json:"id,omitzero" format:"uuid"`
	// The country code for the phone number
	PhoneCountryCode param.Opt[string] `json:"phoneCountryCode,omitzero"`
	// The phone number of the person
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	paramObj
}

func (r IssuingApplicationPersonParam) MarshalJSON() (data []byte, err error) {
	type shadow IssuingApplicationPersonParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IssuingApplicationPersonParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The details of an issuing application.
type IssuingCompany struct {
	// The company's unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// The company's physical address
	Address PhysicalAddress `json:"address" api:"required"`
	// The company's name on the Rain platform
	Name string `json:"name" api:"required"`
	// The company's ultimate beneficial owners (UBOs)
	UltimateBeneficialOwners []IssuingCompanyUltimateBeneficialOwner `json:"ultimateBeneficialOwners"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		Address                  respjson.Field
		Name                     respjson.Field
		UltimateBeneficialOwners respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
	IssuingApplication
}

// Returns the unmodified JSON received from the API
func (r IssuingCompany) RawJSON() string { return r.JSON.raw }
func (r *IssuingCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The details of an issuing application.
type IssuingCompanyUltimateBeneficialOwner struct {
	// The UBO's unique identifier
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
func (r IssuingCompanyUltimateBeneficialOwner) RawJSON() string { return r.JSON.raw }
func (r *IssuingCompanyUltimateBeneficialOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a physical address with components like street, city, region, postal
// code, and country.
type PhysicalAddress struct {
	// The city of the address
	City string `json:"city" api:"required"`
	// The full name of the country
	Country string `json:"country" api:"required"`
	// The 2-letter country code of the address
	CountryCode string `json:"countryCode" api:"required"`
	// The first line of the street address
	Line1 string `json:"line1" api:"required"`
	// The postal or zip code of the address
	PostalCode string `json:"postalCode" api:"required"`
	// The region or state of the address
	Region string `json:"region" api:"required"`
	// The second line of the street address (optional)
	Line2 string `json:"line2"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		Country     respjson.Field
		CountryCode respjson.Field
		Line1       respjson.Field
		PostalCode  respjson.Field
		Region      respjson.Field
		Line2       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhysicalAddress) RawJSON() string { return r.JSON.raw }
func (r *PhysicalAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PhysicalAddress to a PhysicalAddressParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PhysicalAddressParam.Overrides()
func (r PhysicalAddress) ToParam() PhysicalAddressParam {
	return param.Override[PhysicalAddressParam](json.RawMessage(r.RawJSON()))
}

// Represents a physical address with components like street, city, region, postal
// code, and country.
//
// The properties City, Country, CountryCode, Line1, PostalCode, Region are
// required.
type PhysicalAddressParam struct {
	// The city of the address
	City string `json:"city" api:"required"`
	// The full name of the country
	Country string `json:"country" api:"required"`
	// The 2-letter country code of the address
	CountryCode string `json:"countryCode" api:"required"`
	// The first line of the street address
	Line1 string `json:"line1" api:"required"`
	// The postal or zip code of the address
	PostalCode string `json:"postalCode" api:"required"`
	// The region or state of the address
	Region string `json:"region" api:"required"`
	// The second line of the street address (optional)
	Line2 param.Opt[string] `json:"line2,omitzero"`
	paramObj
}

func (r PhysicalAddressParam) MarshalJSON() (data []byte, err error) {
	type shadow PhysicalAddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhysicalAddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The details of an issuing application.
type ApplicationCompanyGetResponse struct {
	// The identifier of the company application
	ID string `json:"id" api:"required" format:"uuid"`
	// The company's ultimate beneficial owners (UBOs)
	UltimateBeneficialOwners []ApplicationCompanyGetResponseUltimateBeneficialOwner `json:"ultimateBeneficialOwners" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		UltimateBeneficialOwners respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
	IssuingApplication
}

// Returns the unmodified JSON received from the API
func (r ApplicationCompanyGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ApplicationCompanyGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The details of an issuing application.
type ApplicationCompanyGetResponseUltimateBeneficialOwner struct {
	// The UBO's unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// The UBO's email address
	Email string `json:"email" format:"email"`
	// The UBO's first name
	FirstName string `json:"firstName"`
	// The UBO's last name
	LastName string `json:"lastName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Email       respjson.Field
		FirstName   respjson.Field
		LastName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	IssuingApplication
}

// Returns the unmodified JSON received from the API
func (r ApplicationCompanyGetResponseUltimateBeneficialOwner) RawJSON() string { return r.JSON.raw }
func (r *ApplicationCompanyGetResponseUltimateBeneficialOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationCompanyNewParams struct {
	// The company's physical address
	Address PhysicalAddressParam `json:"address,omitzero" api:"required"`
	// The company's legal entity details.
	Entity ApplicationCompanyNewParamsEntity `json:"entity,omitzero" api:"required"`
	// The initial user of the company. This user must have a wallet address, and their
	// wallet address will be associated as an owner on the company's Rain smart
	// contract.
	InitialUser ApplicationCompanyNewParamsInitialUser `json:"initialUser,omitzero" api:"required"`
	// The name of the company requesting to create an account
	Name string `json:"name" api:"required"`
	// The company's representatives
	Representatives []IssuingApplicationPersonParam `json:"representatives,omitzero" api:"required"`
	// The company's ultimate beneficial owners (UBOs)
	UltimateBeneficialOwners []IssuingApplicationPersonParam `json:"ultimateBeneficialOwners,omitzero" api:"required"`
	// The chain ID of the external collateral contract, if used. Not required when
	// using Rain's collateral contracts.
	ChainID param.Opt[string] `json:"chainId,omitzero"`
	// The address of the external collateral contract, if used. Not required when
	// using Rain's collateral contracts.
	ContractAddress param.Opt[string] `json:"contractAddress,omitzero"`
	// A unique identifier for the origin of the user
	SourceKey param.Opt[string] `json:"sourceKey,omitzero"`
	paramObj
}

func (r ApplicationCompanyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ApplicationCompanyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplicationCompanyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The company's legal entity details.
//
// The properties Name, RegistrationNumber, TaxID, Website are required.
type ApplicationCompanyNewParamsEntity struct {
	// The legal entity's name
	Name string `json:"name" api:"required"`
	// The legal entity's registration number
	RegistrationNumber string `json:"registrationNumber" api:"required"`
	// The legal entity's national tax id
	TaxID string `json:"taxId" api:"required"`
	// The legal entity's website
	Website string `json:"website" api:"required"`
	// A brief description of the legal entity and its activities
	Description param.Opt[string] `json:"description,omitzero"`
	// The estimated monthly spending by the legal entity
	ExpectedSpend param.Opt[string] `json:"expectedSpend,omitzero"`
	// The type of legal entity (e.g., LLC, S Corp)
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r ApplicationCompanyNewParamsEntity) MarshalJSON() (data []byte, err error) {
	type shadow ApplicationCompanyNewParamsEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplicationCompanyNewParamsEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The initial user of the company. This user must have a wallet address, and their
// wallet address will be associated as an owner on the company's Rain smart
// contract.
type ApplicationCompanyNewParamsInitialUser struct {
	// This user's IP address
	IPAddress string `json:"ipAddress" api:"required"`
	// Indicates whether the user has accepted the terms of service
	IsTermsOfServiceAccepted bool `json:"isTermsOfServiceAccepted,omitzero" api:"required"`
	// This user's role at their company (not their role on the Rain platform)
	Role param.Opt[string] `json:"role,omitzero"`
	// The user's Solana address. Either this or a EVM address is required if using a
	// Rain-managed solution, but optional otherwise.
	SolanaAddress param.Opt[string] `json:"solanaAddress,omitzero"`
	// The user's Ethereum Virtual Machine (EVM) address. Either this or a Solana
	// address is required if using a Rain-managed solution, but optional otherwise.
	WalletAddress param.Opt[string] `json:"walletAddress,omitzero"`
	IssuingApplicationPersonParam
}

func (r ApplicationCompanyNewParamsInitialUser) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*ApplicationCompanyNewParamsInitialUser
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type ApplicationCompanyUpdateParams struct {
	// The name of the company for the corporate application
	Name param.Opt[string] `json:"name,omitzero"`
	// The company's physical address
	Address PhysicalAddressParam `json:"address,omitzero"`
	// The company's legal entity details.
	Entity ApplicationCompanyUpdateParamsEntity `json:"entity,omitzero"`
	paramObj
}

func (r ApplicationCompanyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ApplicationCompanyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplicationCompanyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The company's legal entity details.
type ApplicationCompanyUpdateParamsEntity struct {
	// A brief description of the legal entity and its activities
	Description param.Opt[string] `json:"description,omitzero"`
	// The estimated monthly spending by the legal entity
	ExpectedSpend param.Opt[string] `json:"expectedSpend,omitzero"`
	// The legal entity's registration number
	RegistrationNumber param.Opt[string] `json:"registrationNumber,omitzero"`
	// The legal entity's national tax ID
	TaxID param.Opt[string] `json:"taxId,omitzero"`
	// The type of legal entity (e.g., LLC, S Corp)
	Type param.Opt[string] `json:"type,omitzero"`
	// The legal entity's website
	Website param.Opt[string] `json:"website,omitzero"`
	paramObj
}

func (r ApplicationCompanyUpdateParamsEntity) MarshalJSON() (data []byte, err error) {
	type shadow ApplicationCompanyUpdateParamsEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplicationCompanyUpdateParamsEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationCompanyReapplyParams struct {
	// The company's physical address
	Address PhysicalAddressParam `json:"address,omitzero" api:"required"`
	// The company's legal entity details.
	Entity ApplicationCompanyReapplyParamsEntity `json:"entity,omitzero" api:"required"`
	// The initial user of the company who will be the owner on the Rain smart
	// contract. This user must provide various personal details.
	InitialUser ApplicationCompanyReapplyParamsInitialUser `json:"initialUser,omitzero" api:"required"`
	// The name of the company reapplying for the corporate application
	Name string `json:"name" api:"required"`
	// The company's representatives
	Representatives []IssuingApplicationPersonParam `json:"representatives,omitzero" api:"required"`
	// The company's ultimate beneficial owners (UBOs)
	UltimateBeneficialOwners []IssuingApplicationPersonParam `json:"ultimateBeneficialOwners,omitzero" api:"required"`
	paramObj
}

func (r ApplicationCompanyReapplyParams) MarshalJSON() (data []byte, err error) {
	type shadow ApplicationCompanyReapplyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplicationCompanyReapplyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The company's legal entity details.
//
// The property Website is required.
type ApplicationCompanyReapplyParamsEntity struct {
	// The legal entity's website
	Website string `json:"website" api:"required"`
	// A brief description of the legal entity, and its activities
	Description param.Opt[string] `json:"description,omitzero"`
	// The estimated monthly spending by the legal entity
	ExpectedSpend param.Opt[string] `json:"expectedSpend,omitzero"`
	// The type of legal entity (e.g., LLC, S Corp)
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r ApplicationCompanyReapplyParamsEntity) MarshalJSON() (data []byte, err error) {
	type shadow ApplicationCompanyReapplyParamsEntity
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplicationCompanyReapplyParamsEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The initial user of the company who will be the owner on the Rain smart
// contract. This user must provide various personal details.
//
// The properties Address, BirthDate, CountryOfIssue, IPAddress,
// IsTermsOfServiceAccepted, NationalID are required.
type ApplicationCompanyReapplyParamsInitialUser struct {
	// The user's address
	Address PhysicalAddressParam `json:"address,omitzero" api:"required"`
	// The user's birth date
	BirthDate time.Time `json:"birthDate" api:"required" format:"date"`
	// The 2-digit country code of the user's national ID issuer
	CountryOfIssue string `json:"countryOfIssue" api:"required"`
	// The user's IP address
	IPAddress string `json:"ipAddress" api:"required"`
	// Indicates whether the user has accepted the terms of service
	//
	// Any of true.
	IsTermsOfServiceAccepted bool `json:"isTermsOfServiceAccepted,omitzero" api:"required"`
	// The user's national ID number. For the US, this is a 9-digit SSN
	NationalID string `json:"nationalId" api:"required"`
	// This user's role at their company (not their role on the Rain platform)
	Role param.Opt[string] `json:"role,omitzero"`
	paramObj
}

func (r ApplicationCompanyReapplyParamsInitialUser) MarshalJSON() (data []byte, err error) {
	type shadow ApplicationCompanyReapplyParamsInitialUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplicationCompanyReapplyParamsInitialUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ApplicationCompanyReapplyParamsInitialUser](
		"isTermsOfServiceAccepted", true,
	)
}

type ApplicationCompanyUploadDocumentParams struct {
	// The actual document file to be uploaded. The document must be in binary format,
	// and the maximum allowed size is 20 MB.
	Document io.Reader `json:"document,omitzero" api:"required" format:"binary"`
	// The country where the document was issued
	Country param.Opt[string] `json:"country,omitzero"`
	// The name of the document being uploaded
	Name param.Opt[string] `json:"name,omitzero"`
	// The side of the document being uploaded
	//
	// Any of "front", "back".
	Side ApplicationCompanyUploadDocumentParamsSide `json:"side,omitzero"`
	// The type of the document being uploaded
	//
	// Any of "directorsRegistry", "stateRegistry", "incumbencyCert", "proofOfAddress",
	// "trustAgreement", "informationStatement", "incorporationCert",
	// "incorporationArticles", "shareholderRegistry", "goodStandingCert",
	// "powerOfAttorney", "other".
	Type ApplicationCompanyUploadDocumentParamsType `json:"type,omitzero"`
	paramObj
}

func (r ApplicationCompanyUploadDocumentParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type ApplicationCompanyUploadDocumentParamsSide string

const (
	ApplicationCompanyUploadDocumentParamsSideFront ApplicationCompanyUploadDocumentParamsSide = "front"
	ApplicationCompanyUploadDocumentParamsSideBack  ApplicationCompanyUploadDocumentParamsSide = "back"
)

// The type of the document being uploaded
type ApplicationCompanyUploadDocumentParamsType string

const (
	ApplicationCompanyUploadDocumentParamsTypeDirectorsRegistry     ApplicationCompanyUploadDocumentParamsType = "directorsRegistry"
	ApplicationCompanyUploadDocumentParamsTypeStateRegistry         ApplicationCompanyUploadDocumentParamsType = "stateRegistry"
	ApplicationCompanyUploadDocumentParamsTypeIncumbencyCert        ApplicationCompanyUploadDocumentParamsType = "incumbencyCert"
	ApplicationCompanyUploadDocumentParamsTypeProofOfAddress        ApplicationCompanyUploadDocumentParamsType = "proofOfAddress"
	ApplicationCompanyUploadDocumentParamsTypeTrustAgreement        ApplicationCompanyUploadDocumentParamsType = "trustAgreement"
	ApplicationCompanyUploadDocumentParamsTypeInformationStatement  ApplicationCompanyUploadDocumentParamsType = "informationStatement"
	ApplicationCompanyUploadDocumentParamsTypeIncorporationCert     ApplicationCompanyUploadDocumentParamsType = "incorporationCert"
	ApplicationCompanyUploadDocumentParamsTypeIncorporationArticles ApplicationCompanyUploadDocumentParamsType = "incorporationArticles"
	ApplicationCompanyUploadDocumentParamsTypeShareholderRegistry   ApplicationCompanyUploadDocumentParamsType = "shareholderRegistry"
	ApplicationCompanyUploadDocumentParamsTypeGoodStandingCert      ApplicationCompanyUploadDocumentParamsType = "goodStandingCert"
	ApplicationCompanyUploadDocumentParamsTypePowerOfAttorney       ApplicationCompanyUploadDocumentParamsType = "powerOfAttorney"
	ApplicationCompanyUploadDocumentParamsTypeOther                 ApplicationCompanyUploadDocumentParamsType = "other"
)
