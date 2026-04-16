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

	"github.com/SignifyHQ/rain-sdk-go/internal/apiform"
	"github.com/SignifyHQ/rain-sdk-go/internal/requestconfig"
	"github.com/SignifyHQ/rain-sdk-go/option"
)

// DisputeEvidenceService contains methods and other services that help with
// interacting with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDisputeEvidenceService] method instead.
type DisputeEvidenceService struct {
	options []option.RequestOption
}

// NewDisputeEvidenceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDisputeEvidenceService(opts ...option.RequestOption) (r DisputeEvidenceService) {
	r = DisputeEvidenceService{}
	r.options = opts
	return
}

// Retrieve the file evidence associated with a dispute.
func (r *DisputeEvidenceService) List(ctx context.Context, disputeID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if disputeID == "" {
		err = errors.New("missing required disputeId parameter")
		return nil, err
	}
	path := fmt.Sprintf("disputes/%s/evidence", disputeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Upload a file that will serve as evidence for a dispute.
func (r *DisputeEvidenceService) Upload(ctx context.Context, disputeID string, body DisputeEvidenceUploadParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if disputeID == "" {
		err = errors.New("missing required disputeId parameter")
		return err
	}
	path := fmt.Sprintf("disputes/%s/evidence", disputeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

type DisputeEvidenceUploadParams struct {
	// The evidence to upload
	Evidence io.Reader `json:"evidence,omitzero" api:"required" format:"binary"`
	// The name of the evidence
	Name string `json:"name" api:"required"`
	// The type of evidence
	Type string `json:"type" api:"required"`
	paramObj
}

func (r DisputeEvidenceUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
