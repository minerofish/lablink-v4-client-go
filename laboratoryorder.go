// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/lablink-v4-client-go/internal/apijson"
	"github.com/stainless-sdks/lablink-v4-client-go/internal/apiquery"
	"github.com/stainless-sdks/lablink-v4-client-go/internal/requestconfig"
	"github.com/stainless-sdks/lablink-v4-client-go/option"
	"github.com/stainless-sdks/lablink-v4-client-go/packages/param"
	"github.com/stainless-sdks/lablink-v4-client-go/packages/respjson"
)

// LaboratoryOrderService contains methods and other services that help with
// interacting with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaboratoryOrderService] method instead.
type LaboratoryOrderService struct {
	Options []option.RequestOption
}

// NewLaboratoryOrderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLaboratoryOrderService(opts ...option.RequestOption) (r LaboratoryOrderService) {
	r = LaboratoryOrderService{}
	r.Options = opts
	return
}

// Update the status of orders to PROCESSING or DOWNLOADED
func (r *LaboratoryOrderService) Update(ctx context.Context, laboratoryID string, body LaboratoryOrderUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if laboratoryID == "" {
		err = errors.New("missing required laboratoryId parameter")
		return
	}
	path := fmt.Sprintf("v4/laboratories/%s/orders", laboratoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Retrieve a list of orders
func (r *LaboratoryOrderService) List(ctx context.Context, laboratoryID string, query LaboratoryOrderListParams, opts ...option.RequestOption) (res *[]LaboratoryOrderListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if laboratoryID == "" {
		err = errors.New("missing required laboratoryId parameter")
		return
	}
	path := fmt.Sprintf("v4/laboratories/%s/orders", laboratoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Order struct {
	// The ID of the order
	ID string `json:"id,required" format:"uuid"`
	// The examinations belonging to the order
	Examinations []OrderExamination `json:"examinations,required"`
	// The laboratory ID where the order will be sent
	LaboratoryID string `json:"laboratoryId,required" format:"uuid"`
	// Identifier of the location (client)
	LocationID string `json:"locationId,required" format:"uuid"`
	// The order creation date-time (yyyy-MM-dd'T'HH:mm:ss.SSSZ)
	OrderCreationDateTime time.Time `json:"orderCreationDateTime,required" format:"date-time"`
	// The order state
	//
	// Any of "ACCEPTED", "CONFIRMATORY", "DELETED", "ENTERED", "ERROR", "FINALIZED",
	// "PENDING", "PROCESSING", "WAITING_FOR_MATERIAL".
	State OrderStateType `json:"state,required"`
	// The order type
	//
	// Any of "DONOR", "BONE_MARROW_DONOR", "PERSONAL", "PSEUDONYMIZED".
	Type OrderType `json:"type,required"`
	// The blood donor data when type is DONOR
	BloodDonor BloodDonor `json:"bloodDonor"`
	// The bone-marrow donor data when type is BONE_MARROW_DONOR
	BoneMarrowDonor BoneMarrowDonor `json:"boneMarrowDonor"`
	// The tags belonging to the order
	OrderTags []string `json:"orderTags"`
	// The patient data when type is PERSONAL
	Patient Patient `json:"patient"`
	// The pseudonym data when type is PSEUDONYMIZED
	Pseudonym Pseudonym `json:"pseudonym"`
	// Add information in key:value pairs object array that are stored with the order
	References map[string]string `json:"references"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Examinations          respjson.Field
		LaboratoryID          respjson.Field
		LocationID            respjson.Field
		OrderCreationDateTime respjson.Field
		State                 respjson.Field
		Type                  respjson.Field
		BloodDonor            respjson.Field
		BoneMarrowDonor       respjson.Field
		OrderTags             respjson.Field
		Patient               respjson.Field
		Pseudonym             respjson.Field
		References            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Order) RawJSON() string { return r.JSON.raw }
func (r *Order) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryOrderListResponse struct {
	Examination  LaboratoryOrderListResponseExamination  `json:"examination"`
	Laboratory   LaboratoryOrderListResponseLaboratory   `json:"laboratory"`
	Location     LaboratoryOrderListResponseLocation     `json:"location"`
	Order        Order                                   `json:"Order"`
	Organization LaboratoryOrderListResponseOrganization `json:"organization"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Examination  respjson.Field
		Laboratory   respjson.Field
		Location     respjson.Field
		Order        respjson.Field
		Organization respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaboratoryOrderListResponse) RawJSON() string { return r.JSON.raw }
func (r *LaboratoryOrderListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryOrderListResponseExamination struct {
	ID   string `json:"id" format:"uuid"`
	Code string `json:"code"`
	Name string `json:"name"`
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Code        respjson.Field
		Name        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaboratoryOrderListResponseExamination) RawJSON() string { return r.JSON.raw }
func (r *LaboratoryOrderListResponseExamination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryOrderListResponseLaboratory struct {
	LaboratoryID string `json:"laboratoryId" format:"uuid"`
	Name         string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LaboratoryID respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaboratoryOrderListResponseLaboratory) RawJSON() string { return r.JSON.raw }
func (r *LaboratoryOrderListResponseLaboratory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryOrderListResponseLocation struct {
	ID   string `json:"id" format:"uuid"`
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaboratoryOrderListResponseLocation) RawJSON() string { return r.JSON.raw }
func (r *LaboratoryOrderListResponseLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryOrderListResponseOrganization struct {
	Name           string `json:"name"`
	OrganizationID string `json:"organizationId" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name           respjson.Field
		OrganizationID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaboratoryOrderListResponseOrganization) RawJSON() string { return r.JSON.raw }
func (r *LaboratoryOrderListResponseOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryOrderUpdateParams struct {
	// New state to set for the orders (PROCESSING or DOWNLOADED)
	//
	// Any of "PROCESSING", "DOWNLOADED".
	NewState LaboratoryOrderUpdateParamsNewState `json:"newState,omitzero"`
	// Order IDs to mark as downloaded (required for DOWNLOADED state)
	OrderIDs []string `json:"orderIds,omitzero"`
	// Sample codes for orders to update (required for PROCESSING state)
	SampleCodes []string `json:"sampleCodes,omitzero"`
	paramObj
}

func (r LaboratoryOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow LaboratoryOrderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaboratoryOrderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// New state to set for the orders (PROCESSING or DOWNLOADED)
type LaboratoryOrderUpdateParamsNewState string

const (
	LaboratoryOrderUpdateParamsNewStateProcessing LaboratoryOrderUpdateParamsNewState = "PROCESSING"
	LaboratoryOrderUpdateParamsNewStateDownloaded LaboratoryOrderUpdateParamsNewState = "DOWNLOADED"
)

type LaboratoryOrderListParams struct {
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaboratoryOrderListParams]'s query parameters as
// `url.Values`.
func (r LaboratoryOrderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
