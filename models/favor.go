package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Favor struct {
	Id         int `orm:"column(favorId);auto"`
	UserId     int `orm:"column(userId)"`
	PositionId int `orm:"column(positionId)"`
}

type FavorPositionBO struct {
	Id           int       `orm:"column(favorId);auto"`
	UserId       int       `orm:"column(userId)"`
	PositionId   int       `orm:"column(positionId)"`
	PosId        int       `orm:"column(positionId);auto"`
	Title        string    `orm:"column(title);size(50);null"`
	Requirement  string    `orm:"column(requirement);null"`
	Quantity     int       `orm:"column(quantity);null"`
	WorkCity     string    `orm:"column(workCity);size(50);null"`
	SalaryUp     int       `orm:"column(salaryUp);null"`
	SalaryDown   int       `orm:"column(salaryDown);null"`
	ReleaseDate  time.Time `orm:"column(releaseDate);type(date);null"`
	ValidDate    time.Time `orm:"column(validDate);type(date);null"`
	StatePub     int       `orm:"column(statePub);null"`
	Hits         int       `orm:"column(hits);null"`
	CategoryId   int       `orm:"column(categoryId)"`
	DepartmentId int       `orm:"column(departmentId)"`
	HrIdPub      int       `orm:"column(hrIdPub)"`
}

func (t *Favor) TableName() string {
	return "favor"
}

func init() {
	orm.RegisterModel(new(Favor))
}

// AddFavor insert a new Favor into database and returns
// last inserted Id on success.
func AddFavor(m *Favor) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFavorById retrieves Favor by Id. Returns error if
// Id doesn't exist
func GetFavorById(id int) (v *Favor, err error) {
	o := orm.NewOrm()
	v = &Favor{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFavor retrieves all Favor matches certain condition. Returns empty list if
// no records exist
func GetAllFavor(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Favor))
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

	var l []Favor
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

// UpdateFavor updates Favor by Id and returns error if
// the record to be updated doesn't exist
func UpdateFavorById(m *Favor) (err error) {
	o := orm.NewOrm()
	v := Favor{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFavor deletes Favor by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFavor(uid, posid int) (err error) {
	o := orm.NewOrm()
	v := Favor{
		UserId:     uid,
		PositionId: posid,
	}
	// ascertain id exists in the database
	if err = o.Read(&v, "userId", "positionId"); err == nil {
		var num int64
		if num, err = o.Delete(&v); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetFavorByRePos(uid, posid int) (r *Favor, err error) {
	o := orm.NewOrm()
	if err = o.Raw("select * from favor where userId = ? and positionId = ? limit 1", uid, posid).QueryRow(&r); err == nil {
		return r, err
	}
	return r, err
}

func ListFavorUid(userid int) (listfav []Favor, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select * from favor where userId = ?", userid).QueryRows(&listfav)
	if err != nil {
		return nil, err
	}
	return listfav, nil
}

func ListFavPos(userid int) (listFP []FavorPositionBO, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select favorId,userId,position.* from favor,position where favor.positionId = position.positionId and userId = ?", userid).QueryRows(&listFP)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
