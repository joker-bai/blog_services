package model

import (
	"code.coolops.cn/blog_services/pkg/errcode"
	"github.com/jinzhu/gorm"
)

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

func (a Article) Count(db *gorm.DB) (int, error) {
	var count int
	var err error
	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", 0)
	if err = db.Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), err
}

func (a Article) List(db *gorm.DB) ([]*Article, error) {
	var articles []*Article

	if a.Title != "" {
		db = db.Where("title = ?", a.Title)
	}
	db = db.Where("state = ?", 0)
	if err := db.Where("is_del = ?", 0).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (a Article) Create(db *gorm.DB) error {
	var old Article
	res := db.Where("title = ? AND is_del = ?", a.Title, 0).First(&old)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return db.Create(&a).Error
		}
	}
	return errcode.ErrorTagIsExistFail
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(a).Where("id = ? AND is_del = ?", a.ID, 0).Update(values).Error; err != nil {
		return err
	}
	return nil
}

func (a Article) Delete(db *gorm.DB) error {
	if err := db.Where("id = ? AND is_del = ?", a.ID, 0).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}
