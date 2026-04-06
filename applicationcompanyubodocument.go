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

	"github.com/stainless-sdks/rain-hello-world-go/internal/apiform"
	"github.com/stainless-sdks/rain-hello-world-go/internal/requestconfig"
	"github.com/stainless-sdks/rain-hello-world-go/option"
	"github.com/stainless-sdks/rain-hello-world-go/packages/param"
)

// ApplicationCompanyUboDocumentService contains methods and other services that
// help with interacting with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApplicationCompanyUboDocumentService] method instead.
type ApplicationCompanyUboDocumentService struct {
	options []option.RequestOption
}

// NewApplicationCompanyUboDocumentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewApplicationCompanyUboDocumentService(opts ...option.RequestOption) (r ApplicationCompanyUboDocumentService) {
	r = ApplicationCompanyUboDocumentService{}
	r.options = opts
	return
}

// Uploads a document for a company's Ultimate Beneficial Owner (UBO) to support
// the company's corporate application. This endpoint allows for the submission of
// various legal and identification documents.
func (r *ApplicationCompanyUboDocumentService) Upload(ctx context.Context, uboID string, params ApplicationCompanyUboDocumentUploadParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.CompanyID == "" {
		err = errors.New("missing required companyId parameter")
		return err
	}
	if uboID == "" {
		err = errors.New("missing required uboId parameter")
		return err
	}
	path := fmt.Sprintf("applications/company/%s/ubo/%s/document", params.CompanyID, uboID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return err
}

type ApplicationCompanyUboDocumentUploadParams struct {
	CompanyID string `path:"companyId" api:"required" format:"uuid" json:"-"`
	// The actual document file to be uploaded. The document must be in binary format,
	// and the maximum allowed size is 20 MB.
	Document io.Reader `json:"document,omitzero" api:"required" format:"binary"`
	// The country where the document was issued
	Country param.Opt[string] `json:"country,omitzero"`
	// The side of the document being uploaded
	//
	// Any of "front", "back".
	Side ApplicationCompanyUboDocumentUploadParamsSide `json:"side,omitzero"`
	// The type of the document being uploaded
	//
	// Any of "idCard", "passport", "drivers", "residencePermit", "utilityBill",
	// "selfie", "videoSelfie", "profileImage", "idDocPhoto", "agreement", "contract",
	// "driversTranslation", "investorDoc", "vehicleRegistrationCertificate",
	// "incomeSource", "paymentMethod", "bankCard", "covidVaccinationForm", "other".
	Type ApplicationCompanyUboDocumentUploadParamsType `json:"type,omitzero"`
	paramObj
}

func (r ApplicationCompanyUboDocumentUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type ApplicationCompanyUboDocumentUploadParamsSide string

const (
	ApplicationCompanyUboDocumentUploadParamsSideFront ApplicationCompanyUboDocumentUploadParamsSide = "front"
	ApplicationCompanyUboDocumentUploadParamsSideBack  ApplicationCompanyUboDocumentUploadParamsSide = "back"
)

// The type of the document being uploaded
type ApplicationCompanyUboDocumentUploadParamsType string

const (
	ApplicationCompanyUboDocumentUploadParamsTypeIDCard                         ApplicationCompanyUboDocumentUploadParamsType = "idCard"
	ApplicationCompanyUboDocumentUploadParamsTypePassport                       ApplicationCompanyUboDocumentUploadParamsType = "passport"
	ApplicationCompanyUboDocumentUploadParamsTypeDrivers                        ApplicationCompanyUboDocumentUploadParamsType = "drivers"
	ApplicationCompanyUboDocumentUploadParamsTypeResidencePermit                ApplicationCompanyUboDocumentUploadParamsType = "residencePermit"
	ApplicationCompanyUboDocumentUploadParamsTypeUtilityBill                    ApplicationCompanyUboDocumentUploadParamsType = "utilityBill"
	ApplicationCompanyUboDocumentUploadParamsTypeSelfie                         ApplicationCompanyUboDocumentUploadParamsType = "selfie"
	ApplicationCompanyUboDocumentUploadParamsTypeVideoSelfie                    ApplicationCompanyUboDocumentUploadParamsType = "videoSelfie"
	ApplicationCompanyUboDocumentUploadParamsTypeProfileImage                   ApplicationCompanyUboDocumentUploadParamsType = "profileImage"
	ApplicationCompanyUboDocumentUploadParamsTypeIDDocPhoto                     ApplicationCompanyUboDocumentUploadParamsType = "idDocPhoto"
	ApplicationCompanyUboDocumentUploadParamsTypeAgreement                      ApplicationCompanyUboDocumentUploadParamsType = "agreement"
	ApplicationCompanyUboDocumentUploadParamsTypeContract                       ApplicationCompanyUboDocumentUploadParamsType = "contract"
	ApplicationCompanyUboDocumentUploadParamsTypeDriversTranslation             ApplicationCompanyUboDocumentUploadParamsType = "driversTranslation"
	ApplicationCompanyUboDocumentUploadParamsTypeInvestorDoc                    ApplicationCompanyUboDocumentUploadParamsType = "investorDoc"
	ApplicationCompanyUboDocumentUploadParamsTypeVehicleRegistrationCertificate ApplicationCompanyUboDocumentUploadParamsType = "vehicleRegistrationCertificate"
	ApplicationCompanyUboDocumentUploadParamsTypeIncomeSource                   ApplicationCompanyUboDocumentUploadParamsType = "incomeSource"
	ApplicationCompanyUboDocumentUploadParamsTypePaymentMethod                  ApplicationCompanyUboDocumentUploadParamsType = "paymentMethod"
	ApplicationCompanyUboDocumentUploadParamsTypeBankCard                       ApplicationCompanyUboDocumentUploadParamsType = "bankCard"
	ApplicationCompanyUboDocumentUploadParamsTypeCovidVaccinationForm           ApplicationCompanyUboDocumentUploadParamsType = "covidVaccinationForm"
	ApplicationCompanyUboDocumentUploadParamsTypeOther                          ApplicationCompanyUboDocumentUploadParamsType = "other"
)
