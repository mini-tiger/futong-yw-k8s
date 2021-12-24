package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqUpdateRoleByUserGroup struct {
	TenantID         string
	UserID           string   `json:"user_id"`
	GroupID          string   `json:"group_id"`
	RoleAccountSlice []string `json:"role_account_slice" binding:"required"`
	RoleAssObj       string
}

type ReqReadRoleByUserGroup struct {
	TenantID   string
	UseID      string `form:"user_id"`
	GroupID    string `form:"group_id"`
	RoleAssObj string
}

type ReqUpdateUserGroupByRole struct {
	TenantID         string
	RoleAccount      string   `json:"role_account" binding:"required"`
	UserIDSlice      []string `json:"user_id_slice"`
	GroupIDSlice     []string `json:"group_id_slice"`
	RoleAssObjSlice  []string
	RoleAssObjPrefix string
}

type ReqReadUserGroupByRole struct {
	TenantID    string
	RoleAccount string `form:"role_account" binding:"required"`
}
