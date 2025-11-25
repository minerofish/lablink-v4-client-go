// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"
	"time"

	"github.com/minerofish/lablink-v4-client-go/internal/apijson"
	"github.com/minerofish/lablink-v4-client-go/packages/param"
	"github.com/minerofish/lablink-v4-client-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Result struct {
	// The order-examination ID.
	ExaminationID string `json:"examinationId,required" format:"uuid"`
	// The order-examination item ID
	ItemID string `json:"itemId,required" format:"uuid"`
	// The order ID
	OrderID string `json:"orderId,required" format:"uuid"`
	Result  string `json:"result,required"`
	// Any of "FINAL", "PRELIMINARY", "CORRECTED".
	Status ResultStatus `json:"status,required"`
	// Indicates if the result requires confirmation.
	ConfirmationPending bool `json:"confirmationPending"`
	// The data type of the result
	//
	// Any of "int", "decimal", "string", "pein", "react", "invalid", "enum".
	DataType    ResultDataType `json:"dataType"`
	InfoText    string         `json:"infoText"`
	PerformedAt time.Time      `json:"performedAt" format:"date-time"`
	Unit        string         `json:"unit"`
	ValidatedAt time.Time      `json:"validatedAt" format:"date-time"`
	ValidatedBy string         `json:"validatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExaminationID       respjson.Field
		ItemID              respjson.Field
		OrderID             respjson.Field
		Result              respjson.Field
		Status              respjson.Field
		ConfirmationPending respjson.Field
		DataType            respjson.Field
		InfoText            respjson.Field
		PerformedAt         respjson.Field
		Unit                respjson.Field
		ValidatedAt         respjson.Field
		ValidatedBy         respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Result) RawJSON() string { return r.JSON.raw }
func (r *Result) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Result to a ResultParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ResultParam.Overrides()
func (r Result) ToParam() ResultParam {
	return param.Override[ResultParam](json.RawMessage(r.RawJSON()))
}

type ResultStatus string

const (
	ResultStatusFinal       ResultStatus = "FINAL"
	ResultStatusPreliminary ResultStatus = "PRELIMINARY"
	ResultStatusCorrected   ResultStatus = "CORRECTED"
)

// The data type of the result
type ResultDataType string

const (
	ResultDataTypeInt     ResultDataType = "int"
	ResultDataTypeDecimal ResultDataType = "decimal"
	ResultDataTypeString  ResultDataType = "string"
	ResultDataTypePein    ResultDataType = "pein"
	ResultDataTypeReact   ResultDataType = "react"
	ResultDataTypeInvalid ResultDataType = "invalid"
	ResultDataTypeEnum    ResultDataType = "enum"
)

// The properties ExaminationID, ItemID, OrderID, Result, Status are required.
type ResultParam struct {
	// The order-examination ID.
	ExaminationID string `json:"examinationId,required" format:"uuid"`
	// The order-examination item ID
	ItemID string `json:"itemId,required" format:"uuid"`
	// The order ID
	OrderID string `json:"orderId,required" format:"uuid"`
	Result  string `json:"result,required"`
	// Any of "FINAL", "PRELIMINARY", "CORRECTED".
	Status ResultStatus `json:"status,omitzero,required"`
	// Indicates if the result requires confirmation.
	ConfirmationPending param.Opt[bool]      `json:"confirmationPending,omitzero"`
	InfoText            param.Opt[string]    `json:"infoText,omitzero"`
	PerformedAt         param.Opt[time.Time] `json:"performedAt,omitzero" format:"date-time"`
	Unit                param.Opt[string]    `json:"unit,omitzero"`
	ValidatedAt         param.Opt[time.Time] `json:"validatedAt,omitzero" format:"date-time"`
	ValidatedBy         param.Opt[string]    `json:"validatedBy,omitzero"`
	// The data type of the result
	//
	// Any of "int", "decimal", "string", "pein", "react", "invalid", "enum".
	DataType ResultDataType `json:"dataType,omitzero"`
	paramObj
}

func (r ResultParam) MarshalJSON() (data []byte, err error) {
	type shadow ResultParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResultParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
