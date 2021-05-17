package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"job/utils"
	"time"
)

type User struct {
	Id         int    `orm:"column(userId);auto"`
	Mobile     string `orm:"column(mobile);size(11)"`
	Password   string `orm:"column(password);size(500)"`
	Name       string `orm:"column(name);size(50);null"`
	Gender     int    `orm:"column(gender);null"`
	BirthYear  int    `orm:"column(birthYear);null"`
	Nickname   string `orm:"column(nickname);size(100);null"`
	Email      string `orm:"column(email);size(50);null"`
	Province   string `orm:"column(province);size(50);null"`
	City       string `orm:"column(city);size(50);null"`
	EduDegree  string `orm:"column(eduDegree);size(50);null"`
	Graduation string `orm:"column(graduation);size(100);null"`
	GraYear    int    `orm:"column(graYear);null"`
	Major      string `orm:"column(major);size(50);null"`
	DirDesire  int    `orm:"column(dirDesire);null"`
}

func GetUser(id int) (*User, error) {
	o := orm.NewOrm()
	t := &User{Id: id}
	if err := o.Read(t, "userId"); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return t, nil

}

func (m *User) UpdateUser() bool {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err := o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
			return true
		}
	}
	return false
}

func (m *User) RegisterUser() bool {
	mobile := m.Mobile
	if err := m.GetUserByMobile(mobile); err == nil {
		return false
	} else {
		password := m.Password
		md5pwd := utils.Md5Encode(password)
		user := &User{
			Mobile:   m.Mobile,
			Password: md5pwd,
			Nickname: m.Nickname,
		}
		_, err := AddUser(user)
		if err != nil {
			return false
		} else {
			return true
		}
	}
}

func (t *User) LoginUser(phone, password string) bool {
	if err := t.GetUserByMobile(phone); err != nil {
		return false
	} else {
		if utils.Md5Encode(password) == t.Password {
			return true
		} else {
			return false
		}
	}
}

func (t *User) GetUserByMobile(mobile string) (err error) {
	o := orm.NewOrm()
	t.Mobile = mobile
	if err = o.Read(t, "mobile"); err == nil {

		return nil
	}
	return err
}

type UserComment struct {
	Id          int       `orm:"column(userId);auto"`
	Mobile      string    `orm:"column(mobile);size(11)"`
	Password    string    `orm:"column(password);size(500)"`
	Name        string    `orm:"column(name);size(50);null"`
	Gender      int       `orm:"column(gender);null"`
	BirthYear   int       `orm:"column(birthYear);null"`
	Nickname    string    `orm:"column(nickname);size(100);null"`
	Email       string    `orm:"column(email);size(50);null"`
	Province    string    `orm:"column(province);size(50);null"`
	City        string    `orm:"column(city);size(50);null"`
	EduDegree   string    `orm:"column(eduDegree);size(50);null"`
	Graduation  string    `orm:"column(graduation);size(100);null"`
	GraYear     int       `orm:"column(graYear);null"`
	Major       string    `orm:"column(major);size(50);null"`
	DirDesire   int       `orm:"column(dirDesire);null"`
	ComId       int       `orm:"column(commentId);auto"`
	Type        int       `orm:"column(type);null"`
	Content     string    `orm:"column(content);null"`
	ReleaseTime time.Time `orm:"column(releaseTime);type(datetime);null"`
	UserId      int       `orm:"column(userId)"`
	PositionId  int       `orm:"column(positionId)"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
//func GetUserByMobile(Mobile string) (v *User, err error) {
//
//}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser() (user []User, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	_, err = qs.All(&user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func UserSize() (ret int) {
	o := orm.NewOrm()
	if err := o.Raw("select COUNT(*) from user").QueryRow(&ret); err != nil {
		fmt.Println(err)
		return
	}
	return
}
