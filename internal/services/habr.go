package services

import (
	"context"

	"github.com/AnxVit/ozon_1/api/graphQL/models"
	"github.com/AnxVit/ozon_1/internal/transport/graphql/api"
)

var _ api.IHabrService = HabrService{}

type HabrService struct {
	habrService IHabrRepositories
}

func NewHabrService(habrService IHabrRepositories) HabrService {
	return HabrService{
		habrService: habrService,
	}
}

func (h HabrService) CreatePost(ctx context.Context, post *models.Post) (string, error) {
	id, err := h.habrService.CreatePost(ctx, post)
	if err != nil {
		return "", err
	}

	return id, nil
}
func (h HabrService) GetAllPost(ctx context.Context) ([]*models.Post, error) {
	posts, err := h.habrService.GetAllPost(ctx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
func (h HabrService) GetPostByID(ctx context.Context, id string, limit *int, offset *int) (*models.Post, error) {
	post, err := h.habrService.GetPostByID(ctx, id, limit, offset)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (h HabrService) CreateComment(ctx context.Context, postID *string, comment *models.Comment) error {
	var err error
	if postID != nil {
		err = h.habrService.CreatePostComment(ctx, postID, comment)
	} else {
		err = h.habrService.CreateComment(ctx, comment)
	}
	if err != nil {
		return err
	}

	return nil
}
