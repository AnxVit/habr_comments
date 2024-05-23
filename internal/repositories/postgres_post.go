package repositories

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/AnxVit/ozon_1/api/graphQL/models"
	"github.com/AnxVit/ozon_1/internal/repositories/schema"
	"github.com/AnxVit/ozon_1/internal/repositories/utiles"
	"github.com/AnxVit/ozon_1/internal/services"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ services.IHabrRepositories = HabrRepository{}

type HabrRepository struct {
	client *pgxpool.Pool
}

func NewHabrRepository(client *pgxpool.Pool) HabrRepository {
	return HabrRepository{client: client}
}

func (h HabrRepository) CreatePost(ctx context.Context, post *models.Post) (string, error) {
	id := struct {
		ID int `db:"id"`
	}{}

	if err := pgxscan.Get(ctx, h.client, &id, createPost, post.Content, post.Author, post.Title, post.Commented); err != nil {

		return "", errors.Join(ErrInternalError, err)
	}
	id_s := strconv.Itoa(id.ID)
	return id_s, nil
}
func (h HabrRepository) GetAllPost(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post

	if err := pgxscan.Select(ctx, h.client, &posts, getAllPosts); err != nil {

		return nil, errors.Join(ErrInternalError, err)
	}
	return posts, nil
}
func (h HabrRepository) GetPostByID(ctx context.Context, id string, limit *int, offset *int) (*models.Post, error) {
	post := new(models.Post)

	if err := pgxscan.Get(ctx, h.client, post, getPostByID, id); err != nil {
		return nil, errors.Join(ErrInternalError, err)
	}

	if !post.Commented {
		return post, nil
	}
	var comments_id []int
	if err := pgxscan.Select(ctx, h.client, &comments_id, getMainComments, id); err != nil {
		return nil, errors.Join(ErrInternalError, err)
	}

	var comments []*schema.Comment
	var build strings.Builder
	build.WriteString(getPostByID)
	if limit != nil {
		build.WriteString(" LIMIT " + strconv.Itoa(*limit))
	}
	if offset != nil {
		build.WriteString(" OFFSET " + strconv.Itoa(*offset))
	}
	build.WriteString(";")

	if err := pgxscan.Select(ctx, h.client, &comments, build.String(), comments_id); err != nil {
		return nil, errors.Join(ErrInternalError, err)
	}
	p, err := utiles.CombinePost(comments, post)
	return p, err
}

func (h HabrRepository) CreateComment(ctx context.Context, comment *models.Comment) error {
	_, err := h.client.Exec(ctx, createComment, comment.Content, comment.Author, comment.ParentCommentID)
	if err != nil {
		return errors.Join(ErrInternalError, err)
	}
	return nil
}

func (h HabrRepository) CreatePostComment(ctx context.Context, postID *string, comment *models.Comment) error {
	var commented bool
	if err := pgxscan.Get(ctx, h.client, &commented, checkPost, *postID); err != nil {
		return errors.Join(ErrInternalError, err)
	}

	if !commented {
		return ErrNotCommented
	}

	var comment_id int
	tx, err := h.client.Begin(ctx)
	if err != nil {
		return errors.Join(ErrInternalError, err)
	}
	defer tx.Rollback(ctx)
	err = tx.QueryRow(ctx, createPostComment, comment.Content, comment.Author).Scan(&comment_id)
	if err != nil {
		return errors.Join(ErrInternalError, err)
	}
	_, err = tx.Exec(ctx, postcomment, *postID, comment_id)
	if err != nil {
		return errors.Join(ErrInternalError, err)
	}

	if err = tx.Commit(ctx); err != nil {
		return errors.Join(ErrInternalError, err)
	}
	return nil
}
