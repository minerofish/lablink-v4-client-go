// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/minerofish/lablink-v4-client-go"
	"github.com/minerofish/lablink-v4-client-go/internal/testutil"
	"github.com/minerofish/lablink-v4-client-go/option"
	"github.com/minerofish/lablink-v4-client-go/shared"
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

func TestLaboratoryQueryContractsWithOptionalParams(t *testing.T) {
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
	_, err := client.Laboratories.QueryContracts(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lablinkv4client.LaboratoryQueryContractsParams{
			Page:            lablinkv4client.Int(0),
			PageSize:        lablinkv4client.Int(1),
			Sort:            []string{"validFrom,asc"},
			OrganizationIDs: []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
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
		option.WithAPIKey("My API Key"),
	)
	err := client.Laboratories.UploadDocument(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lablinkv4client.LaboratoryUploadDocumentParams{
			File:    io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			OrderID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Type:    lablinkv4client.LaboratoryUploadDocumentParamsTypeLabReport,
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
		option.WithAPIKey("My API Key"),
	)
	err := client.Laboratories.UploadResults(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lablinkv4client.LaboratoryUploadResultsParams{
			Body: []shared.ResultParam{{
				ExaminationID:       "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				ItemID:              "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				OrderID:             "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				Result:              "result",
				Status:              shared.ResultStatusFinal,
				ConfirmationPending: lablinkv4client.Bool(true),
				DataType:            shared.ResultDataTypeInt,
				InfoText:            lablinkv4client.String("infoText"),
				PerformedAt:         lablinkv4client.Time(time.Now()),
				Unit:                lablinkv4client.String("unit"),
				ValidatedAt:         lablinkv4client.Time(time.Now()),
				ValidatedBy:         lablinkv4client.String("validatedBy"),
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
