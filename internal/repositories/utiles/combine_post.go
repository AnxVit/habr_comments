package utiles

import (
	"strconv"

	"github.com/AnxVit/ozon_1/api/graphQL/models"
	"github.com/AnxVit/ozon_1/internal/repositories/schema"
)

func CombinePost(comments []*schema.Comment, post *models.Post) (*models.Post, error) {

	post.Comments = make([]*models.Comment, 0, len(comments))

	for _, com := range comments {
		comment := new(models.Comment)
		comment.Author = com.Author
		comment.Content = com.Content
		comment.ID = strconv.Itoa(com.ID)
		if len(com.Path) > 1 {
			parent_id := new(string)
			*parent_id = strconv.Itoa(com.Path[len(com.Path)-2])
			comment.ParentCommentID = parent_id
		}
		post.Comments = append(post.Comments, comment)
	}
	return post, nil
}
