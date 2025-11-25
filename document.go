// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"github.com/minerofish/lablink-v4-client-go/internal/apijson"
	"github.com/minerofish/lablink-v4-client-go/option"
	"github.com/minerofish/lablink-v4-client-go/packages/respjson"
)

// DocumentService contains methods and other services that help with interacting
// with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDocumentService] method instead.
type DocumentService struct {
	Options []option.RequestOption
}

// NewDocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentService(opts ...option.RequestOption) (r DocumentService) {
	r = DocumentService{}
	r.Options = opts
	return
}

type Page struct {
	// The actual page number
	CurrentPage int64 `json:"currentPage"`
	// The items
	Items any `json:"items"`
	// The number of items per page
	PageSize int64 `json:"pageSize"`
	// The total count of items
	TotalCount int64 `json:"totalCount"`
	// The total pages
	TotalPages int64 `json:"totalPages"`
	// The transaction identifier
	TransactionID string `json:"transactionId" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CurrentPage   respjson.Field
		Items         respjson.Field
		PageSize      respjson.Field
		TotalCount    respjson.Field
		TotalPages    respjson.Field
		TransactionID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Page) RawJSON() string { return r.JSON.raw }
func (r *Page) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
