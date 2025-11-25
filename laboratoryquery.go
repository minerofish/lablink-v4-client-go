// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/minerofish/lablink-v4-client-go/internal/apijson"
	"github.com/minerofish/lablink-v4-client-go/internal/apiquery"
	"github.com/minerofish/lablink-v4-client-go/internal/requestconfig"
	"github.com/minerofish/lablink-v4-client-go/option"
	"github.com/minerofish/lablink-v4-client-go/packages/param"
	"github.com/minerofish/lablink-v4-client-go/packages/respjson"
)

// LaboratoryQueryService contains methods and other services that help with
// interacting with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaboratoryQueryService] method instead.
type LaboratoryQueryService struct {
	Options []option.RequestOption
}

// NewLaboratoryQueryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLaboratoryQueryService(opts ...option.RequestOption) (r LaboratoryQueryService) {
	r = LaboratoryQueryService{}
	r.Options = opts
	return
}

// Queries laboratories. The caller must be user for the respective laboratories.
func (r *LaboratoryQueryService) Execute(ctx context.Context, params LaboratoryQueryExecuteParams, opts ...option.RequestOption) (res *LaboratoryQueryExecuteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/laboratories-query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type LaboratoryQueryExecuteResponse struct {
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
func (r LaboratoryQueryExecuteResponse) RawJSON() string { return r.JSON.raw }
func (r *LaboratoryQueryExecuteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryQueryExecuteParams struct {
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64]  `query:"pageSize,omitzero" json:"-"`
	Name     param.Opt[string] `json:"name,omitzero"`
	// The sorting parameters in the format of "fieldName,asc/desc". E.g. state,desc
	//
	// Any of "name,asc", "name,desc", "createdAt,asc", "createdAt,desc".
	Sort          []string `query:"sort,omitzero" json:"-"`
	LaboratoryIDs []string `json:"laboratoryIDs,omitzero" format:"uuid"`
	paramObj
}

func (r LaboratoryQueryExecuteParams) MarshalJSON() (data []byte, err error) {
	type shadow LaboratoryQueryExecuteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaboratoryQueryExecuteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [LaboratoryQueryExecuteParams]'s query parameters as
// `url.Values`.
func (r LaboratoryQueryExecuteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
