// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainsdk

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

// TransactionReceiptService contains methods and other services that help with
// interacting with the rain API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionReceiptService] method instead.
type TransactionReceiptService struct {
	options []option.RequestOption
}

// NewTransactionReceiptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransactionReceiptService(opts ...option.RequestOption) (r TransactionReceiptService) {
	r = TransactionReceiptService{}
	r.options = opts
	return
}

// This endpoint retrieves the receipt for a specific transaction. The receipt is
// returned as a binary file, typically in PDF or similar format.
func (r *TransactionReceiptService) Get(ctx context.Context, transactionID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("transactions/%s/receipt", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This endpoint allows you to upload a receipt for a specific transaction. The
// receipt is provided as a binary file, typically in PDF format.
func (r *TransactionReceiptService) Upload(ctx context.Context, transactionID string, body TransactionReceiptUploadParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return err
	}
	path := fmt.Sprintf("transactions/%s/receipt", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, nil, opts...)
	return err
}

type TransactionReceiptUploadParams struct {
	// The receipt file to upload, typically in PDF or other binary formats
	Receipt io.Reader `json:"receipt,omitzero" api:"required" format:"binary"`
	paramObj
}

func (r TransactionReceiptUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
