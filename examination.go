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

// ExaminationService contains methods and other services that help with
// interacting with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExaminationService] method instead.
type ExaminationService struct {
	Options []option.RequestOption
}

// NewExaminationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExaminationService(opts ...option.RequestOption) (r ExaminationService) {
	r = ExaminationService{}
	r.Options = opts
	return
}

// Gets all examination of the authenticated user.
func (r *ExaminationService) List(ctx context.Context, query ExaminationListParams, opts ...option.RequestOption) (res *ExaminationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/examinations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type Examination struct {
	// The organization examination ID
	ID string `json:"id" format:"uuid"`
	// The code
	Code string `json:"code"`
	// The description
	Description string `json:"description"`
	// The laboratory ID
	LaboratoryID string `json:"laboratoryId" format:"uuid"`
	// The name
	Name string `json:"name"`
	// The possible result values
	PossibleValues string `json:"possibleValues"`
	// The unit
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Code           respjson.Field
		Description    respjson.Field
		LaboratoryID   respjson.Field
		Name           respjson.Field
		PossibleValues respjson.Field
		Unit           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Examination) RawJSON() string { return r.JSON.raw }
func (r *Examination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExaminationListResponse struct {
	Items []Examination `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Page
}

// Returns the unmodified JSON received from the API
func (r ExaminationListResponse) RawJSON() string { return r.JSON.raw }
func (r *ExaminationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExaminationListParams struct {
	// The code filter
	Code param.Opt[string] `query:"code,omitzero" json:"-"`
	// The laboratory ID filter
	LaboratoryID param.Opt[string] `query:"laboratoryID,omitzero" format:"uuid" json:"-"`
	// The name filter
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// The organization ID filter
	OrganizationID param.Opt[string] `query:"organizationID,omitzero" format:"uuid" json:"-"`
	// The desired page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// The desired number of items per page
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// The unit filter
	Unit param.Opt[string] `query:"unit,omitzero" json:"-"`
	// The sorting parameters in the format of "fieldName,asc/desc". E.g. type,desc
	//
	// Any of "code,asc", "code,desc", "name,asc", "name,desc", "unit,asc",
	// "unit,desc".
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExaminationListParams]'s query parameters as `url.Values`.
func (r ExaminationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
