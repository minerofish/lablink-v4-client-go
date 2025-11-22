// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/lablink-v4-client-go/internal/apijson"
	"github.com/stainless-sdks/lablink-v4-client-go/internal/requestconfig"
	"github.com/stainless-sdks/lablink-v4-client-go/option"
	"github.com/stainless-sdks/lablink-v4-client-go/packages/param"
)

// TokenService contains methods and other services that help with interacting with
// the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenService] method instead.
type TokenService struct {
	Options []option.RequestOption
}

// NewTokenService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTokenService(opts ...option.RequestOption) (r TokenService) {
	r = TokenService{}
	r.Options = opts
	return
}

// OAuth 2.0 Token endpoint for PKCE flow
func (r *TokenService) New(ctx context.Context, body TokenNewParams, opts ...option.RequestOption) (res *AuthToken, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TokenNewParams struct {
	// The client identifier
	ClientID string `json:"client_id,required"`
	// The authorization code received from the authorization endpoint
	Code string `json:"code,required"`
	// Code verifier for PKCE
	CodeVerifier string `json:"code_verifier,required"`
	// Must be "authorization_code" for PKCE flow
	//
	// Any of "authorization_code".
	GrantType TokenNewParamsGrantType `json:"grant_type,omitzero,required"`
	// The redirect URI used in the authorization request
	RedirectUri string `json:"redirect_uri,required" format:"uri"`
	paramObj
}

func (r TokenNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TokenNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TokenNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Must be "authorization_code" for PKCE flow
type TokenNewParamsGrantType string

const (
	TokenNewParamsGrantTypeAuthorizationCode TokenNewParamsGrantType = "authorization_code"
)
