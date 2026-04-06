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

	"github.com/stainless-sdks/rain-hello-world-go"
	"github.com/stainless-sdks/rain-hello-world-go/internal/testutil"
	"github.com/stainless-sdks/rain-hello-world-go/option"
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
	client := rainhelloworld.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Applications.User.New(context.TODO(), rainhelloworld.ApplicationUserNewParams{
		OfUsingSumsubShareToken: &rainhelloworld.ApplicationUserNewParamsBodyUsingSumsubShareToken{
			AccountPurpose:           "accountPurpose",
			AnnualSalary:             "annualSalary",
			ExpectedMonthlyVolume:    "expectedMonthlyVolume",
			IPAddress:                "ipAddress",
			IsTermsOfServiceAccepted: true,
			Occupation:               "occupation",
			SumsubShareToken:         "sumsubShareToken",
			ChainID:                  rainhelloworld.String("chainId"),
			ContractAddress:          rainhelloworld.String("0xE1CB97d8EBbDbaAae6d9B1ca0D1cFaADcCcbdaDa"),
			HasExistingDocuments:     rainhelloworld.Bool(true),
			SolanaAddress:            rainhelloworld.String("WRktL2iKFTHZg6qNBPzV1b1WLYwfnZ5JSHo2UV8L1R"),
			SourceKey:                rainhelloworld.String("x"),
			WalletAddress:            rainhelloworld.String("0xE1CB97d8EBbDbaAae6d9B1ca0D1cFaADcCcbdaDa"),
		},
	})
	if err != nil {
		var apierr *rainhelloworld.Error
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
	client := rainhelloworld.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Applications.User.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *rainhelloworld.Error
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
	client := rainhelloworld.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Applications.User.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.ApplicationUserUpdateParams{
			AccountPurpose: rainhelloworld.String("accountPurpose"),
			Address: rainhelloworld.PhysicalAddressParam{
				City:        "city",
				Country:     "country",
				CountryCode: "xx",
				Line1:       "line1",
				PostalCode:  "postalCode",
				Region:      "region",
				Line2:       rainhelloworld.String("line2"),
			},
			AnnualSalary:             rainhelloworld.String("annualSalary"),
			BirthDate:                rainhelloworld.Time(time.Now()),
			CountryOfIssue:           rainhelloworld.String("countryOfIssue"),
			ExpectedMonthlyVolume:    rainhelloworld.String("expectedMonthlyVolume"),
			FirstName:                rainhelloworld.String("firstName"),
			HasExistingDocuments:     rainhelloworld.Bool(true),
			IPAddress:                rainhelloworld.String("ipAddress"),
			IsTermsOfServiceAccepted: true,
			LastName:                 rainhelloworld.String("lastName"),
			NationalID:               rainhelloworld.String("nationalId"),
			Occupation:               rainhelloworld.String("occupation"),
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

func TestApplicationUserInitiateWithOptionalParams(t *testing.T) {
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
	_, err := client.Applications.User.Initiate(context.TODO(), rainhelloworld.ApplicationUserInitiateParams{
		Email:         rainhelloworld.String("email"),
		FirstName:     rainhelloworld.String("firstName"),
		LastName:      rainhelloworld.String("lastName"),
		WalletAddress: rainhelloworld.String("0xE1CB97d8EBbDbaAae6d9B1ca0D1cFaADcCcbdaDa"),
	})
	if err != nil {
		var apierr *rainhelloworld.Error
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
	client := rainhelloworld.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Applications.User.Reapply(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.ApplicationUserReapplyParams{
			AccountPurpose: "accountPurpose",
			Address: rainhelloworld.PhysicalAddressParam{
				City:        "city",
				Country:     "country",
				CountryCode: "xx",
				Line1:       "line1",
				PostalCode:  "postalCode",
				Region:      "region",
				Line2:       rainhelloworld.String("line2"),
			},
			AnnualSalary:             "annualSalary",
			BirthDate:                time.Now(),
			CountryOfIssue:           "countryOfIssue",
			ExpectedMonthlyVolume:    "expectedMonthlyVolume",
			IPAddress:                "ipAddress",
			IsTermsOfServiceAccepted: true,
			NationalID:               "nationalId",
			Occupation:               "occupation",
			HasExistingDocuments:     rainhelloworld.Bool(true),
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

func TestApplicationUserUploadDocumentWithOptionalParams(t *testing.T) {
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
	err := client.Applications.User.UploadDocument(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.ApplicationUserUploadDocumentParams{
			Document: io.Reader(bytes.NewBuffer([]byte("Example data"))),
			Country:  rainhelloworld.String("xxx"),
			Name:     rainhelloworld.String("name"),
			Side:     rainhelloworld.ApplicationUserUploadDocumentParamsSideFront,
			Type:     rainhelloworld.ApplicationUserUploadDocumentParamsTypeIDCard,
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
