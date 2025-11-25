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

	"github.com/minerofish/lablink-v4-client-go/internal/apijson"
	"github.com/minerofish/lablink-v4-client-go/internal/apiquery"
	"github.com/minerofish/lablink-v4-client-go/internal/requestconfig"
	"github.com/minerofish/lablink-v4-client-go/option"
	"github.com/minerofish/lablink-v4-client-go/packages/param"
	"github.com/minerofish/lablink-v4-client-go/packages/respjson"
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

// Downloads orders for a laboratory with pagination. With the download there is a
// job-tag received.
func (r *LaboratoryOrderService) List(ctx context.Context, laboratoryID string, query LaboratoryOrderListParams, opts ...option.RequestOption) (res *LaboratoryOrderListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if laboratoryID == "" {
		err = errors.New("missing required laboratoryId parameter")
		return
	}
	path := fmt.Sprintf("v4/laboratories/%s/orders", laboratoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type LaboratoryOrderListResponse struct {
	// The actual page number
	CurrentPage int64 `json:"currentPage"`
	Items       any   `json:"items"`
	// The number of items per page
	PageSize int64 `json:"pageSize"`
	// The total count of items
	TotalCount int64 `json:"totalCount"`
	// The total pages
	TotalPages int64 `json:"totalPages"`
	// The transaction identifier
	TransactionID string `json:"transactionId" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentPage   respjson.Field
		Items         respjson.Field
		PageSize      respjson.Field
		TotalCount    respjson.Field
		TotalPages    respjson.Field
		TransactionID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaboratoryOrderListResponse) RawJSON() string { return r.JSON.raw }
func (r *LaboratoryOrderListResponse) UnmarshalJSON(data []byte) error {
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
	// Only orders updated since this timestamp are returned
	Since param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
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
