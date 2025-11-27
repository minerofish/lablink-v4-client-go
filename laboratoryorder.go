// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package lablinkv4client

import (
	"github.com/minerofish/lablink-v4-client-go/option"
)

// LaboratoryOrderService contains methods and other services that help with
// interacting with the lablink-v4-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLaboratoryOrderService] method instead.
type LaboratoryOrderService struct {
	Options []option.RequestOption
}

// NewLaboratoryOrderService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLaboratoryOrderService(opts ...option.RequestOption) (r LaboratoryOrderService) {
	r = LaboratoryOrderService{}
	r.Options = opts
	return
}
