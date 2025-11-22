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

func TestExaminationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Examinations.List(context.TODO(), lablinkv4client.ExaminationListParams{
		Code:           lablinkv4client.String("code"),
		LaboratoryID:   lablinkv4client.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Name:           lablinkv4client.String("name"),
		OrganizationID: lablinkv4client.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Page:           lablinkv4client.Int(0),
		PageSize:       lablinkv4client.Int(1),
		Sort:           []string{"code,asc"},
		Unit:           lablinkv4client.String("unit"),
	})
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
