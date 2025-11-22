// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/minerofish/lablink-v4-client-go/internal/apiquery"
	"github.com/minerofish/lablink-v4-client-go/internal/requestconfig"
	"github.com/minerofish/lablink-v4-client-go/option"
	"github.com/minerofish/lablink-v4-client-go/packages/param"
)

// AuthorizeService contains methods and other services that help with interacting
// with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthorizeService] method instead.
type AuthorizeService struct {
	Options []option.RequestOption
}

// NewAuthorizeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuthorizeService(opts ...option.RequestOption) (r AuthorizeService) {
	r = AuthorizeService{}
	r.Options = opts
	return
}

// OAuth 2.0 Authorization endpoint for PKCE flow
func (r *AuthorizeService) Get(ctx context.Context, query AuthorizeGetParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/authorize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AuthorizeGetParams struct {
	// The client identifier
	ClientID string `query:"client_id,required" json:"-"`
	// Code challenge for PKCE
	CodeChallenge string `query:"code_challenge,required" json:"-"`
	// Code challenge method, must be S256
	//
	// Any of "S256".
	CodeChallengeMethod AuthorizeGetParamsCodeChallengeMethod `query:"code_challenge_method,omitzero,required" json:"-"`
	// The redirect URI where authorization code will be sent
	RedirectUri string `query:"redirect_uri,required" format:"uri" json:"-"`
	// Must be "code" for PKCE flow
	//
	// Any of "code".
	ResponseType AuthorizeGetParamsResponseType `query:"response_type,omitzero,required" json:"-"`
	// The requested scope
	Scope param.Opt[string] `query:"scope,omitzero" json:"-"`
	// An opaque value to prevent CSRF attacks
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuthorizeGetParams]'s query parameters as `url.Values`.
func (r AuthorizeGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Code challenge method, must be S256
type AuthorizeGetParamsCodeChallengeMethod string

const (
	AuthorizeGetParamsCodeChallengeMethodS256 AuthorizeGetParamsCodeChallengeMethod = "S256"
)

// Must be "code" for PKCE flow
type AuthorizeGetParamsResponseType string

const (
	AuthorizeGetParamsResponseTypeCode AuthorizeGetParamsResponseType = "code"
)
