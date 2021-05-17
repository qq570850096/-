package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Application struct {
	Id         int       `orm:"column(applicationId);auto"`
	State      int       `orm:"column(state);null"`
	RecentTime time.Time `orm:"column(recentTime);type(datetime);null"`
	ResumeId   int       `orm:"column(resumeId)"`
	PositionId int       `orm:"column(positionId)"`
	HrId       int       `orm:"column(hrId);null"`
}

type ApplicationHrBO struct {
	AppId        int       `orm:"column(applicationId);auto"`
	State        int       `orm:"column(state);null"`
	RecentTime   time.Time `orm:"column(recentTime);type(datetime);null"`
	ResumeId     int       `orm:"column(resumeId)"`
	PositionId   int       `orm:"column(positionId)"`
	HrId         int       `orm:"column(hrId);null"`
	Id           int       `orm:"column(hrId);auto"`
	HrMobile     string    `orm:"column(hrMobile);size(11)"`
	HrPassword   string    `orm:"column(hrPassword);size(500)"`
	HrName       string    `orm:"column(hrName);size(50);null"`
	HrEmail      string    `orm:"column(hrEmail);size(50);null"`
	Description  string    `orm:"column(description);null"`
	DepartmentId int       `orm:"column(departmentId)"`
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
	HrIdPub      int       `orm:"column(hrIdPub)"`
}

type ExcelContent struct {
	AppId       int       `orm:"column(applicationId);auto"`
	State       int       `orm:"column(state);null"`
	RecentTime  time.Time `orm:"column(recentTime);type(datetime);null"`
	Title       string    `orm:"column(title);size(50);null"`
	WorkCity    string    `orm:"column(workCity);size(50);null"`
	SalaryUp    int       `orm:"column(salaryUp);null"`
	SalaryDown  int       `orm:"column(salaryDown);null"`
	Mobile      string    `orm:"column(mobile);size(11)"`
	CompanyName string    `orm:"column(companyName);size(100);null"`
	Name        string    `orm:"column(name);size(50);null"`
	Gender      int       `orm:"column(gender);null"`
	Email       string    `orm:"column(email);size(50);null"`
	EduDegree   string    `orm:"column(eduDegree);size(50);null"`
	Graduation  string    `orm:"column(graduation);size(100);null"`
	GraYear     int       `orm:"column(graYear);null"`
	Major       string    `orm:"column(major);size(50);null"`
}

func (t *Application) TableName() string {
	return "application"
}

func init() {
	orm.RegisterModel(new(Application))
}

