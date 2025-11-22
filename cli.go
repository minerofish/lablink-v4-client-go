// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/lablink-v4-client-go/internal/apijson"
	"github.com/stainless-sdks/lablink-v4-client-go/internal/requestconfig"
	"github.com/stainless-sdks/lablink-v4-client-go/option"
	"github.com/stainless-sdks/lablink-v4-client-go/packages/respjson"
)

// CliService contains methods and other services that help with interacting with
// the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCliService] method instead.
type CliService struct {
	Options []option.RequestOption
}

// NewCliService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCliService(opts ...option.RequestOption) (r CliService) {
	r = CliService{}
	r.Options = opts
	return
}

// Return order results by the tag
func (r *CliService) GetOrdersByTag(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *[]CliGetOrdersByTagResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organizationId parameter")
		return
	}
	path := fmt.Sprintf("v4/cli/orders/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CliGetOrdersByTagResponse struct {
	// The blood donor data when type is DONOR
	BloodDonor BloodDonor `json:"bloodDonor"`
	// The bone-marrow donor data when type is BONE_MARROW_DONOR
	BoneMarrowDonor BoneMarrowDonor `json:"boneMarrowDonor"`
	// The location ID
	LocationID string `json:"location_id" format:"uuid"`
	// The order ID
	OrderID string `json:"order_id" format:"uuid"`
	// The patient data when type is PERSONAL
	Patient Patient `json:"patient"`
	// The pseudonym data when type is PSEUDONYMIZED
	Pseudonym Pseudonym `json:"pseudonym"`
	// The reference of the order created by Lablink
	Reference string `json:"reference"`
	// The results belonging to the order
	Results []CliGetOrdersByTagResponseResult `json:"results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BloodDonor      respjson.Field
		BoneMarrowDonor respjson.Field
		LocationID      respjson.Field
		OrderID         respjson.Field
		Patient         respjson.Field
		Pseudonym       respjson.Field
		Reference       respjson.Field
		Results         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CliGetOrdersByTagResponse) RawJSON() string { return r.JSON.raw }
func (r *CliGetOrdersByTagResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CliGetOrdersByTagResponseResult struct {
	// The ID
	ID string `json:"id" format:"uuid"`
	// The analyte name
	Analyte string `json:"analyte"`
	// The organization examination code
	ExaminationCode string `json:"examinationCode"`
	// The result
	Result string `json:"result"`
	// The result type
	ResultType string `json:"resultType"`
	// The result yield date-time
	ResultYieldDateTime time.Time `json:"resultYieldDateTime" format:"date-time"`
	// The result status
	//
	// Any of "FIN", "PRE".
	Status OrderResultStatus `json:"status"`
	// The unit
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Analyte             respjson.Field
		ExaminationCode     respjson.Field
		Result              respjson.Field
		ResultType          respjson.Field
		ResultYieldDateTime respjson.Field
		Status              respjson.Field
		Unit                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CliGetOrdersByTagResponseResult) RawJSON() string { return r.JSON.raw }
func (r *CliGetOrdersByTagResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
