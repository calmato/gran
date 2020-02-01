package validation

import (
	"reflect"
	"testing"

	"github.com/16francs/gran/api/user/internal/application/request"
	"github.com/16francs/gran/api/user/internal/domain"
)

func TestUserRequestValidation_Createuser(t *testing.T) {
	target := NewUserRequestValidation()

	want := []*domain.ValidationError(nil)

	u := &request.CreateUser{
		Email:                "hoge@hoge.com",
		Password:             "12345678",
		PasswordConfirmation: "12345678",
	}

	got := target.CreateUser(u)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %#v, but %#v", want, got)
	}
}
