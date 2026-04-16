// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rainhelloworld_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/SignifyHQ/rain-sdk-go"
	"github.com/SignifyHQ/rain-sdk-go/internal/testutil"
	"github.com/SignifyHQ/rain-sdk-go/option"
)

func TestTransactionGet(t *testing.T) {
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
	_, err := client.Transactions.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *rainhelloworld.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionUpdateWithOptionalParams(t *testing.T) {
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
	err := client.Transactions.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.TransactionUpdateParams{
			Memo: rainhelloworld.String("memo"),
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

func TestTransactionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.List(context.TODO(), rainhelloworld.TransactionListParams{
		AuthorizedAfter:  rainhelloworld.Time(time.Now()),
		AuthorizedBefore: rainhelloworld.Time(time.Now()),
		CardID:           rainhelloworld.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		CompanyID:        rainhelloworld.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Cursor:           rainhelloworld.String("cursor"),
		Limit:            rainhelloworld.Int(1),
		PostedAfter:      rainhelloworld.Time(time.Now()),
		PostedBefore:     rainhelloworld.Time(time.Now()),
		TransactionHash:  rainhelloworld.String("transactionHash"),
		Type:             []string{"spend"},
		UserID:           rainhelloworld.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *rainhelloworld.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTransactionNewDisputeWithOptionalParams(t *testing.T) {
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
	_, err := client.Transactions.NewDispute(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		rainhelloworld.TransactionNewDisputeParams{
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
