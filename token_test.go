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

func TestTokenNew(t *testing.T) {
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
	_, err := client.Token.New(context.TODO(), lablinkv4client.TokenNewParams{
		ClientID:     "client_id",
		Code:         "code",
		CodeVerifier: "code_verifier",
		GrantType:    lablinkv4client.TokenNewParamsGrantTypeAuthorizationCode,
		RedirectUri:  "https://example.com",
	})
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
