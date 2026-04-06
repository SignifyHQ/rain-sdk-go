// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainhelloworld

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/rain-hello-world-go/internal/apijson"
	"github.com/stainless-sdks/rain-hello-world-go/internal/requestconfig"
	"github.com/stainless-sdks/rain-hello-world-go/option"
	"github.com/stainless-sdks/rain-hello-world-go/packages/param"
	"github.com/stainless-sdks/rain-hello-world-go/packages/respjson"
)

// KeyService contains methods and other services that help with interacting with
// the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKeyService] method instead.
type KeyService struct {
	options []option.RequestOption
}

// NewKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewKeyService(opts ...option.RequestOption) (r KeyService) {
	r = KeyService{}
	r.options = opts
	return
}

// This endpoint allows for the creation of a new key with a specified name and
// expiration time. The key is used for various security operations within the
// system.
func (r *KeyService) New(ctx context.Context, body KeyNewParams, opts ...option.RequestOption) (res *KeyNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint allows for the deletion of a specific key using its unique ID.
func (r *KeyService) Delete(ctx context.Context, keyID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if keyID == "" {
		err = errors.New("missing required keyId parameter")
		return err
	}
	path := fmt.Sprintf("keys/%s", keyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Represents a unique key used for issuing, typically used for authentication or
// encryption purposes.
type KeyNewResponse struct {
	// The key's unique identifier
	ID string `json:"id" api:"required" format:"uuid"`
	// The expiration date and time of the key
	ExpiresAt time.Time `json:"expiresAt" api:"required" format:"date-time"`
	// The actual key used for the issuing process
	Key string `json:"key" api:"required"`
	// The name assigned to the key
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExpiresAt   respjson.Field
		Key         respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KeyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *KeyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KeyNewParams struct {
	// The expiration timestamp of the key
	ExpiresAt string `json:"expiresAt" api:"required"`
	// The name of the key being created
	Name string `json:"name" api:"required"`
	paramObj
}

func (r KeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow KeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
