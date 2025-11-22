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

	"github.com/stainless-sdks/lablink-v4-client-go/internal/apijson"
	"github.com/stainless-sdks/lablink-v4-client-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/lablink-v4-client-go/internal/encoding/json"
	"github.com/stainless-sdks/lablink-v4-client-go/internal/requestconfig"
	"github.com/stainless-sdks/lablink-v4-client-go/option"
	"github.com/stainless-sdks/lablink-v4-client-go/packages/param"
	"github.com/stainless-sdks/lablink-v4-client-go/packages/respjson"
)

// OrganizationService contains methods and other services that help with
// interacting with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options []option.RequestOption
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r OrganizationService) {
	r = OrganizationService{}
	r.Options = opts
	return
}

// Creates one or more organizations for the authenticated user.
func (r *OrganizationService) New(ctx context.Context, body OrganizationNewParams, opts ...option.RequestOption) (res *[]Organization, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates a single organization owned by the authenticated user.
func (r *OrganizationService) Update(ctx context.Context, organizationID string, body OrganizationUpdateParams, opts ...option.RequestOption) (res *Organization, err error) {
	opts = slices.Concat(r.Options, opts)
	if organizationID == "" {
		err = errors.New("missing required organizationId parameter")
		return
	}
	path := fmt.Sprintf("v4/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieves all organizations for the authenticated user.
func (r *OrganizationService) List(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) (res *[]Organization, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Query Organizations with filters in request body
func (r *OrganizationService) Query(ctx context.Context, params OrganizationQueryParams, opts ...option.RequestOption) (res *OrganizationQueryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/organizations/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type Organization struct {
	ID                   string                           `json:"id,required" format:"uuid"`
	Name                 string                           `json:"name,required"`
	ContractLaboratories []OrganizationContractLaboratory `json:"contract_laboratories"`
	CreatedBy            string                           `json:"created_by"`
	Locations            []Location                       `json:"locations"`
	SampleCodeRanges     []OrganizationSampleCodeRange    `json:"sample_code_ranges"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		Name                 respjson.Field
		ContractLaboratories respjson.Field
		CreatedBy            respjson.Field
		Locations            respjson.Field
		SampleCodeRanges     respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Organization) RawJSON() string { return r.JSON.raw }
func (r *Organization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationContractLaboratory struct {
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
func (r OrganizationContractLaboratory) RawJSON() string { return r.JSON.raw }
func (r *OrganizationContractLaboratory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationSampleCodeRange struct {
	Max            int64  `json:"max,required"`
	Min            int64  `json:"min,required"`
	OrganizationID string `json:"organizationId,required" format:"uuid"`
	Prefix         string `json:"prefix,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Max            respjson.Field
		Min            respjson.Field
		OrganizationID respjson.Field
		Prefix         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrganizationSampleCodeRange) RawJSON() string { return r.JSON.raw }
func (r *OrganizationSampleCodeRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationRole string

const (
	OrganizationRoleOrganizationAdmin OrganizationRole = "ORGANIZATION_ADMIN"
	OrganizationRoleUser              OrganizationRole = "USER"
	OrganizationRoleAPIUser           OrganizationRole = "API_USER"
)

type OrganizationQueryResponse struct {
	Items []Organization `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Page
}

// Returns the unmodified JSON received from the API
func (r OrganizationQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *OrganizationQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationNewParams struct {
	Body []OrganizationNewParamsBody
	paramObj
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *OrganizationNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The property Name is required.
type OrganizationNewParamsBody struct {
	// The name of the organization
	Name string `json:"name,required"`
	paramObj
}

func (r OrganizationNewParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationNewParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationNewParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationUpdateParams struct {
	Name string `json:"name,required"`
	paramObj
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrganizationListParams struct {
	Name           param.Opt[string]  `query:"name,omitzero" json:"-"`
	OrganizationID param.Opt[string]  `query:"organizationID,omitzero" format:"uuid" json:"-"`
	Page           param.Opt[int64]   `query:"page,omitzero" json:"-"`
	PageSize       param.Opt[int64]   `query:"pageSize,omitzero" json:"-"`
	Email          []string           `query:"email,omitzero" format:"email" json:"-"`
	Roles          []OrganizationRole `query:"roles,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OrganizationListParams]'s query parameters as `url.Values`.
func (r OrganizationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationQueryParams struct {
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

func (r OrganizationQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow OrganizationQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrganizationQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [OrganizationQueryParams]'s query parameters as
// `url.Values`.
func (r OrganizationQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
