// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/minerofish/lablink-v4-client-go/internal/apijson"
	shimjson "github.com/minerofish/lablink-v4-client-go/internal/encoding/json"
	"github.com/minerofish/lablink-v4-client-go/internal/requestconfig"
	"github.com/minerofish/lablink-v4-client-go/option"
	"github.com/minerofish/lablink-v4-client-go/packages/param"
)

// LaboratoryContractExaminationService contains methods and other services that
// help with interacting with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaboratoryContractExaminationService] method instead.
type LaboratoryContractExaminationService struct {
	Options []option.RequestOption
}

// NewLaboratoryContractExaminationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLaboratoryContractExaminationService(opts ...option.RequestOption) (r LaboratoryContractExaminationService) {
	r = LaboratoryContractExaminationService{}
	r.Options = opts
	return
}

// Updates one or more existing contracts for the specified laboratory.
func (r *LaboratoryContractExaminationService) Update(ctx context.Context, contractID string, params LaboratoryContractExaminationUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.LaboratoryID == "" {
		err = errors.New("missing required laboratoryId parameter")
		return
	}
	if contractID == "" {
		err = errors.New("missing required contractId parameter")
		return
	}
	path := fmt.Sprintf("v4/laboratories/%s/contracts/%s/examinations", params.LaboratoryID, contractID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, nil, opts...)
	return
}

// Adds a new examination to an existing contract between a laboratory and an
// organization.
func (r *LaboratoryContractExaminationService) Add(ctx context.Context, contractID string, params LaboratoryContractExaminationAddParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.LaboratoryID == "" {
		err = errors.New("missing required laboratoryId parameter")
		return
	}
	if contractID == "" {
		err = errors.New("missing required contractId parameter")
		return
	}
	path := fmt.Sprintf("v4/laboratories/%s/contracts/%s/examinations", params.LaboratoryID, contractID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Removes one or more examinations from the specified contract.
func (r *LaboratoryContractExaminationService) Remove(ctx context.Context, contractID string, params LaboratoryContractExaminationRemoveParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.LaboratoryID == "" {
		err = errors.New("missing required laboratoryId parameter")
		return
	}
	if contractID == "" {
		err = errors.New("missing required contractId parameter")
		return
	}
	path := fmt.Sprintf("v4/laboratories/%s/contracts/%s/examinations", params.LaboratoryID, contractID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

type LaboratoryContractExaminationUpdateParams struct {
	LaboratoryID string `path:"laboratoryId,required" format:"uuid" json:"-"`
	Body         []LaboratoryContractExaminationUpdateParamsBody
	paramObj
}

func (r LaboratoryContractExaminationUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *LaboratoryContractExaminationUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The property ExaminationID is required.
type LaboratoryContractExaminationUpdateParamsBody struct {
	ExaminationID   string            `json:"examinationId,required" format:"uuid"`
	ExaminationCode param.Opt[string] `json:"examinationCode,omitzero"`
	paramObj
}

func (r LaboratoryContractExaminationUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LaboratoryContractExaminationUpdateParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaboratoryContractExaminationUpdateParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryContractExaminationAddParams struct {
	LaboratoryID string `path:"laboratoryId,required" format:"uuid" json:"-"`
	// Define Procedures to add to the contract
	Body []LaboratoryContractExaminationAddParamsBody
	paramObj
}

func (r LaboratoryContractExaminationAddParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *LaboratoryContractExaminationAddParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The property ProcedureID is required.
type LaboratoryContractExaminationAddParamsBody struct {
	ProcedureID string `json:"procedureId,required" format:"uuid"`
	// Customer specific code of the examination. If omitted, the suggestion from the
	// procedure is used
	ExaminationCode param.Opt[string] `json:"examinationCode,omitzero"`
	paramObj
}

func (r LaboratoryContractExaminationAddParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LaboratoryContractExaminationAddParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaboratoryContractExaminationAddParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LaboratoryContractExaminationRemoveParams struct {
	LaboratoryID string `path:"laboratoryId,required" format:"uuid" json:"-"`
	Body         []LaboratoryContractExaminationRemoveParamsBody
	paramObj
}

func (r LaboratoryContractExaminationRemoveParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *LaboratoryContractExaminationRemoveParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The property ExaminationID is required.
type LaboratoryContractExaminationRemoveParamsBody struct {
	// Identifiers of the examinations to delete
	ExaminationID string `json:"examinationId,required" format:"uuid"`
	paramObj
}

func (r LaboratoryContractExaminationRemoveParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow LaboratoryContractExaminationRemoveParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LaboratoryContractExaminationRemoveParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
