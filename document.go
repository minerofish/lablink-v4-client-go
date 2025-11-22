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

// DocumentService contains methods and other services that help with interacting
// with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentService] method instead.
type DocumentService struct {
	Options []option.RequestOption
}

// NewDocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentService(opts ...option.RequestOption) (r DocumentService) {
	r = DocumentService{}
	r.Options = opts
	return
}

// Gets a document of the authenticated user.
func (r *DocumentService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v4/documents/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets all document metadata of the authenticated user.
func (r *DocumentService) List(ctx context.Context, query DocumentListParams, opts ...option.RequestOption) (res *DocumentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/documents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Page struct {
	// The actual page number
	CurrentPage int64 `json:"currentPage"`
	// The items
	Items any `json:"items"`
	// The number of items per page
	PageSize int64 `json:"pageSize"`
	// The total count of items
	TotalCount int64 `json:"totalCount"`
	// The total pages
	TotalPages int64 `json:"totalPages"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentPage respjson.Field
		Items       respjson.Field
		PageSize    respjson.Field
		TotalCount  respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Page) RawJSON() string { return r.JSON.raw }
func (r *Page) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentListResponse struct {
	Items []DocumentListResponseItem `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Page
}

// Returns the unmodified JSON received from the API
func (r DocumentListResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentListResponseItem struct {
	// The document ID
	ID string `json:"id" format:"uuid"`
	// The filename
	Filename string `json:"filename"`
	// The file size (in bytes)
	FileSize int64 `json:"fileSize"`
	// The file type
	FileType string `json:"fileType"`
	// The links to the file
	Links Link `json:"links"`
	// The time when it got stored
	StoredAt time.Time `json:"storedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Filename    respjson.Field
		FileSize    respjson.Field
		FileType    respjson.Field
		Links       respjson.Field
		StoredAt    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *DocumentListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentListParams struct {
	// The desired page number
	Page int64 `query:"page,required" json:"-"`
	// The desired number of items per page
	PageSize int64 `query:"pageSize,required" json:"-"`
	// The document storage date-time filter
	DocumentStoredAfter param.Opt[string] `query:"documentStoredAfter,omitzero" json:"-"`
	// The order creation date-time filter
	OrderCreatedAfter param.Opt[string] `query:"orderCreatedAfter,omitzero" json:"-"`
	// The order examination ID filter
	OrderExaminationID param.Opt[string] `query:"orderExaminationId,omitzero" json:"-"`
	// The order ID filter
	OrderID param.Opt[string] `query:"orderId,omitzero" json:"-"`
	// The state filter
	//
	// Any of "CONFIRMATORY", "DELETED", "ENTERED", "FINALIZED", "PROCESSING",
	// "WAITING_FOR_MATERIAL".
	OrderStateType DocumentListParamsOrderStateType `query:"orderStateType,omitzero" json:"-"`
	// The type filter
	//
	// Any of "DONOR", "BONE_MARROW_DONOR", "PERSONAL", "PSEUDONYMIZED".
	OrderType DocumentListParamsOrderType `query:"orderType,omitzero" json:"-"`
	// The sorting parameters in the format of "fieldName,asc/desc". E.g. type,desc
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DocumentListParams]'s query parameters as `url.Values`.
func (r DocumentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The state filter
type DocumentListParamsOrderStateType string

const (
	DocumentListParamsOrderStateTypeConfirmatory       DocumentListParamsOrderStateType = "CONFIRMATORY"
	DocumentListParamsOrderStateTypeDeleted            DocumentListParamsOrderStateType = "DELETED"
	DocumentListParamsOrderStateTypeEntered            DocumentListParamsOrderStateType = "ENTERED"
	DocumentListParamsOrderStateTypeFinalized          DocumentListParamsOrderStateType = "FINALIZED"
	DocumentListParamsOrderStateTypeProcessing         DocumentListParamsOrderStateType = "PROCESSING"
	DocumentListParamsOrderStateTypeWaitingForMaterial DocumentListParamsOrderStateType = "WAITING_FOR_MATERIAL"
)

// The type filter
type DocumentListParamsOrderType string

const (
	DocumentListParamsOrderTypeDonor           DocumentListParamsOrderType = "DONOR"
	DocumentListParamsOrderTypeBoneMarrowDonor DocumentListParamsOrderType = "BONE_MARROW_DONOR"
	DocumentListParamsOrderTypePersonal        DocumentListParamsOrderType = "PERSONAL"
	DocumentListParamsOrderTypePseudonymized   DocumentListParamsOrderType = "PSEUDONYMIZED"
)
