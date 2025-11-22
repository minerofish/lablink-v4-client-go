// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/lablink-v4-client-go/internal/apijson"
	"github.com/stainless-sdks/lablink-v4-client-go/internal/requestconfig"
	"github.com/stainless-sdks/lablink-v4-client-go/option"
	"github.com/stainless-sdks/lablink-v4-client-go/packages/respjson"
)

// BloodlabService contains methods and other services that help with interacting
// with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBloodlabService] method instead.
type BloodlabService struct {
	Options []option.RequestOption
	Orders  BloodlabOrderService
}

// NewBloodlabService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBloodlabService(opts ...option.RequestOption) (r BloodlabService) {
	r = BloodlabService{}
	r.Options = opts
	r.Orders = NewBloodlabOrderService(opts...)
	return
}

// Lablink-client uses this to get all examinations for making the Lablink catalog
// This function is redundant to the api get examinations and should be eliminated
// in the future
func (r *BloodlabService) ListExaminations(ctx context.Context, opts ...option.RequestOption) (res *[]BloodlabOrganizationRelation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/bloodlab/examinations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets all organizations' relations including locations and examinations.
func (r *BloodlabService) ListOrganizationRelations(ctx context.Context, opts ...option.RequestOption) (res *[]BloodlabOrganizationRelation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/bloodlab/organization-relations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BloodlabOrganizationRelation struct {
	// The organization's examinations
	Examinations []BloodlabOrganizationRelationExamination `json:"examinations"`
	// The organization's locations
	Locations []BloodlabOrganizationRelationLocation `json:"locations"`
	// The organization ID
	OrganizationID string `json:"organizationId" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Examinations   respjson.Field
		Locations      respjson.Field
		OrganizationID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BloodlabOrganizationRelation) RawJSON() string { return r.JSON.raw }
func (r *BloodlabOrganizationRelation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BloodlabOrganizationRelationExamination struct {
	// The ID
	ID string `json:"id" format:"uuid"`
	// The code
	Code string `json:"code"`
	// The description
	Description string `json:"description"`
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
		Name           respjson.Field
		PossibleValues respjson.Field
		Unit           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BloodlabOrganizationRelationExamination) RawJSON() string { return r.JSON.raw }
func (r *BloodlabOrganizationRelationExamination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BloodlabOrganizationRelationLocation struct {
	// The ID
	ID string `json:"id" format:"uuid"`
	// The name
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BloodlabOrganizationRelationLocation) RawJSON() string { return r.JSON.raw }
func (r *BloodlabOrganizationRelationLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
