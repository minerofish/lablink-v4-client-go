package main

import (
	"context"
	"errors"
	"fmt"

	lablink "github.com/minerofish/lablink-v4-client-go"
	"github.com/minerofish/lablink-v4-client-go/option"
)

func main() {
	baseURL := "https://lablink.bloodlab.org"
	apiKey := "your-api-key"

	testWithAPIKeyClient(baseURL, apiKey)
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

	if response == nil {
		fmt.Println("response is nil - something went wrong")
		return
	}

	for i, examination := range *response {
		fmt.Printf("%d: %+v\n", i, examination)
	}
}
