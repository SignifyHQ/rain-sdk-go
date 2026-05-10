// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainsdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/SignifyHQ/rain-sdk-go/internal/apijson"
	"github.com/SignifyHQ/rain-sdk-go/internal/apiquery"
	shimjson "github.com/SignifyHQ/rain-sdk-go/internal/encoding/json"
	"github.com/SignifyHQ/rain-sdk-go/internal/requestconfig"
	"github.com/SignifyHQ/rain-sdk-go/option"
	"github.com/SignifyHQ/rain-sdk-go/packages/param"
	"github.com/SignifyHQ/rain-sdk-go/packages/respjson"
)

// UserService contains methods and other services that help with interacting with
// the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	options    []option.RequestOption
	Signatures UserSignatureService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.options = opts
	r.Signatures = NewUserSignatureService(opts...)
	return
}

// This endpoint allows the creation of a new authorized user with the necessary
// personal details, including contact and wallet information.
func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (res *IssuingUser, err error) {
	opts = slices.Concat(r.options, opts)
	path := "users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint retrieves a specific user by their unique ID. The user details,
// including contact and wallet information, are returned.
func (r *UserService) Get(ctx context.Context, userID string, opts ...option.RequestOption) (res *IssuingUser, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This endpoint allows the update of a specific user's information, such as their
// name, email, and contact details.
func (r *UserService) Update(ctx context.Context, userID string, body UserUpdateParams, opts ...option.RequestOption) (res *IssuingUser, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// This endpoint retrieves all users associated with a specific company. The
// response will return a list of users, including their personal and contact
// details.
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *[]IssuingUser, err error) {
	opts = slices.Concat(r.options, opts)
	path := "users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// This endpoint deletes a user by their unique ID. Once deleted, the user will no
// longer have access to the system.
func (r *UserService) Delete(ctx context.Context, userID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return err
	}
	path := fmt.Sprintf("users/%s", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// This endpoint allows the creation of a card for a specific user. The card can
// either be physical or virtual and can include various configuration options such
// as the display name and limit.
func (r *UserService) NewCard(ctx context.Context, userID string, body UserNewCardParams, opts ...option.RequestOption) (res *IssuingCard, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/%s/cards", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint allows creating a custom fee charge for a specific user. The fee
// can be defined in the request body along with its details.
func (r *UserService) NewCharge(ctx context.Context, userID string, body UserNewChargeParams, opts ...option.RequestOption) (res *IssuingChargeCreateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/%s/charges", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint allows initiating a payment for a specific user. The payment is
// made from the user's wallet address to the specified recipient address.
func (r *UserService) InitiatePayment(ctx context.Context, userID string, body UserInitiatePaymentParams, opts ...option.RequestOption) (res *UserInitiatePaymentResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/%s/payments", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint retrieves the credit balances for a specific user. It provides
// information about the user's credit limit, pending charges, posted charges,
// balance due, and spending power.
func (r *UserService) GetBalances(ctx context.Context, userID string, opts ...option.RequestOption) (res *UserGetBalancesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/%s/balances", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This endpoint retrieves smart contract information for a specific user. It
// returns a list of contracts associated with the user.
func (r *UserService) GetContracts(ctx context.Context, userID string, opts ...option.RequestOption) (res *[]IssuingContract, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/%s/contracts", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type UserInitiatePaymentResponse struct {
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
func (r UserInitiatePaymentResponse) RawJSON() string { return r.JSON.raw }
func (r *UserInitiatePaymentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetBalancesResponse struct {
	// Balance due of the user, in cents
	BalanceDue int64 `json:"balanceDue" api:"required"`
	// Credit limit of the user, in cents
	CreditLimit int64 `json:"creditLimit" api:"required"`
	// Pending charges of the user, in cents
	PendingCharges int64 `json:"pendingCharges" api:"required"`
	// Posted charges of the user, in cents
	PostedCharges int64 `json:"postedCharges" api:"required"`
	// The amount of money the user can spend, in cents
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
func (r UserGetBalancesResponse) RawJSON() string { return r.JSON.raw }
func (r *UserGetBalancesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewParams struct {
	// The user's email address
	Email string `json:"email" api:"required"`
	// The user's first name
	FirstName string `json:"firstName" api:"required"`
	// The user's last name
	LastName string `json:"lastName" api:"required"`
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

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserUpdateParams struct {
	// The user's email address
	Email param.Opt[string] `json:"email,omitzero"`
	// The user's first name
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// Indicates whether the user is active
	IsActive param.Opt[bool] `json:"isActive,omitzero"`
	// Indicates whether the user has accepted the terms of service
	IsTermsOfServiceAccepted param.Opt[bool] `json:"isTermsOfServiceAccepted,omitzero"`
	// The user's last name
	LastName param.Opt[string] `json:"lastName,omitzero"`
	// The user's phone country code
	PhoneCountryCode param.Opt[string] `json:"phoneCountryCode,omitzero"`
	// The user's phone number
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	// The user's address
	Address PhysicalAddressParam `json:"address,omitzero"`
	paramObj
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListParams struct {
	// For corporate cards, the identifier of the company to get users for
	CompanyID param.Opt[string] `query:"companyId,omitzero" format:"uuid" json:"-"`
	// The ID of the resource after which to start fetching
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// The number of resources to fetch
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserNewCardParams struct {
	// Any of "physical", "virtual".
	Type UserNewCardParamsType `json:"type,omitzero" api:"required"`
	// The address that will serve as the billing address for the card. Defaults to the
	// shipping address or team address if not explicitly provided.
	Billing PhysicalAddressParam `json:"billing,omitzero"`
	// Configuration details for the card, including display name, product ID, and
	// virtual card art
	Configuration UserNewCardParamsConfiguration `json:"configuration,omitzero"`
	// The initial credit limit for the card, expressed in cents
	Limit IssuingCardLimitParam `json:"limit,omitzero"`
	// The address to ship the card, if it is a physical card
	Shipping UserNewCardParamsShipping `json:"shipping,omitzero"`
	// The initial status of the card
	//
	// Any of "notActivated", "active", "locked", "canceled".
	Status IssuingCardStatus `json:"status,omitzero"`
	paramObj
}

func (r UserNewCardParams) MarshalJSON() (data []byte, err error) {
	type shadow UserNewCardParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewCardParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewCardParamsType string

const (
	UserNewCardParamsTypePhysical UserNewCardParamsType = "physical"
	UserNewCardParamsTypeVirtual  UserNewCardParamsType = "virtual"
)

// Configuration details for the card, including display name, product ID, and
// virtual card art
type UserNewCardParamsConfiguration struct {
	// The name to display on the card. If omitted, it will be the user's full name
	// trimmed to 30 characters.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// The product ID to use for the card, denoting its BIN range. Only relevant if
	// custom product IDs or product references are part of the user's contract.
	ProductID param.Opt[string] `json:"productId,omitzero"`
	// The product reference to use for the card, denoting its appearance. Only
	// relevant if custom product IDs or product references are part of the user's
	// contract.
	ProductRef param.Opt[string] `json:"productRef,omitzero"`
	// The virtual card art ID to use for the card, denoting its virtual appearance.
	// Only relevant if custom virtual card art IDs are part of the user's contract.
	VirtualCardArt param.Opt[string] `json:"virtualCardArt,omitzero"`
	paramObj
}

func (r UserNewCardParamsConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow UserNewCardParamsConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewCardParamsConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The address to ship the card, if it is a physical card
type UserNewCardParamsShipping struct {
	// The phone number to use for shipping, if it is a physical card
	PhoneNumber string `json:"phoneNumber" api:"required"`
	// The shipping method to use for the card. Standard and express are for US
	// addresses, while international is for non-US addresses.
	Method string `json:"method,omitzero"`
	PhysicalAddressParam
}

func (r UserNewCardParamsShipping) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*UserNewCardParamsShipping
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type UserNewChargeParams struct {
	// Represents the body of a request to create a charge, including the amount and a
	// description of the charge.
	IssuingChargeCreateBody IssuingChargeCreateBodyParam
	paramObj
}

func (r UserNewChargeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.IssuingChargeCreateBody)
}
func (r *UserNewChargeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserInitiatePaymentParams struct {
	// The amount of the payment, in cents
	Amount int64 `json:"amount" api:"required"`
	// The wallet address the payment is being sent from
	WalletAddress string `json:"walletAddress" api:"required"`
	// The chain ID (base-10 number) that the payment transaction is on
	ChainID param.Opt[int64] `json:"chainId,omitzero"`
	paramObj
}

func (r UserInitiatePaymentParams) MarshalJSON() (data []byte, err error) {
	type shadow UserInitiatePaymentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserInitiatePaymentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
