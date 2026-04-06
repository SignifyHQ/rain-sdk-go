// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainhelloworld

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/rain-hello-world-go/internal/apiquery"
	"github.com/stainless-sdks/rain-hello-world-go/internal/requestconfig"
	"github.com/stainless-sdks/rain-hello-world-go/option"
	"github.com/stainless-sdks/rain-hello-world-go/packages/param"
)

// UserSignatureService contains methods and other services that help with
// interacting with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserSignatureService] method instead.
type UserSignatureService struct {
	options []option.RequestOption
}

// NewUserSignatureService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserSignatureService(opts ...option.RequestOption) (r UserSignatureService) {
	r = UserSignatureService{}
	r.options = opts
	return
}

// This endpoint retrieves the payment signature for a specific user. The signature
// is used to authorize a payment transaction for the user.
func (r *UserSignatureService) GetPaymentSignature(ctx context.Context, userID string, query UserSignatureGetPaymentSignatureParams, opts ...option.RequestOption) (res *IssuingSignatureUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/%s/signatures/payments", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// This endpoint retrieves the withdrawal signature for a specific user. The
// signature is required to authorize a withdrawal transaction for the user.
func (r *UserSignatureService) GetWithdrawalSignature(ctx context.Context, userID string, query UserSignatureGetWithdrawalSignatureParams, opts ...option.RequestOption) (res *IssuingSignatureUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if userID == "" {
		err = errors.New("missing required userId parameter")
		return nil, err
	}
	path := fmt.Sprintf("users/%s/signatures/withdrawals", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type UserSignatureGetPaymentSignatureParams struct {
	// The address of the token that the payment should be made in, as a hex string for
	// EVM or base58 string for Solana
	Token string `query:"token" api:"required" json:"-"`
	// The address of the admin making the payment, as a hex string for EVM or base58
	// string for Solana
	AdminAddress string `query:"adminAddress" api:"required" json:"-"`
	// The amount of token that is being paid
	Amount string `query:"amount" api:"required" json:"-"`
	// The chain ID (base-10 number) that the smart contract is deployed on
	ChainID param.Opt[int64] `query:"chainId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserSignatureGetPaymentSignatureParams]'s query parameters
// as `url.Values`.
func (r UserSignatureGetPaymentSignatureParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserSignatureGetWithdrawalSignatureParams struct {
	// The address of the token that the withdrawal should be made in, as a hex string
	// for EVM or base58 string for Solana
	Token string `query:"token" api:"required" json:"-"`
	// The address of the admin making the payment, as a hex string for EVM or base58
	// string for Solana
	AdminAddress string `query:"adminAddress" api:"required" json:"-"`
	// The amount of token that is being withdrawn
	Amount string `query:"amount" api:"required" json:"-"`
	// The address the withdrawal should be sent to, as a hex string for EVM or base58
	// string for Solana
	RecipientAddress string `query:"recipientAddress" api:"required" json:"-"`
	// The chain ID (base-10 number) that the smart contract is deployed on
	ChainID param.Opt[int64] `query:"chainId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserSignatureGetWithdrawalSignatureParams]'s query
// parameters as `url.Values`.
func (r UserSignatureGetWithdrawalSignatureParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
