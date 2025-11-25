// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/minerofish/lablink-v4-client-go/internal/apijson"
	"github.com/minerofish/lablink-v4-client-go/internal/apiquery"
	shimjson "github.com/minerofish/lablink-v4-client-go/internal/encoding/json"
	"github.com/minerofish/lablink-v4-client-go/internal/requestconfig"
	"github.com/minerofish/lablink-v4-client-go/option"
	"github.com/minerofish/lablink-v4-client-go/packages/param"
	"github.com/minerofish/lablink-v4-client-go/packages/respjson"
)

// OrderExaminationService contains methods and other services that help with
// interacting with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrderExaminationService] method instead.
type OrderExaminationService struct {
	Options []option.RequestOption
}

// NewOrderExaminationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrderExaminationService(opts ...option.RequestOption) (r OrderExaminationService) {
	r = OrderExaminationService{}
	r.Options = opts
	return
}

// Get all items of an order
func (r *OrderExaminationService) List(ctx context.Context, orderID string, opts ...option.RequestOption) (res *[]OrderExaminationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return
	}
	path := fmt.Sprintf("v4/orders/%s/examinations", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Add an item to an order
func (r *OrderExaminationService) Add(ctx context.Context, orderID string, body OrderExaminationAddParams, opts ...option.RequestOption) (res *[]OrderExaminationAddResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return
	}
	path := fmt.Sprintf("v4/orders/%s/examinations", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Remove an item from an order. When the last item is removed, the whole order is
// deleted.
func (r *OrderExaminationService) Remove(ctx context.Context, orderID string, body OrderExaminationRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return
	}
	path := fmt.Sprintf("v4/orders/%s/examinations", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type OrderExaminationListResponse struct {
	// Custom code for this examination
	Code string `json:"code,required"`
	// The sample time
	CollectionTime time.Time `json:"collectionTime,required" format:"date-time"`
	// The description of the organization examination
	Description string `json:"description,required"`
	// Reference to Examination.
	ExaminationID string `json:"examinationId,required" format:"uuid"`
	// Reference to Order Item.
	ItemID string `json:"itemId,required" format:"uuid"`
	// The name of the organization examination
	Name string `json:"name,required"`
	// The sample codes
	SampleCode string `json:"sampleCode,required"`
	// The unit of the organization examination
	Unit string `json:"unit,required"`
	// The result if available
	Result OrderExaminationListResponseResult `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code           respjson.Field
		CollectionTime respjson.Field
		Description    respjson.Field
		ExaminationID  respjson.Field
		ItemID         respjson.Field
		Name           respjson.Field
		SampleCode     respjson.Field
		Unit           respjson.Field
		Result         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderExaminationListResponse) RawJSON() string { return r.JSON.raw }
func (r *OrderExaminationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The result if available
type OrderExaminationListResponseResult struct {
	// The order-examination ID.
	ExaminationID string `json:"examinationId,required" format:"uuid"`
	// The order-examination item ID
	ItemID string `json:"itemId,required" format:"uuid"`
	// The order ID
	OrderID string `json:"orderId,required" format:"uuid"`
	Result  string `json:"result,required"`
	// Any of "FINAL", "PRELIMINARY", "CORRECTED".
	Status string `json:"status,required"`
	// Indicates if the result requires confirmation.
	ConfirmationPending bool `json:"confirmationPending"`
	// The data type of the result
	//
	// Any of "int", "decimal", "string", "pein", "react", "invalid", "enum".
	DataType    string    `json:"dataType"`
	InfoText    string    `json:"infoText"`
	PerformedAt time.Time `json:"performedAt" format:"date-time"`
	Unit        string    `json:"unit"`
	ValidatedAt time.Time `json:"validatedAt" format:"date-time"`
	ValidatedBy string    `json:"validatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExaminationID       respjson.Field
		ItemID              respjson.Field
		OrderID             respjson.Field
		Result              respjson.Field
		Status              respjson.Field
		ConfirmationPending respjson.Field
		DataType            respjson.Field
		InfoText            respjson.Field
		PerformedAt         respjson.Field
		Unit                respjson.Field
		ValidatedAt         respjson.Field
		ValidatedBy         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderExaminationListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *OrderExaminationListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderExaminationAddResponse struct {
	// Reference to Order Item.
	ID string `json:"id,required" format:"uuid"`
	// Custom code for this examination
	Code string `json:"code,required"`
	// The sample time
	CollectionTime time.Time `json:"collectionTime,required" format:"date-time"`
	// The description of the organization examination
	Description string `json:"description,required"`
	// The name of the organization examination
	Name string `json:"name,required"`
	// The results belonging to the order
	Results []OrderExaminationAddResponseResult `json:"results,required"`
	// The sample codes
	SampleCode string `json:"sampleCode,required"`
	// The unit of the organization examination
	Unit string `json:"unit,required"`
	// Reference to Examination.
	ExaminationID string `json:"examinationId" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Code           respjson.Field
		CollectionTime respjson.Field
		Description    respjson.Field
		Name           respjson.Field
		Results        respjson.Field
		SampleCode     respjson.Field
		Unit           respjson.Field
		ExaminationID  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderExaminationAddResponse) RawJSON() string { return r.JSON.raw }
func (r *OrderExaminationAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderExaminationAddResponseResult struct {
	// The order-examination ID.
	ExaminationID string `json:"examinationId,required" format:"uuid"`
	// The order-examination item ID
	ItemID string `json:"itemId,required" format:"uuid"`
	// The order ID
	OrderID string `json:"orderId,required" format:"uuid"`
	Result  string `json:"result,required"`
	// Any of "FINAL", "PRELIMINARY", "CORRECTED".
	Status string `json:"status,required"`
	// Indicates if the result requires confirmation.
	ConfirmationPending bool `json:"confirmationPending"`
	// The data type of the result
	//
	// Any of "int", "decimal", "string", "pein", "react", "invalid", "enum".
	DataType    string    `json:"dataType"`
	InfoText    string    `json:"infoText"`
	PerformedAt time.Time `json:"performedAt" format:"date-time"`
	Unit        string    `json:"unit"`
	ValidatedAt time.Time `json:"validatedAt" format:"date-time"`
	ValidatedBy string    `json:"validatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExaminationID       respjson.Field
		ItemID              respjson.Field
		OrderID             respjson.Field
		Result              respjson.Field
		Status              respjson.Field
		ConfirmationPending respjson.Field
		DataType            respjson.Field
		InfoText            respjson.Field
		PerformedAt         respjson.Field
		Unit                respjson.Field
		ValidatedAt         respjson.Field
		ValidatedBy         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderExaminationAddResponseResult) RawJSON() string { return r.JSON.raw }
func (r *OrderExaminationAddResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderExaminationAddParams struct {
	Body []OrderExaminationAddParamsBody
	paramObj
}

func (r OrderExaminationAddParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *OrderExaminationAddParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The properties CollectionTime, SampleCode are required.
type OrderExaminationAddParamsBody struct {
	// The collection time of the sample
	CollectionTime time.Time `json:"collectionTime,required" format:"date-time"`
	// The sample code
	SampleCode string `json:"sampleCode,required"`
	// The code of the Examination (you have to use either examinationId or code)
	ExaminationCode param.Opt[string] `json:"examinationCode,omitzero"`
	// Reference to Examination ID (you have to use either examinationId or code)
	ExaminationID param.Opt[string] `json:"examinationId,omitzero" format:"uuid"`
	// The external unique identifier of the order's item
	Reference param.Opt[string] `json:"reference,omitzero"`
	paramObj
}

func (r OrderExaminationAddParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow OrderExaminationAddParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderExaminationAddParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderExaminationRemoveParams struct {
	// List of items to be removed
	ItemID []string `query:"itemId,omitzero,required" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [OrderExaminationRemoveParams]'s query parameters as
// `url.Values`.
func (r OrderExaminationRemoveParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
