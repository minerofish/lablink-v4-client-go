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

func TestOrderTagDelete(t *testing.T) {
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
	err := client.Orders.Tags.Delete(context.TODO(), lablinkv4client.OrderTagDeleteParams{
		OrderIDs: []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		Tags:     []string{"-_-k  W2K-1V"},
	})
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrderTagTag(t *testing.T) {
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
	err := client.Orders.Tags.Tag(context.TODO(), lablinkv4client.OrderTagTagParams{
		OrderIDs: []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		Tags:     []string{"-_-k  W2K-1V"},
	})
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
