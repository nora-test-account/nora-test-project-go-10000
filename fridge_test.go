// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package noratestproject10000_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nora-test-account/nora-test-project-go-10000"
	"github.com/nora-test-account/nora-test-project-go-10000/internal/testutil"
	"github.com/nora-test-account/nora-test-project-go-10000/option"
)

func TestFridgeAddItem(t *testing.T) {
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
	_, err := client.Fridge.AddItem(context.TODO(), noratestproject10000.FridgeAddItemParams{
		Food: noratestproject10000.FoodParam{
			Name:     "name",
			Quantity: "quantity",
		},
	})
	if err != nil {
		var apierr *noratestproject10000.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFridgeDeleteItem(t *testing.T) {
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
	_, err := client.Fridge.DeleteItem(context.TODO(), "name")
	if err != nil {
		var apierr *noratestproject10000.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFridgeListItems(t *testing.T) {
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
	_, err := client.Fridge.ListItems(context.TODO())
	if err != nil {
		var apierr *noratestproject10000.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFridgeGetItem(t *testing.T) {
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
	_, err := client.Fridge.GetItem(context.TODO(), "name")
	if err != nil {
		var apierr *noratestproject10000.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFridgeUpdateItem(t *testing.T) {
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
	_, err := client.Fridge.UpdateItem(
		context.TODO(),
		"name",
		noratestproject10000.FridgeUpdateItemParams{
			Food: noratestproject10000.FoodParam{
				Name:     "name",
				Quantity: "quantity",
			},
		},
	)
	if err != nil {
		var apierr *noratestproject10000.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
