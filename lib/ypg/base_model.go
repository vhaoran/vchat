package ypg

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/vhaoran/vchat/common/ytime"
)

type BaseModel struct {
	CreatedAt ytime.Date `json:"created_at"`
	UpdatedAt ytime.Date `json:"update_at"`
	//DelTime     ytime.Date
	Ver int64 `json:"ver"`
}

// // 注册新建钩子在持久化之前
func createCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		//nowTime := time.Now().Unix()
		now := ytime.OfNow()
		if fd, ok := scope.FieldByName("CreatedAt"); ok {
			if fd.IsBlank {
				_ = fd.Set(now)
			}
		}

		if fd, ok := scope.FieldByName("UpdatedAt"); ok {
			if fd.IsBlank {
				_ = fd.Set(now)
			}
		}

		if fd, ok := scope.FieldByName("Ver"); ok {
			if fd.IsBlank {
				_ = fd.Set(1)
			}
		}

	}
}

// 注册更新钩子在持久化之前
func updateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("UpdatedAt", ytime.OfNow())
	}
}

// 注册删除钩子在删除之前
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnField, hasDeletedOnField := scope.FieldByName("DelTime")

		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
