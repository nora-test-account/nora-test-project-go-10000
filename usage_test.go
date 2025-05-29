// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package noratestproject10000_test

import (
	"context"
	"os"
	"testing"

	"github.com/nora-test-account/nora-test-project-go-10000"
	"github.com/nora-test-account/nora-test-project-go-10000/internal/testutil"
	"github.com/nora-test-account/nora-test-project-go-10000/option"
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
