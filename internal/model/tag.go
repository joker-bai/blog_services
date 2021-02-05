package model

import (
	"code.coolops.cn/blog_services/pkg/app"
	"code.coolops.cn/blog_services/pkg/errcode"
	"github.com/jinzhu/gorm"
)

// 标签
type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

// 统计Tag数量
func (t Tag) Count(db *gorm.DB) (int, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

// 列出多个Tag
func (t Tag) List(db *gorm.DB, pageOffSet, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffSet > 0 && pageSize > 0 {
		db = db.Offset(pageOffSet).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	// 判断数据库中是否已存在相同标签
	var old Tag
	res := db.Where("name = ?",t.Name).First(&old)
	if res.Error !=nil{
		if res.Error == gorm.ErrRecordNotFound{
			return db.Create(&t).Error
		}
	}

	return errcode.ErrorTagIsExistFail
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	//db = db.Model(t).Where("id = ? AND is_del = ?", t.ID, 0)
	if err := db.Model(t).Where("id = ? AND is_del = ?", t.ID, 0).Update(values).Error; err != nil {
		return err
	}
	//return db.Updates(t).Error
	return nil
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", t.ID, 0).Delete(&t).Error
}
