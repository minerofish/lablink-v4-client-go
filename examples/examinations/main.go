package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	lablink "github.com/minerofish/lablink-v4-client-go"
	"github.com/minerofish/lablink-v4-client-go/option"
)

func main() {
	baseURL := "https://lablink.bloodlab.org"

	// test With API Key Client
	// testWithAPIKeyClient(baseURL, "your-api-key")
	// return

	// test With Authenticated Client
	username := "your-username"
	password := "your-password"

	if os.Getenv("BASE_URL") != "" {
		baseURL = os.Getenv("BASE_URL")
	}
	if os.Getenv("USERNAME") != "" {
		username = os.Getenv("USERNAME")
	}
	if os.Getenv("PASSWORD") != "" {
		password = os.Getenv("PASSWORD")
	}

	testWithAuthenticatedClient(baseURL, username, password)

}

func testWithAuthenticatedClient(baseURL, username, password string) {
	authorizedClient, err := lablink.NewAuthenticatedClient(baseURL, username, password)
	if err != nil {
		panic("failed to create authenticated client: " + err.Error())
	}

	// -- Paginated list of examinations --

	response, err := authorizedClient.Do().Examinations.List(context.TODO(), lablink.ExaminationListParams{
		OrganizationID: lablink.String("3a1aedae-26d9-4888-bfe7-3c1ce75e79a5"),
		Page:           lablink.Int(0),
		PageSize:       lablink.Int(200),
		Sort:           []string{"code,asc"},
		// Unit:     lablink.String("unit"),
		// Code:           lablink.String("code"),
		// LaboratoryID:   lablink.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		// Name:           lablink.String("name"),
	})
	if err != nil {
		var apierr *lablink.Error
		if errors.As(err, &apierr) {
			fmt.Println(string(apierr.DumpRequest(true)))
		}
		fmt.Printf("err should be nil: %s", err.Error())
	}

	for i, examination := range response.Items {
		fmt.Printf("%d: %+v\n", i, examination)
	}

	// -- Non-paginated list of examinations in an order --

	response2, err := authorizedClient.Do().Laboratories.Orders.List(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", lablink.LaboratoryOrderListParams{
		Page:     lablink.Int(0),
		PageSize: lablink.Int(200),
	})
	if err != nil {
		var apierr *lablink.Error
		if errors.As(err, &apierr) {
			fmt.Println(string(apierr.DumpRequest(true)))
		}
		fmt.Printf("err should be nil: %s", err.Error())
	}

	for i, examination := range *response2 {
		fmt.Printf("%d: %+v\n", i, examination)
	}
}

func testWithAPIKeyClient(baseURL, apiKey string) {
	apiKeyClient := lablink.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey(apiKey),
	)

	response, err := apiKeyClient.Laboratories.Orders.List(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e", lablink.LaboratoryOrderListParams{
		Page:     lablink.Int(0),
		PageSize: lablink.Int(200),
	})
	if err != nil {
		var apierr *lablink.Error
		if errors.As(err, &apierr) {
			fmt.Println(string(apierr.DumpRequest(true)))
		}
		fmt.Printf("err should be nil: %s", err.Error())
	}

	for i, examination := range *response {
		fmt.Printf("%d: %+v\n", i, examination)
	}
}
