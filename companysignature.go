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
	"time"

	"github.com/SignifyHQ/rain-sdk-go/internal/apijson"
	"github.com/SignifyHQ/rain-sdk-go/internal/apiquery"
	"github.com/SignifyHQ/rain-sdk-go/internal/requestconfig"
	"github.com/SignifyHQ/rain-sdk-go/option"
	"github.com/SignifyHQ/rain-sdk-go/packages/param"
	"github.com/SignifyHQ/rain-sdk-go/packages/respjson"
)

// CompanySignatureService contains methods and other services that help with
// interacting with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCompanySignatureService] method instead.
type CompanySignatureService struct {
	options []option.RequestOption
}

// NewCompanySignatureService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCompanySignatureService(opts ...option.RequestOption) (r CompanySignatureService) {
	r = CompanySignatureService{}
	r.options = opts
	return
}

// Retrieve the payment signature for a company, which is required for completing
// payment transactions.
func (r *CompanySignatureService) GetPaymentSignature(ctx context.Context, companyID string, query CompanySignatureGetPaymentSignatureParams, opts ...option.RequestOption) (res *IssuingSignatureUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if companyID == "" {
		err = errors.New("missing required companyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("companies/%s/signatures/payments", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve the withdrawal signature for a company, which is required for
// processing withdrawal requests.
func (r *CompanySignatureService) GetWithdrawalSignature(ctx context.Context, companyID string, query CompanySignatureGetWithdrawalSignatureParams, opts ...option.RequestOption) (res *IssuingSignatureUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if companyID == "" {
		err = errors.New("missing required companyId parameter")
		return nil, err
	}
	path := fmt.Sprintf("companies/%s/signatures/withdrawals", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// IssuingSignatureUnion contains all possible properties and values from
// [IssuingSignatureIfSignatureIsPending], [IssuingSignatureIfSignatureIsReady].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type IssuingSignatureUnion struct {
	// This field is from variant [IssuingSignatureIfSignatureIsPending].
	RetryAfter int64  `json:"retryAfter"`
	Status     string `json:"status"`
	// This field is from variant [IssuingSignatureIfSignatureIsReady].
	Signature IssuingSignatureIfSignatureIsReadySignature `json:"signature"`
	// This field is from variant [IssuingSignatureIfSignatureIsReady].
	ExpiresAt time.Time `json:"expiresAt"`
	JSON      struct {
		RetryAfter respjson.Field
		Status     respjson.Field
		Signature  respjson.Field
		ExpiresAt  respjson.Field
		raw        string
	} `json:"-"`
}

func (u IssuingSignatureUnion) AsIfSignatureIsPending() (v IssuingSignatureIfSignatureIsPending) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IssuingSignatureUnion) AsIfSignatureIsReady() (v IssuingSignatureIfSignatureIsReady) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u IssuingSignatureUnion) RawJSON() string { return u.JSON.raw }

func (r *IssuingSignatureUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the signature is pending and provides the time after which a retry is
// possible.
type IssuingSignatureIfSignatureIsPending struct {
	// The number of seconds after which the signature can be retried
	RetryAfter int64 `json:"retryAfter" api:"required"`
	// The status of the signature
	//
	// Any of "pending".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RetryAfter  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingSignatureIfSignatureIsPending) RawJSON() string { return r.JSON.raw }
func (r *IssuingSignatureIfSignatureIsPending) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates that the signature is ready and includes the signature data and
// expiration time.
type IssuingSignatureIfSignatureIsReady struct {
	Signature IssuingSignatureIfSignatureIsReadySignature `json:"signature" api:"required"`
	// The status of the signature
	//
	// Any of "ready".
	Status string `json:"status" api:"required"`
	// The time at which the signature will expire
	ExpiresAt time.Time `json:"expiresAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Signature   respjson.Field
		Status      respjson.Field
		ExpiresAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingSignatureIfSignatureIsReady) RawJSON() string { return r.JSON.raw }
func (r *IssuingSignatureIfSignatureIsReady) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IssuingSignatureIfSignatureIsReadySignature struct {
	// The actual signature data
	Data string `json:"data" api:"required"`
	// The salt used to generate the signature
	Salt string `json:"salt" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Salt        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingSignatureIfSignatureIsReadySignature) RawJSON() string { return r.JSON.raw }
func (r *IssuingSignatureIfSignatureIsReadySignature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanySignatureGetPaymentSignatureParams struct {
	// The address of the token that the payment should be made in (as a hex string for
	// EVM or base58 string for Solana)
	Token string `query:"token" api:"required" json:"-"`
	// The address of the admin making the payment, as a hex string for EVM or base58
	// string for Solana
	AdminAddress string `query:"adminAddress" api:"required" json:"-"`
	// The amount of tokens that are being paid
	Amount string `query:"amount" api:"required" json:"-"`
	// The chain ID (base-10 number) that the smart contract is deployed on
	ChainID param.Opt[int64] `query:"chainId,omitzero" json:"-"`
	// Indicates whether the amount is in the asset's native units. If false, the
	// amount is expressed in cents.
	IsAmountNative param.Opt[bool] `query:"isAmountNative,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CompanySignatureGetPaymentSignatureParams]'s query
// parameters as `url.Values`.
func (r CompanySignatureGetPaymentSignatureParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CompanySignatureGetWithdrawalSignatureParams struct {
	// The address of the token that the withdrawal should be made in, as a hex string
	// for EVM or base58 string for Solana
	Token string `query:"token" api:"required" json:"-"`
	// The address of the admin making the payment, as a hex string for EVM or base58
	// string for Solana
	AdminAddress string `query:"adminAddress" api:"required" json:"-"`
	// The amount of tokens being withdrawn
	Amount string `query:"amount" api:"required" json:"-"`
	// The address the withdrawal should be sent to, as a hex string for EVM or base58
	// string for Solana
	RecipientAddress string `query:"recipientAddress" api:"required" json:"-"`
	// The chain ID (base-10 number) that the smart contract is deployed on
	ChainID param.Opt[int64] `query:"chainId,omitzero" json:"-"`
	// Indicates whether the amount is in the asset's native units. If false, the
	// amount is expressed in cents.
	IsAmountNative param.Opt[bool] `query:"isAmountNative,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CompanySignatureGetWithdrawalSignatureParams]'s query
// parameters as `url.Values`.
func (r CompanySignatureGetWithdrawalSignatureParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
