package api

import (
	"context"

	"github.com/AnxVit/ozon_1/api/graphQL/models"
)

type IHabrService interface {
	CreatePost(ctx context.Context, post *models.Post) (string, error)
	GetAllPost(ctx context.Context) ([]*models.Post, error)
	GetPostByID(ctx context.Context, id string, limit *int, offset *int) (*models.Post, error)
	CreateComment(ctx context.Context, postID *string, comment *models.Comment) error
}
