// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"github.com/minerofish/lablink-v4-client-go/internal/apierror"
	"github.com/minerofish/lablink-v4-client-go/packages/param"
	"github.com/minerofish/lablink-v4-client-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type Result = shared.Result

// This is an alias to an internal type.
type ResultStatus = shared.ResultStatus

// Equals "FINAL"
const ResultStatusFinal = shared.ResultStatusFinal

// Equals "PRELIMINARY"
const ResultStatusPreliminary = shared.ResultStatusPreliminary

// Equals "CORRECTED"
const ResultStatusCorrected = shared.ResultStatusCorrected

// The data type of the result
//
// This is an alias to an internal type.
type ResultDataType = shared.ResultDataType

// Equals "int"
const ResultDataTypeInt = shared.ResultDataTypeInt

// Equals "decimal"
const ResultDataTypeDecimal = shared.ResultDataTypeDecimal

// Equals "string"
const ResultDataTypeString = shared.ResultDataTypeString

// Equals "pein"
const ResultDataTypePein = shared.ResultDataTypePein

// Equals "react"
const ResultDataTypeReact = shared.ResultDataTypeReact

// Equals "invalid"
const ResultDataTypeInvalid = shared.ResultDataTypeInvalid

// Equals "enum"
const ResultDataTypeEnum = shared.ResultDataTypeEnum

// This is an alias to an internal type.
type ResultParam = shared.ResultParam
