// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainsdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/SignifyHQ/rain-sdk-go/internal/apijson"
	"github.com/SignifyHQ/rain-sdk-go/internal/requestconfig"
	"github.com/SignifyHQ/rain-sdk-go/option"
	"github.com/SignifyHQ/rain-sdk-go/packages/param"
	"github.com/SignifyHQ/rain-sdk-go/packages/respjson"
)

// CardPinService contains methods and other services that help with interacting
// with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCardPinService] method instead.
type CardPinService struct {
	options []option.RequestOption
}

// NewCardPinService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCardPinService(opts ...option.RequestOption) (r CardPinService) {
	r = CardPinService{}
	r.options = opts
	return
}

// Retrieve the encrypted PIN for a specific card
func (r *CardPinService) Get(ctx context.Context, cardID string, query CardPinGetParams, opts ...option.RequestOption) (res *CardPinGetResponse, err error) {
	if !param.IsOmitted(query.SessionID) {
		opts = append(opts, option.WithHeader("SessionId", fmt.Sprintf("%v", query.SessionID)))
	}
	opts = slices.Concat(r.options, opts)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return nil, err
	}
	path := fmt.Sprintf("cards/%s/pin", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates the PIN of a specific card by setting the encrypted PIN
func (r *CardPinService) Update(ctx context.Context, cardID string, params CardPinUpdateParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.SessionID) {
		opts = append(opts, option.WithHeader("SessionId", fmt.Sprintf("%v", params.SessionID)))
	}
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if cardID == "" {
		err = errors.New("missing required cardId parameter")
		return err
	}
	path := fmt.Sprintf("cards/%s/pin", cardID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return err
}

// The encrypted pin
type CardPinGetResponse struct {
	// The encrypted pin
	EncryptedPin CardPinGetResponseEncryptedPin `json:"encryptedPin" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EncryptedPin respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CardPinGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CardPinGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The encrypted pin
type CardPinGetResponseEncryptedPin struct {
	// The encrypted pin in base64
	Data string `json:"data" api:"required"`
	// The initialization vector in base64
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
func (r CardPinGetResponseEncryptedPin) RawJSON() string { return r.JSON.raw }
func (r *CardPinGetResponseEncryptedPin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CardPinGetParams struct {
	SessionID string `header:"SessionId" api:"required" json:"-"`
	paramObj
}

type CardPinUpdateParams struct {
	// The encrypted pin
	EncryptedPin CardPinUpdateParamsEncryptedPin `json:"encryptedPin,omitzero" api:"required"`
	SessionID    string                          `header:"SessionId" api:"required" json:"-"`
	paramObj
}

func (r CardPinUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CardPinUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardPinUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The encrypted pin
//
// The properties Data, Iv are required.
type CardPinUpdateParamsEncryptedPin struct {
	// The encrypted PIN data
	Data string `json:"data" api:"required"`
	// The initialization vector for encryption
	Iv string `json:"iv" api:"required"`
	paramObj
}

func (r CardPinUpdateParamsEncryptedPin) MarshalJSON() (data []byte, err error) {
	type shadow CardPinUpdateParamsEncryptedPin
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CardPinUpdateParamsEncryptedPin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
