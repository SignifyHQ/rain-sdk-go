// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainsdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/SignifyHQ/rain-sdk-go/internal/apijson"
	"github.com/SignifyHQ/rain-sdk-go/internal/apiquery"
	"github.com/SignifyHQ/rain-sdk-go/internal/requestconfig"
	"github.com/SignifyHQ/rain-sdk-go/option"
	"github.com/SignifyHQ/rain-sdk-go/packages/param"
	"github.com/SignifyHQ/rain-sdk-go/packages/respjson"
)

// CardService contains methods and other services that help with interacting with
// the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardService] method instead.
type CardService struct {
	options []option.RequestOption
	Pin     CardPinService
}

// NewCardService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCardService(opts ...option.RequestOption) (r CardService) {
	r = CardService{}
	r.options = opts
	r.Pin = NewCardPinService(opts...)
	return
}

// Retrieve detailed information for a specific card by its unique ID
func (r *CardService) Get(ctx context.Context, cardID string, opts ...option.RequestOption) (res *IssuingCard, err error) {
	opts = slices.Concat(r.options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cards/%s", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update details for an existing card, such as status, limit, billing address, and
// configuration.
func (r *CardService) Update(ctx context.Context, cardID string, body CardUpdateParams, opts ...option.RequestOption) (res *IssuingCard, err error) {
	opts = slices.Concat(r.options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cards/%s", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Retrieves all cards associated with a user or company. You can filter by user or
// company ID and card status.
func (r *CardService) List(ctx context.Context, query CardListParams, opts ...option.RequestOption) (res *[]IssuingCard, err error) {
	opts = slices.Concat(r.options, opts)
	path := "cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the encrypted data for a specific card, including the encrypted PAN and
// CVC
func (r *CardService) GetSecrets(ctx context.Context, cardID string, query CardGetSecretsParams, opts ...option.RequestOption) (res *CardGetSecretsResponse, err error) {
	if !param.IsOmitted(query.SessionID) {
		opts = append(opts, option.WithHeader("SessionId", fmt.Sprintf("%v", query.SessionID)))
	}
	opts = slices.Concat(r.options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cards/%s/secrets", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type IssuingCard struct {
	// The card's ID
	ID string `json:"id" api:"required" format:"uuid"`
	// The ID of the company that issued the card
	CompanyID string `json:"companyId" api:"required" format:"uuid"`
	// The expiration month of the card
	ExpirationMonth string `json:"expirationMonth" api:"required"`
	// The expiration year of the card
	ExpirationYear string `json:"expirationYear" api:"required"`
	// The last four digits of the card number
	Last4 string `json:"last4" api:"required"`
	// The card's current status
	//
	// Any of "notActivated", "active", "locked", "canceled".
	Status IssuingCardStatus `json:"status" api:"required"`
	// The type of the card (physical or virtual)
	//
	// Any of "physical", "virtual".
	Type IssuingCardType `json:"type" api:"required"`
	// The userID to whom the card was issued
	UserID string `json:"userId" api:"required" format:"uuid"`
	// The card's spending limit
	Limit        IssuingCardLimit `json:"limit"`
	TokenWallets []string         `json:"tokenWallets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CompanyID       respjson.Field
		ExpirationMonth respjson.Field
		ExpirationYear  respjson.Field
		Last4           respjson.Field
		Status          respjson.Field
		Type            respjson.Field
		UserID          respjson.Field
		Limit           respjson.Field
		TokenWallets    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingCard) RawJSON() string { return r.JSON.raw }
func (r *IssuingCard) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the card (physical or virtual)
type IssuingCardType string

const (
	IssuingCardTypePhysical IssuingCardType = "physical"
	IssuingCardTypeVirtual  IssuingCardType = "virtual"
)

// Represents the spending limit and frequency for a card.
type IssuingCardLimit struct {
	// The maximum spending amount in cents
	Amount int64 `json:"amount" api:"required"`
	// The frequency at which the spending limit resets
	//
	// Any of "per24HourPeriod", "per7DayPeriod", "per30DayPeriod", "perYearPeriod",
	// "allTime", "perAuthorization".
	Frequency IssuingCardLimitFrequency `json:"frequency" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Frequency   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingCardLimit) RawJSON() string { return r.JSON.raw }
func (r *IssuingCardLimit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this IssuingCardLimit to a IssuingCardLimitParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// IssuingCardLimitParam.Overrides()
func (r IssuingCardLimit) ToParam() IssuingCardLimitParam {
	return param.Override[IssuingCardLimitParam](json.RawMessage(r.RawJSON()))
}

// The frequency at which the spending limit resets
type IssuingCardLimitFrequency string

const (
	IssuingCardLimitFrequencyPer24HourPeriod  IssuingCardLimitFrequency = "per24HourPeriod"
	IssuingCardLimitFrequencyPer7DayPeriod    IssuingCardLimitFrequency = "per7DayPeriod"
	IssuingCardLimitFrequencyPer30DayPeriod   IssuingCardLimitFrequency = "per30DayPeriod"
	IssuingCardLimitFrequencyPerYearPeriod    IssuingCardLimitFrequency = "perYearPeriod"
	IssuingCardLimitFrequencyAllTime          IssuingCardLimitFrequency = "allTime"
	IssuingCardLimitFrequencyPerAuthorization IssuingCardLimitFrequency = "perAuthorization"
)

// Represents the spending limit and frequency for a card.
//
// The properties Amount, Frequency are required.
type IssuingCardLimitParam struct {
	// The maximum spending amount in cents
	Amount int64 `json:"amount" api:"required"`
	// The frequency at which the spending limit resets
	//
	// Any of "per24HourPeriod", "per7DayPeriod", "per30DayPeriod", "perYearPeriod",
	// "allTime", "perAuthorization".
	Frequency IssuingCardLimitFrequency `json:"frequency,omitzero" api:"required"`
	paramObj
}

func (r IssuingCardLimitParam) MarshalJSON() (data []byte, err error) {
	type shadow IssuingCardLimitParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IssuingCardLimitParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the card
type IssuingCardStatus string

const (
	IssuingCardStatusNotActivated IssuingCardStatus = "notActivated"
	IssuingCardStatusActive       IssuingCardStatus = "active"
	IssuingCardStatusLocked       IssuingCardStatus = "locked"
	IssuingCardStatusCanceled     IssuingCardStatus = "canceled"
)

// The encrypted data for the card
type CardGetSecretsResponse struct {
	// The encrypted CVC
	EncryptedCvc CardGetSecretsResponseEncryptedCvc `json:"encryptedCvc" api:"required"`
	// The encrypted PAN
	EncryptedPan CardGetSecretsResponseEncryptedPan `json:"encryptedPan" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EncryptedCvc respjson.Field
		EncryptedPan respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardGetSecretsResponse) RawJSON() string { return r.JSON.raw }
func (r *CardGetSecretsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The encrypted CVC
type CardGetSecretsResponseEncryptedCvc struct {
	// The encrypted data
	Data string `json:"data" api:"required"`
	// The initialization vector
	Iv string `json:"iv" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Iv          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardGetSecretsResponseEncryptedCvc) RawJSON() string { return r.JSON.raw }
func (r *CardGetSecretsResponseEncryptedCvc) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The encrypted PAN
type CardGetSecretsResponseEncryptedPan struct {
	// The encrypted data
	Data string `json:"data" api:"required"`
	// The initialization vector
	Iv string `json:"iv" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Iv          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardGetSecretsResponseEncryptedPan) RawJSON() string { return r.JSON.raw }
func (r *CardGetSecretsResponseEncryptedPan) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardUpdateParams struct {
	// The billing address associated with the card.
	Billing PhysicalAddressParam `json:"billing,omitzero"`
	// Configuration for the card, such as virtual card art
	Configuration CardUpdateParamsConfiguration `json:"configuration,omitzero"`
	// The limit associated with the card
	Limit IssuingCardLimitParam `json:"limit,omitzero"`
	// The current status of the card
	//
	// Any of "notActivated", "active", "locked", "canceled".
	Status IssuingCardStatus `json:"status,omitzero"`
	paramObj
}

func (r CardUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CardUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the card, such as virtual card art
type CardUpdateParamsConfiguration struct {
	// The virtual card art ID used to customize the card's appearance, if applicable
	VirtualCardArt param.Opt[string] `json:"virtualCardArt,omitzero"`
	paramObj
}

func (r CardUpdateParamsConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow CardUpdateParamsConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardUpdateParamsConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardListParams struct {
	// For corporate cards, the identifier of the company to get cards for
	CompanyID param.Opt[string] `query:"companyId,omitzero" format:"uuid" json:"-"`
	// The ID of the resource after which to start fetching
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// The number of resources to fetch
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The identifier of the user to get cards for
	UserID param.Opt[string] `query:"userId,omitzero" format:"uuid" json:"-"`
	// Filter cards by status
	//
	// Any of "notActivated", "active", "locked", "canceled".
	Status IssuingCardStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CardListParams]'s query parameters as `url.Values`.
func (r CardListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CardGetSecretsParams struct {
	SessionID string `header:"SessionId" api:"required" json:"-"`
	paramObj
}
