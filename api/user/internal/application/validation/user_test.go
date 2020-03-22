package validation

import (
	"reflect"
	"testing"

	"github.com/16francs/gran/api/user/internal/application/request"
	"github.com/16francs/gran/api/user/internal/domain"
)

func TestUserRequestValidation_CreateUser(t *testing.T) {
	// Defined variables
	u := &request.CreateUser{
		Email:                "hoge@hoge.com",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	}

	// Start test
	target := NewUserRequestValidation()

	want := []*domain.ValidationError(nil)

	got := target.CreateUser(u)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
