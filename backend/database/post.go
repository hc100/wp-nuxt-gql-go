package database

import (
	"context"
	"errors"
	"sort"
	"time"

	"github.com/hc100/wp-nuxt-gql-go/backend/graph/model"
	"github.com/hc100/wp-nuxt-gql-go/backend/util"
	"github.com/jinzhu/gorm"
)

type Post struct {
	ID           string    `gorm:"column:ID;primary_key"`
	PostDate     time.Time `gorm:"column:post_date"`
	PostContent  string    `gorm:"column:post_content"`
	PostTitle    string    `gorm:"column:post_title"`
	PostExcerpt  string    `gorm:"column:post_excerpt"`
	PostModified time.Time `gorm:"column:post_modified"`
}

func (u *Post) TableName() string {
	return "wp_posts"
}

type PostDao interface {
	FindAll() ([]*Post, error)
	FindOne(id string) (*Post, error)
	CountByTextFilter(ctx context.Context, filterWord *model.TextFilterCondition) (int, error)
	FindByCondition(ctx context.Context, filterWord *model.TextFilterCondition, pageCondition *model.PageCondition, edgeOrder *model.EdgeOrder) ([]*Post, error)
}

type postDao struct {
	db *gorm.DB
}

func NewPostDao(db *gorm.DB) PostDao {
	return &postDao{db: db}
}

func DefaultQuery(d *postDao) *gorm.DB {
	return d.db.Model(&Post{}).Where("post_type = 'post' AND post_status='publish' AND post_date < NOW()")
}

