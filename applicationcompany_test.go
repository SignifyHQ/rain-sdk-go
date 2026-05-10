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

func TestApplicationCompanyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Applications.Company.New(context.TODO(), rainsdk.ApplicationCompanyNewParams{
		Address: rainsdk.PhysicalAddressParam{
			City:        "city",
			Country:     "country",
			CountryCode: "xx",
			Line1:       "line1",
			PostalCode:  "postalCode",
			Region:      "region",
			Line2:       rainsdk.String("line2"),
		},
		Entity: rainsdk.ApplicationCompanyNewParamsEntity{
			Name:               "name",
			RegistrationNumber: "registrationNumber",
			TaxID:              "taxId",
			Website:            "website",
			Description:        rainsdk.String("description"),
			ExpectedSpend:      rainsdk.String("expectedSpend"),
			Type:               rainsdk.String("type"),
		},
		InitialUser: rainsdk.ApplicationCompanyNewParamsInitialUser{
			IssuingApplicationPersonParam: rainsdk.IssuingApplicationPersonParam{
				Address: rainsdk.PhysicalAddressParam{
					City:        "city",
					Country:     "country",
					CountryCode: "xx",
					Line1:       "line1",
					PostalCode:  "postalCode",
					Region:      "region",
					Line2:       rainsdk.String("line2"),
				},
				BirthDate:        time.Now(),
				CountryOfIssue:   "xx",
				Email:            "email",
				FirstName:        "firstName",
				LastName:         "lastName",
				NationalID:       "nationalId",
				ID:               rainsdk.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				PhoneCountryCode: rainsdk.String("1"),
				PhoneNumber:      rainsdk.String("5555555555"),
			},
			IPAddress:                "ipAddress",
			IsTermsOfServiceAccepted: true,
			Role:                     rainsdk.String("role"),
			SolanaAddress:            rainsdk.String("WRktL2iKFTHZg6qNBPzV1b1WLYwfnZ5JSHo2UV8L1R"),
			WalletAddress:            rainsdk.String("0xE1CB97d8EBbDbaAae6d9B1ca0D1cFaADcCcbdaDa"),
		},
		Name: "name",
		Representatives: []rainsdk.IssuingApplicationPersonParam{{
			Address: rainsdk.PhysicalAddressParam{
				City:        "city",
				Country:     "country",
				CountryCode: "xx",
				Line1:       "line1",
				PostalCode:  "postalCode",
				Region:      "region",
				Line2:       rainsdk.String("line2"),
			},
			BirthDate:        time.Now(),
			CountryOfIssue:   "xx",
			Email:            "email",
			FirstName:        "firstName",
			LastName:         "lastName",
			NationalID:       "nationalId",
			ID:               rainsdk.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			PhoneCountryCode: rainsdk.String("1"),
			PhoneNumber:      rainsdk.String("5555555555"),
		}},
		UltimateBeneficialOwners: []rainsdk.IssuingApplicationPersonParam{{
			Address: rainsdk.PhysicalAddressParam{
				City:        "city",
				Country:     "country",
				CountryCode: "xx",
				Line1:       "line1",
				PostalCode:  "postalCode",
				Region:      "region",
				Line2:       rainsdk.String("line2"),
			},
			BirthDate:        time.Now(),
			CountryOfIssue:   "xx",
			Email:            "email",
			FirstName:        "firstName",
			LastName:         "lastName",
			NationalID:       "nationalId",
			ID:               rainsdk.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			PhoneCountryCode: rainsdk.String("1"),
			PhoneNumber:      rainsdk.String("5555555555"),
		}},
		ChainID:         rainsdk.String("chainId"),
		ContractAddress: rainsdk.String("0xE1CB97d8EBbDbaAae6d9B1ca0D1cFaADcCcbdaDa"),
		SourceKey:       rainsdk.String("x"),
	})
	if err != nil {
		var apierr *rainsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationCompanyGet(t *testing.T) {
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
	_, err := client.Applications.Company.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *rainsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationCompanyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Applications.Company.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainsdk.ApplicationCompanyUpdateParams{
			Address: rainsdk.PhysicalAddressParam{
				City:        "city",
				Country:     "country",
				CountryCode: "xx",
				Line1:       "line1",
				PostalCode:  "postalCode",
				Region:      "region",
				Line2:       rainsdk.String("line2"),
			},
			Entity: rainsdk.ApplicationCompanyUpdateParamsEntity{
				Description:        rainsdk.String("description"),
				ExpectedSpend:      rainsdk.String("expectedSpend"),
				RegistrationNumber: rainsdk.String("registrationNumber"),
				TaxID:              rainsdk.String("taxId"),
				Type:               rainsdk.String("type"),
				Website:            rainsdk.String("website"),
			},
			Name: rainsdk.String("name"),
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

func TestApplicationCompanyReapplyWithOptionalParams(t *testing.T) {
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
	_, err := client.Applications.Company.Reapply(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainsdk.ApplicationCompanyReapplyParams{
			Address: rainsdk.PhysicalAddressParam{
				City:        "city",
				Country:     "country",
				CountryCode: "xx",
				Line1:       "line1",
				PostalCode:  "postalCode",
				Region:      "region",
				Line2:       rainsdk.String("line2"),
			},
			Entity: rainsdk.ApplicationCompanyReapplyParamsEntity{
				Website:       "website",
				Description:   rainsdk.String("description"),
				ExpectedSpend: rainsdk.String("expectedSpend"),
				Type:          rainsdk.String("type"),
			},
			InitialUser: rainsdk.ApplicationCompanyReapplyParamsInitialUser{
				Address: rainsdk.PhysicalAddressParam{
					City:        "city",
					Country:     "country",
					CountryCode: "xx",
					Line1:       "line1",
					PostalCode:  "postalCode",
					Region:      "region",
					Line2:       rainsdk.String("line2"),
				},
				BirthDate:                time.Now(),
				CountryOfIssue:           "countryOfIssue",
				IPAddress:                "ipAddress",
				IsTermsOfServiceAccepted: true,
				NationalID:               "nationalId",
				Role:                     rainsdk.String("role"),
			},
			Name: "name",
			Representatives: []rainsdk.IssuingApplicationPersonParam{{
				Address: rainsdk.PhysicalAddressParam{
					City:        "city",
					Country:     "country",
					CountryCode: "xx",
					Line1:       "line1",
					PostalCode:  "postalCode",
					Region:      "region",
					Line2:       rainsdk.String("line2"),
				},
				BirthDate:        time.Now(),
				CountryOfIssue:   "xx",
				Email:            "email",
				FirstName:        "firstName",
				LastName:         "lastName",
				NationalID:       "nationalId",
				ID:               rainsdk.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				PhoneCountryCode: rainsdk.String("1"),
				PhoneNumber:      rainsdk.String("5555555555"),
			}},
			UltimateBeneficialOwners: []rainsdk.IssuingApplicationPersonParam{{
				Address: rainsdk.PhysicalAddressParam{
					City:        "city",
					Country:     "country",
					CountryCode: "xx",
					Line1:       "line1",
					PostalCode:  "postalCode",
					Region:      "region",
					Line2:       rainsdk.String("line2"),
				},
				BirthDate:        time.Now(),
				CountryOfIssue:   "xx",
				Email:            "email",
				FirstName:        "firstName",
				LastName:         "lastName",
				NationalID:       "nationalId",
				ID:               rainsdk.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				PhoneCountryCode: rainsdk.String("1"),
				PhoneNumber:      rainsdk.String("5555555555"),
			}},
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

func TestApplicationCompanyUploadDocumentWithOptionalParams(t *testing.T) {
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
	err := client.Applications.Company.UploadDocument(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainsdk.ApplicationCompanyUploadDocumentParams{
			Document: io.Reader(bytes.NewBuffer([]byte("Example data"))),
			Country:  rainsdk.String("xxx"),
			Name:     rainsdk.String("name"),
			Side:     rainsdk.ApplicationCompanyUploadDocumentParamsSideFront,
			Type:     rainsdk.ApplicationCompanyUploadDocumentParamsTypeDirectorsRegistry,
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
