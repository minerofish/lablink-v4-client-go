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

func TestOrganizationQueryNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Organizations.Query.New(context.TODO(), lablinkv4client.OrganizationQueryNewParams{
		Page:            lablinkv4client.Int(0),
		PageSize:        lablinkv4client.Int(1),
		Sort:            []string{"name,asc"},
		Emails:          []string{"user@example.com"},
		Name:            lablinkv4client.String("John Doe Clinic"),
		OrganizationIDs: []string{"123e4567-e89b-12d3-a456-426614174000"},
		Roles:           []lablinkv4client.OrganizationRole{lablinkv4client.OrganizationRoleOrganizationAdmin},
	})
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
