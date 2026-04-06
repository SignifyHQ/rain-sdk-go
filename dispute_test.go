// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainhelloworld_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/rain-hello-world-go"
	"github.com/stainless-sdks/rain-hello-world-go/internal/testutil"
	"github.com/stainless-sdks/rain-hello-world-go/option"
)

func TestDisputeGet(t *testing.T) {
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
	_, err := client.Disputes.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *rainhelloworld.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDisputeUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Disputes.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.DisputeUpdateParams{
			Status:       rainhelloworld.DisputeUpdateParamsStatusCanceled,
			TextEvidence: rainhelloworld.String("textEvidence"),
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

func TestDisputeListWithOptionalParams(t *testing.T) {
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
	_, err := client.Disputes.List(context.TODO(), rainhelloworld.DisputeListParams{
		CompanyID:     rainhelloworld.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Cursor:        rainhelloworld.String("cursor"),
		Limit:         rainhelloworld.Int(1),
		TransactionID: rainhelloworld.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		UserID:        rainhelloworld.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *rainhelloworld.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
