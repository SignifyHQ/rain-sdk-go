// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainhelloworld

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/rain-hello-world-go/internal/apijson"
	"github.com/stainless-sdks/rain-hello-world-go/internal/requestconfig"
	"github.com/stainless-sdks/rain-hello-world-go/option"
	"github.com/stainless-sdks/rain-hello-world-go/packages/param"
	"github.com/stainless-sdks/rain-hello-world-go/packages/respjson"
)

// PaymentService contains methods and other services that help with interacting
// with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPaymentService] method instead.
type PaymentService struct {
	options []option.RequestOption
}

// NewPaymentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPaymentService(opts ...option.RequestOption) (r PaymentService) {
	r = PaymentService{}
	r.options = opts
	return
}

// This endpoint initiates a payment for an authorized user tenant. The request
// includes the amount to be transferred and the wallet address to send the payment
// from.
func (r *PaymentService) Initiate(ctx context.Context, body PaymentInitiateParams, opts ...option.RequestOption) (res *PaymentInitiateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type PaymentInitiateResponse struct {
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
func (r PaymentInitiateResponse) RawJSON() string { return r.JSON.raw }
func (r *PaymentInitiateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PaymentInitiateParams struct {
	// The amount of the payment, in cents
	Amount int64 `json:"amount" api:"required"`
	// The wallet address the payment is being sent from
	WalletAddress string `json:"walletAddress" api:"required"`
	// The chain ID (base-10 number) that the payment transaction is on
	ChainID param.Opt[int64] `json:"chainId,omitzero"`
	paramObj
}

func (r PaymentInitiateParams) MarshalJSON() (data []byte, err error) {
	type shadow PaymentInitiateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PaymentInitiateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
