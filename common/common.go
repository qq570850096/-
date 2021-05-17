package common

// 用户权限
const (
	// 超级管理员.
	MemberSuperRole = 0
	//普通管理员.
	MemberAdminRole = 1
	//普通用户.
	MemberGeneralRole = 2
	School_Email      = "570850096@qq.com"
)

func Role(role int) string {
	if role == MemberSuperRole {
		return "超级管理员"
	} else if role == MemberAdminRole {
		return "管理员"
	} else if role == MemberGeneralRole {
		return "普通用户"
	} else {
		return ""
	}
}

//图书关系
const (
	USER_RPATH = "user"
	HR_PATH    = "hr"
	ADMIN_PATH = "admin"

	LOGIN_USER_KEY = "user"
	Login_HR_key   = "hr"

	MAX_PAGE = 100

	MAX_ARRAY_SIZE = 100

	SESSION_KEY = "user"
)
