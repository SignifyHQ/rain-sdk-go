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

func TestUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.New(context.TODO(), rainhelloworld.UserNewParams{
		Email:     "email",
		FirstName: "firstName",
		LastName:  "lastName",
		Address: rainhelloworld.PhysicalAddressParam{
			City:        "city",
			Country:     "country",
			CountryCode: "xx",
			Line1:       "line1",
			PostalCode:  "postalCode",
			Region:      "region",
			Line2:       rainhelloworld.String("line2"),
		},
		PhoneCountryCode: rainhelloworld.String("1"),
		PhoneNumber:      rainhelloworld.String("5555555555"),
		WalletAddress:    rainhelloworld.String("0xE1CB97d8EBbDbaAae6d9B1ca0D1cFaADcCcbdaDa"),
	})
	if err != nil {
		var apierr *rainhelloworld.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserGet(t *testing.T) {
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
	_, err := client.Users.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *rainhelloworld.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.UserUpdateParams{
			Address: rainhelloworld.PhysicalAddressParam{
				City:        "city",
				Country:     "country",
				CountryCode: "xx",
				Line1:       "line1",
				PostalCode:  "postalCode",
				Region:      "region",
				Line2:       rainhelloworld.String("line2"),
			},
			Email:                    rainhelloworld.String("email"),
			FirstName:                rainhelloworld.String("firstName"),
			IsActive:                 rainhelloworld.Bool(true),
			IsTermsOfServiceAccepted: rainhelloworld.Bool(true),
			LastName:                 rainhelloworld.String("lastName"),
			PhoneCountryCode:         rainhelloworld.String("1"),
			PhoneNumber:              rainhelloworld.String("5555555555"),
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

func TestUserListWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.List(context.TODO(), rainhelloworld.UserListParams{
		CompanyID: rainhelloworld.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Cursor:    rainhelloworld.String("cursor"),
		Limit:     rainhelloworld.Int(1),
	})
	if err != nil {
		var apierr *rainhelloworld.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserDelete(t *testing.T) {
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
	err := client.Users.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *rainhelloworld.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserNewCardWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.NewCard(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.UserNewCardParams{
			Type: rainhelloworld.UserNewCardParamsTypePhysical,
			Billing: rainhelloworld.PhysicalAddressParam{
				City:        "city",
				Country:     "country",
				CountryCode: "xx",
				Line1:       "line1",
				PostalCode:  "postalCode",
				Region:      "region",
				Line2:       rainhelloworld.String("line2"),
			},
			Configuration: rainhelloworld.UserNewCardParamsConfiguration{
				DisplayName:    rainhelloworld.String("displayName"),
				ProductID:      rainhelloworld.String("productId"),
				ProductRef:     rainhelloworld.String("productRef"),
				VirtualCardArt: rainhelloworld.String("virtualCardArt"),
			},
			Limit: rainhelloworld.IssuingCardLimitParam{
				Amount:    0,
				Frequency: rainhelloworld.IssuingCardLimitFrequencyPer24HourPeriod,
			},
			Shipping: rainhelloworld.UserNewCardParamsShipping{
				PhysicalAddressParam: rainhelloworld.PhysicalAddressParam{
					City:        "city",
					Country:     "country",
					CountryCode: "xx",
					Line1:       "line1",
					PostalCode:  "postalCode",
					Region:      "region",
					Line2:       rainhelloworld.String("line2"),
				},
				PhoneNumber: "phoneNumber",
				Method:      "standard",
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

func TestUserNewCharge(t *testing.T) {
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
	_, err := client.Users.NewCharge(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.UserNewChargeParams{
			IssuingChargeCreateBody: rainhelloworld.IssuingChargeCreateBodyParam{
				Amount:      1,
				Description: "description",
			},
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

func TestUserInitiatePaymentWithOptionalParams(t *testing.T) {
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
	_, err := client.Users.InitiatePayment(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.UserInitiatePaymentParams{
			Amount:        0,
			WalletAddress: "0xE1CB97d8EBbDbaAae6d9B1ca0D1cFaADcCcbdaDa",
			ChainID:       rainhelloworld.Int(0),
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

func TestUserGetBalances(t *testing.T) {
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
	_, err := client.Users.GetBalances(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *rainhelloworld.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserGetContracts(t *testing.T) {
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
	_, err := client.Users.GetContracts(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *rainhelloworld.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
