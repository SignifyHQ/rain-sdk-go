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
	"github.com/SignifyHQ/rain-sdk-go/shared/constant"
)

// TransactionService contains methods and other services that help with
// interacting with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
type TransactionService struct {
	options []option.RequestOption
	Receipt TransactionReceiptService
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r TransactionService) {
	r = TransactionService{}
	r.options = opts
	r.Receipt = NewTransactionReceiptService(opts...)
	return
}

// This endpoint retrieves a transaction by its unique ID. The transaction
// information returned includes details such as the transaction type, amount, and
// status.
func (r *TransactionService) Get(ctx context.Context, transactionID string, opts ...option.RequestOption) (res *IssuingTransactionUnion, err error) {
	opts = slices.Concat(r.options, opts)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("transactions/%s", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This endpoint allows updating a specific transaction by its ID. You can modify
// the transaction's memo or other editable fields.
func (r *TransactionService) Update(ctx context.Context, transactionID string, body TransactionUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return err
	}
	path := fmt.Sprintf("transactions/%s", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// This endpoint retrieves all transactions associated with corporate cards, users,
// or specific cards.
func (r *TransactionService) List(ctx context.Context, query TransactionListParams, opts ...option.RequestOption) (res *[]IssuingTransactionUnion, err error) {
	opts = slices.Concat(r.options, opts)
	path := "transactions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// This endpoint allows the creation of a dispute for a specific transaction. The
// dispute can include textual evidence to support the claim.
func (r *TransactionService) NewDispute(ctx context.Context, transactionID string, body TransactionNewDisputeParams, opts ...option.RequestOption) (res *IssuingDispute, err error) {
	opts = slices.Concat(r.options, opts)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("transactions/%s/disputes", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// IssuingTransactionUnion contains all possible properties and values from
// [IssuingTransactionSpend], [IssuingTransactionCollateral],
// [IssuingTransactionPayment], [IssuingTransactionFee].
//
// Use the [IssuingTransactionUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type IssuingTransactionUnion struct {
	ID string `json:"id"`
	// This field is from variant [IssuingTransactionSpend].
	Spend IssuingTransactionSpendSpend `json:"spend"`
	// Any of "spend", "collateral", "payment", "fee".
	Type string `json:"type"`
	// This field is from variant [IssuingTransactionCollateral].
	Collateral IssuingTransactionCollateralCollateral `json:"collateral"`
	// This field is from variant [IssuingTransactionPayment].
	Payment IssuingTransactionPaymentPayment `json:"payment"`
	// This field is from variant [IssuingTransactionFee].
	Fee  IssuingTransactionFeeFee `json:"fee"`
	JSON struct {
		ID         respjson.Field
		Spend      respjson.Field
		Type       respjson.Field
		Collateral respjson.Field
		Payment    respjson.Field
		Fee        respjson.Field
		raw        string
	} `json:"-"`
}

// anyIssuingTransaction is implemented by each variant of
// [IssuingTransactionUnion] to add type safety for the return type of
// [IssuingTransactionUnion.AsAny]
type anyIssuingTransaction interface {
	implIssuingTransactionUnion()
}

func (IssuingTransactionSpend) implIssuingTransactionUnion()      {}
func (IssuingTransactionCollateral) implIssuingTransactionUnion() {}
func (IssuingTransactionPayment) implIssuingTransactionUnion()    {}
func (IssuingTransactionFee) implIssuingTransactionUnion()        {}

// Use the following switch statement to find the correct variant
//
//	switch variant := IssuingTransactionUnion.AsAny().(type) {
//	case rainsdk.IssuingTransactionSpend:
//	case rainsdk.IssuingTransactionCollateral:
//	case rainsdk.IssuingTransactionPayment:
//	case rainsdk.IssuingTransactionFee:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u IssuingTransactionUnion) AsAny() anyIssuingTransaction {
	switch u.Type {
	case "spend":
		return u.AsSpend()
	case "collateral":
		return u.AsCollateral()
	case "payment":
		return u.AsPayment()
	case "fee":
		return u.AsFee()
	}
	return nil
}

func (u IssuingTransactionUnion) AsSpend() (v IssuingTransactionSpend) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IssuingTransactionUnion) AsCollateral() (v IssuingTransactionCollateral) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IssuingTransactionUnion) AsPayment() (v IssuingTransactionPayment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IssuingTransactionUnion) AsFee() (v IssuingTransactionFee) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u IssuingTransactionUnion) RawJSON() string { return u.JSON.raw }

func (r *IssuingTransactionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a transaction of type 'spend'. This includes details such as the
// transaction amount, merchant, and the associated user.
type IssuingTransactionSpend struct {
	// The unique identifier of the transaction
	ID string `json:"id" api:"required" format:"uuid"`
	// Details specific to a spend transaction, including merchant, amount, and user
	// information.
	Spend IssuingTransactionSpendSpend `json:"spend" api:"required"`
	// The type of transaction
	Type constant.Spend `json:"type" default:"spend"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Spend       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingTransactionSpend) RawJSON() string { return r.JSON.raw }
func (r *IssuingTransactionSpend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details specific to a spend transaction, including merchant, amount, and user
// information.
type IssuingTransactionSpendSpend struct {
	// The amount of the transaction, in cents
	Amount int64 `json:"amount" api:"required"`
	// The time at which the transaction was authorized
	AuthorizedAt string `json:"authorizedAt" api:"required"`
	// The unique identifier of the card used for the transaction
	CardID string `json:"cardId" api:"required" format:"uuid"`
	// The type of card used for the transaction
	//
	// Any of "physical", "virtual".
	CardType string `json:"cardType" api:"required"`
	// The currency of the transaction
	Currency string `json:"currency" api:"required"`
	// The category of the merchant (e.g., electronics, food)
	MerchantCategory string `json:"merchantCategory" api:"required"`
	// The merchant's category code
	MerchantCategoryCode string `json:"merchantCategoryCode" api:"required"`
	// The name of the merchant where the transaction took place
	MerchantName string `json:"merchantName" api:"required"`
	// Indicates whether a receipt was generated for the transaction
	Receipt bool `json:"receipt" api:"required"`
	// The status of the transaction
	//
	// Any of "pending", "reversed", "declined", "completed".
	Status string `json:"status" api:"required"`
	// The email address of the user who made the transaction
	UserEmail string `json:"userEmail" api:"required"`
	// The first name of the user who made the transaction
	UserFirstName string `json:"userFirstName" api:"required"`
	// The identifier of the user who made the transaction
	UserID string `json:"userId" api:"required" format:"uuid"`
	// The last name of the user who made the transaction
	UserLastName string `json:"userLastName" api:"required"`
	// The method used for authorization of the transaction
	AuthorizationMethod string `json:"authorizationMethod"`
	// The authorized amount of the transaction, in cents
	AuthorizedAmount int64 `json:"authorizedAmount"`
	// The identifier of the company under which the transaction was made
	CompanyID string `json:"companyId" format:"uuid"`
	// The reason why the transaction was declined
	DeclinedReason string `json:"declinedReason"`
	// An enriched category of the merchant, providing further details
	EnrichedMerchantCategory string `json:"enrichedMerchantCategory"`
	// The enriched icon of the merchant
	EnrichedMerchantIcon string `json:"enrichedMerchantIcon"`
	// An enriched name of the merchant, possibly with additional information
	EnrichedMerchantName string `json:"enrichedMerchantName"`
	// The amount of the transaction in local currency, in cents
	LocalAmount int64 `json:"localAmount"`
	// The local currency of the transaction
	LocalCurrency string `json:"localCurrency"`
	// A memo or note associated with the transaction
	Memo string `json:"memo"`
	// The time at which the transaction was posted
	PostedAt string `json:"postedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount                   respjson.Field
		AuthorizedAt             respjson.Field
		CardID                   respjson.Field
		CardType                 respjson.Field
		Currency                 respjson.Field
		MerchantCategory         respjson.Field
		MerchantCategoryCode     respjson.Field
		MerchantName             respjson.Field
		Receipt                  respjson.Field
		Status                   respjson.Field
		UserEmail                respjson.Field
		UserFirstName            respjson.Field
		UserID                   respjson.Field
		UserLastName             respjson.Field
		AuthorizationMethod      respjson.Field
		AuthorizedAmount         respjson.Field
		CompanyID                respjson.Field
		DeclinedReason           respjson.Field
		EnrichedMerchantCategory respjson.Field
		EnrichedMerchantIcon     respjson.Field
		EnrichedMerchantName     respjson.Field
		LocalAmount              respjson.Field
		LocalCurrency            respjson.Field
		Memo                     respjson.Field
		PostedAt                 respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingTransactionSpendSpend) RawJSON() string { return r.JSON.raw }
func (r *IssuingTransactionSpendSpend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a collateral transaction, where a user provides collateral for a
// transaction.
type IssuingTransactionCollateral struct {
	// The unique identifier of the transaction
	ID string `json:"id" api:"required" format:"uuid"`
	// Details of the collateral transaction, including amount, currency, and
	// transaction details.
	Collateral IssuingTransactionCollateralCollateral `json:"collateral" api:"required"`
	// The type of transaction, in this case, a collateral transaction
	Type constant.Collateral `json:"type" default:"collateral"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Collateral  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingTransactionCollateral) RawJSON() string { return r.JSON.raw }
func (r *IssuingTransactionCollateral) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the collateral transaction, including amount, currency, and
// transaction details.
type IssuingTransactionCollateralCollateral struct {
	// The amount of the collateral transaction, in cents
	Amount float64 `json:"amount" api:"required"`
	// The chain ID (base-10 number) that the collateral transaction is on
	ChainID int64 `json:"chainId" api:"required"`
	// The currency of the collateral transaction
	Currency string `json:"currency" api:"required"`
	// The hash of the collateral transaction
	TransactionHash string `json:"transactionHash" api:"required"`
	// The wallet address the collateral was added from
	WalletAddress string `json:"walletAddress" api:"required"`
	// The identifier of the company under which the collateral transaction was made
	CompanyID string `json:"companyId" format:"uuid"`
	// A memo or note associated with the collateral transaction
	Memo string `json:"memo"`
	// The time at which the collateral transaction was posted
	PostedAt time.Time `json:"postedAt" format:"date-time"`
	// The identifier of the user who provided the collateral
	UserID string `json:"userId" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount          respjson.Field
		ChainID         respjson.Field
		Currency        respjson.Field
		TransactionHash respjson.Field
		WalletAddress   respjson.Field
		CompanyID       respjson.Field
		Memo            respjson.Field
		PostedAt        respjson.Field
		UserID          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingTransactionCollateralCollateral) RawJSON() string { return r.JSON.raw }
func (r *IssuingTransactionCollateralCollateral) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a payment transaction, where a payment is made for a particular
// service or product.
type IssuingTransactionPayment struct {
	// The unique identifier of the payment transaction
	ID string `json:"id" api:"required" format:"uuid"`
	// Details of the payment transaction, including amount, currency, and status.
	Payment IssuingTransactionPaymentPayment `json:"payment" api:"required"`
	// The type of transaction
	Type constant.Payment `json:"type" default:"payment"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Payment     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingTransactionPayment) RawJSON() string { return r.JSON.raw }
func (r *IssuingTransactionPayment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the payment transaction, including amount, currency, and status.
type IssuingTransactionPaymentPayment struct {
	// The amount of the transaction, in cents
	Amount int64 `json:"amount" api:"required"`
	// The currency of the transaction
	Currency string `json:"currency" api:"required"`
	// The status of the transaction
	//
	// Any of "pending", "completed".
	Status string `json:"status" api:"required"`
	// The chain ID (base-10 number) that the payment transaction is on
	ChainID int64 `json:"chainId"`
	// The identifier of the company under which the payment transaction was made
	CompanyID string `json:"companyId" format:"uuid"`
	// The memo or note associated with the transaction
	Memo string `json:"memo"`
	// The time at which the payment transaction was posted
	PostedAt string `json:"postedAt"`
	// The hash of the payment transaction
	TransactionHash string `json:"transactionHash"`
	// The identifier of the user who made the payment
	UserID string `json:"userId" format:"uuid"`
	// The wallet address from which the payment was made
	WalletAddress string `json:"walletAddress"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount          respjson.Field
		Currency        respjson.Field
		Status          respjson.Field
		ChainID         respjson.Field
		CompanyID       respjson.Field
		Memo            respjson.Field
		PostedAt        respjson.Field
		TransactionHash respjson.Field
		UserID          respjson.Field
		WalletAddress   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingTransactionPaymentPayment) RawJSON() string { return r.JSON.raw }
func (r *IssuingTransactionPaymentPayment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a fee transaction, where a fee is charged for a service or product.
type IssuingTransactionFee struct {
	// The identifier of the fee transaction
	ID string `json:"id" api:"required" format:"uuid"`
	// Details of the fee transaction, including amount, description, and status.
	Fee IssuingTransactionFeeFee `json:"fee" api:"required"`
	// The type of transaction, in this case, a fee transaction
	Type constant.Fee `json:"type" default:"fee"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Fee         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingTransactionFee) RawJSON() string { return r.JSON.raw }
func (r *IssuingTransactionFee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the fee transaction, including amount, description, and status.
type IssuingTransactionFeeFee struct {
	// The amount of the fee, in cents
	Amount int64 `json:"amount" api:"required"`
	// The identifier of the company to which the fee was charged
	CompanyID string `json:"companyId" format:"uuid"`
	// The description of the fee
	Description string `json:"description"`
	// The time at which the fee was posted
	PostedAt time.Time `json:"postedAt" format:"date-time"`
	// The identifier of the user to whom the fee was charged
	UserID string `json:"userId" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		CompanyID   respjson.Field
		Description respjson.Field
		PostedAt    respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingTransactionFeeFee) RawJSON() string { return r.JSON.raw }
func (r *IssuingTransactionFeeFee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionUpdateParams struct {
	// A memo or note to attach to the transaction, providing additional information or
	// context
	Memo param.Opt[string] `json:"memo,omitzero"`
	paramObj
}

func (r TransactionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TransactionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionListParams struct {
	// Filter transactions authorized after a specific date
	AuthorizedAfter param.Opt[time.Time] `query:"authorizedAfter,omitzero" format:"date-time" json:"-"`
	// Filter transactions authorized before a specific date
	AuthorizedBefore param.Opt[time.Time] `query:"authorizedBefore,omitzero" format:"date-time" json:"-"`
	// The ID of the card to get transactions for
	CardID param.Opt[string] `query:"cardId,omitzero" format:"uuid" json:"-"`
	// For corporate cards, the identifier of the company to get transactions for
	CompanyID param.Opt[string] `query:"companyId,omitzero" format:"uuid" json:"-"`
	// The ID of the resource after which to start fetching
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// The number of resources to fetch
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter transactions posted after a specific date
	PostedAfter param.Opt[time.Time] `query:"postedAfter,omitzero" format:"date-time" json:"-"`
	// Filter transactions posted before a specific date
	PostedBefore param.Opt[time.Time] `query:"postedBefore,omitzero" format:"date-time" json:"-"`
	// The hash of the transaction to retrieve
	TransactionHash param.Opt[string] `query:"transactionHash,omitzero" json:"-"`
	// The ID of the user to get transactions for
	UserID param.Opt[string] `query:"userId,omitzero" format:"uuid" json:"-"`
	// The type of transactions to retrieve
	//
	// Any of "spend", "collateral", "payment", "fee".
	Type []string `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TransactionListParams]'s query parameters as `url.Values`.
func (r TransactionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TransactionNewDisputeParams struct {
	// The textual evidence that supports the dispute
	TextEvidence param.Opt[string] `json:"textEvidence,omitzero"`
	paramObj
}

func (r TransactionNewDisputeParams) MarshalJSON() (data []byte, err error) {
	type shadow TransactionNewDisputeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransactionNewDisputeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
