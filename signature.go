// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainhelloworld

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/SignifyHQ/rain-sdk-go/internal/apiquery"
	"github.com/SignifyHQ/rain-sdk-go/internal/requestconfig"
	"github.com/SignifyHQ/rain-sdk-go/option"
	"github.com/SignifyHQ/rain-sdk-go/packages/param"
)

// SignatureService contains methods and other services that help with interacting
// with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSignatureService] method instead.
type SignatureService struct {
	options []option.RequestOption
}

// NewSignatureService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSignatureService(opts ...option.RequestOption) (r SignatureService) {
	r = SignatureService{}
	r.options = opts
	return
}

// This endpoint retrieves the payment signature for an authorized user tenant. The
// signature is used to authorize a payment transaction on a blockchain.
func (r *SignatureService) GetPaymentSignature(ctx context.Context, query SignatureGetPaymentSignatureParams, opts ...option.RequestOption) (res *IssuingSignatureUnion, err error) {
	opts = slices.Concat(r.options, opts)
	path := "signatures/payments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// This endpoint retrieves the withdrawal signature for an authorized user tenant.
// The signature is used to authorize a withdrawal transaction on a blockchain.
func (r *SignatureService) GetWithdrawalSignature(ctx context.Context, query SignatureGetWithdrawalSignatureParams, opts ...option.RequestOption) (res *IssuingSignatureUnion, err error) {
	opts = slices.Concat(r.options, opts)
	path := "signatures/withdrawals"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type SignatureGetPaymentSignatureParams struct {
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

// URLQuery serializes [SignatureGetPaymentSignatureParams]'s query parameters as
// `url.Values`.
func (r SignatureGetPaymentSignatureParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SignatureGetWithdrawalSignatureParams struct {
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

// URLQuery serializes [SignatureGetWithdrawalSignatureParams]'s query parameters
// as `url.Values`.
func (r SignatureGetWithdrawalSignatureParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
