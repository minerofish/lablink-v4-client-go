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

func TestAuthorizeGetWithOptionalParams(t *testing.T) {
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
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Authorize.Get(context.TODO(), lablinkv4client.AuthorizeGetParams{
		ClientID:            "client_id",
		CodeChallenge:       "code_challenge",
		CodeChallengeMethod: lablinkv4client.AuthorizeGetParamsCodeChallengeMethodS256,
		RedirectUri:         "https://example.com",
		ResponseType:        lablinkv4client.AuthorizeGetParamsResponseTypeCode,
		Scope:               lablinkv4client.String("scope"),
		State:               lablinkv4client.String("state"),
	})
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
