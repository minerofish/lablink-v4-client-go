// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/minerofish/lablink-v4-client-go/internal/apijson"
	"github.com/minerofish/lablink-v4-client-go/internal/apiquery"
	"github.com/minerofish/lablink-v4-client-go/internal/requestconfig"
	"github.com/minerofish/lablink-v4-client-go/option"
	"github.com/minerofish/lablink-v4-client-go/packages/param"
)

// OrderTagService contains methods and other services that help with interacting
// with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrderTagService] method instead.
type OrderTagService struct {
	Options []option.RequestOption
}

// NewOrderTagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrderTagService(opts ...option.RequestOption) (r OrderTagService) {
	r = OrderTagService{}
	r.Options = opts
	return
}

// Remove tag(s) from order(s) - n:n relation
func (r *OrderTagService) Delete(ctx context.Context, body OrderTagDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v4/orders/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return
}

// Tags order or orders with tag or tags. n:n relation
func (r *OrderTagService) Tag(ctx context.Context, body OrderTagTagParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v4/orders/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type OrderTagDeleteParams struct {
	// List of order IDs to remove tags from
	OrderIDs []string `query:"orderIds,omitzero,required" format:"uuid" json:"-"`
	// List of tags to be removed
	Tags []string `query:"tags,omitzero,required" json:"-"`
	paramObj
}

// URLQuery serializes [OrderTagDeleteParams]'s query parameters as `url.Values`.
func (r OrderTagDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrderTagTagParams struct {
	OrderIDs []string `json:"orderIds,omitzero,required" format:"uuid"`
	Tags     []string `json:"tags,omitzero,required"`
	paramObj
}

func (r OrderTagTagParams) MarshalJSON() (data []byte, err error) {
	type shadow OrderTagTagParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OrderTagTagParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
