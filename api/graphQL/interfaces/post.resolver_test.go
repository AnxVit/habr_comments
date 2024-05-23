package interfaces_test

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/AnxVit/ozon_1/api/graphQL/generated"
	"github.com/AnxVit/ozon_1/api/graphQL/interfaces"
	"github.com/AnxVit/ozon_1/api/graphQL/models"
	"github.com/AnxVit/ozon_1/internal/transport/graphql/api"
	"github.com/stretchr/testify/assert"
)

type fakePostService struct{}

var fakePost api.IHabrService = &fakePostService{}

func TestCreatePost_Success(t *testing.T) {

	var srv = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		HabrService: fakePost,
	}})))

	//We dont call the domain method, we swap it with this
	CreatePostFn = func(question *models.Post) (string, error) {
		return "1", nil
	}

	var resp struct {
		CreatePost struct {
			Message string
			Status  int
			ID      string
		}
	}

	srv.MustPost(`mutation { CreatePost(post:{title:"d", content:"", author:""}) { message, status, id }}`, &resp)

	assert.Equal(t, 201, resp.CreatePost.Status)
	assert.Equal(t, "Successfully created post", resp.CreatePost.Message)
	assert.Equal(t, "1", resp.CreatePost.ID)
}

func TestCreateComment_Success(t *testing.T) {

	var srv = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		HabrService: fakePost,
	}})))

	//We dont call the domain method, we swap it with this
	CreateCommentFn = func(comment *models.Comment) error {
		return nil
	}

	var resp struct {
		CreateComment struct {
			Message string
			Status  int
		}
	}

	srv.MustPost(`mutation { CreateComment(postId:1, comment:{author:"", content:""}) { message, status }}`, &resp)

	assert.Equal(t, 201, resp.CreateComment.Status)
	assert.Equal(t, "Successfully created comment", resp.CreateComment.Message)
}

func TestGetAllPost_Success(t *testing.T) {

	var srv = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		HabrService: fakePost,
	}})))

	//We dont call the domain method, we swap it with this
	GetAllPostFn = func() ([]*models.Post, error) {
		return []*models.Post{
			{
				ID:        "1",
				Author:    "Tom",
				Title:     "Title_1",
				Content:   "Content_1",
				Commented: true,
			},
			{
				ID:        "2",
				Author:    "Kate",
				Title:     "Title_2",
				Content:   "Content_2",
				Commented: false,
			},
		}, nil
	}

	var resp struct {
		GetAllPost struct {
			Message  string
			Status   int
			DataList []*models.Post
		}
	}

	srv.MustPost(`query { GetAllPost {
				message,
				status,
				dataList {
							author
							title
							content
						}
				}
		}`, &resp)

	assert.Equal(t, 200, resp.GetAllPost.Status)
	assert.Equal(t, "Successfully retrived all posts", resp.GetAllPost.Message)
	assert.Equal(t, 2, len(resp.GetAllPost.DataList))
}

func TestGetPostByID_Success(t *testing.T) {

	var srv = client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		HabrService: fakePost,
	}})))

	//We dont call the domain method, we swap it with this
	GetPostByIDFn = func(string) (*models.Post, error) {
		return &models.Post{
			ID:        "1",
			Author:    "Tom",
			Title:     "Title_1",
			Content:   "Content_1",
			Commented: true,
		}, nil
	}

	var resp struct {
		GetPost struct {
			Message string
			Status  int
			Data    *models.Post
		}
	}

	srv.MustPost(`query { GetPost(id: "1") {
				message,
				status,
				data {
						author
						title
						content
					}
				}
		}`, &resp)

	assert.Equal(t, 200, resp.GetPost.Status)
	assert.Equal(t, "Successfully retrived post", resp.GetPost.Message)
	assert.Equal(t, "Tom", resp.GetPost.Data.Author)
}
