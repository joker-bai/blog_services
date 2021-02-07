package dao

import (
	"code.coolops.cn/blog_services/internal/model"
)

// 处理文章模块的Dao

func (d *Dao) ArticleList(title string, state uint8) ([]*model.Article, error) {
	article := &model.Article{Title: title, State: state}
	return article.List(d.engine)
}

func (d *Dao) ArticleCount(title string, state uint8) (int, error) {
	article := &model.Article{Title: title, State: state}
	return article.Count(d.engine)
}

func (d *Dao) ArticleCreate(title string, state uint8, createdBy string) error {
	article := model.Article{
		Model: &model.Model{CreatedBy: createdBy},
		Title: title,
		State: state,
	}
	return article.Create(d.engine)
}

func (d *Dao) ArticleUpdate(id uint32, state uint8, title, modifiedBy string) error {
	article := model.Article{Model: &model.Model{ID: id, ModifiedBy: modifiedBy}}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if title != "" {
		values["title"] = title
	}
	return article.Update(d.engine, values)
}

func (d *Dao) ArticleDelete(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}
