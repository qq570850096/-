package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Resume struct {
	Id             int    `orm:"column(resumeId);auto" `
	Ability        string `orm:"column(ability);null" `
	Internship     string `orm:"column(internship);null"`
	WorkExperience string `orm:"column(workExperience);null"`
	Certificate    string `orm:"column(certificate);null"`
	JobDesire      string `orm:"column(jobDesire);null"`
	UserId         int    `orm:"column(userId)"`
}

func (t *Resume) TableName() string {
	return "resume"
}

func init() {
	orm.RegisterModel(new(Resume))
}

// AddResume insert a new Resume into database and returns
// last inserted Id on success.
func AddResume(m *Resume) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetResumeById retrieves Resume by Id. Returns error if
// Id doesn't exist
func GetResumeByUId(uid int) (v *Resume, err error) {
	o := orm.NewOrm()
	v = &Resume{UserId: uid}
	if err = o.Read(v, "userId"); err == nil {
		return v, nil
	}
	return nil, err
}
func GetResumeById(id int) (v *Resume, err error) {
	o := orm.NewOrm()
	v = &Resume{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllResume retrieves all Resume matches certain condition. Returns empty list if
// no records exist
func GetAllResume(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Resume))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Resume
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateResume updates Resume by Id and returns error if
// the record to be updated doesn't exist
func UpdateResumeById(m *Resume) (err error) {
	o := orm.NewOrm()
	v := Resume{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		v.Ability = m.Ability
		v.Certificate = m.Certificate
		v.Internship = m.Internship
		v.WorkExperience = m.WorkExperience
		v.JobDesire = m.JobDesire
		if num, err = o.Update(&v); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteResume deletes Resume by Id and returns error if
// the record to be deleted doesn't exist
func DeleteResume(id int) (err error) {
	o := orm.NewOrm()
	v := Resume{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Resume{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
