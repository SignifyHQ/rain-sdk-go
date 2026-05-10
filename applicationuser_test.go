// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainsdk_test

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

func TestApplicationUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Applications.User.New(context.TODO(), rainsdk.ApplicationUserNewParams{
		OfUsingSumsubShareToken: &rainsdk.ApplicationUserNewParamsBodyUsingSumsubShareToken{
			AccountPurpose:           "accountPurpose",
			AnnualSalary:             "annualSalary",
			ExpectedMonthlyVolume:    "expectedMonthlyVolume",
			IPAddress:                "ipAddress",
			IsTermsOfServiceAccepted: true,
			Occupation:               "occupation",
			SumsubShareToken:         "sumsubShareToken",
			ChainID:                  rainsdk.String("chainId"),
			ContractAddress:          rainsdk.String("0xE1CB97d8EBbDbaAae6d9B1ca0D1cFaADcCcbdaDa"),
			HasExistingDocuments:     rainsdk.Bool(true),
			SolanaAddress:            rainsdk.String("WRktL2iKFTHZg6qNBPzV1b1WLYwfnZ5JSHo2UV8L1R"),
			SourceKey:                rainsdk.String("x"),
			WalletAddress:            rainsdk.String("0xE1CB97d8EBbDbaAae6d9B1ca0D1cFaADcCcbdaDa"),
		},
	})
	if err != nil {
		var apierr *rainsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationUserGet(t *testing.T) {
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
	_, err := client.Applications.User.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *rainsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationUserUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Applications.User.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainsdk.ApplicationUserUpdateParams{
			AccountPurpose: rainsdk.String("accountPurpose"),
			Address: rainsdk.PhysicalAddressParam{
				City:        "city",
				Country:     "country",
				CountryCode: "xx",
				Line1:       "line1",
				PostalCode:  "postalCode",
				Region:      "region",
				Line2:       rainsdk.String("line2"),
			},
			AnnualSalary:             rainsdk.String("annualSalary"),
			BirthDate:                rainsdk.Time(time.Now()),
			CountryOfIssue:           rainsdk.String("countryOfIssue"),
			ExpectedMonthlyVolume:    rainsdk.String("expectedMonthlyVolume"),
			FirstName:                rainsdk.String("firstName"),
			HasExistingDocuments:     rainsdk.Bool(true),
			IPAddress:                rainsdk.String("ipAddress"),
			IsTermsOfServiceAccepted: true,
			LastName:                 rainsdk.String("lastName"),
			NationalID:               rainsdk.String("nationalId"),
			Occupation:               rainsdk.String("occupation"),
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

func TestApplicationUserInitiateWithOptionalParams(t *testing.T) {
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
	_, err := client.Applications.User.Initiate(context.TODO(), rainsdk.ApplicationUserInitiateParams{
		Email:         rainsdk.String("email"),
		FirstName:     rainsdk.String("firstName"),
		LastName:      rainsdk.String("lastName"),
		WalletAddress: rainsdk.String("0xE1CB97d8EBbDbaAae6d9B1ca0D1cFaADcCcbdaDa"),
	})
	if err != nil {
		var apierr *rainsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationUserReapplyWithOptionalParams(t *testing.T) {
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
	_, err := client.Applications.User.Reapply(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainsdk.ApplicationUserReapplyParams{
			AccountPurpose: "accountPurpose",
			Address: rainsdk.PhysicalAddressParam{
				City:        "city",
				Country:     "country",
				CountryCode: "xx",
				Line1:       "line1",
				PostalCode:  "postalCode",
				Region:      "region",
				Line2:       rainsdk.String("line2"),
			},
			AnnualSalary:             "annualSalary",
			BirthDate:                time.Now(),
			CountryOfIssue:           "countryOfIssue",
			ExpectedMonthlyVolume:    "expectedMonthlyVolume",
			IPAddress:                "ipAddress",
			IsTermsOfServiceAccepted: true,
			NationalID:               "nationalId",
			Occupation:               "occupation",
			HasExistingDocuments:     rainsdk.Bool(true),
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

func TestApplicationUserUploadDocumentWithOptionalParams(t *testing.T) {
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
	err := client.Applications.User.UploadDocument(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainsdk.ApplicationUserUploadDocumentParams{
			Document: io.Reader(bytes.NewBuffer([]byte("Example data"))),
			Country:  rainsdk.String("xxx"),
			Name:     rainsdk.String("name"),
			Side:     rainsdk.ApplicationUserUploadDocumentParamsSideFront,
			Type:     rainsdk.ApplicationUserUploadDocumentParamsTypeIDCard,
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
