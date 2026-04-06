// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainhelloworld

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/rain-hello-world-go/internal/apiform"
	"github.com/stainless-sdks/rain-hello-world-go/internal/apijson"
	"github.com/stainless-sdks/rain-hello-world-go/internal/requestconfig"
	"github.com/stainless-sdks/rain-hello-world-go/option"
	"github.com/stainless-sdks/rain-hello-world-go/packages/param"
)

// ApplicationCompanyUboService contains methods and other services that help with
// interacting with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApplicationCompanyUboService] method instead.
type ApplicationCompanyUboService struct {
	options  []option.RequestOption
	Document ApplicationCompanyUboDocumentService
}

// NewApplicationCompanyUboService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewApplicationCompanyUboService(opts ...option.RequestOption) (r ApplicationCompanyUboService) {
	r = ApplicationCompanyUboService{}
	r.options = opts
	r.Document = NewApplicationCompanyUboDocumentService(opts...)
	return
}

// Updates the application information for a company's Ultimate Beneficial Owner
// (UBO). This allows modification of the UBO's personal details such as name,
// birth date, national ID, and address.
func (r *ApplicationCompanyUboService) Update(ctx context.Context, uboID string, params ApplicationCompanyUboUpdateParams, opts ...option.RequestOption) (res *IssuingCompany, err error) {
	opts = slices.Concat(r.options, opts)
	if params.CompanyID == "" {
		err = errors.New("missing required companyId parameter")
		return nil, err
	}
	if uboID == "" {
		err = errors.New("missing required uboId parameter")
		return nil, err
	}
	path := fmt.Sprintf("applications/company/%s/ubo/%s", params.CompanyID, uboID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// This deprecated endpoint allows the upload of a document for a UBO to support a
// company's corporate application. It is recommended to use the newer endpoint for
// document uploads.
//
// Deprecated: deprecated
func (r *ApplicationCompanyUboService) UploadDocument(ctx context.Context, companyID string, body ApplicationCompanyUboUploadDocumentParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if companyID == "" {
		err = errors.New("missing required companyId parameter")
		return err
	}
	path := fmt.Sprintf("applications/company/%s/ubo/document", companyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

type ApplicationCompanyUboUpdateParams struct {
	CompanyID string `path:"companyId" api:"required" format:"uuid" json:"-"`
	// The UBO's birth date
	BirthDate param.Opt[time.Time] `json:"birthDate,omitzero" format:"date"`
	// The 2-digit country code of the user's national ID issuer
	CountryOfIssue param.Opt[string] `json:"countryOfIssue,omitzero"`
	// The UBO's email address
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// The UBO's first name
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	// The UBO's last name
	LastName param.Opt[string] `json:"lastName,omitzero"`
	// The UBO's national ID number. For the US, this is a 9-digit SSN
	NationalID param.Opt[string] `json:"nationalId,omitzero"`
	// The UBO's address
	Address PhysicalAddressParam `json:"address,omitzero"`
	paramObj
}

func (r ApplicationCompanyUboUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ApplicationCompanyUboUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplicationCompanyUboUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationCompanyUboUploadDocumentParams struct {
	// The actual document file to be uploaded. The document must be in binary format,
	// and the maximum allowed size is 20 MB.
	Document io.Reader `json:"document,omitzero" api:"required" format:"binary"`
	// The UBO's email address
	Email string `json:"email" api:"required"`
	// The country where the document was issued
	Country param.Opt[string] `json:"country,omitzero"`
	// The side of the document being uploaded
	//
	// Any of "front", "back".
	Side ApplicationCompanyUboUploadDocumentParamsSide `json:"side,omitzero"`
	// The type of the document being uploaded
	//
	// Any of "idCard", "passport", "drivers", "residencePermit", "utilityBill",
	// "selfie", "videoSelfie", "profileImage", "idDocPhoto", "agreement", "contract",
	// "driversTranslation", "investorDoc", "vehicleRegistrationCertificate",
	// "incomeSource", "paymentMethod", "bankCard", "covidVaccinationForm", "other".
	Type ApplicationCompanyUboUploadDocumentParamsType `json:"type,omitzero"`
	paramObj
}

func (r ApplicationCompanyUboUploadDocumentParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// The side of the document being uploaded
type ApplicationCompanyUboUploadDocumentParamsSide string

const (
	ApplicationCompanyUboUploadDocumentParamsSideFront ApplicationCompanyUboUploadDocumentParamsSide = "front"
	ApplicationCompanyUboUploadDocumentParamsSideBack  ApplicationCompanyUboUploadDocumentParamsSide = "back"
)

// The type of the document being uploaded
type ApplicationCompanyUboUploadDocumentParamsType string

const (
	ApplicationCompanyUboUploadDocumentParamsTypeIDCard                         ApplicationCompanyUboUploadDocumentParamsType = "idCard"
	ApplicationCompanyUboUploadDocumentParamsTypePassport                       ApplicationCompanyUboUploadDocumentParamsType = "passport"
	ApplicationCompanyUboUploadDocumentParamsTypeDrivers                        ApplicationCompanyUboUploadDocumentParamsType = "drivers"
	ApplicationCompanyUboUploadDocumentParamsTypeResidencePermit                ApplicationCompanyUboUploadDocumentParamsType = "residencePermit"
	ApplicationCompanyUboUploadDocumentParamsTypeUtilityBill                    ApplicationCompanyUboUploadDocumentParamsType = "utilityBill"
	ApplicationCompanyUboUploadDocumentParamsTypeSelfie                         ApplicationCompanyUboUploadDocumentParamsType = "selfie"
	ApplicationCompanyUboUploadDocumentParamsTypeVideoSelfie                    ApplicationCompanyUboUploadDocumentParamsType = "videoSelfie"
	ApplicationCompanyUboUploadDocumentParamsTypeProfileImage                   ApplicationCompanyUboUploadDocumentParamsType = "profileImage"
	ApplicationCompanyUboUploadDocumentParamsTypeIDDocPhoto                     ApplicationCompanyUboUploadDocumentParamsType = "idDocPhoto"
	ApplicationCompanyUboUploadDocumentParamsTypeAgreement                      ApplicationCompanyUboUploadDocumentParamsType = "agreement"
	ApplicationCompanyUboUploadDocumentParamsTypeContract                       ApplicationCompanyUboUploadDocumentParamsType = "contract"
	ApplicationCompanyUboUploadDocumentParamsTypeDriversTranslation             ApplicationCompanyUboUploadDocumentParamsType = "driversTranslation"
	ApplicationCompanyUboUploadDocumentParamsTypeInvestorDoc                    ApplicationCompanyUboUploadDocumentParamsType = "investorDoc"
	ApplicationCompanyUboUploadDocumentParamsTypeVehicleRegistrationCertificate ApplicationCompanyUboUploadDocumentParamsType = "vehicleRegistrationCertificate"
	ApplicationCompanyUboUploadDocumentParamsTypeIncomeSource                   ApplicationCompanyUboUploadDocumentParamsType = "incomeSource"
	ApplicationCompanyUboUploadDocumentParamsTypePaymentMethod                  ApplicationCompanyUboUploadDocumentParamsType = "paymentMethod"
	ApplicationCompanyUboUploadDocumentParamsTypeBankCard                       ApplicationCompanyUboUploadDocumentParamsType = "bankCard"
	ApplicationCompanyUboUploadDocumentParamsTypeCovidVaccinationForm           ApplicationCompanyUboUploadDocumentParamsType = "covidVaccinationForm"
	ApplicationCompanyUboUploadDocumentParamsTypeOther                          ApplicationCompanyUboUploadDocumentParamsType = "other"
)
