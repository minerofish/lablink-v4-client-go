// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/lablink-v4-client-go"
	"github.com/stainless-sdks/lablink-v4-client-go/internal/testutil"
	"github.com/stainless-sdks/lablink-v4-client-go/option"
)

func TestLaboratoryNew(t *testing.T) {
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
	_, err := client.Laboratories.New(context.TODO(), lablinkv4client.LaboratoryNewParams{
		Body: []lablinkv4client.LaboratoryNewParamsBody{{
			InstanceName: "central-lab-instance",
			Name:         "Central Lab",
			Address:      lablinkv4client.String("123 Lab St, Lab City"),
			City:         lablinkv4client.String("Lab City"),
			Country:      lablinkv4client.String("Labland"),
			Email:        lablinkv4client.String("lab@labland.com"),
			Phone:        lablinkv4client.String("+1234567890"),
			Postcode:     lablinkv4client.String("12345"),
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

func TestLaboratoryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Laboratories.List(context.TODO(), lablinkv4client.LaboratoryListParams{
		LaboratoryID: lablinkv4client.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaboratoryDelete(t *testing.T) {
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
	err := client.Laboratories.Delete(context.TODO(), lablinkv4client.LaboratoryDeleteParams{
		LaboratoryID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
	})
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaboratoryQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Laboratories.Query(context.TODO(), lablinkv4client.LaboratoryQueryParams{
		Page:          lablinkv4client.Int(0),
		PageSize:      lablinkv4client.Int(1),
		Sort:          []string{"name,asc"},
		LaboratoryIDs: []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		Name:          lablinkv4client.String("name"),
	})
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLaboratoryUploadConfirmatoryResults(t *testing.T) {
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
	_, err := client.Laboratories.UploadConfirmatoryResults(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lablinkv4client.LaboratoryUploadConfirmatoryResultsParams{
			Body: []lablinkv4client.LaboratoryUploadConfirmatoryResultsParamsBody{{
				Examination:         "examination",
				SampleCode:          "sampleCode",
				YieldDate:           "yieldDate",
				ConfirmatoryGroupID: lablinkv4client.Int(0),
				Customer:            lablinkv4client.String("customer"),
				QualifiedResult:     lablinkv4client.String("qualifiedResult"),
				QuantitativeResult:  lablinkv4client.String("quantitativeResult"),
			}},
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

func TestLaboratoryUploadDocument(t *testing.T) {
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
	err := client.Laboratories.UploadDocument(
		context.TODO(),
		"laboratoryId",
		lablinkv4client.LaboratoryUploadDocumentParams{
			Examination: "examination",
			File:        io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			SampleCode:  "sampleCode",
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

func TestLaboratoryUploadResults(t *testing.T) {
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
	_, err := client.Laboratories.UploadResults(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lablinkv4client.LaboratoryUploadResultsParams{
			Body: []lablinkv4client.LaboratoryUploadResultsParamsBody{{
				Examination:        "examination",
				SampleCode:         "sampleCode",
				YieldDate:          "yieldDate",
				Customer:           lablinkv4client.String("customer"),
				QualifiedResult:    lablinkv4client.String("qualifiedResult"),
				QuantitativeResult: lablinkv4client.String("quantitativeResult"),
			}},
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
