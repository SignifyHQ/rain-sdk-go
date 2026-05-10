// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainsdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/SignifyHQ/rain-sdk-go"
	"github.com/SignifyHQ/rain-sdk-go/internal/testutil"
	"github.com/SignifyHQ/rain-sdk-go/option"
)

func TestBalanceGet(t *testing.T) {
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
	_, err := client.Balances.Get(context.TODO())
	if err != nil {
		var apierr *rainsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