// AddApplication insert a new Application into database and returns
// last inserted Id on success.
func AddApplication(m *Application) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func UpdateApplicationById(m *Application) (err error) {
	o := orm.NewOrm()
	v := Application{Id: m.Id}
	m.RecentTime = time.Now()
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// GetApplicationById retrieves Application by Id. Returns error if
// Id doesn't exist
func GetApplicationById(id int) (v *Application, err error) {
	o := orm.NewOrm()
	v = &Application{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllApplication retrieves all Application matches certain condition. Returns empty list if
// no records exist
func GetAllApplicationByResu(resumeid int) (applist []Application) {
	o := orm.NewOrm()
	_, err := o.Raw("select * from application where resumeId = ?", resumeid).QueryRows(&applist)
	if err != nil {
		return nil
	}
	return applist
}

func GetApplicationByRePos(reid, posid int) (r *Application, err error) {
	o := orm.NewOrm()
	if reid == 0 {
		if err = o.Raw("select * from application where positionId = ? limit 1", posid).QueryRow(&r); err == nil {
			return r, err
		}
	} else {
		if err = o.Raw("select * from application where resumeId = ? and positionId = ? limit 1", reid, posid).QueryRow(&r); err == nil {
			return r, err
		}
	}
	return r, err
}

func ListApplication(resid int) (applist []Application, err error) {
	o := orm.NewOrm()
	if _, err = o.Raw("select * from application where resumeId = ?", resid).QueryRows(&applist); err == nil {
		return
	}
	return
}

func SaveApplication(time2 time.Time, resumeid, positionid int) error {
	o := orm.NewOrm()
	app := &Application{
		RecentTime: time2,
		ResumeId:   resumeid,
		PositionId: positionid,
	}
	_, err := o.Insert(&app)
	return err
}

func ListAppPosHR(reid int) (applist []ApplicationHrBO, err error) {
	o := orm.NewOrm()
	if _, err = o.Raw("select a.applicationId,a.state,a.recentTime,a.resumeId,p.*,h.hrId,h.hrMobile,h.hrName,h.hrEmail "+
		"from application a,position p,hr h "+
		"where a.positionId = p.positionId and a.hrId = h.hrId "+
		"and a.hrId is not null and a.resumeId = ? "+
		"order by recentTime;", reid).QueryRows(&applist); err == nil {
		return
	}
	fmt.Println(err)
	return
}

func ListAppPosHRPub(reid int) (applist []ApplicationHrBO, err error) {
	o := orm.NewOrm()
	if _, err = o.Raw("select a.applicationId,a.state,a.recentTime,a.resumeId,p.*,h.hrId,h.hrMobile,h.hrName,h.hrEmail\n"+
		"from application a,position p,hr h\n"+
		"where a.positionId = p.positionId and p.hrIdPub = h.hrId \n"+
		"and a.hrId = 0 and a.resumeId = ?\n"+
		"order by p.releaseDate;", reid).QueryRows(&applist); err == nil {
		return
	}
	fmt.Println(err)
	return
}

func ListAppByHr(hrid int) (applist []ApplicationHrBO, err error) {
	o := orm.NewOrm()
	if _, err = o.Raw("select a.applicationId,a.state,a.recentTime,a.resumeId,p.*,h.hrId,h.hrMobile,h.hrName,h.hrEmail\n"+
		"from application a,position p,hr h\n"+
		"where a.positionId = p.positionId and a.hrId = h.hrId\n"+
		"and a.hrId = ?\n"+
		"order by recentTime", hrid).QueryRows(&applist); err != nil {
		return
	}
	fmt.Println(err)
	return
}

func ListAppByHrId(hrid int) (applist []ApplicationHrBO, err error) {
	o := orm.NewOrm()
	if _, err = o.Raw("select a.applicationId,a.state,a.recentTime,a.resumeId,p.*\n"+
		"from application a,position p,hr h\n"+
		"where a.positionId = p.positionId and a.hrId = h.hrId\n"+
		"and a.hrId = ? and a.state = 1\n"+
		"order by recentTime", hrid).QueryRows(&applist); err != nil {
		return
	}
	fmt.Println(err)
	return
}

func ListAppByState(state int) (excel []ExcelContent, err error) {
	o := orm.NewOrm()
	if _, err = o.Raw("select a.applicationId,a.state,a.recentTime,p.title,p.`salaryDown`,p.`salaryUp`,p.`workCity`,u.graduation,u.`graYear`,u.name,u.mobile,u.email, c.`companyName`,u.gender,u.major\n"+
		"from application a,position p,hr h, `user` u, resume r, company c,department d\n"+
		"where a.positionId = p.positionId and a.hrId = h.hrId and r.`resumeId` = a.`resumeId` and d.`departmentId` = p.`departmentId`\n"+
		"and a.state = ? and r.`userId` = u.`userId` and d.`companyId` = c.`companyId`\n"+
		"order by recentTime", state).QueryRows(&excel); err != nil {
		return
	}
	fmt.Println(err)
	return
}

func ListPositionByHr(hrid int) (pos []Position, err error) {
	o := orm.NewOrm()
	if _, err = o.Raw("select * "+
		"from position "+
		"where hrIdPub = ? and statePub = 1 order by releaseDate DESC", hrid).QueryRows(&pos); err != nil {
		return
	}
	fmt.Println(err)
	return
}
