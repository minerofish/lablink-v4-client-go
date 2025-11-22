// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/lablink-v4-client-go"
	"github.com/stainless-sdks/lablink-v4-client-go/internal/testutil"
	"github.com/stainless-sdks/lablink-v4-client-go/option"
)

func TestOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Orders.New(context.TODO(), lablinkv4client.OrderNewParams{
		FailOnError:  lablinkv4client.Bool(true),
		LaboratoryID: lablinkv4client.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		LocationID:   lablinkv4client.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		Body: []lablinkv4client.OrderNewParamsBody{{
			Examinations: []lablinkv4client.OrderNewParamsBodyExamination{{
				CollectionTime:  time.Now(),
				SampleCode:      "sampleCode",
				ExaminationCode: lablinkv4client.String("examinationCode"),
				ExaminationID:   lablinkv4client.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
				Reference:       lablinkv4client.String("reference"),
			}},
			LaboratoryID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			LocationID:   "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Type:         lablinkv4client.OrderTypeDonor,
			BloodDonor: lablinkv4client.BloodDonorParam{
				DateOfBirth:  time.Now(),
				DonationCode: "donationCode",
				DonorCode:    "donorCode",
				Gender:       lablinkv4client.GenderF,
				BloodType:    lablinkv4client.BloodTypeA,
				Cde:          lablinkv4client.String("cde"),
				DonorType:    lablinkv4client.String("donorType"),
				Kell:         lablinkv4client.KellPos,
				Rhesus:       lablinkv4client.RhesusPos,
			},
			BoneMarrowDonor: lablinkv4client.BoneMarrowDonorParam{
				DonorCode: "donorCode",
			},
			Patient: lablinkv4client.PatientParam{
				DateOfBirth:          time.Now(),
				FirstName:            "firstName",
				Gender:               lablinkv4client.GenderF,
				LastName:             "lastName",
				Address:              lablinkv4client.String("address"),
				City:                 lablinkv4client.String("city"),
				Country:              lablinkv4client.String("country"),
				PhoneNumberPrimary:   lablinkv4client.String("phoneNumberPrimary"),
				PhoneNumberSecondary: lablinkv4client.String("phoneNumberSecondary"),
				PostCode:             lablinkv4client.String("postCode"),
			},
			Pseudonym: lablinkv4client.PseudonymParam{
				Code:        "code",
				DateOfBirth: lablinkv4client.Time(time.Now()),
				Gender:      lablinkv4client.GenderF,
			},
			References: map[string]string{
				"foo": "string",
			},
			Tags: []string{"-_-k  W2K-1V"},
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

func TestOrderGet(t *testing.T) {
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
	_, err := client.Orders.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrderDelete(t *testing.T) {
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
	err := client.Orders.Delete(context.TODO(), lablinkv4client.OrderDeleteParams{
		OrderIDs: []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
	})
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrderNewCustom(t *testing.T) {
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
	_, err := client.Orders.NewCustom(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		lablinkv4client.OrderNewCustomParams{
			LaboratoryID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			Body:         "body",
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

func TestOrderDeleteOrder(t *testing.T) {
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
	err := client.Orders.DeleteOrder(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrderGetState(t *testing.T) {
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
	_, err := client.Orders.GetState(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrderQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Orders.Query(context.TODO(), lablinkv4client.OrderQueryParams{
		Page:            lablinkv4client.Int(0),
		PageSize:        lablinkv4client.Int(1),
		Sort:            []string{"state,asc"},
		IDs:             []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		LaboratoryIDs:   []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		LocationIDs:     []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		OrganizationIDs: []string{"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"},
		SampleCodes:     []string{"string"},
		State:           []string{"CONFIRMATORY"},
		WithoutTags:     []string{"string"},
		WithTags:        []string{"string"},
	})
	if err != nil {
		var apierr *lablinkv4client.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
