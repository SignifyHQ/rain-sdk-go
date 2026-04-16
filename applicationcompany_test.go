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

func TestApplicationCompanyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Applications.Company.New(context.TODO(), rainhelloworld.ApplicationCompanyNewParams{
		Address: rainhelloworld.PhysicalAddressParam{
			City:        "city",
			Country:     "country",
			CountryCode: "xx",
			Line1:       "line1",
			PostalCode:  "postalCode",
			Region:      "region",
			Line2:       rainhelloworld.String("line2"),
		},
		Entity: rainhelloworld.ApplicationCompanyNewParamsEntity{
			Name:               "name",
			RegistrationNumber: "registrationNumber",
			TaxID:              "taxId",
			Website:            "website",
			Description:        rainhelloworld.String("description"),
			ExpectedSpend:      rainhelloworld.String("expectedSpend"),
			Type:               rainhelloworld.String("type"),
		},
		InitialUser: rainhelloworld.ApplicationCompanyNewParamsInitialUser{
			IssuingApplicationPersonParam: rainhelloworld.IssuingApplicationPersonParam{
				Address: rainhelloworld.PhysicalAddressParam{
					City:        "city",
					Country:     "country",
					CountryCode: "xx",
					Line1:       "line1",
					PostalCode:  "postalCode",
					Region:      "region",
					Line2:       rainhelloworld.String("line2"),
				},
				BirthDate:        time.Now(),
				CountryOfIssue:   "xx",
				Email:            "email",
				FirstName:        "firstName",
				LastName:         "lastName",
				NationalID:       "nationalId",
				ID:               rainhelloworld.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				PhoneCountryCode: rainhelloworld.String("1"),
				PhoneNumber:      rainhelloworld.String("5555555555"),
			},
			IPAddress:                "ipAddress",
			IsTermsOfServiceAccepted: true,
			Role:                     rainhelloworld.String("role"),
			SolanaAddress:            rainhelloworld.String("WRktL2iKFTHZg6qNBPzV1b1WLYwfnZ5JSHo2UV8L1R"),
			WalletAddress:            rainhelloworld.String("0xE1CB97d8EBbDbaAae6d9B1ca0D1cFaADcCcbdaDa"),
		},
		Name: "name",
		Representatives: []rainhelloworld.IssuingApplicationPersonParam{{
			Address: rainhelloworld.PhysicalAddressParam{
				City:        "city",
				Country:     "country",
				CountryCode: "xx",
				Line1:       "line1",
				PostalCode:  "postalCode",
				Region:      "region",
				Line2:       rainhelloworld.String("line2"),
			},
			BirthDate:        time.Now(),
			CountryOfIssue:   "xx",
			Email:            "email",
			FirstName:        "firstName",
			LastName:         "lastName",
			NationalID:       "nationalId",
			ID:               rainhelloworld.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			PhoneCountryCode: rainhelloworld.String("1"),
			PhoneNumber:      rainhelloworld.String("5555555555"),
		}},
		UltimateBeneficialOwners: []rainhelloworld.IssuingApplicationPersonParam{{
			Address: rainhelloworld.PhysicalAddressParam{
				City:        "city",
				Country:     "country",
				CountryCode: "xx",
				Line1:       "line1",
				PostalCode:  "postalCode",
				Region:      "region",
				Line2:       rainhelloworld.String("line2"),
			},
			BirthDate:        time.Now(),
			CountryOfIssue:   "xx",
			Email:            "email",
			FirstName:        "firstName",
			LastName:         "lastName",
			NationalID:       "nationalId",
			ID:               rainhelloworld.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			PhoneCountryCode: rainhelloworld.String("1"),
			PhoneNumber:      rainhelloworld.String("5555555555"),
		}},
		ChainID:         rainhelloworld.String("chainId"),
		ContractAddress: rainhelloworld.String("0xE1CB97d8EBbDbaAae6d9B1ca0D1cFaADcCcbdaDa"),
		SourceKey:       rainhelloworld.String("x"),
	})
	if err != nil {
		var apierr *rainhelloworld.Error
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
	client := rainhelloworld.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Applications.Company.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *rainhelloworld.Error
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
	client := rainhelloworld.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Applications.Company.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.ApplicationCompanyUpdateParams{
			Address: rainhelloworld.PhysicalAddressParam{
				City:        "city",
				Country:     "country",
				CountryCode: "xx",
				Line1:       "line1",
				PostalCode:  "postalCode",
				Region:      "region",
				Line2:       rainhelloworld.String("line2"),
			},
			Entity: rainhelloworld.ApplicationCompanyUpdateParamsEntity{
				Description:        rainhelloworld.String("description"),
				ExpectedSpend:      rainhelloworld.String("expectedSpend"),
				RegistrationNumber: rainhelloworld.String("registrationNumber"),
				TaxID:              rainhelloworld.String("taxId"),
				Type:               rainhelloworld.String("type"),
				Website:            rainhelloworld.String("website"),
			},
			Name: rainhelloworld.String("name"),
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

func TestApplicationCompanyReapplyWithOptionalParams(t *testing.T) {
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
	_, err := client.Applications.Company.Reapply(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.ApplicationCompanyReapplyParams{
			Address: rainhelloworld.PhysicalAddressParam{
				City:        "city",
				Country:     "country",
				CountryCode: "xx",
				Line1:       "line1",
				PostalCode:  "postalCode",
				Region:      "region",
				Line2:       rainhelloworld.String("line2"),
			},
			Entity: rainhelloworld.ApplicationCompanyReapplyParamsEntity{
				Website:       "website",
				Description:   rainhelloworld.String("description"),
				ExpectedSpend: rainhelloworld.String("expectedSpend"),
				Type:          rainhelloworld.String("type"),
			},
			InitialUser: rainhelloworld.ApplicationCompanyReapplyParamsInitialUser{
				Address: rainhelloworld.PhysicalAddressParam{
					City:        "city",
					Country:     "country",
					CountryCode: "xx",
					Line1:       "line1",
					PostalCode:  "postalCode",
					Region:      "region",
					Line2:       rainhelloworld.String("line2"),
				},
				BirthDate:                time.Now(),
				CountryOfIssue:           "countryOfIssue",
				IPAddress:                "ipAddress",
				IsTermsOfServiceAccepted: true,
				NationalID:               "nationalId",
				Role:                     rainhelloworld.String("role"),
			},
			Name: "name",
			Representatives: []rainhelloworld.IssuingApplicationPersonParam{{
				Address: rainhelloworld.PhysicalAddressParam{
					City:        "city",
					Country:     "country",
					CountryCode: "xx",
					Line1:       "line1",
					PostalCode:  "postalCode",
					Region:      "region",
					Line2:       rainhelloworld.String("line2"),
				},
				BirthDate:        time.Now(),
				CountryOfIssue:   "xx",
				Email:            "email",
				FirstName:        "firstName",
				LastName:         "lastName",
				NationalID:       "nationalId",
				ID:               rainhelloworld.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				PhoneCountryCode: rainhelloworld.String("1"),
				PhoneNumber:      rainhelloworld.String("5555555555"),
			}},
			UltimateBeneficialOwners: []rainhelloworld.IssuingApplicationPersonParam{{
				Address: rainhelloworld.PhysicalAddressParam{
					City:        "city",
					Country:     "country",
					CountryCode: "xx",
					Line1:       "line1",
					PostalCode:  "postalCode",
					Region:      "region",
					Line2:       rainhelloworld.String("line2"),
				},
				BirthDate:        time.Now(),
				CountryOfIssue:   "xx",
				Email:            "email",
				FirstName:        "firstName",
				LastName:         "lastName",
				NationalID:       "nationalId",
				ID:               rainhelloworld.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				PhoneCountryCode: rainhelloworld.String("1"),
				PhoneNumber:      rainhelloworld.String("5555555555"),
			}},
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

func TestApplicationCompanyUploadDocumentWithOptionalParams(t *testing.T) {
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
	err := client.Applications.Company.UploadDocument(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.ApplicationCompanyUploadDocumentParams{
			Document: io.Reader(bytes.NewBuffer([]byte("Example data"))),
			Country:  rainhelloworld.String("xxx"),
			Name:     rainhelloworld.String("name"),
			Side:     rainhelloworld.ApplicationCompanyUploadDocumentParamsSideFront,
			Type:     rainhelloworld.ApplicationCompanyUploadDocumentParamsTypeDirectorsRegistry,
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
