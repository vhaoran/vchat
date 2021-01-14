package ypg

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/jinzhu/gorm"

	"github.com/vhaoran/vchat/common/reflectUtils"
	"github.com/vhaoran/vchat/common/ypage"
)

//cols只需要输入一个，用逗号分隔如： in,name,age
func FindMany(db *gorm.DB, l interface{}, where bson.M, cols ...string) error {
	if !reflectUtils.IsSlice(l) {
		return errors.New("passed l is not slice")
	}

	selectStr := ""
	if len(cols) > 0 {
		selectStr = cols[0]
	}

	//-------- -----------------------------
	exp, p := new(ypage.SqlWhereMap).GetWhere(where)

	var err error
	if len(selectStr) > 0 {
		err = db.
			Select(selectStr).
			Where(exp, p...).
			Find(l).Error
	} else {
		err = db.
			Where(exp, p...).
			Find(l).Error
	}

	if err != nil {
		return err
	}

	return nil
}
func FindOne(db *gorm.DB, retPtr interface{}, where bson.M, cols ...string) error {
	if !reflectUtils.IsPointer(retPtr) {
		return errors.New("passed retPtr is not pointer")
	}

	selectStr := ""
	if len(cols) > 0 {
		selectStr = cols[0]
	}

	//-------- -----------------------------
	exp, p := new(ypage.SqlWhereMap).GetWhere(where)

	var err error
	if len(selectStr) > 0 {
		err = db.
			Select(selectStr).
			Where(exp, p...).
			First(retPtr).Error
	} else {
		err = db.
			Where(exp, p...).
			First(retPtr).Error
	}

	if err != nil {
		return err
	}

	return nil
}
