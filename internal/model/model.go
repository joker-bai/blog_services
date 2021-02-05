package model

import (
	"code.coolops.cn/blog_services/global"
	setting2 "code.coolops.cn/blog_services/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// 公共
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

// 初始化数据库
func NewDBEngine(databaseSetting *setting2.DatabaseSettingS) (*gorm.DB, error) {
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	db, err := gorm.Open(databaseSetting.DBType, s)
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	// 注册回调函数
	db.Callback().Create().Replace("gorm:update_time_stamp",updateTimeStampForUpdateCallback)
	db.Callback().Create().Replace("gorm:update_time_stamp",updateTimeStampForCreateCallback)
	db.Callback().Create().Replace("gorm:delete",deleteCallBack)

	db.DB().SetMaxIdleConns(global.DatabaseSetting.MaxIdleTime)
	db.DB().SetMaxOpenConns(global.DatabaseSetting.MaxOpenConn)
	return db, nil
}

// Model回调函数，实现对公共字段进行处理
// 新增行为的回调函数
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		// 获取当前是否包含CreatedOn字段
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			// 判断字段是否为空
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}
		// 获取当前是否包含ModifiedOn字段
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			// 判断字段是否为空
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

// 更新行为的回调函数
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	// 获取字段属性
	if _, ok := scope.Get("gorm:update_column"); !ok {
		// 如果不存在，则更新值
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// 删除行为的回调函数
func deleteCallBack(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		// 获取字段属性
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		deleteOnField, hasDeleteOnField := scope.FieldByName("DeletedOn")
		isDeField, hasIsDelField := scope.FieldByName("IsDel")
		// 判断是否有DeletedOn和IsDel字段，若存在，则进行软删除，否则进行硬删除
		if scope.Search.Unscoped && hasDeleteOnField && hasIsDelField {
			nowTime := time.Now().Unix()
			scope.Raw(fmt.Sprintf("UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deleteOnField.DBName),
				scope.AddToVars(nowTime),
				scope.Quote(isDeField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()

		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			))
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