func (d *postDao) FindAll() ([]*Post, error) {
	var posts []*Post
	res := DefaultQuery(d).Order("post_date desc").Limit(3).Find(&posts)
	if err := res.Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (d *postDao) FindOne(id string) (*Post, error) {
	var posts []*Post
	res := d.db.Where("ID = ?", id).Find(&posts)
	if err := res.Error; err != nil {
		return nil, err
	}
	if len(posts) < 1 {
		return nil, nil
	}
	return posts[0], nil
}

func (d *postDao) CountByTextFilter(ctx context.Context, filterWord *model.TextFilterCondition) (int, error) {
	if filterWord == nil || filterWord.FilterWord == "" {
		var cnt int
		if err := DefaultQuery(d).Count(&cnt).Error; err != nil {
			return 0, err
		}
		return cnt, nil
	}

	matchStr := "%" + filterWord.FilterWord + "%"
	if filterWord.MatchingPattern != nil && *filterWord.MatchingPattern == model.MatchingPatternExactMatch {
		matchStr = filterWord.FilterWord
	}

	var cnt int

	res := DefaultQuery(d).
		Where("post_content LIKE ? OR post_title LIKE ?", matchStr, matchStr).
		Count(&cnt)
	if res.Error != nil {
		return 0, res.Error
	}

	return cnt, nil
}

func (d *postDao) FindByCondition(ctx context.Context, filterCondition *model.TextFilterCondition, pageCondition *model.PageCondition, edgeOrder *model.EdgeOrder) ([]*Post, error) {

	base := DefaultQuery(d)

	if filterCondition.ExistsFilter() {
		matchStr := filterCondition.MatchString()
		base = base.Where("post_content LIKE ? OR post_title LIKE ?", matchStr, matchStr)
	}

	if pageCondition.IsInitialPageView() {
		if pageCondition.HasInitialLimit() {
			if edgeOrder.ExistsOrder() {
				switch edgeOrder.Direction {
				case model.OrderDirectionAsc:
					base = base.Order(col_ASC(edgeOrder)).Limit(*pageCondition.InitialLimit)
				case model.OrderDirectionDesc:
					base = base.Order(col_DESC(edgeOrder)).Limit(*pageCondition.InitialLimit)
				}
			} else {
				base = base.Limit(*pageCondition.InitialLimit)
			}
		}
	}

	if pageCondition.ExistsPaging() && edgeOrder.ExistsOrder() {
		switch edgeOrder.Direction {
		case model.OrderDirectionAsc:
			if pageCondition.Forward != nil {
				target, err := d.getCompareTarget(pageCondition.Forward.After)
				if err != nil {
					return nil, err
				}
				targetValue := getTargetValueByOrderKey(*edgeOrder.Key.PostOrderKey, target)
				if targetValue == nil {
					return nil, errors.New("no target value")
				}
				base = base.Where("post_date > ?", targetValue).Order(col_ASC(edgeOrder)).Limit(pageCondition.Forward.First)
			}

			if pageCondition.Backward != nil {
				target, err := d.getCompareTarget(pageCondition.Backward.Before)
				if err != nil {
					return nil, err
				}
				targetValue := getTargetValueByOrderKey(*edgeOrder.Key.PostOrderKey, target)
				if targetValue == nil {
					return nil, errors.New("no target value")
				}
				base = base.Where("post_date < ?", targetValue).Order(col_DESC(edgeOrder)).Limit(pageCondition.Backward.Last)
			}
		case model.OrderDirectionDesc:
			if pageCondition.Forward != nil {
				target, err := d.getCompareTarget(pageCondition.Forward.After)
				if err != nil {
					return nil, err
				}
				targetValue := getTargetValueByOrderKey(*edgeOrder.Key.PostOrderKey, target)
				if targetValue == nil {
					return nil, errors.New("no target value")
				}
				base = base.Where("post_date < ?", targetValue).Order(col_DESC(edgeOrder)).Limit(pageCondition.Forward.First)
			}

			if pageCondition.Backward != nil {
				target, err := d.getCompareTarget(pageCondition.Backward.Before)
				if err != nil {
					return nil, err
				}
				targetValue := getTargetValueByOrderKey(*edgeOrder.Key.PostOrderKey, target)
				if targetValue == nil {
					return nil, errors.New("no target value")
				}
				base = base.Where("post_date > ?", targetValue).Order(col_ASC(edgeOrder)).Limit(pageCondition.Backward.Last)
			}
		}
	}

	var results []*Post
	if err := base.Find(&results).Error; err != nil {
		return nil, err
	}

	if edgeOrder.ExistsOrder() {
		reOrder(results, edgeOrder)
	}

	return results, nil
}

func reOrder(results []*Post, edgeOrder *model.EdgeOrder) {
	if results == nil {
		return
	}
	if len(results) == 0 {
		return
	}
	if edgeOrder.Key.PostOrderKey == nil {
		return
	}
	switch *edgeOrder.Key.PostOrderKey {
	case model.PostOrderKeyPostDate:
		if edgeOrder.Direction == model.OrderDirectionAsc {
			sort.Slice(results, func(i int, j int) bool {
				return results[i].PostDate.UnixNano() < results[j].PostDate.UnixNano()
			})
		}
		if edgeOrder.Direction == model.OrderDirectionDesc {
			sort.Slice(results, func(i int, j int) bool {
				return results[i].PostDate.UnixNano() > results[j].PostDate.UnixNano()
			})
		}
	}
}

func (d *postDao) getCompareTarget(cursor *string) (*Post, error) {
	if cursor == nil {
		return nil, errors.New("cursor is nil")
	}
	_, postID, err := util.DecodeCursor(*cursor)
	if err != nil {
		return nil, err
	}

	var target Post
	if err := d.db.Where(&Post{ID: postID}).First(&target).Error; err != nil {
		return nil, err
	}
	return &target, nil
}

func getTargetValueByOrderKey(postOrderKey model.PostOrderKey, post *Post) interface{} {
	switch postOrderKey {
	case model.PostOrderKeyPostDate:
		return post.PostDate
	default:
		return nil
	}
}

func col_ASC(o *model.EdgeOrder) string {
	return "post_date asc"
}

func col_DESC(o *model.EdgeOrder) string {
	return "post_date desc"
}
