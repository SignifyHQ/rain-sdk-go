// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainhelloworld_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/rain-hello-world-go"
	"github.com/stainless-sdks/rain-hello-world-go/internal/testutil"
	"github.com/stainless-sdks/rain-hello-world-go/option"
)

func TestUsage(t *testing.T) {
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
	issuingChargeCreateResponse, err := client.Companies.Charge(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.CompanyChargeParams{
			IssuingChargeCreateBody: rainhelloworld.IssuingChargeCreateBodyParam{
				Amount:      1,
				Description: "description",
			},
		},
	)
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", issuingChargeCreateResponse.ID)
}
