package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"job/utils"
)

type Hr struct {
	Id           int    `orm:"column(hrId);auto"`
	HrMobile     string `orm:"column(hrMobile);size(11)"`
	HrPassword   string `orm:"column(hrPassword);size(500)"`
	HrName       string `orm:"column(hrName);size(50);null"`
	HrEmail      string `orm:"column(hrEmail);size(50);null"`
	Description  string `orm:"column(description);null"`
	DepartmentId int    `orm:"column(departmentId)"`
}

func (t *Hr) TableName() string {
	return "hr"
}

func init() {
	orm.RegisterModel(new(Hr))
}

// AddHr insert a new Hr into database and returns
// last inserted Id on success.
func AddHr(m *Hr) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetHrById retrieves Hr by Id. Returns error if
// Id doesn't exist
func GetHrByMobile(mobile string) (v *Hr, err error) {
	o := orm.NewOrm()
	v = &Hr{
		HrMobile: mobile,
	}
	if err = o.Read(v, "hrMobile"); err == nil {
		return v, nil
	}
	return nil, err
}
func GetHrById(id int) (v *Hr, err error) {
	o := orm.NewOrm()
	v = &Hr{
		Id: id,
	}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateHr updates Hr by Id and returns error if
// the record to be updated doesn't exist
func UpdateHrById(m *Hr) (err error) {
	o := orm.NewOrm()
	v := Hr{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func LoginHr(username, password string) (v *Hr, flag bool) {
	if v, _ := GetHrByMobile(username); v != nil {
		if utils.Md5Encode(password) == v.HrPassword {
			return v, true
		} else {
			return nil, false
		}
	} else {
		return nil, false
	}
}
