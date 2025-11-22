// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/minerofish/lablink-v4-client-go/internal/apiform"
	"github.com/minerofish/lablink-v4-client-go/internal/apijson"
	"github.com/minerofish/lablink-v4-client-go/internal/apiquery"
	shimjson "github.com/minerofish/lablink-v4-client-go/internal/encoding/json"
	"github.com/minerofish/lablink-v4-client-go/internal/requestconfig"
	"github.com/minerofish/lablink-v4-client-go/option"
	"github.com/minerofish/lablink-v4-client-go/packages/param"
	"github.com/minerofish/lablink-v4-client-go/packages/respjson"
)

// LaboratoryService contains methods and other services that help with interacting
// with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaboratoryService] method instead.
type LaboratoryService struct {
	Options   []option.RequestOption
	Orders    LaboratoryOrderService
	Contracts LaboratoryContractService
}

// NewLaboratoryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLaboratoryService(opts ...option.RequestOption) (r LaboratoryService) {
	r = LaboratoryService{}
	r.Options = opts
	r.Orders = NewLaboratoryOrderService(opts...)
	r.Contracts = NewLaboratoryContractService(opts...)
	return
}

// Creates new laboratories for the authenticated user.
func (r *LaboratoryService) New(ctx context.Context, body LaboratoryNewParams, opts ...option.RequestOption) (res *[]LaboratoryNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/laboratories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets a laboratory by its id.
func (r *LaboratoryService) List(ctx context.Context, query LaboratoryListParams, opts ...option.RequestOption) (res *[]Laboratory, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/laboratories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a laboratory for the authenticated user. User must have the
// laboratory_admin role to delete. This operation is audited.
func (r *LaboratoryService) Delete(ctx context.Context, body LaboratoryDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v4/laboratories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Queries laboratories. The caller must be user for the respective laboratories.
func (r *LaboratoryService) Query(ctx context.Context, params LaboratoryQueryParams, opts ...option.RequestOption) (res *LaboratoryQueryResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/laboratories/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Upload confirmatory laboratory results from Glims system
func (r *LaboratoryService) UploadConfirmatoryResults(ctx context.Context, laboratoryID string, body LaboratoryUploadConfirmatoryResultsParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	if laboratoryID == "" {
		err = errors.New("missing required laboratoryId parameter")
		return
	}
	path := fmt.Sprintf("v4/laboratories/%s/confirmatory", laboratoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Uploads a document, including a compressed one (zip).
func (r *LaboratoryService) UploadDocument(ctx context.Context, laboratoryID string, body LaboratoryUploadDocumentParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if laboratoryID == "" {
		err = errors.New("missing required laboratoryId parameter")
		return
	}
	path := fmt.Sprintf("v4/laboratories/%s/document", laboratoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Upload laboratory results from various cusotm systems
func (r *LaboratoryService) UploadResults(ctx context.Context, laboratoryID string, body LaboratoryUploadResultsParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	if laboratoryID == "" {
		err = errors.New("missing required laboratoryId parameter")
		return
	}
	path := fmt.Sprintf("v4/laboratories/%s/results", laboratoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Laboratory struct {
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
func (r Laboratory) RawJSON() string { return r.JSON.raw }
func (r *Laboratory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryUserRelations struct {
	Email string `json:"email,required"`
	// Any of "ORGANIZATION_ADMIN", "USER", "API_USER", "LABORATORY".
	Role LaboratoryUserRelationsRole `json:"role,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaboratoryUserRelations) RawJSON() string { return r.JSON.raw }
func (r *LaboratoryUserRelations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryUserRelationsRole string

const (
	LaboratoryUserRelationsRoleOrganizationAdmin LaboratoryUserRelationsRole = "ORGANIZATION_ADMIN"
	LaboratoryUserRelationsRoleUser              LaboratoryUserRelationsRole = "USER"
	LaboratoryUserRelationsRoleAPIUser           LaboratoryUserRelationsRole = "API_USER"
	LaboratoryUserRelationsRoleLaboratory        LaboratoryUserRelationsRole = "LABORATORY"
)

type LaboratoryNewResponse struct {
	ID string `json:"id" format:"uuid"`
	// API access token for the laboratory
	APIAccessToken string `json:"api_access_token"`
	InstanceName   string `json:"instanceName"`
	Name           string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		APIAccessToken respjson.Field
		InstanceName   respjson.Field
		Name           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LaboratoryNewResponse) RawJSON() string { return r.JSON.raw }
func (r *LaboratoryNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryQueryResponse struct {
	Items []Laboratory `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	Page
}

// Returns the unmodified JSON received from the API
func (r LaboratoryQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *LaboratoryQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryNewParams struct {
	Body []LaboratoryNewParamsBody
	paramObj
}

func (r LaboratoryNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *LaboratoryNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// Request body for creating a new laboratory
//
// The properties InstanceName, Name are required.
type LaboratoryNewParamsBody struct {
	// Bloodlab's laboratory instance name
	InstanceName string `json:"instance_name,required"`
	// Name of the laboratory
	Name     string            `json:"name,required"`
	Address  param.Opt[string] `json:"address,omitzero"`
	City     param.Opt[string] `json:"city,omitzero"`
	Country  param.Opt[string] `json:"country,omitzero"`
	Email    param.Opt[string] `json:"email,omitzero"`
	Phone    param.Opt[string] `json:"phone,omitzero"`
	Postcode param.Opt[string] `json:"postcode,omitzero"`
	paramObj
}

func (r LaboratoryNewParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LaboratoryNewParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaboratoryNewParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryListParams struct {
	LaboratoryID param.Opt[string] `query:"laboratoryId,omitzero" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [LaboratoryListParams]'s query parameters as `url.Values`.
func (r LaboratoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaboratoryDeleteParams struct {
	// The laboratory ID
	LaboratoryID string `query:"laboratoryID,required" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [LaboratoryDeleteParams]'s query parameters as `url.Values`.
func (r LaboratoryDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaboratoryQueryParams struct {
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

func (r LaboratoryQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow LaboratoryQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaboratoryQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [LaboratoryQueryParams]'s query parameters as `url.Values`.
func (r LaboratoryQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LaboratoryUploadConfirmatoryResultsParams struct {
	Body []LaboratoryUploadConfirmatoryResultsParamsBody
	paramObj
}

func (r LaboratoryUploadConfirmatoryResultsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *LaboratoryUploadConfirmatoryResultsParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The properties Examination, SampleCode, YieldDate are required.
type LaboratoryUploadConfirmatoryResultsParamsBody struct {
	Examination         string            `json:"examination,required" format:"int32"`
	SampleCode          string            `json:"sampleCode,required" format:"int32"`
	YieldDate           string            `json:"yieldDate,required" format:"int32"`
	ConfirmatoryGroupID param.Opt[int64]  `json:"confirmatoryGroupId,omitzero"`
	Customer            param.Opt[string] `json:"customer,omitzero" format:"int32"`
	QualifiedResult     param.Opt[string] `json:"qualifiedResult,omitzero" format:"int32"`
	QuantitativeResult  param.Opt[string] `json:"quantitativeResult,omitzero" format:"int32"`
	paramObj
}

func (r LaboratoryUploadConfirmatoryResultsParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LaboratoryUploadConfirmatoryResultsParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaboratoryUploadConfirmatoryResultsParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryUploadDocumentParams struct {
	Examination string `json:"examination,required"`
	// The file
	File       io.Reader `json:"file,omitzero,required" format:"binary"`
	SampleCode string    `json:"sampleCode,required"`
	paramObj
}

func (r LaboratoryUploadDocumentParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type LaboratoryUploadResultsParams struct {
	// application/glims+json upload glims structured results in json format
	Body []LaboratoryUploadResultsParamsBody
	paramObj
}

func (r LaboratoryUploadResultsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *LaboratoryUploadResultsParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The properties Examination, SampleCode, YieldDate are required.
type LaboratoryUploadResultsParamsBody struct {
	Examination        string            `json:"examination,required"`
	SampleCode         string            `json:"sampleCode,required"`
	YieldDate          string            `json:"yieldDate,required"`
	Customer           param.Opt[string] `json:"customer,omitzero"`
	QualifiedResult    param.Opt[string] `json:"qualifiedResult,omitzero"`
	QuantitativeResult param.Opt[string] `json:"quantitativeResult,omitzero"`
	paramObj
}

func (r LaboratoryUploadResultsParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LaboratoryUploadResultsParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaboratoryUploadResultsParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
