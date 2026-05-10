// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainsdk

import (
	"context"
	"net/http"
	"slices"

	"github.com/SignifyHQ/rain-sdk-go/internal/apijson"
	"github.com/SignifyHQ/rain-sdk-go/internal/requestconfig"
	"github.com/SignifyHQ/rain-sdk-go/option"
	"github.com/SignifyHQ/rain-sdk-go/packages/respjson"
)

// BalanceService contains methods and other services that help with interacting
// with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBalanceService] method instead.
type BalanceService struct {
	options []option.RequestOption
}

// NewBalanceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBalanceService(opts ...option.RequestOption) (r BalanceService) {
	r = BalanceService{}
	r.options = opts
	return
}

// Retrieves the credit balances for an authorized user tenant. This includes
// details such as credit limit, pending charges, posted charges, balance due, and
// spending power.
func (r *BalanceService) Get(ctx context.Context, opts ...option.RequestOption) (res *BalanceGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "balances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type BalanceGetResponse struct {
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
func (r BalanceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BalanceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
