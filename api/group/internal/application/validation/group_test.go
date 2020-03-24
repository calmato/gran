package validation

import (
	"reflect"
	"testing"

	"github.com/16francs/gran/api/group/internal/application/request"
	"github.com/16francs/gran/api/group/internal/domain"
)

func TestGroupRequestValidation_CreateGroup(t *testing.T) {
	testCases := map[string]struct {
		Request  *request.CreateGroup
		Expected []*domain.ValidationError
	}{
		"ok": {
			Request: &request.CreateGroup{
				Name:        "テストグループ",
				Description: "説明",
			},
			Expected: make([]*domain.ValidationError, 0),
		},
	}

	for result, testCase := range testCases {
		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupRequestValidation()

			got := target.CreateGroup(testCase.Request)
			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}

func TestGroupRequestValidation_UpdateGroup(t *testing.T) {
	testCases := map[string]struct {
		Request  *request.UpdateGroup
		Expected []*domain.ValidationError
	}{
		"ok": {
			Request: &request.UpdateGroup{
				Name:        "テストグループ",
				Description: "説明",
			},
			Expected: make([]*domain.ValidationError, 0),
		},
	}

	for result, testCase := range testCases {
		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupRequestValidation()

			got := target.UpdateGroup(testCase.Request)
			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}
