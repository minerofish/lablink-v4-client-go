// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/lablink-v4-client-go/internal/apijson"
	"github.com/stainless-sdks/lablink-v4-client-go/internal/apiquery"
	"github.com/stainless-sdks/lablink-v4-client-go/internal/requestconfig"
	"github.com/stainless-sdks/lablink-v4-client-go/option"
	"github.com/stainless-sdks/lablink-v4-client-go/packages/param"
	"github.com/stainless-sdks/lablink-v4-client-go/packages/respjson"
)

// OrderStateService contains methods and other services that help with interacting
// with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrderStateService] method instead.
type OrderStateService struct {
	Options []option.RequestOption
}

// NewOrderStateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrderStateService(opts ...option.RequestOption) (r OrderStateService) {
	r = OrderStateService{}
	r.Options = opts
	return
}

// Gets the state of all orders.
func (r *OrderStateService) List(ctx context.Context, query OrderStateListParams, opts ...option.RequestOption) (res *OrderStateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/order-states"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OrderState struct {
	// The external unique identifier of the order
	ID string `json:"id,required"`
	// The order state
	//
	// Any of "ACCEPTED", "CONFIRMATORY", "DELETED", "ENTERED", "ERROR", "FINALIZED",
	// "PENDING", "PROCESSING", "WAITING_FOR_MATERIAL".
	State OrderStateType `json:"state,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderState) RawJSON() string { return r.JSON.raw }
func (r *OrderState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderStateListResponse struct {
	Items []OrderState `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Page
}

// Returns the unmodified JSON received from the API
func (r OrderStateListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrderStateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderStateListParams struct {
	// The desired page number
	Page int64 `query:"page,required" json:"-"`
	// The desired number of items per page
	PageSize int64 `query:"pageSize,required" json:"-"`
	// The created-from filter (inclusive). Format: yyyy-MM-dd.
	CreatedFrom param.Opt[string] `query:"createdFrom,omitzero" json:"-"`
	// The created-to filter (exclusive). Format: yyyy-MM-dd.
	CreatedTo param.Opt[string] `query:"createdTo,omitzero" json:"-"`
	// The searched content in the different dataset
	SearchContent param.Opt[string] `query:"searchContent,omitzero" json:"-"`
	// The sorting parameters in the format of "fieldName,asc/desc". E.g. type,desc
	Sort []string `query:"sort,omitzero" json:"-"`
	// The state filter
	//
	// Any of "CONFIRMATORY", "DELETED", "ENTERED", "FINALIZED", "PROCESSING",
	// "WAITING_FOR_MATERIAL".
	State OrderStateListParamsState `query:"state,omitzero" json:"-"`
	// The type filter
	//
	// Any of "DONOR", "BONE_MARROW_DONOR", "PERSONAL", "PSEUDONYMIZED".
	Type OrderStateListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrderStateListParams]'s query parameters as `url.Values`.
func (r OrderStateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The state filter
type OrderStateListParamsState string

const (
	OrderStateListParamsStateConfirmatory       OrderStateListParamsState = "CONFIRMATORY"
	OrderStateListParamsStateDeleted            OrderStateListParamsState = "DELETED"
	OrderStateListParamsStateEntered            OrderStateListParamsState = "ENTERED"
	OrderStateListParamsStateFinalized          OrderStateListParamsState = "FINALIZED"
	OrderStateListParamsStateProcessing         OrderStateListParamsState = "PROCESSING"
	OrderStateListParamsStateWaitingForMaterial OrderStateListParamsState = "WAITING_FOR_MATERIAL"
)

// The type filter
type OrderStateListParamsType string

const (
	OrderStateListParamsTypeDonor           OrderStateListParamsType = "DONOR"
	OrderStateListParamsTypeBoneMarrowDonor OrderStateListParamsType = "BONE_MARROW_DONOR"
	OrderStateListParamsTypePersonal        OrderStateListParamsType = "PERSONAL"
	OrderStateListParamsTypePseudonymized   OrderStateListParamsType = "PSEUDONYMIZED"
)
