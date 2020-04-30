package validation

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/calmato/gran/api/todo/internal/domain"
)

func TestGroupDomainValidation_Group(t *testing.T) {
	current := time.Now()

	testCases := map[string]struct {
		Group    *domain.Group
		Expected []*domain.ValidationError
	}{
		"ok": {
			Group: &domain.Group{
				ID:          "group-id",
				Name:        "テストグループ",
				Description: "グループの説明",
				UserIDs:     make([]string, 0),
				BoardIDs:    make([]string, 0),
				CreatedAt:   current,
				UpdatedAt:   current,
			},
			Expected: make([]*domain.ValidationError, 0),
		},
	}

	for result, testCase := range testCases {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Start test
		t.Run(result, func(t *testing.T) {
			target := NewGroupDomainValidation()

			got := target.Group(ctx, testCase.Group)
			if !reflect.DeepEqual(got, testCase.Expected) {
				t.Fatalf("want %#v, but %#v", testCase.Expected, got)
				return
			}
		})
	}
}
