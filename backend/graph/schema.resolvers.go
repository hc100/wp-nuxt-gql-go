package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/hc100/wp-nuxt-gql-go/backend/database"
	"github.com/hc100/wp-nuxt-gql-go/backend/graph/generated"
	"github.com/hc100/wp-nuxt-gql-go/backend/graph/model"
	"github.com/hc100/wp-nuxt-gql-go/backend/util"
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

func (r *queryResolver) PostConnection(ctx context.Context, filterWord *model.TextFilterCondition, pageCondition *model.PageCondition, edgeOrder *model.EdgeOrder) (*model.PostConnection, error) {
	log.Println("[queryResolver.PostConnection]")

	dao := database.NewPostDao(r.DB)

	totalCount, err := dao.CountByTextFilter(ctx, filterWord)
	if err != nil {
		return nil, err
	}
	if totalCount == 0 {
		return EmptyPostConnection(), nil
	}

	totalPage := pageCondition.TotalPage(totalCount)

	mtp := pageCondition.MoveToPageNo()
	mtpi64 := int(mtp)
	hnp := totalPage - mtpi64
	fmt.Println(hnp)

	pageInfo := &model.PageInfo{
		HasNextPage:     (totalPage - int(pageCondition.MoveToPageNo())) >= 1,
		HasPreviousPage: pageCondition.MoveToPageNo() > 1,
	}

	posts, err := dao.FindByCondition(ctx, filterWord, pageCondition, getOrder(edgeOrder))
	if err != nil {
		return nil, err
	}
	if posts == nil || len(posts) == 0 {
		return EmptyPostConnection(), nil
	}

	var edges []*model.PostEdge
	for idx, post := range posts {
		cursor := util.CreateCursor("post", post.ID)
		edges = append(edges, &model.PostEdge{
			Cursor: cursor,
			Node: &model.Post{
				ID:           post.ID,
				PostDate:     post.PostDate.Format("2006-01-02 15:04:05"),
				PostContent:  post.PostContent,
				PostTitle:    post.PostTitle,
				PostExcerpt:  post.PostExcerpt,
				PostModified: post.PostModified.Format("2006-01-02 15:04:05"),
			},
		})

		if idx == 0 {
			pageInfo.StartCursor = cursor
		}
		if idx == len(posts)-1 {
			pageInfo.EndCursor = cursor
		}
	}

	return &model.PostConnection{
		PageInfo:   pageInfo,
		Edges:      edges,
		TotalCount: totalCount,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func getOrder(edgeOrder *model.EdgeOrder) *model.EdgeOrder {
	var order *model.EdgeOrder
	if edgeOrder == nil {
		createdAt := model.PostOrderKeyPostDate
		order = &model.EdgeOrder{
			Key:       &model.OrderKey{PostOrderKey: &createdAt},
			Direction: model.OrderDirectionDesc,
		}
	} else {
		order = edgeOrder
	}
	return order
}
func EmptyPostConnection() *model.PostConnection {
	return &model.PostConnection{
		PageInfo:   &model.PageInfo{},
		Edges:      []*model.PostEdge{},
		TotalCount: 0,
	}
}
