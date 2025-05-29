// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package noratestproject10000

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/nora-test-project-10000-go/internal/requestconfig"
	"github.com/stainless-sdks/nora-test-project-10000-go/option"
)

// StoreService contains methods and other services that help with interacting with
// the nora-test-project-10000 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStoreService] method instead.
type StoreService struct {
	Options []option.RequestOption
	Order   StoreOrderService
}

// NewStoreService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStoreService(opts ...option.RequestOption) (r StoreService) {
	r = StoreService{}
	r.Options = opts
	r.Order = NewStoreOrderService(opts...)
	return
}

// Returns a map of status codes to quantities
func (r *StoreService) ListInventory(ctx context.Context, opts ...option.RequestOption) (res *StoreListInventoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "store/inventory"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type StoreListInventoryResponse map[string]int64
