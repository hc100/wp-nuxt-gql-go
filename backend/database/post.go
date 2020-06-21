package database

import (
	"time"

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
}

type postDao struct {
	db *gorm.DB
}

func NewPostDao(db *gorm.DB) PostDao {
	return &postDao{db: db}
}

func (d *postDao) FindAll() ([]*Post, error) {
	var posts []*Post
	res := d.db.Where("post_type = 'post' AND post_status='publish' AND post_date < NOW()").Order("post_date desc").Limit(3).Find(&posts)
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
