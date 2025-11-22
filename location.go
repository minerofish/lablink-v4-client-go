// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"

	"github.com/minerofish/lablink-v4-client-go/internal/apijson"
	"github.com/minerofish/lablink-v4-client-go/internal/apiquery"
	shimjson "github.com/minerofish/lablink-v4-client-go/internal/encoding/json"
	"github.com/minerofish/lablink-v4-client-go/internal/requestconfig"
	"github.com/minerofish/lablink-v4-client-go/option"
	"github.com/minerofish/lablink-v4-client-go/packages/param"
	"github.com/minerofish/lablink-v4-client-go/packages/respjson"
)

// LocationService contains methods and other services that help with interacting
// with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLocationService] method instead.
type LocationService struct {
	Options []option.RequestOption
}

// NewLocationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLocationService(opts ...option.RequestOption) (r LocationService) {
	r = LocationService{}
	r.Options = opts
	return
}

// Creates new locations for the authenticated user.
func (r *LocationService) New(ctx context.Context, body LocationNewParams, opts ...option.RequestOption) (res *[]Location, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets all location of the authenticated user.
func (r *LocationService) List(ctx context.Context, query LocationListParams, opts ...option.RequestOption) (res *LocationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a location for the authenticated user.
func (r *LocationService) Delete(ctx context.Context, body LocationDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v4/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

type Location struct {
	// The ID
	ID string `json:"id,required" format:"uuid"`
	// The name
	Name string `json:"name,required"`
	// The organization ID which the location belongs to
	OrganizationID string `json:"organizationId,required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Name           respjson.Field
		OrganizationID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Location) RawJSON() string { return r.JSON.raw }
func (r *Location) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationListResponse struct {
	Items []Location `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Page
}

// Returns the unmodified JSON received from the API
func (r LocationListResponse) RawJSON() string { return r.JSON.raw }
func (r *LocationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationNewParams struct {
	Body []LocationNewParamsBody
	paramObj
}

func (r LocationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *LocationNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Request body for creating a new location
//
// The properties Name, OrganizationID are required.
type LocationNewParamsBody struct {
	Name           string `json:"name,required"`
	OrganizationID string `json:"organizationId,required" format:"uuid"`
	paramObj
}

func (r LocationNewParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LocationNewParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LocationNewParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LocationListParams struct {
	// The name filter
	FilterName param.Opt[string] `query:"filterName,omitzero" json:"-"`
	// The desired page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// The desired number of items per page
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// The sorting parameters in the format of "fieldName,asc/desc". E.g. type,desc
	//
	// Any of "name,asc", "name,desc".
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LocationListParams]'s query parameters as `url.Values`.
func (r LocationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LocationDeleteParams struct {
	// The location ID
	LocationID string `query:"locationId,required" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [LocationDeleteParams]'s query parameters as `url.Values`.
func (r LocationDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
