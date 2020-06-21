package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hc100/wp-nuxt-gql-go/backend/database"
	"github.com/hc100/wp-nuxt-gql-go/backend/graph/generated"
	"github.com/hc100/wp-nuxt-gql-go/backend/graph/model"
)

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	posts, err := database.NewPostDao(r.DB).FindAll()
	if err != nil {
		return nil, err
	}
	var results []*model.Post
	for _, post := range posts {
		results = append(results, &model.Post{
			ID:           post.ID,
			PostDate:     post.PostDate.Format("2006-01-02 15:04:05"),
			PostContent:  post.PostContent,
			PostTitle:    post.PostTitle,
			PostExcerpt:  post.PostExcerpt,
			PostModified: post.PostModified.Format("2006-01-02 15:04:05"),
		})
	}
	return results, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
