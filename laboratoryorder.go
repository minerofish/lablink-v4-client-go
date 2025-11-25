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
	CurrentPage int64                             `json:"currentPage"`
	Items       []LaboratoryOrderListResponseItem `json:"items"`
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

type LaboratoryOrderListResponseItem struct {
	// The ID of the order
	ID string `json:"id,required" format:"uuid"`
	// The blood donor data when type is DONOR
	BloodDonor BloodDonor `json:"bloodDonor"`
	// The bone-marrow donor data when type is BONE_MARROW_DONOR
	BoneMarrowDonor BoneMarrowDonor `json:"boneMarrowDonor"`
	// The documents associated with the order
	Documents []LaboratoryOrderListResponseItemDocument `json:"documents"`
	// The items belonging to the order. Each item represents one examination.
	Items []OrderExamination `json:"items"`
	// The laboratory ID where the order will be sent
	LaboratoryID string `json:"laboratoryId" format:"uuid"`
	// Identifier of the location
	LocationID string `json:"locationId" format:"uuid"`
	// The order creation date-time (yyyy-MM-dd'T'HH:mm:ss.SSSZ)
	OrderCreationDateTime time.Time `json:"orderCreationDateTime" format:"date-time"`
	// The tags belonging to the order
	OrderTags []string `json:"orderTags"`
	// The patient data when type is PERSONAL
	Patient Patient `json:"patient"`
	// The pseudonym data when type is PSEUDONYM
	Pseudonym Pseudonym `json:"pseudonym"`
	// Add information in key:value pairs object array that are stored with the order
	References map[string]string `json:"references"`
	// The order status model
	//
	// Any of "ENTERED", "WAITING_FOR_MATERIAL", "PROCESSING", "CONFIRMATION_PENDING",
	// "FINAL", "DELETED", "ERROR".
	State OrderStateType `json:"state"`
	// The order type
	//
	// Any of "DONOR", "BONE_MARROW_DONOR", "PERSONAL", "PSEUDONYM".
	Type OrderType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		BloodDonor            respjson.Field
		BoneMarrowDonor       respjson.Field
		Documents             respjson.Field
		Items                 respjson.Field
		LaboratoryID          respjson.Field
		LocationID            respjson.Field
		OrderCreationDateTime respjson.Field
		OrderTags             respjson.Field
		Patient               respjson.Field
		Pseudonym             respjson.Field
		References            respjson.Field
		State                 respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaboratoryOrderListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *LaboratoryOrderListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryOrderListResponseItemDocument struct {
	// The document ID
	DocumentID string `json:"documentId,required" format:"uuid"`
	// The filename
	FileName string `json:"fileName,required"`
	// The file size (in bytes)
	FileSize int64 `json:"fileSize,required"`
	// The file type
	FileType string `json:"fileType,required"`
	// The links to the file
	Link Link `json:"link,required"`
	// The time when it got stored
	StoredAt time.Time `json:"storedAt,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentID  respjson.Field
		FileName    respjson.Field
		FileSize    respjson.Field
		FileType    respjson.Field
		Link        respjson.Field
		StoredAt    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaboratoryOrderListResponseItemDocument) RawJSON() string { return r.JSON.raw }
func (r *LaboratoryOrderListResponseItemDocument) UnmarshalJSON(data []byte) error {
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
