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

	"github.com/minerofish/lablink-v4-client-go/internal/apijson"
	"github.com/minerofish/lablink-v4-client-go/internal/apiquery"
	shimjson "github.com/minerofish/lablink-v4-client-go/internal/encoding/json"
	"github.com/minerofish/lablink-v4-client-go/internal/requestconfig"
	"github.com/minerofish/lablink-v4-client-go/option"
	"github.com/minerofish/lablink-v4-client-go/packages/param"
	"github.com/minerofish/lablink-v4-client-go/packages/respjson"
)

// OrderService contains methods and other services that help with interacting with
// the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrderService] method instead.
type OrderService struct {
	Options      []option.RequestOption
	Tags         OrderTagService
	Examinations OrderExaminationService
}

// NewOrderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrderService(opts ...option.RequestOption) (r OrderService) {
	r = OrderService{}
	r.Options = opts
	r.Tags = NewOrderTagService(opts...)
	r.Examinations = NewOrderExaminationService(opts...)
	return
}

// Creates orders in Lablink. Orders can be created by providing JSON or text/exlab
// payloads directly. A list of the created orderIDs is returned.
func (r *OrderService) New(ctx context.Context, params OrderNewParams, opts ...option.RequestOption) (res *[]OrderMetadata, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Gets an order.
func (r *OrderService) Get(ctx context.Context, orderID string, opts ...option.RequestOption) (res *OrderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return
	}
	path := fmt.Sprintf("v4/orders/%s", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes all order with the provided ids.
func (r *OrderService) Delete(ctx context.Context, body OrderDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v4/orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Deletes an order.
func (r *OrderService) DeleteOrder(ctx context.Context, orderID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if orderID == "" {
		err = errors.New("missing required orderId parameter")
		return
	}
	path := fmt.Sprintf("v4/orders/%s", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type BloodDonor struct {
	// The date of birth (yyyy-MM-dd)
	DateOfBirth time.Time `json:"dateOfBirth,required" format:"date"`
	// The donation code
	DonationCode string `json:"donationCode,required"`
	// The donor code
	DonorCode string `json:"donorCode,required"`
	// The gender (Female, Male, Unknown)
	//
	// Any of "F", "M", "U".
	Gender Gender `json:"gender,required"`
	// The blood type
	//
	// Any of "A", "B", "0", "0h", "AB".
	BloodType BloodType `json:"bloodType"`
	// The CDE
	Cde string `json:"cde"`
	// The donor type (E: first-time donor, 2: second-time donor, M: multi-time donor)
	DonorType string `json:"donorType"`
	// The Kell factor
	//
	// Any of "pos", "neg".
	Kell Kell `json:"kell"`
	// The rhesus factor
	//
	// Any of "pos", "neg", "1", "2".
	Rhesus Rhesus `json:"rhesus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateOfBirth  respjson.Field
		DonationCode respjson.Field
		DonorCode    respjson.Field
		Gender       respjson.Field
		BloodType    respjson.Field
		Cde          respjson.Field
		DonorType    respjson.Field
		Kell         respjson.Field
		Rhesus       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BloodDonor) RawJSON() string { return r.JSON.raw }
func (r *BloodDonor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BloodDonor to a BloodDonorParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BloodDonorParam.Overrides()
func (r BloodDonor) ToParam() BloodDonorParam {
	return param.Override[BloodDonorParam](json.RawMessage(r.RawJSON()))
}

// The properties DateOfBirth, DonationCode, DonorCode, Gender are required.
type BloodDonorParam struct {
	// The date of birth (yyyy-MM-dd)
	DateOfBirth time.Time `json:"dateOfBirth,required" format:"date"`
	// The donation code
	DonationCode string `json:"donationCode,required"`
	// The donor code
	DonorCode string `json:"donorCode,required"`
	// The gender (Female, Male, Unknown)
	//
	// Any of "F", "M", "U".
	Gender Gender `json:"gender,omitzero,required"`
	// The CDE
	Cde param.Opt[string] `json:"cde,omitzero"`
	// The donor type (E: first-time donor, 2: second-time donor, M: multi-time donor)
	DonorType param.Opt[string] `json:"donorType,omitzero"`
	// The blood type
	//
	// Any of "A", "B", "0", "0h", "AB".
	BloodType BloodType `json:"bloodType,omitzero"`
	// The Kell factor
	//
	// Any of "pos", "neg".
	Kell Kell `json:"kell,omitzero"`
	// The rhesus factor
	//
	// Any of "pos", "neg", "1", "2".
	Rhesus Rhesus `json:"rhesus,omitzero"`
	paramObj
}

func (r BloodDonorParam) MarshalJSON() (data []byte, err error) {
	type shadow BloodDonorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BloodDonorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BloodType string

const (
	BloodTypeA  BloodType = "A"
	BloodTypeB  BloodType = "B"
	BloodType0  BloodType = "0"
	BloodType0h BloodType = "0h"
	BloodTypeAb BloodType = "AB"
)

type BoneMarrowDonor struct {
	// The donor code
	DonorCode string `json:"donorCode,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DonorCode   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BoneMarrowDonor) RawJSON() string { return r.JSON.raw }
func (r *BoneMarrowDonor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BoneMarrowDonor to a BoneMarrowDonorParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BoneMarrowDonorParam.Overrides()
func (r BoneMarrowDonor) ToParam() BoneMarrowDonorParam {
	return param.Override[BoneMarrowDonorParam](json.RawMessage(r.RawJSON()))
}

// The property DonorCode is required.
type BoneMarrowDonorParam struct {
	// The donor code
	DonorCode string `json:"donorCode,required"`
	paramObj
}

func (r BoneMarrowDonorParam) MarshalJSON() (data []byte, err error) {
	type shadow BoneMarrowDonorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BoneMarrowDonorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Gender string

const (
	GenderF Gender = "F"
	GenderM Gender = "M"
	GenderU Gender = "U"
)

type Kell string

const (
	KellPos Kell = "pos"
	KellNeg Kell = "neg"
)

type Link struct {
	// The related link to the order
	Related string `json:"related" format:"uri"`
	// The link to order itself
	Self string `json:"self" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Related     respjson.Field
		Self        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Link) RawJSON() string { return r.JSON.raw }
func (r *Link) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderExamination struct {
	// The sample time
	CollectionTime time.Time `json:"collectionTime,required" format:"date-time"`
	// The custom code of the examination that is used in the order item. This is
	// customer specific and can be differ from the procedure code.
	ExaminationCode string `json:"examinationCode,required"`
	// Reference to examination.
	ExaminationID string `json:"examinationId,required" format:"uuid"`
	// Reference to Order Item.
	ItemID string `json:"itemId,required" format:"uuid"`
	// The sample codes
	SampleCode string `json:"sampleCode,required"`
	// The external unique identifier of the order
	Reference string `json:"reference"`
	// The result belonging to the order
	Result OrderExaminationResult `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CollectionTime  respjson.Field
		ExaminationCode respjson.Field
		ExaminationID   respjson.Field
		ItemID          respjson.Field
		SampleCode      respjson.Field
		Reference       respjson.Field
		Result          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderExamination) RawJSON() string { return r.JSON.raw }
func (r *OrderExamination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The result belonging to the order
type OrderExaminationResult struct {
	// The order-examination ID.
	ExaminationID string `json:"examinationId,required" format:"uuid"`
	// The order-examination item ID
	ItemID string `json:"itemId,required" format:"uuid"`
	// The order ID
	OrderID string `json:"orderId,required" format:"uuid"`
	Result  string `json:"result,required"`
	// Any of "FINAL", "PRELIMINARY", "CORRECTED".
	Status string `json:"status,required"`
	// Indicates if the result requires confirmation.
	ConfirmationPending bool `json:"confirmationPending"`
	// The data type of the result
	//
	// Any of "int", "decimal", "string", "pein", "react", "invalid", "enum".
	DataType    string    `json:"dataType"`
	InfoText    string    `json:"infoText"`
	PerformedAt time.Time `json:"performedAt" format:"date-time"`
	Unit        string    `json:"unit"`
	ValidatedAt time.Time `json:"validatedAt" format:"date-time"`
	ValidatedBy string    `json:"validatedBy"`
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
func (r OrderExaminationResult) RawJSON() string { return r.JSON.raw }
func (r *OrderExaminationResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderMetadata struct {
	// The unique identifier of the order
	ID string `json:"id" format:"uuid"`
	// Error message for failed orders
	Error string `json:"error"`
	// The links to the order
	Links Link `json:"links"`
	// The reference of the order
	Reference string `json:"reference"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Error       respjson.Field
		Links       respjson.Field
		Reference   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderMetadata) RawJSON() string { return r.JSON.raw }
func (r *OrderMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The order status model
type OrderStateType string

const (
	OrderStateTypeEntered             OrderStateType = "ENTERED"
	OrderStateTypeWaitingForMaterial  OrderStateType = "WAITING_FOR_MATERIAL"
	OrderStateTypeProcessing          OrderStateType = "PROCESSING"
	OrderStateTypeConfirmationPending OrderStateType = "CONFIRMATION_PENDING"
	OrderStateTypeFinal               OrderStateType = "FINAL"
	OrderStateTypeDeleted             OrderStateType = "DELETED"
	OrderStateTypeError               OrderStateType = "ERROR"
)

type OrderType string

const (
	OrderTypeDonor           OrderType = "DONOR"
	OrderTypeBoneMarrowDonor OrderType = "BONE_MARROW_DONOR"
	OrderTypePersonal        OrderType = "PERSONAL"
	OrderTypePseudonym       OrderType = "PSEUDONYM"
)

type Patient struct {
	// The date of birth (yyyy-MM-dd)
	DateOfBirth time.Time `json:"dateOfBirth,required" format:"date"`
	// The first name
	FirstName string `json:"firstName,required"`
	// The gender (Female, Male, Unknown)
	//
	// Any of "F", "M", "U".
	Gender Gender `json:"gender,required"`
	// The last name
	LastName string `json:"lastName,required"`
	// The address
	Address string `json:"address"`
	// The city
	City string `json:"city"`
	// The country
	Country string `json:"country"`
	// The primary phone number
	PhoneNumberPrimary string `json:"phoneNumberPrimary"`
	// The secondary phone number
	PhoneNumberSecondary string `json:"phoneNumberSecondary"`
	// The post
	PostCode string `json:"postCode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateOfBirth          respjson.Field
		FirstName            respjson.Field
		Gender               respjson.Field
		LastName             respjson.Field
		Address              respjson.Field
		City                 respjson.Field
		Country              respjson.Field
		PhoneNumberPrimary   respjson.Field
		PhoneNumberSecondary respjson.Field
		PostCode             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Patient) RawJSON() string { return r.JSON.raw }
func (r *Patient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Patient to a PatientParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PatientParam.Overrides()
func (r Patient) ToParam() PatientParam {
	return param.Override[PatientParam](json.RawMessage(r.RawJSON()))
}

// The properties DateOfBirth, FirstName, Gender, LastName are required.
type PatientParam struct {
	// The date of birth (yyyy-MM-dd)
	DateOfBirth time.Time `json:"dateOfBirth,required" format:"date"`
	// The first name
	FirstName string `json:"firstName,required"`
	// The gender (Female, Male, Unknown)
	//
	// Any of "F", "M", "U".
	Gender Gender `json:"gender,omitzero,required"`
	// The last name
	LastName string `json:"lastName,required"`
	// The address
	Address param.Opt[string] `json:"address,omitzero"`
	// The city
	City param.Opt[string] `json:"city,omitzero"`
	// The country
	Country param.Opt[string] `json:"country,omitzero"`
	// The primary phone number
	PhoneNumberPrimary param.Opt[string] `json:"phoneNumberPrimary,omitzero"`
	// The secondary phone number
	PhoneNumberSecondary param.Opt[string] `json:"phoneNumberSecondary,omitzero"`
	// The post
	PostCode param.Opt[string] `json:"postCode,omitzero"`
	paramObj
}

func (r PatientParam) MarshalJSON() (data []byte, err error) {
	type shadow PatientParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PatientParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Pseudonym struct {
	// The code
	Code string `json:"code,required"`
	// The date of birth (yyyy-MM-dd)
	DateOfBirth time.Time `json:"dateOfBirth" format:"date"`
	// Any of "F", "M", "U".
	Gender Gender `json:"gender"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		DateOfBirth respjson.Field
		Gender      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Pseudonym) RawJSON() string { return r.JSON.raw }
func (r *Pseudonym) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Pseudonym to a PseudonymParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PseudonymParam.Overrides()
func (r Pseudonym) ToParam() PseudonymParam {
	return param.Override[PseudonymParam](json.RawMessage(r.RawJSON()))
}

// The property Code is required.
type PseudonymParam struct {
	// The code
	Code string `json:"code,required"`
	// The date of birth (yyyy-MM-dd)
	DateOfBirth param.Opt[time.Time] `json:"dateOfBirth,omitzero" format:"date"`
	// Any of "F", "M", "U".
	Gender Gender `json:"gender,omitzero"`
	paramObj
}

func (r PseudonymParam) MarshalJSON() (data []byte, err error) {
	type shadow PseudonymParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PseudonymParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Rhesus string

const (
	RhesusPos Rhesus = "pos"
	RhesusNeg Rhesus = "neg"
	Rhesus1   Rhesus = "1"
	Rhesus2   Rhesus = "2"
)

type OrderGetResponse struct {
	// The ID of the order
	ID string `json:"id,required" format:"uuid"`
	// The examinations belonging to the order
	Examinations []OrderExamination `json:"examinations,required"`
	// The laboratory ID where the order will be sent
	LaboratoryID string `json:"laboratoryId,required" format:"uuid"`
	// Identifier of the location (client)
	LocationID string `json:"locationId,required" format:"uuid"`
	// The order creation date-time (yyyy-MM-dd'T'HH:mm:ss.SSSZ)
	OrderCreationDateTime time.Time `json:"orderCreationDateTime,required" format:"date-time"`
	// The order status model
	//
	// Any of "ENTERED", "WAITING_FOR_MATERIAL", "PROCESSING", "CONFIRMATION_PENDING",
	// "FINAL", "DELETED", "ERROR".
	State OrderStateType `json:"state,required"`
	// The order type
	//
	// Any of "DONOR", "BONE_MARROW_DONOR", "PERSONAL", "PSEUDONYM".
	Type OrderType `json:"type,required"`
	// The blood donor data when type is DONOR
	BloodDonor BloodDonor `json:"bloodDonor"`
	// The bone-marrow donor data when type is BONE_MARROW_DONOR
	BoneMarrowDonor BoneMarrowDonor `json:"boneMarrowDonor"`
	// The tags belonging to the order
	OrderTags []string `json:"orderTags"`
	// The patient data when type is PERSONAL
	Patient Patient `json:"patient"`
	// The pseudonym data when type is PSEUDONYM
	Pseudonym Pseudonym `json:"pseudonym"`
	// Add information in key:value pairs object array that are stored with the order
	References map[string]string `json:"references"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Examinations          respjson.Field
		LaboratoryID          respjson.Field
		LocationID            respjson.Field
		OrderCreationDateTime respjson.Field
		State                 respjson.Field
		Type                  respjson.Field
		BloodDonor            respjson.Field
		BoneMarrowDonor       respjson.Field
		OrderTags             respjson.Field
		Patient               respjson.Field
		Pseudonym             respjson.Field
		References            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OrderGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OrderGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderNewParams struct {
	// Fail entire batch on error
	FailOnError param.Opt[bool] `query:"failOnError,omitzero" json:"-"`
	// A laboratory ID that is used to complete missing ids in the submitted order
	LaboratoryID param.Opt[string] `query:"laboratoryId,omitzero" format:"uuid" json:"-"`
	// A location ID that is used to complete missing IDs in the submitted order
	LocationID param.Opt[string] `query:"locationId,omitzero" format:"uuid" json:"-"`
	Body       []OrderNewParamsBody
	paramObj
}

func (r OrderNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *OrderNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// URLQuery serializes [OrderNewParams]'s query parameters as `url.Values`.
func (r OrderNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The properties LaboratoryID, LocationID, Type are required.
type OrderNewParamsBody struct {
	// The laboratory ID where the order will be sent
	LaboratoryID string `json:"laboratoryId,required" format:"uuid"`
	// Identifier of the location (client)
	LocationID string `json:"locationId,required" format:"uuid"`
	// The order type
	//
	// Any of "DONOR", "BONE_MARROW_DONOR", "PERSONAL", "PSEUDONYM".
	Type OrderType `json:"type,omitzero,required"`
	// The blood donor data when type is DONOR
	BloodDonor BloodDonorParam `json:"bloodDonor,omitzero"`
	// The bone-marrow donor data when type is BONE_MARROW_DONOR
	BoneMarrowDonor BoneMarrowDonorParam `json:"boneMarrowDonor,omitzero"`
	// The items belonging to the order
	Items []OrderNewParamsBodyItem `json:"items,omitzero"`
	// The patient data when type is PERSONAL
	Patient PatientParam `json:"patient,omitzero"`
	// The pseudonym data when type is PSEUDONYM
	Pseudonym PseudonymParam `json:"pseudonym,omitzero"`
	// Information in key:value pairs object array. Information stored here is
	// reflected back when reading the orders again thus allows for storing keys.
	References map[string]string `json:"references,omitzero"`
	// The tags belonging to the order
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r OrderNewParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CollectionTime, SampleCode are required.
type OrderNewParamsBodyItem struct {
	// The collection time of the sample. It is recommended to use the precise time. If
	// only the Date is availabe opt for 12 o'clock to avoid misterpretations in
	// timezones.
	CollectionTime time.Time `json:"collectionTime,required" format:"date-time"`
	// The sample codes for this order. They must be permitted by the laboratory.
	// Lablink verifies that the order is not colliding with any other orders to that
	// laboratory. In such cases it is possible that the order gets declined with a
	// Conflict error.
	SampleCode string `json:"sampleCode,required"`
	// The code of the examination as configured. As an alternative id can be used.
	ExaminationCode param.Opt[string] `json:"examinationCode,omitzero"`
	// ExaminationID as provided by the GetExaminations endpoint. Aas an alternative id
	// can be used.
	ExaminationID param.Opt[string] `json:"examinationId,omitzero" format:"uuid"`
	// A field to hold an external reference for the examination. Those references will
	// be reflected back when reading the orders again thus allows for storing keys.
	Reference param.Opt[string] `json:"reference,omitzero"`
	paramObj
}

func (r OrderNewParamsBodyItem) MarshalJSON() (data []byte, err error) {
	type shadow OrderNewParamsBodyItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderNewParamsBodyItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OrderDeleteParams struct {
	// List of order IDs to delete
	OrderIDs []string `query:"orderIds,omitzero,required" format:"uuid" json:"-"`
	paramObj
}

// URLQuery serializes [OrderDeleteParams]'s query parameters as `url.Values`.
func (r OrderDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
