// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/lablink-v4-client-go"
	"github.com/stainless-sdks/lablink-v4-client-go/internal/testutil"
	"github.com/stainless-sdks/lablink-v4-client-go/option"
)

func TestLaboratoryOrderUpdateWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	err := client.Laboratories.Orders.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lablinkv4client.LaboratoryOrderUpdateParams{
			NewState:    lablinkv4client.LaboratoryOrderUpdateParamsNewStateProcessing,
			OrderIDs:    []string{`["550e8400-e29b-41d4-a716-446655440000"]`},
			SampleCodes: []string{`["SAMPLE001"`, `"SAMPLE002"]`},
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

func TestLaboratoryOrderListWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Laboratories.Orders.List(
		context.TODO(),
		"laboratoryId",
		lablinkv4client.LaboratoryOrderListParams{
			Page:     lablinkv4client.Int(0),
			PageSize: lablinkv4client.Int(1),
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
