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

	"github.com/stainless-sdks/lablink-v4-client-go/internal/apijson"
	"github.com/stainless-sdks/lablink-v4-client-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/lablink-v4-client-go/internal/encoding/json"
	"github.com/stainless-sdks/lablink-v4-client-go/internal/requestconfig"
	"github.com/stainless-sdks/lablink-v4-client-go/option"
	"github.com/stainless-sdks/lablink-v4-client-go/packages/param"
	"github.com/stainless-sdks/lablink-v4-client-go/packages/respjson"
)

// LaboratoryContractService contains methods and other services that help with
// interacting with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaboratoryContractService] method instead.
type LaboratoryContractService struct {
	Options []option.RequestOption
}

// NewLaboratoryContractService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLaboratoryContractService(opts ...option.RequestOption) (r LaboratoryContractService) {
	r = LaboratoryContractService{}
	r.Options = opts
	return
}

// Use this to
//
//   - initiate the dataexchange. The contract needs to define which examinations are
//     exchanged.
func (r *LaboratoryContractService) New(ctx context.Context, laboratoryID string, body LaboratoryContractNewParams, opts ...option.RequestOption) (res *Contract, err error) {
	opts = slices.Concat(r.Options, opts)
	if laboratoryID == "" {
		err = errors.New("missing required laboratoryId parameter")
		return
	}
	path := fmt.Sprintf("v4/laboratories/%s/contracts", laboratoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retreieves a contract by laboratory and or organization
func (r *LaboratoryContractService) Get(ctx context.Context, laboratoryID string, query LaboratoryContractGetParams, opts ...option.RequestOption) (res *LaboratoryContractGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if laboratoryID == "" {
		err = errors.New("missing required laboratoryId parameter")
		return
	}
	path := fmt.Sprintf("v4/laboratories/%s/contracts", laboratoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates one or more existing contracts for the specified laboratory.
func (r *LaboratoryContractService) Update(ctx context.Context, laboratoryID string, body LaboratoryContractUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if laboratoryID == "" {
		err = errors.New("missing required laboratoryId parameter")
		return
	}
	path := fmt.Sprintf("v4/laboratories/%s/contracts", laboratoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, nil, opts...)
	return
}

// Removes one or more contracts for the specified laboratory.
func (r *LaboratoryContractService) Delete(ctx context.Context, laboratoryID string, body LaboratoryContractDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if laboratoryID == "" {
		err = errors.New("missing required laboratoryId parameter")
		return
	}
	path := fmt.Sprintf("v4/laboratories/%s/contracts", laboratoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Retrieves a list of contracts along with organizations and laboratories that
// have contracts.
func (r *LaboratoryContractService) Query(ctx context.Context, laboratoryID string, params LaboratoryContractQueryParams, opts ...option.RequestOption) (res *LaboratoryContractQueryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if laboratoryID == "" {
		err = errors.New("missing required laboratoryId parameter")
		return
	}
	path := fmt.Sprintf("v4/laboratories/%s/contracts/query", laboratoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type Contract struct {
	// The contract ID
	ID string `json:"id,required" format:"uuid"`
	// The examinations included in the contract
	Examinations []Examination        `json:"examinations,required"`
	Laboratory   ContractLaboratory   `json:"laboratory,required"`
	Organization ContractOrganization `json:"organization,required"`
	// The contract creation date-time
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The contract last modification date-time
	ModifiedAt time.Time `json:"modifiedAt" format:"date-time"`
	// First day of validity
	ValidFrom time.Time `json:"validFrom" format:"date-time"`
	// The contract valid until date
	ValidUntil time.Time `json:"validUntil" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Examinations respjson.Field
		Laboratory   respjson.Field
		Organization respjson.Field
		CreatedAt    respjson.Field
		ModifiedAt   respjson.Field
		ValidFrom    respjson.Field
		ValidUntil   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Contract) RawJSON() string { return r.JSON.raw }
func (r *Contract) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractLaboratory struct {
	// The laboratory ID
	ID   string `json:"id,required" format:"uuid"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractLaboratory) RawJSON() string { return r.JSON.raw }
func (r *ContractLaboratory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContractOrganization struct {
	// The organization ID
	ID string `json:"id,required" format:"uuid"`
	// The organization name
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractOrganization) RawJSON() string { return r.JSON.raw }
func (r *ContractOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryContractGetResponse struct {
	Items []Contract `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Page
}

// Returns the unmodified JSON received from the API
func (r LaboratoryContractGetResponse) RawJSON() string { return r.JSON.raw }
func (r *LaboratoryContractGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryContractQueryResponse struct {
	Items []Contract `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Page
}

// Returns the unmodified JSON received from the API
func (r LaboratoryContractQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *LaboratoryContractQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryContractNewParams struct {
	// application/glims+json upload glims structured results in json format
	Body []LaboratoryContractNewParamsBody
	paramObj
}

func (r LaboratoryContractNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *LaboratoryContractNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Request body for creating a new contract between a laboratory and an
// organization
//
// The properties OrganizationID, SampleCodeRanges are required.
type LaboratoryContractNewParamsBody struct {
	// The ID of the organization involved in the contract
	OrganizationID   string                                           `json:"organizationId,required" format:"uuid"`
	SampleCodeRanges []LaboratoryContractNewParamsBodySampleCodeRange `json:"sampleCodeRanges,omitzero,required"`
	// When the contract was established
	ValidFrom param.Opt[time.Time] `json:"validFrom,omitzero" format:"date"`
	// When the contract ended
	ValidTo      param.Opt[time.Time]                         `json:"validTo,omitzero" format:"date"`
	Examinations []LaboratoryContractNewParamsBodyExamination `json:"examinations,omitzero"`
	paramObj
}

func (r LaboratoryContractNewParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LaboratoryContractNewParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaboratoryContractNewParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties End, Start are required.
type LaboratoryContractNewParamsBodySampleCodeRange struct {
	// End of the sample code range
	End int64 `json:"end,required"`
	// Start of the sample code range
	Start int64 `json:"start,required"`
	// Number of digits reserved for the counter. Shorter numbers will be filled with
	// leading zeros.
	Digits param.Opt[int64] `json:"digits,omitzero"`
	// Amount of digits ignored at the end of the code, often used to encode material
	// etc.
	PostfixDigits param.Opt[int64] `json:"postfixDigits,omitzero"`
	// Prefix of the sample code range (optional)
	Prefix param.Opt[string] `json:"prefix,omitzero"`
	paramObj
}

func (r LaboratoryContractNewParamsBodySampleCodeRange) MarshalJSON() (data []byte, err error) {
	type shadow LaboratoryContractNewParamsBodySampleCodeRange
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaboratoryContractNewParamsBodySampleCodeRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ProcedureID is required.
type LaboratoryContractNewParamsBodyExamination struct {
	ProcedureID string `json:"procedureId,required" format:"uuid"`
	// Customer specific code of the examination. If omitted, the suggestion from the
	// procedure is used
	Customercode param.Opt[string] `json:"customercode,omitzero"`
	// Customer specific name of the examination, otherwise the procedure is used
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r LaboratoryContractNewParamsBodyExamination) MarshalJSON() (data []byte, err error) {
	type shadow LaboratoryContractNewParamsBodyExamination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaboratoryContractNewParamsBodyExamination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryContractGetParams struct {
	Organizationid param.Opt[string] `query:"organizationid,omitzero" format:"uuid" json:"-"`
	Page           param.Opt[int64]  `query:"Page,omitzero" json:"-"`
	PageSize       param.Opt[int64]  `query:"PageSize,omitzero" json:"-"`
	// The sorting parameters in the format of "fieldName,asc/desc". E.g. type,desc
	//
	// Any of "code,asc", "code,desc", "organzationName,asc", "organzationName,desc",
	// "procedureName,asc", "procedureName,desc", "organizationCreated,asc",
	// "organizationCreated,desc".
	Sort []string `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LaboratoryContractGetParams]'s query parameters as
// `url.Values`.
func (r LaboratoryContractGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaboratoryContractUpdateParams struct {
	Body []LaboratoryContractUpdateParamsBody
	paramObj
}

func (r LaboratoryContractUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *LaboratoryContractUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The property ContractID is required.
type LaboratoryContractUpdateParamsBody struct {
	// Identifier of the contract to update
	ContractID string `json:"contractId,required" format:"uuid"`
	// Updated organization identifier. If omitted, the existing value is kept.
	OrganizationID param.Opt[string] `json:"organizationId,omitzero" format:"uuid"`
	// Updated contract valid-from date
	ValidFrom param.Opt[time.Time] `json:"validFrom,omitzero" format:"date"`
	// Updated contract valid-to date
	ValidTo param.Opt[time.Time] `json:"validTo,omitzero" format:"date"`
	// Updated examinations assigned to the contract
	Examinations []LaboratoryContractUpdateParamsBodyExamination `json:"examinations,omitzero"`
	paramObj
}

func (r LaboratoryContractUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LaboratoryContractUpdateParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaboratoryContractUpdateParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ProcedureID is required.
type LaboratoryContractUpdateParamsBodyExamination struct {
	ProcedureID  string            `json:"procedureId,required" format:"uuid"`
	Customercode param.Opt[string] `json:"customercode,omitzero"`
	Name         param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r LaboratoryContractUpdateParamsBodyExamination) MarshalJSON() (data []byte, err error) {
	type shadow LaboratoryContractUpdateParamsBodyExamination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaboratoryContractUpdateParamsBodyExamination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryContractDeleteParams struct {
	// Identifiers of the contracts to delete
	ContractIDs []string `json:"contractIds,omitzero,required" format:"uuid"`
	paramObj
}

func (r LaboratoryContractDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow LaboratoryContractDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaboratoryContractDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryContractQueryParams struct {
	Page     param.Opt[int64] `query:"Page,omitzero" json:"-"`
	PageSize param.Opt[int64] `query:"PageSize,omitzero" json:"-"`
	// The sorting parameters in the format of "fieldName,asc/desc". E.g. type,desc
	//
	// Any of "validFrom,asc", "validFrom,desc".
	Sort []string `query:"sort,omitzero" json:"-"`
	// Filter by Organization IDs
	OrganizationIDs []string `json:"organizationIds,omitzero" format:"uuid"`
	paramObj
}

func (r LaboratoryContractQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow LaboratoryContractQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaboratoryContractQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [LaboratoryContractQueryParams]'s query parameters as
// `url.Values`.
func (r LaboratoryContractQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
