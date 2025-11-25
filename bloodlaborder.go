// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/minerofish/lablink-v4-client-go/internal/apijson"
	shimjson "github.com/minerofish/lablink-v4-client-go/internal/encoding/json"
	"github.com/minerofish/lablink-v4-client-go/internal/requestconfig"
	"github.com/minerofish/lablink-v4-client-go/option"
	"github.com/minerofish/lablink-v4-client-go/packages/param"
	"github.com/minerofish/lablink-v4-client-go/packages/respjson"
)

// BloodlabOrderService contains methods and other services that help with
// interacting with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBloodlabOrderService] method instead.
type BloodlabOrderService struct {
	Options []option.RequestOption
}

// NewBloodlabOrderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBloodlabOrderService(opts ...option.RequestOption) (r BloodlabOrderService) {
	r = BloodlabOrderService{}
	r.Options = opts
	return
}

// Gets all order which are not downloaded yet.
func (r *BloodlabOrderService) ListNewOrders(ctx context.Context, opts ...option.RequestOption) (res *[]BloodlabOrderListNewOrdersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/bloodlab/orders"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Mark orders as downloaded.
func (r *BloodlabOrderService) MarkAsDownloaded(ctx context.Context, body BloodlabOrderMarkAsDownloadedParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v4/bloodlab/orders/downloaded"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Set orders state. (unruly)
func (r *BloodlabOrderService) SetState(ctx context.Context, body BloodlabOrderSetStateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v4/bloodlab/orders/state"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type BloodlabOrderListNewOrdersResponse struct {
	// The examinations belonging to the order
	Examinations []BloodlabOrderListNewOrdersResponseExamination `json:"examinations,required"`
	// The laboratory ID
	LaboratoryID string `json:"laboratoryId,required"`
	// Identifier of the location (client)
	LocationID string `json:"locationId,required"`
	// The order type
	//
	// Any of "DONOR", "BONE_MARROW_DONOR", "PERSONAL", "PSEUDONYM".
	Type OrderType `json:"type,required"`
	// The ID of the order
	ID string `json:"id" format:"uuid"`
	// The blood donor data when type is DONOR
	BloodDonor BloodlabOrderListNewOrdersResponseBloodDonor `json:"bloodDonor"`
	// The bone-marrow donor data when type is BONE_MARROW_DONOR
	BoneMarrowDonor BloodlabOrderListNewOrdersResponseBoneMarrowDonor `json:"boneMarrowDonor"`
	// The order creation date-time (yyyy-MM-dd'T'HH:mm:ss.SSSZ)
	OrderCreationDateTime time.Time `json:"orderCreationDateTime" format:"date-time"`
	// The patient data when type is PERSONAL
	Patient BloodlabOrderListNewOrdersResponsePatient `json:"patient"`
	// The pseudonym data when type is PSEUDONYM
	Pseudonym BloodlabOrderListNewOrdersResponsePseudonym `json:"pseudonym"`
	// The external unique identifier of the order
	Reference string `json:"reference"`
	// Additional external information in key:value pairs object array
	ReferenceMap []map[string]string `json:"referenceMap"`
	// The order status model
	//
	// Any of "ENTERED", "WAITING_FOR_MATERIAL", "PROCESSING", "CONFIRMATION_PENDING",
	// "FINAL", "DELETED", "ERROR".
	State OrderStateType `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Examinations          respjson.Field
		LaboratoryID          respjson.Field
		LocationID            respjson.Field
		Type                  respjson.Field
		ID                    respjson.Field
		BloodDonor            respjson.Field
		BoneMarrowDonor       respjson.Field
		OrderCreationDateTime respjson.Field
		Patient               respjson.Field
		Pseudonym             respjson.Field
		Reference             respjson.Field
		ReferenceMap          respjson.Field
		State                 respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BloodlabOrderListNewOrdersResponse) RawJSON() string { return r.JSON.raw }
func (r *BloodlabOrderListNewOrdersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BloodlabOrderListNewOrdersResponseExamination struct {
	// The ID of the order-examination
	ID string `json:"id,required"`
	// The ID of the examination
	ExaminationID string `json:"examinationId,required"`
	// The ID of the examination
	OrganizationExaminationID string `json:"organizationExaminationId,required"`
	// The external unique identifier of the order
	Reference string `json:"reference"`
	// The results belonging to the order
	Results []BloodlabOrderListNewOrdersResponseExaminationResult `json:"results"`
	// The sample codes
	SampleCodes []string `json:"sampleCodes"`
	// The sample date
	SampleDate time.Time `json:"sampleDate" format:"date"`
	// The sample time
	SampleTime string `json:"sampleTime" format:"time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		ExaminationID             respjson.Field
		OrganizationExaminationID respjson.Field
		Reference                 respjson.Field
		Results                   respjson.Field
		SampleCodes               respjson.Field
		SampleDate                respjson.Field
		SampleTime                respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BloodlabOrderListNewOrdersResponseExamination) RawJSON() string { return r.JSON.raw }
func (r *BloodlabOrderListNewOrdersResponseExamination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BloodlabOrderListNewOrdersResponseExaminationResult struct {
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
func (r BloodlabOrderListNewOrdersResponseExaminationResult) RawJSON() string { return r.JSON.raw }
func (r *BloodlabOrderListNewOrdersResponseExaminationResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The blood donor data when type is DONOR
type BloodlabOrderListNewOrdersResponseBloodDonor struct {
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
func (r BloodlabOrderListNewOrdersResponseBloodDonor) RawJSON() string { return r.JSON.raw }
func (r *BloodlabOrderListNewOrdersResponseBloodDonor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The bone-marrow donor data when type is BONE_MARROW_DONOR
type BloodlabOrderListNewOrdersResponseBoneMarrowDonor struct {
	// The donor code
	DonorCode string `json:"donorCode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DonorCode   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BloodlabOrderListNewOrdersResponseBoneMarrowDonor) RawJSON() string { return r.JSON.raw }
func (r *BloodlabOrderListNewOrdersResponseBoneMarrowDonor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The patient data when type is PERSONAL
type BloodlabOrderListNewOrdersResponsePatient struct {
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
func (r BloodlabOrderListNewOrdersResponsePatient) RawJSON() string { return r.JSON.raw }
func (r *BloodlabOrderListNewOrdersResponsePatient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The pseudonym data when type is PSEUDONYM
type BloodlabOrderListNewOrdersResponsePseudonym struct {
	// The code
	Code string `json:"code"`
	// The date of birth (yyyy-MM-dd)
	DateOfBirth time.Time `json:"dateOfBirth" format:"date"`
	// The gender (Female, Male, Unknown)
	//
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
func (r BloodlabOrderListNewOrdersResponsePseudonym) RawJSON() string { return r.JSON.raw }
func (r *BloodlabOrderListNewOrdersResponsePseudonym) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BloodlabOrderMarkAsDownloadedParams struct {
	Body []string
	paramObj
}

func (r BloodlabOrderMarkAsDownloadedParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *BloodlabOrderMarkAsDownloadedParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

type BloodlabOrderSetStateParams struct {
	Body []BloodlabOrderSetStateParamsBody
	paramObj
}

func (r BloodlabOrderSetStateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *BloodlabOrderSetStateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// The properties OrderID, State are required.
type BloodlabOrderSetStateParamsBody struct {
	// The order ID which state will be changed
	OrderID string `json:"orderId,required" format:"uuid"`
	// The order status model
	//
	// Any of "ENTERED", "WAITING_FOR_MATERIAL", "PROCESSING", "CONFIRMATION_PENDING",
	// "FINAL", "DELETED", "ERROR".
	State OrderStateType `json:"state,omitzero,required"`
	paramObj
}

func (r BloodlabOrderSetStateParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow BloodlabOrderSetStateParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BloodlabOrderSetStateParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
