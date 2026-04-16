// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainhelloworld_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/SignifyHQ/rain-sdk-go"
	"github.com/SignifyHQ/rain-sdk-go/internal/testutil"
	"github.com/SignifyHQ/rain-sdk-go/option"
)

func TestApplicationCompanyUboUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Applications.Company.Ubo.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.ApplicationCompanyUboUpdateParams{
			CompanyID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Address: rainhelloworld.PhysicalAddressParam{
				City:        "city",
				Country:     "country",
				CountryCode: "xx",
				Line1:       "line1",
				PostalCode:  "postalCode",
				Region:      "region",
				Line2:       rainhelloworld.String("line2"),
			},
			BirthDate:      rainhelloworld.Time(time.Now()),
			CountryOfIssue: rainhelloworld.String("countryOfIssue"),
			Email:          rainhelloworld.String("dev@stainless.com"),
			FirstName:      rainhelloworld.String("firstName"),
			LastName:       rainhelloworld.String("lastName"),
			NationalID:     rainhelloworld.String("nationalId"),
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

func TestApplicationCompanyUboUploadDocumentWithOptionalParams(t *testing.T) {
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
	err := client.Applications.Company.Ubo.UploadDocument(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.ApplicationCompanyUboUploadDocumentParams{
			Document: io.Reader(bytes.NewBuffer([]byte("Example data"))),
			Email:    "email",
			Country:  rainhelloworld.String("xxx"),
			Side:     rainhelloworld.ApplicationCompanyUboUploadDocumentParamsSideFront,
			Type:     rainhelloworld.ApplicationCompanyUboUploadDocumentParamsTypeIDCard,
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
