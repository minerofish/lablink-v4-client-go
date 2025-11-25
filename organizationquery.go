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

// OrganizationQueryService contains methods and other services that help with
// interacting with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationQueryService] method instead.
type OrganizationQueryService struct {
	Options []option.RequestOption
}

// NewOrganizationQueryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationQueryService(opts ...option.RequestOption) (r OrganizationQueryService) {
	r = OrganizationQueryService{}
	r.Options = opts
	return
}

// Query Organizations with filters in request body
func (r *OrganizationQueryService) New(ctx context.Context, params OrganizationQueryNewParams, opts ...option.RequestOption) (res *OrganizationQueryNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/organizations-query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type OrganizationQueryNewResponse struct {
	// The actual page number
	CurrentPage int64                              `json:"currentPage"`
	Items       []OrganizationQueryNewResponseItem `json:"items"`
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
func (r OrganizationQueryNewResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationQueryNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationQueryNewResponseItem struct {
	ID                   string                                               `json:"id,required" format:"uuid"`
	Name                 string                                               `json:"name,required"`
	ContractLaboratories []OrganizationQueryNewResponseItemContractLaboratory `json:"contract_laboratories"`
	CreatedBy            string                                               `json:"created_by"`
	Locations            []Location                                           `json:"locations"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Name                 respjson.Field
		ContractLaboratories respjson.Field
		CreatedBy            respjson.Field
		Locations            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationQueryNewResponseItem) RawJSON() string { return r.JSON.raw }
func (r *OrganizationQueryNewResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationQueryNewResponseItemContractLaboratory struct {
	ID           string                    `json:"id" format:"uuid"`
	Address      string                    `json:"address"`
	City         string                    `json:"city"`
	Country      string                    `json:"country"`
	CreatedBy    string                    `json:"createdBy"`
	Email        string                    `json:"email"`
	InstanceName string                    `json:"instanceName"`
	Name         string                    `json:"name"`
	Phone        string                    `json:"phone"`
	Postcode     string                    `json:"postcode"`
	UserRoles    []LaboratoryUserRelations `json:"userRoles"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Address      respjson.Field
		City         respjson.Field
		Country      respjson.Field
		CreatedBy    respjson.Field
		Email        respjson.Field
		InstanceName respjson.Field
		Name         respjson.Field
		Phone        respjson.Field
		Postcode     respjson.Field
		UserRoles    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationQueryNewResponseItemContractLaboratory) RawJSON() string { return r.JSON.raw }
func (r *OrganizationQueryNewResponseItemContractLaboratory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationQueryNewParams struct {
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64]  `query:"pageSize,omitzero" json:"-"`
	Name     param.Opt[string] `json:"name,omitzero"`
	// The sorting parameters in the format of "fieldName,asc/desc". E.g. state,desc
	//
	// Any of "name,asc", "name,desc", "created_at,asc", "created_at,desc".
	Sort   []string `query:"sort,omitzero" json:"-"`
	Emails []string `json:"emails,omitzero" format:"email"`
	// Filter by order IDs
	OrganizationIDs []string           `json:"organizationIDs,omitzero" format:"uuid"`
	Roles           []OrganizationRole `json:"roles,omitzero"`
	paramObj
}

func (r OrganizationQueryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationQueryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationQueryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OrganizationQueryNewParams]'s query parameters as
// `url.Values`.
func (r OrganizationQueryNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
