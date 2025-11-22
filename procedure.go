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

	"github.com/minerofish/lablink-v4-client-go/internal/apijson"
	"github.com/minerofish/lablink-v4-client-go/internal/apiquery"
	shimjson "github.com/minerofish/lablink-v4-client-go/internal/encoding/json"
	"github.com/minerofish/lablink-v4-client-go/internal/requestconfig"
	"github.com/minerofish/lablink-v4-client-go/option"
	"github.com/minerofish/lablink-v4-client-go/packages/param"
	"github.com/minerofish/lablink-v4-client-go/packages/respjson"
)

// ProcedureService contains methods and other services that help with interacting
// with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProcedureService] method instead.
type ProcedureService struct {
	Options []option.RequestOption
}

// NewProcedureService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProcedureService(opts ...option.RequestOption) (r ProcedureService) {
	r = ProcedureService{}
	r.Options = opts
	return
}

// Creates a new procedure in the global catalog.
func (r *ProcedureService) New(ctx context.Context, body ProcedureNewParams, opts ...option.RequestOption) (res *ProcedureResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/procedures"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an existing procedure in the global catalog.
func (r *ProcedureService) Update(ctx context.Context, procedureID string, body ProcedureUpdateParams, opts ...option.RequestOption) (res *ProcedureUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if procedureID == "" {
		err = errors.New("missing required procedureId parameter")
		return
	}
	path := fmt.Sprintf("v4/procedures/%s", procedureID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Procedures can be single or multiple laboratory-tests performed in a certain
// manner. Procecdures can be seen as a global catalog of all valid examination
// procedures offered by laboratories.
func (r *ProcedureService) List(ctx context.Context, query ProcedureListParams, opts ...option.RequestOption) (res *[]ProcedureResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/procedures"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a procedure from the global catalog.
func (r *ProcedureService) Delete(ctx context.Context, procedureID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if procedureID == "" {
		err = errors.New("missing required procedureId parameter")
		return
	}
	path := fmt.Sprintf("v4/procedures/%s", procedureID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// The coding system. Currently supported are
//
//   - LOINC see https://loing.org
//   - SNOMED-CT see https://snomed.org
//   - ABO Phenotype Standard (AB0,Rh,Kell,...)
//   - AB0 DRK with the DRK additons (D-1/D-2)
//     https://www.drk-haemotherapie.de/beitraege/molekulare-diagnostik-in-der-immunhamatologie/herunterladen
//   - p_group P-Level-Proteine Group https://www.ebi.ac.uk/ipd
//   - g_group G-Level-Proteine Group https://www.ebi.ac.uk/ipd Leave empty for
//     internal code
type CodingSystem string

const (
	CodingSystemLoinc        CodingSystem = "loinc"
	CodingSystemSnomed       CodingSystem = "snomed"
	CodingSystemAb0Phenotype CodingSystem = "ab0_phenotype"
	CodingSystemAb0Drk       CodingSystem = "ab0_drk"
	CodingSystemPGroup       CodingSystem = "p_group"
	CodingSystemGGroup       CodingSystem = "g_group"
)

type ProcedureResponse struct {
	// Unique identifier for the procedure
	ID string `json:"id" format:"uuid"`
	// A default code suggested for simple reference. This is not a Primary key and may
	// not be unique
	Code string `json:"code"`
	// The coding system. Currently supported are loinc and snomed, abo_phenotype
	// Standards, p_group, g_group
	CodeSystems []ProcedureResponseCodeSystem `json:"codeSystems"`
	// The description of the procedure
	Description string `json:"description"`
	// A distinct and name of the offered procedure
	Name string `json:"name"`
	// The unit
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Code        respjson.Field
		CodeSystems respjson.Field
		Description respjson.Field
		Name        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProcedureResponse) RawJSON() string { return r.JSON.raw }
func (r *ProcedureResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProcedureResponseCodeSystem struct {
	Code string `json:"code"`
	// The coding system. Currently supported are
	//
	//   - LOINC see https://loing.org
	//   - SNOMED-CT see https://snomed.org
	//   - ABO Phenotype Standard (AB0,Rh,Kell,...)
	//   - AB0 DRK with the DRK additons (D-1/D-2)
	//     https://www.drk-haemotherapie.de/beitraege/molekulare-diagnostik-in-der-immunhamatologie/herunterladen
	//   - p_group P-Level-Proteine Group https://www.ebi.ac.uk/ipd
	//   - g_group G-Level-Proteine Group https://www.ebi.ac.uk/ipd Leave empty for
	//     internal code
	//
	// Any of "loinc", "snomed", "ab0_phenotype", "ab0_drk", "p_group", "g_group".
	CodingSystem CodingSystem `json:"codingSystem"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code         respjson.Field
		CodingSystem respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProcedureResponseCodeSystem) RawJSON() string { return r.JSON.raw }
func (r *ProcedureResponseCodeSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProcedureUpdateResponse struct {
	// Unique identifier for the procedure
	ID string `json:"id" format:"uuid"`
	// A default code suggested for simple reference. This is not a Primary key and may
	// not be unique
	Code string `json:"code"`
	// The coding system. Currently supported are loinc and snomed, abo_phenotype
	// Standards, p_group, g_group
	CodeSystems []ProcedureUpdateResponseCodeSystem `json:"codeSystems"`
	// The description of the procedure
	Description string `json:"description"`
	// A distinct and name of the offered procedure
	Name string `json:"name"`
	// The unit
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Code        respjson.Field
		CodeSystems respjson.Field
		Description respjson.Field
		Name        respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProcedureUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ProcedureUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProcedureUpdateResponseCodeSystem struct {
	Code string `json:"code"`
	// Any of "loinc", "snomed", "abo_phenotype", "p_group", "g_group".
	CodingSystem string `json:"codingSystem"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code         respjson.Field
		CodingSystem respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProcedureUpdateResponseCodeSystem) RawJSON() string { return r.JSON.raw }
func (r *ProcedureUpdateResponseCodeSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProcedureNewParams struct {
	Body []ProcedureNewParamsBody
	paramObj
}

func (r ProcedureNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ProcedureNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The property Name is required.
type ProcedureNewParamsBody struct {
	// A distinct and name of the offered procedure
	Name string `json:"name,required"`
	// A default code suggested for simple reference. This is not a Primary key and may
	// not be unique
	Code param.Opt[string] `json:"code,omitzero"`
	// The description of the procedure
	Description param.Opt[string] `json:"description,omitzero"`
	// The unit
	Unit param.Opt[string] `json:"unit,omitzero"`
	// The coding systems used for the procedure
	CodeSystems []ProcedureNewParamsBodyCodeSystem `json:"codeSystems,omitzero"`
	paramObj
}

func (r ProcedureNewParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ProcedureNewParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProcedureNewParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProcedureNewParamsBodyCodeSystem struct {
	// The code within the system. Refer to the respective coding szstem for details
	Code param.Opt[string] `json:"code,omitzero"`
	// The coding system. Currently supported are
	//
	//   - LOINC see https://loing.org
	//   - SNOMED-CT see https://snomed.org
	//   - ABO Phenotype Standard (AB0,Rh,Kell,...)
	//   - AB0 DRK with the DRK additons (D-1/D-2)
	//     https://www.drk-haemotherapie.de/beitraege/molekulare-diagnostik-in-der-immunhamatologie/herunterladen
	//   - p_group P-Level-Proteine Group https://www.ebi.ac.uk/ipd
	//   - g_group G-Level-Proteine Group https://www.ebi.ac.uk/ipd Leave empty for
	//     internal code
	//
	// Any of "loinc", "snomed", "ab0_phenotype", "ab0_drk", "p_group", "g_group".
	CodingSystem CodingSystem `json:"codingSystem,omitzero"`
	paramObj
}

func (r ProcedureNewParamsBodyCodeSystem) MarshalJSON() (data []byte, err error) {
	type shadow ProcedureNewParamsBodyCodeSystem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProcedureNewParamsBodyCodeSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProcedureUpdateParams struct {
	Code        param.Opt[string]                 `json:"code,omitzero"`
	Description param.Opt[string]                 `json:"description,omitzero"`
	Name        param.Opt[string]                 `json:"name,omitzero"`
	Unit        param.Opt[string]                 `json:"unit,omitzero"`
	CodeSystems []ProcedureUpdateParamsCodeSystem `json:"codeSystems,omitzero"`
	paramObj
}

func (r ProcedureUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProcedureUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProcedureUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProcedureUpdateParamsCodeSystem struct {
	Code param.Opt[string] `json:"code,omitzero"`
	// The coding system. Currently supported are
	//
	//   - LOINC see https://loing.org
	//   - SNOMED-CT see https://snomed.org
	//   - ABO Phenotype Standard (AB0,Rh,Kell,...)
	//   - AB0 DRK with the DRK additons (D-1/D-2)
	//     https://www.drk-haemotherapie.de/beitraege/molekulare-diagnostik-in-der-immunhamatologie/herunterladen
	//   - p_group P-Level-Proteine Group https://www.ebi.ac.uk/ipd
	//   - g_group G-Level-Proteine Group https://www.ebi.ac.uk/ipd Leave empty for
	//     internal code
	//
	// Any of "loinc", "snomed", "ab0_phenotype", "ab0_drk", "p_group", "g_group".
	CodingSystem CodingSystem `json:"codingSystem,omitzero"`
	paramObj
}

func (r ProcedureUpdateParamsCodeSystem) MarshalJSON() (data []byte, err error) {
	type shadow ProcedureUpdateParamsCodeSystem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProcedureUpdateParamsCodeSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProcedureListParams struct {
	// Filter by name. This search is executed as a case-insensitive partial match.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Page number
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Filter for unit. This search is executed as a case-insensitive partial match.
	Unit param.Opt[string] `query:"unit,omitzero" json:"-"`
	// The code filter. Multiple codes can be provided. For querying a particular
	// codesystem use codeSystem parameter.
	Codes []string `query:"codes,omitzero" json:"-"`
	// The code system filter. Use one of loinc, snomed, abo_phenotype, p_group,
	// g_group
	//
	// Any of "loinc", "snomed", "ab0_phenotype", "ab0_drk", "p_group", "g_group".
	CodeSystem CodingSystem `query:"codeSystem,omitzero" json:"-"`
	// The sorting parameters in the format of "fieldName,asc/desc". E.g. type,desc
	//
	// Any of "code,asc", "code,desc", "name,asc", "name,desc", "unit,asc",
	// "unit,desc".
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ProcedureListParams]'s query parameters as `url.Values`.
func (r ProcedureListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
