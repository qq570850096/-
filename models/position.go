package models

import (
	"fmt"
	"math"
	"sort"
	"time"

	"github.com/astaxie/beego/orm"
)

type Position struct {
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

type PositionCompanyBO struct {
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
	Id           int       `orm:"column(companyId);auto"`
	CompanyName  string    `orm:"column(companyName);size(100);null"`
	CompanyLogo  int       `orm:"column(companyLogo);null"`
	Description  string    `orm:"column(description);null"`
	State        int       `orm:"column(state);null"`
	CompanyCode  string    `orm:"column(companyCode);size(50);null"`
	Mobile       string    `orm:"column(mobile);size(11)"`
}

func (t *Position) TableName() string {
	return "position"
}

func init() {
	orm.RegisterModel(new(Position))
}

// AddPosition insert a new Position into database and returns
// last inserted Id on success.
func AddPosition(m *Position) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPositionById retrieves Position by Id. Returns error if
// Id doesn't exist
func GetPositionById(id int) (v *Position, err error) {
	o := orm.NewOrm()
	v = &Position{PosId: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPosition retrieves all Position matches certain condition. Returns empty list if
// no records exist
func GetAllPosition() (pos []Position, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Position))
	// query k=v
	_, err = qs.All(&pos)
	if err != nil {
		return nil, err
	}
	return
}

// UpdatePosition updates Position by Id and returns error if
// the record to be updated doesn't exist
func UpdatePositionById(m *Position) (err error) {
	o := orm.NewOrm()
	v := Position{PosId: m.PosId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePosition deletes Position by Id and returns error if
// the record to be deleted doesn't exist
func DeletePosition(id int) (err error) {
	o := orm.NewOrm()
	v := Position{PosId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Position{PosId: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func RecPosition(user User, page, limit int) (rec []PositionCompanyBO) {
	posList, _ := SynergyItemRec(user)
	if posList == nil {
		return nil
	}
	h := HelpMapList{}
	for _, v := range posList {
		temp := HelpMap{
			Key:   v.PosId,
			Value: v.Hits,
		}
		h = append(h, temp)
	}

	reclist, _ := Recommend(h, user)
	//if err != nil {
	//	return nil
	//}
	return reclist[limit*page-limit : limit*page]
}

func ListPosCompany(posid int) (pc PositionCompanyBO, err error) {
	o := orm.NewOrm()

	err = o.Raw("select p.*,c.* from position p,department d,company c "+
		"where p.departmentId = d.departmentId and d.companyId = c.companyId and p.positionId = ? limit 1", posid).QueryRow(&pc)
	if err != nil {
		return pc, err
	}
	return pc, nil
}

func ListCatePos(cateid int) (pbo []PositionCompanyBO, err error) {
	o := orm.NewOrm()
	if _, err = o.Raw("select p.*,c.* from position p,department d,company c\n"+
		"where p.departmentId = d.departmentId and d.companyId = c.companyId \n"+
		"and categoryId = ? and statePub = 1 \n"+
		"order by releaseDate DESC", cateid).QueryRows(&pbo); err == nil {
		return
	} else {
		fmt.Println(err)
		return
	}
}

func ListSearchPos(keyword, order string) (pbo []PositionCompanyBO, err error) {
	o := orm.NewOrm()
	if order == "salaryUp" {
		if _, err = o.Raw("select p.*,c.* from position p,department d,company c \n"+
			"where p.departmentId = d.departmentId and d.companyId = c.companyId \n"+
			"and title like ? and statePub = 1 \n"+
			"order by salaryUp DESC", keyword).QueryRows(&pbo); err == nil {
			return
		} else {
			fmt.Println(err)
			return
		}
	} else if order == "hits" {
		if _, err = o.Raw("select p.*,c.* from position p,department d,company c \n"+
			"where p.departmentId = d.departmentId and d.companyId = c.companyId \n"+
			"and title like ? and statePub = 1 \n"+
			"order by hits DESC", keyword).QueryRows(&pbo); err == nil {
			return
		} else {
			fmt.Println(err)
			return
		}
	} else {
		if _, err = o.Raw("select p.*,c.* from position p,department d,company c \n"+
			"where p.departmentId = d.departmentId and d.companyId = c.companyId \n"+
			"and title like ? and statePub = 1 \n"+
			"order by releaseDate DESC", keyword).QueryRows(&pbo); err == nil {
			return
		} else {
			fmt.Println(err)
			return
		}
	}
}

func HrPos(hrid int) (plist []Position, err error) {
	o := orm.NewOrm()
	if _, err = o.Raw("select * from position where hrIdPub = ? and "+
		"statePub = 1 order by releaseDate DESC", hrid).QueryRows(&plist); err == nil {
		return
	} else {
		fmt.Println(err)
		return
	}
}

// 推荐职位的协同过滤算法
func SynergyItemRec(user User) (PBO []PositionCompanyBO, err error) {
	posAll, err := GetAllPosition()
	//posAll = posAll[0:300]
	if err != nil {
		return nil, err
	}
	userAll, err := GetAllUser()
	userAll = userAll[15:]
	if err != nil {
		return nil, err
	}
	userid := user.Id
	var resumeid int
	resume, err := GetResumeByUId(userid)
	if resume != nil {
		resumeid = resume.Id
	}

	//定义键为职业Id值，值为每个用户对该职业好感度的键值对图HashMap
	favorMap := make(map[int][]int)

	//定义键为 职位Id，值为 当前用户对选定职位的好感度 的键值对图HashMap
	defaultmap := HelpMapList{}

	//用户总数
	userSize := len(userAll)

	//定义好感度评分数组
	itemList := make([]int, userSize)

	//定义单个职位Id，单个用户Id,单个简历Id
	var posItemId int
	var userItemId int
	var resumeItemId int

	//遍历所有职位，填充好感度键值对图
	for _, v := range posAll {
		posItemId = v.PosId

		for j, k := range userAll {
			userItemId = k.Id
			temp, _ := GetResumeByUId(userItemId)
			if temp != nil {
				resumeItemId = temp.Id
			}

			if _, err = GetApplicationByRePos(resumeItemId, posItemId); err == nil {
				itemList[j] = 3
			} else {
				itemList[j] = 0
			}

			if _, err = GetFavorByRePos(userItemId, posItemId); err == nil {
				itemList[j] += 3
			}
			com, errs := ListComment(userItemId, posItemId)
			if errs == nil {
				switch com.Type {
				case 1:
					itemList[j] += 1
				case 2:
					itemList[j] += 2
				case 3:
					itemList[j] += 3
				default:
					break
				}
			}

			if userItemId == userid {
				//键为职位Id 值为当前用户好感度 的键值对图TreeMap
				temp := HelpMap{
					Key:   posItemId,
					Value: itemList[j],
				}
				defaultmap = append(defaultmap, temp)
			}
		}
		favorMap[posItemId] = itemList
		itemList = make([]int, userSize)
	}

	//构造当前用户好感度降序排列的职位ArrayList
	sort.Sort(defaultmap)

	//当前用户好感度第一职位
	favorPosId := defaultmap[0].Key

	//好感度第一职位对应数组向量
	favorVector := favorMap[favorPosId]

	//与好感度第一职位的相似度键值对HashMap
	similarMap := HelpMapDoubleList{}

	//计数器
	var (
		m int
		n int
	)

	//favorVector,itemVector元素和,均值,模的平方,模，向量积
	var (
		favorSum      float64
		itemSum       float64
		favorAvg      float64
		itemAvg       float64
		favorMod      float64
		itemMod       float64
		favorModSqrt  float64
		itemModSqrt   float64
		vectorProduct float64
	)

	for i, _ := range favorVector {
		favorSum = favorSum + float64(favorVector[i])
	}
	favorAvg = favorSum / float64(len(favorVector))

	for i, _ := range favorVector {
		favorMod += (float64(favorVector[i]) - favorAvg) * (float64(favorVector[i]) - favorAvg)
	}
	favorModSqrt = math.Sqrt(favorMod)

	//求每一职位与当前用户好感度第一的职位的相似度

	for k, v := range favorMap {

		itemSum = 0
		itemAvg = 0
		itemMod = 0
		itemModSqrt = 0
		vectorProduct = 0

		for m = 0; m < len(v); m++ {
			itemSum += float64(v[m])
		}
		itemAvg = itemSum / float64(len(v))

		for n = 0; n < len(v); n++ {
			vectorProduct += (float64(favorVector[n]) - favorAvg) * (float64(v[n]) - itemAvg)
			itemMod += (float64(v[n]) - itemAvg) * (float64(v[n]) - itemAvg)
		}
		itemModSqrt = math.Sqrt(itemMod)

		//相似度键值对,键为职业Id,值为 选定职业与当前职业 皮尔逊相关系数
		temp := HelpMapDouble{
			Key:   k,
			Value: vectorProduct / (favorModSqrt * itemModSqrt),
		}
		similarMap = append(similarMap, temp)
	}

	//按与选中职业相似度皮尔逊系数降序排序 的元素为职位 的ArrayList
	similarPosList := make([]PositionCompanyBO, 0)
	sort.Sort(similarMap)

	for _, v := range similarMap {
		if _, errs := GetApplicationByRePos(resumeid, v.Key); errs != nil {
			pbo, err := ListPosCompany(v.Key)
			if err == nil {
				similarPosList = append(similarPosList, pbo)
			}

		}
	}

	//返回根据皮尔逊系数降序排列的 元素为职位的 职位列表
	return similarPosList, err
}

// map          ->  存在ServletContext中所有职位当日PV数
// user         ->  当前登录用户
// listPosAll   ->  所有职位 列表
// listPos      ->  当前用户应聘过的所有评论 列表
// listCom      ->  当前用户评论过的所有评论 列表
// listCol      ->  当前用户收藏过的所有评论 列表

func Recommend(h []HelpMap, user User) (pcbo []PositionCompanyBO, err error) {
	uid := user.Id
	resumeid := 0
	if v, err := GetResumeByUId(uid); err == nil {
		resumeid = v.Id
	}
	// 应聘记录
	listApp := GetAllApplicationByResu(resumeid)
	// 评论记录
	listCom, _ := ListCommentUid(uid)
	// 收藏记录
	listFav, _ := ListFavorUid(uid)
	var (
		//计算用户活跃度，判断采取哪种匹配方式
		activation float64
		// 活跃度标准
		actStandard float64 = 3.0
		// 应聘 评论 收藏得分
		pointPos = 1.0
		pointCom = 0.5
		pointCol = 3.0
	)
	//推荐Position列表
	posList := make([]PositionCompanyBO, 0)
	valuePos := float64(len(listApp)) * pointPos
	valueCom := float64(len(listCom)) * pointCom
	valueCol := float64(len(listFav)) * pointCol
	activation = valuePos + valueCom + valueCol
	//根据活跃度标准判断调用的推荐算法
	if activation < actStandard {
		posList = popularityRec(h, user)
	} else {
		posList, _ = SynergyItemRec(user)
		//if err != nil {
		//	return nil, err
		//}
	}
	return posList, nil
}

func popularityRec(h HelpMapList, user User) (pcbo []PositionCompanyBO) {
	//所有Position,所有用户
	posiAll, _ := GetAllPosition()

	//职位推荐程度值
	var matchDegree float64 = 0
	var cityRate float64 = 0.3
	var categoryRate float64 = 0.7
	var pvRate float64 = 1
	//城市、职位种类、pv 对应计算推荐程度值权重
	mapOrder := HelpMapDoubleList{}

	for _, v := range posiAll {
		//定义该职位当日pv数
		pv := 0
		//遍历寻找
		//赋值对应该职位title对应sessionContext内的当日pv数
		for _, j := range h {
			key := j.Key
			if key == v.PosId {
				pv = j.Value
			}
		}
		//判断计算推荐程度
		if v.WorkCity == user.City {
			if v.CategoryId == user.DirDesire {
				matchDegree = float64(pv+1) * pvRate
			} else {
				matchDegree = (float64(pv+1) * pvRate) * categoryRate
			}
		} else {
			if v.CategoryId == user.DirDesire {
				matchDegree = (float64(pv+1) * pvRate) * cityRate
			} else {
				matchDegree = ((float64(pv+1) * pvRate) * categoryRate) * cityRate
			}
		}
		temp := HelpMapDouble{
			Key:   v.PosId,
			Value: matchDegree,
		}
		mapOrder = append(mapOrder, temp)
	}
	//将mapOrder按照value值(pv数)降序排序
	sort.Sort(mapOrder)
	for _, v := range mapOrder {
		pos, err := ListPosCompany(v.Key)
		if err == nil {
			pcbo = append(pcbo, pos)
		} else {
			continue
		}
	}
	return pcbo
}
