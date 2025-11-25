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

func TestLaboratoryContractExaminationUpdate(t *testing.T) {
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
	err := client.Laboratories.Contracts.Examinations.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lablinkv4client.LaboratoryContractExaminationUpdateParams{
			LaboratoryID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Body: []lablinkv4client.LaboratoryContractExaminationUpdateParamsBody{{
				ExaminationID:   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				ExaminationCode: lablinkv4client.String("examinationCode"),
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

func TestLaboratoryContractExaminationAdd(t *testing.T) {
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
	err := client.Laboratories.Contracts.Examinations.Add(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lablinkv4client.LaboratoryContractExaminationAddParams{
			LaboratoryID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Body: []lablinkv4client.LaboratoryContractExaminationAddParamsBody{{
				ProcedureID:     "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				ExaminationCode: lablinkv4client.String("examinationCode"),
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

func TestLaboratoryContractExaminationRemove(t *testing.T) {
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
	err := client.Laboratories.Contracts.Examinations.Remove(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lablinkv4client.LaboratoryContractExaminationRemoveParams{
			LaboratoryID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Body: []lablinkv4client.LaboratoryContractExaminationRemoveParamsBody{{
				ExaminationID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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
