// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainhelloworld

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/rain-hello-world-go/internal/requestconfig"
	"github.com/stainless-sdks/rain-hello-world-go/option"
)

// ContractService contains methods and other services that help with interacting
// with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContractService] method instead.
type ContractService struct {
	options []option.RequestOption
}

// NewContractService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContractService(opts ...option.RequestOption) (r ContractService) {
	r = ContractService{}
	r.options = opts
	return
}

// Retrieve the smart contract information for a specific authorized user tenant.
func (r *ContractService) List(ctx context.Context, opts ...option.RequestOption) (res *[]IssuingContract, err error) {
	opts = slices.Concat(r.options, opts)
	path := "contracts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}
