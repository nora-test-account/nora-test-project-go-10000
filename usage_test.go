// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package noratestproject10000_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/nora-test-project-10000-go"
	"github.com/stainless-sdks/nora-test-project-10000-go/internal/testutil"
	"github.com/stainless-sdks/nora-test-project-10000-go/option"
)

func TestUsage(t *testing.T) {
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
	foods, err := client.Fridge.ListItems(context.TODO())
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", foods)
}
