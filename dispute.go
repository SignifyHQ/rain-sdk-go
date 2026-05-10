// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainsdk

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
	"github.com/SignifyHQ/rain-sdk-go/internal/requestconfig"
	"github.com/SignifyHQ/rain-sdk-go/option"
	"github.com/SignifyHQ/rain-sdk-go/packages/param"
	"github.com/SignifyHQ/rain-sdk-go/packages/respjson"
)

// DisputeService contains methods and other services that help with interacting
// with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDisputeService] method instead.
type DisputeService struct {
	options  []option.RequestOption
	Evidence DisputeEvidenceService
}

// NewDisputeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDisputeService(opts ...option.RequestOption) (r DisputeService) {
	r = DisputeService{}
	r.options = opts
	r.Evidence = NewDisputeEvidenceService(opts...)
	return
}

// Retrieve details of a specific dispute using its unique ID.
func (r *DisputeService) Get(ctx context.Context, disputeID string, opts ...option.RequestOption) (res *IssuingDispute, err error) {
	opts = slices.Concat(r.options, opts)
	if disputeID == "" {
		err = errors.New("missing required disputeId parameter")
		return nil, err
	}
	path := fmt.Sprintf("disputes/%s", disputeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update the status or evidence of a dispute, typically to mark it as canceled or
// add new evidence.
func (r *DisputeService) Update(ctx context.Context, disputeID string, body DisputeUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if disputeID == "" {
		err = errors.New("missing required disputeId parameter")
		return err
	}
	path := fmt.Sprintf("disputes/%s", disputeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return err
}

// Retrieve all disputes, optionally filtered by company, user, or transaction ID.
func (r *DisputeService) List(ctx context.Context, query DisputeListParams, opts ...option.RequestOption) (res *[]IssuingDispute, err error) {
	opts = slices.Concat(r.options, opts)
	path := "disputes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Represents a dispute related to an issuing transaction.
type IssuingDispute struct {
	// The dispute's unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// The date and time when the dispute was created
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// The current status of the dispute
	//
	// Any of "pending", "inReview", "accepted", "rejected", "canceled".
	Status IssuingDisputeStatus `json:"status" api:"required"`
	// The transaction's unique identifier
	TransactionID string `json:"transactionId" api:"required" format:"uuid"`
	// The date and time when the dispute was resolved, if applicable
	ResolvedAt time.Time `json:"resolvedAt" format:"date-time"`
	// Textual evidence provided by the parties involved in the dispute
	TextEvidence string `json:"textEvidence"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Status        respjson.Field
		TransactionID respjson.Field
		ResolvedAt    respjson.Field
		TextEvidence  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IssuingDispute) RawJSON() string { return r.JSON.raw }
func (r *IssuingDispute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the dispute
type IssuingDisputeStatus string

const (
	IssuingDisputeStatusPending  IssuingDisputeStatus = "pending"
	IssuingDisputeStatusInReview IssuingDisputeStatus = "inReview"
	IssuingDisputeStatusAccepted IssuingDisputeStatus = "accepted"
	IssuingDisputeStatusRejected IssuingDisputeStatus = "rejected"
	IssuingDisputeStatusCanceled IssuingDisputeStatus = "canceled"
)

type DisputeUpdateParams struct {
	// The textual evidence to add to the dispute
	TextEvidence param.Opt[string] `json:"textEvidence,omitzero"`
	// The new status of the dispute. Can only be set to 'canceled'.
	//
	// Any of "canceled".
	Status DisputeUpdateParamsStatus `json:"status,omitzero"`
	paramObj
}

func (r DisputeUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DisputeUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DisputeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The new status of the dispute. Can only be set to 'canceled'.
type DisputeUpdateParamsStatus string

const (
	DisputeUpdateParamsStatusCanceled DisputeUpdateParamsStatus = "canceled"
)

type DisputeListParams struct {
	// For corporate cards, the identifier of the company to get disputes for
	CompanyID param.Opt[string] `query:"companyId,omitzero" format:"uuid" json:"-"`
	// The ID of the resource after which to start fetching
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// The number of resources to fetch
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The ID of the transaction to get disputes for
	TransactionID param.Opt[string] `query:"transactionId,omitzero" format:"uuid" json:"-"`
	// The ID of the user to get disputes for
	UserID param.Opt[string] `query:"userId,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [DisputeListParams]'s query parameters as `url.Values`.
func (r DisputeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
