// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainhelloworld_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/SignifyHQ/rain-sdk-go"
	"github.com/SignifyHQ/rain-sdk-go/internal/testutil"
	"github.com/SignifyHQ/rain-sdk-go/option"
)

func TestCardGet(t *testing.T) {
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
	_, err := client.Cards.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *rainhelloworld.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Cards.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.CardUpdateParams{
			Billing: rainhelloworld.PhysicalAddressParam{
				City:        "city",
				Country:     "country",
				CountryCode: "xx",
				Line1:       "line1",
				PostalCode:  "postalCode",
				Region:      "region",
				Line2:       rainhelloworld.String("line2"),
			},
			Configuration: rainhelloworld.CardUpdateParamsConfiguration{
				VirtualCardArt: rainhelloworld.String("virtualCardArt"),
			},
			Limit: rainhelloworld.IssuingCardLimitParam{
				Amount:    0,
				Frequency: rainhelloworld.IssuingCardLimitFrequencyPer24HourPeriod,
			},
			Status: rainhelloworld.IssuingCardStatusNotActivated,
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

func TestCardListWithOptionalParams(t *testing.T) {
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
	_, err := client.Cards.List(context.TODO(), rainhelloworld.CardListParams{
		CompanyID: rainhelloworld.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Cursor:    rainhelloworld.String("cursor"),
		Limit:     rainhelloworld.Int(1),
		Status:    rainhelloworld.IssuingCardStatusNotActivated,
		UserID:    rainhelloworld.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *rainhelloworld.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardGetSecrets(t *testing.T) {
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
	_, err := client.Cards.GetSecrets(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.CardGetSecretsParams{
			SessionID: "x",
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
