package apijson_test

import (
	"encoding/json"
	"github.com/nora-test-account/nora-test-project-go-10000/internal/apijson"
	"github.com/nora-test-account/nora-test-project-go-10000/packages/respjson"
	"testing"
)

type StructWithNullExtraField struct {
	Results []string `json:"results,required"`
	JSON    struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

func (r *StructWithNullExtraField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func TestDecodeWithNullExtraField(t *testing.T) {
	raw := `{"something_else":null}`
	var dst *StructWithNullExtraField
	err := json.Unmarshal([]byte(raw), &dst)
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}
}
