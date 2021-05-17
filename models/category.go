package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id           int    `orm:"column(categoryId);auto"`
	CategoryName string `orm:"column(categoryName);size(50)"`
	Description  string `orm:"column(description);null"`
}

func (t *Category) TableName() string {
	return "category"
}

func init() {
	orm.RegisterModel(new(Category))
}

// AddCategory insert a new Category into database and returns
// last inserted Id on success.
func AddCategory(m *Category) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCategoryById retrieves Category by Id. Returns error if
// Id doesn't exist
func GetCategoryById(id int) (v *Category, err error) {
	o := orm.NewOrm()
	v = &Category{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCategory retrieves all Category matches certain condition. Returns empty list if
// no records exist
func GetAllCategory() (c []Category, err error) {
	o := orm.NewOrm()
	if _, err = o.Raw("select * from category").QueryRows(&c); err == nil {
		return
	}
	fmt.Println(err)
	return
}

// UpdateCategory updates Category by Id and returns error if
// the record to be updated doesn't exist
func UpdateCategoryById(m *Category) (err error) {
	o := orm.NewOrm()
	v := Category{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCategory deletes Category by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCategory(id int) (err error) {
	o := orm.NewOrm()
	v := Category{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Category{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
