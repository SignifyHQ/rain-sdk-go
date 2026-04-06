// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainhelloworld_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/rain-hello-world-go"
	"github.com/stainless-sdks/rain-hello-world-go/internal/testutil"
	"github.com/stainless-sdks/rain-hello-world-go/option"
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
	client := rainhelloworld.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Applications.Company.Ubo.Document.Upload(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.ApplicationCompanyUboDocumentUploadParams{
			CompanyID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Document:  io.Reader(bytes.NewBuffer([]byte("Example data"))),
			Country:   rainhelloworld.String("xxx"),
			Side:      rainhelloworld.ApplicationCompanyUboDocumentUploadParamsSideFront,
			Type:      rainhelloworld.ApplicationCompanyUboDocumentUploadParamsTypeIDCard,
		},
	)
	if err != nil {
		var apierr *rainhelloworld.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
