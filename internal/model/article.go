package model

import "github.com/jinzhu/gorm"

// 文章
type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Count(db *gorm.DB) int {
	var count int

	return count
}

func (a Article) List(db *gorm.DB) ([]*Article, error) {
	var articles []*Article

	return articles, nil
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) Update(db *gorm.DB) error {
	return nil
}

func (a Article) Delete(db *gorm.DB) error {
	return nil
}
