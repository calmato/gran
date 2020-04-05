package application

import (
	"context"
	"encoding/base64"
	"strings"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/user/internal/application/request"
	"github.com/16francs/gran/api/user/internal/application/validation"
	"github.com/16francs/gran/api/user/internal/domain"
	"github.com/16francs/gran/api/user/internal/domain/service"
)

// UserApplication - UserApplicationインターフェース
type UserApplication interface {
	Create(ctx context.Context, req *request.CreateUser) error
	ShowProfile(ctx context.Context) (*domain.User, error)
	UpdateProfile(ctx context.Context, req *request.UpdateProfile) error
}

type userApplication struct {
	userRequestValidation validation.UserRequestValidation
	userService           service.UserService
}

// NewUserApplication - UserApplicationの生成
func NewUserApplication(urv validation.UserRequestValidation, us service.UserService) UserApplication {
	return &userApplication{
		userRequestValidation: urv,
		userService:           us,
	}
}

func (ua *userApplication) Create(ctx context.Context, req *request.CreateUser) error {
	if ves := ua.userRequestValidation.CreateUser(req); len(ves) > 0 {
		err := xerrors.New("Failed to Application/RequestValidation")
		return domain.InvalidRequestValidation.New(err, ves...)
	}

	u := &domain.User{
		Email:    req.Email,
		Password: req.Password,
	}

	if _, err := ua.userService.Create(ctx, u); err != nil {
		return err
	}

	return nil
}

func (ua *userApplication) ShowProfile(ctx context.Context) (*domain.User, error) {
	u, err := ua.userService.Authentication(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ua *userApplication) UpdateProfile(ctx context.Context, req *request.UpdateProfile) error {
	u, err := ua.userService.Authentication(ctx)
	if err != nil {
		return err
	}

	if ves := ua.userRequestValidation.UpdateProfile(req); len(ves) > 0 {
		err := xerrors.New("Failed to Application/RequestValidation")
		return domain.InvalidRequestValidation.New(err, ves...)
	}

	thumbnailURL := ""

	if req.Thumbnail != "" {
		// data:image/png;base64,iVBORw0KGgoAAAA... みたいなのうちの
		// `data:image/png;base64,` の部分を無くした []byte を取得
		b64data := req.Thumbnail[strings.IndexByte(req.Thumbnail, ',')+1:]

		data, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			err = xerrors.Errorf("Failed to Application: %w", err)
			return domain.Unknown.New(err)
		}

		thumbnailURL, err = ua.userService.UploadThumbnail(ctx, data)
		if err != nil {
			return err
		}
	}

	u.Name = req.Name
	u.DisplayName = req.DisplayName
	u.Email = req.Email
	u.PhoneNumber = req.PhoneNumber
	u.ThumbnailURL = thumbnailURL
	u.Biography = req.Biography

	if _, err := ua.userService.Update(ctx, u); err != nil {
		return err
	}

	return nil
}
