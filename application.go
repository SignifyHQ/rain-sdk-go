// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainhelloworld

import (
	"github.com/SignifyHQ/rain-sdk-go/option"
)

// ApplicationService contains methods and other services that help with
// interacting with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApplicationService] method instead.
type ApplicationService struct {
	options []option.RequestOption
	Company ApplicationCompanyService
	User    ApplicationUserService
}

// NewApplicationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewApplicationService(opts ...option.RequestOption) (r ApplicationService) {
	r = ApplicationService{}
	r.options = opts
	r.Company = NewApplicationCompanyService(opts...)
	r.User = NewApplicationUserService(opts...)
	return
}
