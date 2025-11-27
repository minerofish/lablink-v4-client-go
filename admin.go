// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"context"
	"net/http"
	"slices"

	"github.com/minerofish/lablink-v4-client-go/internal/apijson"
	"github.com/minerofish/lablink-v4-client-go/internal/requestconfig"
	"github.com/minerofish/lablink-v4-client-go/option"
	"github.com/minerofish/lablink-v4-client-go/packages/param"
)

// AdminService contains methods and other services that help with interacting with
// the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminService] method instead.
type AdminService struct {
	Options []option.RequestOption
}

// NewAdminService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAdminService(opts ...option.RequestOption) (r AdminService) {
	r = AdminService{}
	r.Options = opts
	return
}

// Invite a new user to the organization.
func (r *AdminService) InviteUser(ctx context.Context, body AdminInviteUserParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v4/admin/invite"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type AdminInviteUserParams struct {
	// The email of the user to invite
	Email string `json:"email,required" format:"email"`
	// The ID of the organization to which the user is invited
	ForOrganizationID string `json:"forOrganizationId,required" format:"uuid"`
	// Any of "ORGANIZATION_ADMIN", "ORGANIZATION_USER".
	Role AdminInviteUserParamsRole `json:"role,omitzero,required"`
	// The IDs of the locations the user will have access to
	ForLocations []string `json:"forLocations,omitzero" format:"uuid"`
	paramObj
}

func (r AdminInviteUserParams) MarshalJSON() (data []byte, err error) {
	type shadow AdminInviteUserParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminInviteUserParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminInviteUserParamsRole string

const (
	AdminInviteUserParamsRoleOrganizationAdmin AdminInviteUserParamsRole = "ORGANIZATION_ADMIN"
	AdminInviteUserParamsRoleOrganizationUser  AdminInviteUserParamsRole = "ORGANIZATION_USER"
)
