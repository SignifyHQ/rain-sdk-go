// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainsdk_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/SignifyHQ/rain-sdk-go"
	"github.com/SignifyHQ/rain-sdk-go/internal/testutil"
	"github.com/SignifyHQ/rain-sdk-go/option"
)

func TestApplicationCompanyUboDocumentUploadWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := rainsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Applications.Company.Ubo.Document.Upload(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainsdk.ApplicationCompanyUboDocumentUploadParams{
			CompanyID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Document:  io.Reader(bytes.NewBuffer([]byte("Example data"))),
			Country:   rainsdk.String("xxx"),
			Side:      rainsdk.ApplicationCompanyUboDocumentUploadParamsSideFront,
			Type:      rainsdk.ApplicationCompanyUboDocumentUploadParamsTypeIDCard,
		},
	)
	if err != nil {
		var apierr *rainsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
