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

// OrderQueryService contains methods and other services that help with interacting
// with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrderQueryService] method instead.
type OrderQueryService struct {
	Options []option.RequestOption
}

// NewOrderQueryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrderQueryService(opts ...option.RequestOption) (r OrderQueryService) {
	r = OrderQueryService{}
	r.Options = opts
	return
}

// Query orders with filters in request body
func (r *OrderQueryService) New(ctx context.Context, params OrderQueryNewParams, opts ...option.RequestOption) (res *OrderQueryNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/orders-query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Commit or rollback a transaction started with /query?transaction=.... Calling
// with SetTags or RemoveTags the order can be tagged.
func (r *OrderQueryService) CommitTransaction(ctx context.Context, transactionID string, body OrderQueryCommitTransactionParams, opts ...option.RequestOption) (res *[]string, err error) {
	opts = slices.Concat(r.Options, opts)
	if transactionID == "" {
		err = errors.New("missing required transactionId parameter")
		return
	}
	path := fmt.Sprintf("v4/orders-query/%s", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type OrderQueryNewResponse struct {
	// The actual page number
	CurrentPage int64                       `json:"currentPage"`
	Items       []OrderQueryNewResponseItem `json:"items"`
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
func (r OrderQueryNewResponse) RawJSON() string { return r.JSON.raw }
func (r *OrderQueryNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderQueryNewResponseItem struct {
	// The ID of the order
	ID string `json:"id,required" format:"uuid"`
	// The blood donor data when type is DONOR
	BloodDonor BloodDonor `json:"bloodDonor"`
	// The bone-marrow donor data when type is BONE_MARROW_DONOR
	BoneMarrowDonor BoneMarrowDonor `json:"boneMarrowDonor"`
	// The documents associated with the order
	Documents []OrderQueryNewResponseItemDocument `json:"documents"`
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
func (r OrderQueryNewResponseItem) RawJSON() string { return r.JSON.raw }
func (r *OrderQueryNewResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderQueryNewResponseItemDocument struct {
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
func (r OrderQueryNewResponseItemDocument) RawJSON() string { return r.JSON.raw }
func (r *OrderQueryNewResponseItemDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderQueryNewParams struct {
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Fields to expand in the response
	//
	// Any of "documents", "state", "items", "details", "full".
	Expand []string `query:"expand,omitzero" json:"-"`
	// The sorting parameters in the format of "fieldName,asc/desc". E.g. state,desc
	//
	// Any of "state,asc", "state,desc", "reference,asc", "reference,desc",
	// "createdAt,asc", "createdAt,desc".
	Sort []string `query:"sort,omitzero" json:"-"`
	// Transaction mode. Use the two-stage mode for transactions that require approval.
	// Two-Stage is not locking the order and the last reader will win.
	//
	// Any of "none", "twostage".
	Transaction OrderQueryNewParamsTransaction `query:"transaction,omitzero" json:"-"`
	// Filter by order IDs
	IDs []string `json:"ids,omitzero" format:"uuid"`
	// Filter by laboratory IDs
	LaboratoryIDs []string `json:"laboratoryIds,omitzero" format:"uuid"`
	// Filter by location IDs
	LocationIDs []string `json:"locationIds,omitzero" format:"uuid"`
	// Filter by organization IDs
	OrganizationIDs []string `json:"organizationIds,omitzero" format:"uuid"`
	// Filter by sample codes
	SampleCodes []string `json:"sampleCodes,omitzero"`
	// Filter by order states
	State []OrderStateType `json:"state,omitzero"`
	// Orders must not have any of these tags
	WithoutTags []string `json:"withoutTags,omitzero"`
	// Orders must have all of these tags
	WithTags []string `json:"withTags,omitzero"`
	paramObj
}

func (r OrderQueryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OrderQueryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderQueryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OrderQueryNewParams]'s query parameters as `url.Values`.
func (r OrderQueryNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Transaction mode. Use the two-stage mode for transactions that require approval.
// Two-Stage is not locking the order and the last reader will win.
type OrderQueryNewParamsTransaction string

const (
	OrderQueryNewParamsTransactionNone     OrderQueryNewParamsTransaction = "none"
	OrderQueryNewParamsTransactionTwostage OrderQueryNewParamsTransaction = "twostage"
)

type OrderQueryCommitTransactionParams struct {
	// Tags to be added to all orders
	AddTags []string `json:"addTags,omitzero"`
	// Tags to be removed from all orders
	RemoveTags []string `json:"removeTags,omitzero"`
	// Update order Status (requires permissions)
	//
	// Any of "ENTERED", "WAITING_FOR_MATERIAL", "PROCESSING", "CONFIRMATION_PENDING",
	// "FINAL", "DELETED", "ERROR".
	Setstate OrderStateType `json:"setstate,omitzero"`
	paramObj
}

func (r OrderQueryCommitTransactionParams) MarshalJSON() (data []byte, err error) {
	type shadow OrderQueryCommitTransactionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderQueryCommitTransactionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
