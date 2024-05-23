package interfaces_test

import (
	"context"

	"github.com/AnxVit/ozon_1/api/graphQL/models"
)

var (
	CreateCommentFn func(*models.Comment) error
	CreatePostFn    func(*models.Post) (string, error)
	GetAllPostFn    func() ([]*models.Post, error)
	GetPostByIDFn   func(string) (*models.Post, error)
)

func (q *fakePostService) CreateComment(ctx context.Context, postID *string, comment *models.Comment) error {
	return CreateCommentFn(comment)
}

func (q *fakePostService) CreatePost(ctx context.Context, post *models.Post) (string, error) {
	return CreatePostFn(post)
}

func (q *fakePostService) GetAllPost(ctx context.Context) ([]*models.Post, error) {
	return GetAllPostFn()
}

func (q *fakePostService) GetPostByID(context context.Context, id string, limit *int, offset *int) (*models.Post, error) {
	return GetPostByIDFn(id)
}
