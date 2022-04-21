package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//查询分类是否存在
func CheckCategory(name string) (code int) {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATENAME_USED //2001
	}
	return errmsg.SUCCSE
}

//新增分类
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE //200
}

// 查询分类列表
func GetCategory(pageSize int, pageNum int) []Category {
	var categorys []Category
	var offset int
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	} else {
		offset = (pageNum - 1) * pageSize
	}
	err = db.Offset(offset).Limit(pageSize).Find(&categorys).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return categorys
}

// 查询分类下的所有文章

// 编辑分类
func EditCategory(id int, data *Category) int {
	var category Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&category).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//删除分类
func DeleteCategory(id int) int {
	var category Category
	err := db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
