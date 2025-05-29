// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package noratestproject10000_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/nora-test-project-10000-go"
	"github.com/stainless-sdks/nora-test-project-10000-go/internal/testutil"
	"github.com/stainless-sdks/nora-test-project-10000-go/option"
)

func TestStoreListInventory(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := noratestproject10000.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Store.ListInventory(context.TODO())
	if err != nil {
		var apierr *noratestproject10000.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
