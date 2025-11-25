// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"github.com/minerofish/lablink-v4-client-go/option"
)

// OrderStateService contains methods and other services that help with interacting
// with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrderStateService] method instead.
type OrderStateService struct {
	Options []option.RequestOption
}

// NewOrderStateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrderStateService(opts ...option.RequestOption) (r OrderStateService) {
	r = OrderStateService{}
	r.Options = opts
	return
}
