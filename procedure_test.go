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

func TestProcedureNew(t *testing.T) {
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
	_, err := client.Procedures.New(context.TODO(), lablinkv4client.ProcedureNewParams{
		Body: []lablinkv4client.ProcedureNewParamsBody{{
			Name: "name",
			Code: lablinkv4client.String("code"),
			CodeSystems: []lablinkv4client.ProcedureNewParamsBodyCodeSystem{{
				Code:         lablinkv4client.String("code"),
				CodingSystem: lablinkv4client.CodingSystemLoinc,
			}},
			Description: lablinkv4client.String("description"),
			Unit:        lablinkv4client.String("unit"),
		}},
	})
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProcedureUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Procedures.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lablinkv4client.ProcedureUpdateParams{
			Code: lablinkv4client.String("code"),
			CodeSystems: []lablinkv4client.ProcedureUpdateParamsCodeSystem{{
				Code:         lablinkv4client.String("code"),
				CodingSystem: lablinkv4client.CodingSystemLoinc,
			}},
			Description: lablinkv4client.String("description"),
			Name:        lablinkv4client.String("name"),
			Unit:        lablinkv4client.String("unit"),
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

func TestProcedureListWithOptionalParams(t *testing.T) {
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
	_, err := client.Procedures.List(context.TODO(), lablinkv4client.ProcedureListParams{
		Codes:      []string{"string"},
		CodeSystem: lablinkv4client.CodingSystemLoinc,
		Name:       lablinkv4client.String("name"),
		Page:       lablinkv4client.Int(0),
		PageSize:   lablinkv4client.Int(1),
		Sort:       []string{"code,asc"},
		Unit:       lablinkv4client.String("unit"),
	})
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProcedureDelete(t *testing.T) {
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
	err := client.Procedures.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
