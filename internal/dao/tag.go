package dao

import (
	"code.coolops.cn/blog_services/internal/model"
	"code.coolops.cn/blog_services/pkg/app"
)

// 处理标签模块的Dao操作
// 对数据访问对象进行封装

func (d *Dao) TagCount(name string, state uint8) (int, error) {
	tag := model.Tag{
		Model: nil,
		Name:  name,
		State: state,
	}
	return tag.Count(d.engine)
}

func (d *Dao) TagList(name string, state uint8, page, pageSize int) ([]*model.Tag, error) {
	tag := model.Tag{
		Model: nil,
		Name:  name,
		State: state,
	}
	pageOffSet := app.GetPageOffSet(page, pageSize)
	return tag.List(d.engine, pageOffSet, pageSize)
}

func (d *Dao) TagCreate(name string, state uint8, createBy string) error {
	tag := model.Tag{
		Model: &model.Model{
			CreatedBy: createBy,
		},
		Name:  name,
		State: state,
	}
	return tag.Create(d.engine)
}

func (d *Dao) TagUpdate(id uint32, state uint8, name, modifiedBy string) error {
	tag := model.Tag{
		Model: &model.Model{
			ID:         id,
			ModifiedBy: modifiedBy,
		},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}
	return tag.Update(d.engine, values)
}

func (d *Dao) TagDelete(id uint32) error {
	tag := model.Tag{
		Model: &model.Model{ID: id},
	}
	return tag.Delete(d.engine)
}
