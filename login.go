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
	"github.com/stainless-sdks/lablink-v4-client-go/packages/respjson"
)

// LoginService contains methods and other services that help with interacting with
// the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLoginService] method instead.
type LoginService struct {
	Options []option.RequestOption
}

// NewLoginService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLoginService(opts ...option.RequestOption) (r LoginService) {
	r = LoginService{}
	r.Options = opts
	return
}

// OAuth 2.0 Ressource Owner Password Credentials Grant
func (r *LoginService) Authenticate(ctx context.Context, body LoginAuthenticateParams, opts ...option.RequestOption) (res *AuthToken, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v4/login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AuthToken struct {
	// The access token
	AccessToken string `json:"access_token"`
	// The access token's expiry time window
	ExpiresIn int64 `json:"expires_in"`
	// The date-time when the token becomes valid
	NotBeforePolicy int64 `json:"not-before-policy"`
	// The refresh token's expiry time window
	RefreshExpiresIn int64 `json:"refresh_expires_in"`
	// The refresh token
	RefreshToken string `json:"refresh_token"`
	// The session scope (api only)
	Scope string `json:"scope"`
	// The session state
	SessionState string `json:"session_state"`
	// The token type (Bearer only)
	TokenType string `json:"token_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken      respjson.Field
		ExpiresIn        respjson.Field
		NotBeforePolicy  respjson.Field
		RefreshExpiresIn respjson.Field
		RefreshToken     respjson.Field
		Scope            respjson.Field
		SessionState     respjson.Field
		TokenType        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthToken) RawJSON() string { return r.JSON.raw }
func (r *AuthToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LoginAuthenticateParams struct {
	// The password
	Password string `json:"password,required" format:"password"`
	// The username
	Username string `json:"username,required"`
	paramObj
}

func (r LoginAuthenticateParams) MarshalJSON() (data []byte, err error) {
	type shadow LoginAuthenticateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LoginAuthenticateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
