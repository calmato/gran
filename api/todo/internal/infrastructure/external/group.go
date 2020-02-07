package external

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/xerrors"

	"github.com/16francs/gran/api/todo/internal/domain"
	"github.com/16francs/gran/api/todo/internal/domain/repository"
	"github.com/16francs/gran/api/todo/internal/infrastructure/external/response"
	"github.com/16francs/gran/api/todo/middleware"
)

type groupAPI struct {
	url    string
	client http.Client
}

// NewGroupAPI - GroupAPIの生成
func NewGroupAPI(url string) repository.GroupRepository {
	return &groupAPI{
		url:    url,
		client: http.Client{},
	}
}

func (ga *groupAPI) Show(ctx context.Context, groupID string) (*domain.Group, error) {
	t, err := getToken(ctx)
	if err != nil {
		return nil, err
	}

	url := strings.Join([]string{ga.url, groupID}, "/")

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", t)

	res, err := ga.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	gr := &response.Group{}

	err = json.NewDecoder(res.Body).Decode(&gr)
	if err != nil {
		return nil, err
	}

	g := &domain.Group{
		ID:          gr.ID,
		Name:        gr.Name,
		Description: gr.Description,
		UserRefs:    gr.UserRefs,
		CreatedAt:   gr.CreatedAt,
		UpdatedAt:   gr.UpdatedAt,
	}

	return g, nil
}

func getToken(ctx context.Context) (string, error) {
	gc, err := middleware.GinContextFromContext(ctx)
	if err != nil {
		return "", xerrors.New("Cannot convert to gin.Context")
	}

	// t <- Bearer含めたToken
	t := gc.GetHeader("Authorization")
	if t == "" {
		return "", xerrors.New("Authorization Header is not contain.")
	}

	return t, nil
}
