// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/minerofish/lablink-v4-client-go"
	"github.com/minerofish/lablink-v4-client-go/internal/testutil"
	"github.com/minerofish/lablink-v4-client-go/option"
)

func TestOrderQueryNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lablinkv4client.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Orders.Query.New(context.TODO(), lablinkv4client.OrderQueryNewParams{
		Expand:          []string{"documents"},
		Page:            lablinkv4client.Int(0),
		PageSize:        lablinkv4client.Int(1),
		Sort:            []string{"state,asc"},
		Transaction:     lablinkv4client.OrderQueryNewParamsTransactionNone,
		IDs:             []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		LaboratoryIDs:   []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		LocationIDs:     []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		OrganizationIDs: []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		SampleCodes:     []string{"string"},
		State:           []lablinkv4client.OrderStateType{lablinkv4client.OrderStateTypeEntered},
		WithoutTags:     []string{"string"},
		WithTags:        []string{"string"},
	})
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrderQueryCommitTransactionWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := lablinkv4client.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Orders.Query.CommitTransaction(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lablinkv4client.OrderQueryCommitTransactionParams{
			AddTags:    []string{"-_-k  W2K-1V"},
			RemoveTags: []string{"-_-k  W2K-1V"},
			Setstate:   lablinkv4client.OrderStateTypeEntered,
		},
	)
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
