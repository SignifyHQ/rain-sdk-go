// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainhelloworld

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/SignifyHQ/rain-sdk-go/internal/apijson"
	"github.com/SignifyHQ/rain-sdk-go/internal/apiquery"
	shimjson "github.com/SignifyHQ/rain-sdk-go/internal/encoding/json"
	"github.com/SignifyHQ/rain-sdk-go/internal/requestconfig"
	"github.com/SignifyHQ/rain-sdk-go/option"
	"github.com/SignifyHQ/rain-sdk-go/packages/param"
	"github.com/SignifyHQ/rain-sdk-go/packages/respjson"
)

// CompanyService contains methods and other services that help with interacting
// with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCompanyService] method instead.
type CompanyService struct {
	options    []option.RequestOption
	Signatures CompanySignatureService
}

// NewCompanyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCompanyService(opts ...option.RequestOption) (r CompanyService) {
	r = CompanyService{}
	r.options = opts
	r.Signatures = NewCompanySignatureService(opts...)
	return
}

// Retrieve detailed information about a specific company using its unique ID
func (r *CompanyService) Get(ctx context.Context, companyID string, opts ...option.RequestOption) (res *IssuingCompany, err error) {
	opts = slices.Concat(r.options, opts)
	if companyID == "" {
		err = errors.New("missing required companyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("companies/%s", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update the details of an existing company such as its name and address
func (r *CompanyService) Update(ctx context.Context, companyID string, body CompanyUpdateParams, opts ...option.RequestOption) (res *IssuingCompany, err error) {
	opts = slices.Concat(r.options, opts)
	if companyID == "" {
		err = errors.New("missing required companyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("companies/%s", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Retrieves a list of all companies registered in the system
func (r *CompanyService) List(ctx context.Context, query CompanyListParams, opts ...option.RequestOption) (res *[]IssuingCompany, err error) {
	opts = slices.Concat(r.options, opts)
	path := "companies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Initiate a custom fee charge for a company.
func (r *CompanyService) Charge(ctx context.Context, companyID string, body CompanyChargeParams, opts ...option.RequestOption) (res *IssuingChargeCreateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if companyID == "" {
		err = errors.New("missing required companyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("companies/%s/charges", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Creates a new user within a specific company. The user must provide details such
// as their name, birthdate, and contact information.
func (r *CompanyService) NewUser(ctx context.Context, companyID string, body CompanyNewUserParams, opts ...option.RequestOption) (res *IssuingUser, err error) {
	opts = slices.Concat(r.options, opts)
	if companyID == "" {
		err = errors.New("missing required companyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("companies/%s/users", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Initiate a payment for a specific company by providing the payment amount and
// wallet address.
func (r *CompanyService) InitiatePayment(ctx context.Context, companyID string, body CompanyInitiatePaymentParams, opts ...option.RequestOption) (res *CompanyInitiatePaymentResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if companyID == "" {
		err = errors.New("missing required companyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("companies/%s/payments", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve the current credit balances of a company, including credit limits,
// pending charges, and the amount due.
func (r *CompanyService) GetBalances(ctx context.Context, companyID string, opts ...option.RequestOption) (res *CompanyGetBalancesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if companyID == "" {
		err = errors.New("missing required companyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("companies/%s/balances", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve the smart contract details associated with a company
func (r *CompanyService) GetContracts(ctx context.Context, companyID string, opts ...option.RequestOption) (res *[]IssuingContract, err error) {
	opts = slices.Concat(r.options, opts)
	if companyID == "" {
		err = errors.New("missing required companyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("companies/%s/contracts", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Represents the body of a request to create a charge, including the amount and a
// description of the charge.
//
// The properties Amount, Description are required.
type IssuingChargeCreateBodyParam struct {
	// The amount of the charge, in cents
	Amount int64 `json:"amount" api:"required"`
	// The description of the charge
	Description string `json:"description" api:"required"`
	paramObj
}

func (r IssuingChargeCreateBodyParam) MarshalJSON() (data []byte, err error) {
	type shadow IssuingChargeCreateBodyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IssuingChargeCreateBodyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents the response body returned after a charge is created, including the
// charge ID, creation timestamp, and charge details.
type IssuingChargeCreateResponse struct {
	// The identifier of the charge
	ID string `json:"id" api:"required" format:"uuid"`
	// The date and time when the charge was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The amount of the charge, in cents
	Amount int64 `json:"amount"`
	// The description of the charge
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Amount      respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingChargeCreateResponse) RawJSON() string { return r.JSON.raw }
func (r *IssuingChargeCreateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents an issuing contract with details about its deployment and token
// handling.
type IssuingContract struct {
	// The contract's unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// The chain ID (base-10 number) that the smart contract is deployed on
	ChainID int64 `json:"chainId" api:"required"`
	// Version of the contract
	ContractVersion int64 `json:"contractVersion" api:"required"`
	// The address of the contract's controller
	ControllerAddress string `json:"controllerAddress" api:"required"`
	// The proxy address of the contract
	ProxyAddress string `json:"proxyAddress" api:"required"`
	// Tokens that the contract accepts for transactions
	Tokens []IssuingContractToken `json:"tokens" api:"required"`
	// The address where funds should be deposited
	DepositAddress string `json:"depositAddress"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		ChainID           respjson.Field
		ContractVersion   respjson.Field
		ControllerAddress respjson.Field
		ProxyAddress      respjson.Field
		Tokens            respjson.Field
		DepositAddress    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingContract) RawJSON() string { return r.JSON.raw }
func (r *IssuingContract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IssuingContractToken struct {
	// The address of the token contract
	Address string `json:"address" api:"required"`
	// The advance rate for the token
	AdvanceRate float64 `json:"advanceRate"`
	// The balance of the token
	Balance string `json:"balance"`
	// The exchange rate for the token
	ExchangeRate float64 `json:"exchangeRate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address      respjson.Field
		AdvanceRate  respjson.Field
		Balance      respjson.Field
		ExchangeRate respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingContractToken) RawJSON() string { return r.JSON.raw }
func (r *IssuingContractToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyInitiatePaymentResponse struct {
	// The address to send the payment to
	Address string `json:"address" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyInitiatePaymentResponse) RawJSON() string { return r.JSON.raw }
func (r *CompanyInitiatePaymentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyGetBalancesResponse struct {
	// Balance due of the company, in cents
	BalanceDue int64 `json:"balanceDue" api:"required"`
	// Credit limit of the company, in cents
	CreditLimit int64 `json:"creditLimit" api:"required"`
	// Pending charges of the company, in cents
	PendingCharges int64 `json:"pendingCharges" api:"required"`
	// Posted charges of the company, in cents
	PostedCharges int64 `json:"postedCharges" api:"required"`
	// The amount of money the company can spend, in cents
	SpendingPower int64 `json:"spendingPower" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BalanceDue     respjson.Field
		CreditLimit    respjson.Field
		PendingCharges respjson.Field
		PostedCharges  respjson.Field
		SpendingPower  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompanyGetBalancesResponse) RawJSON() string { return r.JSON.raw }
func (r *CompanyGetBalancesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyUpdateParams struct {
	// The company's name on the Rain platform
	Name param.Opt[string] `json:"name,omitzero"`
	// The company's physical address
	Address PhysicalAddressParam `json:"address,omitzero"`
	paramObj
}

func (r CompanyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CompanyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompanyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyListParams struct {
	// The ID of the resource after which to start fetching
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// The number of resources to fetch
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CompanyListParams]'s query parameters as `url.Values`.
func (r CompanyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CompanyChargeParams struct {
	// Represents the body of a request to create a charge, including the amount and a
	// description of the charge.
	IssuingChargeCreateBody IssuingChargeCreateBodyParam
	paramObj
}

func (r CompanyChargeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.IssuingChargeCreateBody)
}
func (r *CompanyChargeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyNewUserParams struct {
	// The user's email address
	Email string `json:"email" api:"required"`
	// The user's first name
	FirstName string `json:"firstName" api:"required"`
	// Indicates whether the user has accepted the terms of service
	IsTermsOfServiceAccepted bool `json:"isTermsOfServiceAccepted" api:"required"`
	// The user's last name
	LastName string `json:"lastName" api:"required"`
	// The user's birth date
	BirthDate param.Opt[time.Time] `json:"birthDate,omitzero" format:"date"`
	// The user's phone country code
	PhoneCountryCode param.Opt[string] `json:"phoneCountryCode,omitzero"`
	// The user's phone number
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	// The user's wallet address
	WalletAddress param.Opt[string] `json:"walletAddress,omitzero"`
	// The user's address
	Address PhysicalAddressParam `json:"address,omitzero"`
	paramObj
}

func (r CompanyNewUserParams) MarshalJSON() (data []byte, err error) {
	type shadow CompanyNewUserParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompanyNewUserParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyInitiatePaymentParams struct {
	// The amount of the payment, in cents
	Amount int64 `json:"amount" api:"required"`
	// The wallet address the payment is being sent from
	WalletAddress string `json:"walletAddress" api:"required"`
	// The chain ID (base-10 number) that the payment transaction is on
	ChainID param.Opt[int64] `json:"chainId,omitzero"`
	paramObj
}

func (r CompanyInitiatePaymentParams) MarshalJSON() (data []byte, err error) {
	type shadow CompanyInitiatePaymentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompanyInitiatePaymentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
