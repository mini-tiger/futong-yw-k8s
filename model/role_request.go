package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqCreateRoleWithPermission struct {
	TenantID          string   `json:"tenant_id"`
	RoleAccount       string   `json:"role_account" binding:"required,gt=1"`
	RoleName          string   `json:"role_name" binding:"required,gt=1"`
	Description       string   `json:"description"`
	PermissionIDSlice []string `json:"permission_id_slice" binding:"required"`
	WebpermIDSlice    []string `json:"webperm_id_slice" binding:"required"`

	Role Role
}

type ReqUpdateRoleWithPermission struct {
	RoleID            int      `json:"role_id" binding:"required"`
	RoleName          string   `json:"role_name" binding:"required,gt=1"`
	Description       string   `json:"description"`
	PermissionIDSlice []string `json:"permission_id_slice" binding:"required"`
	WebpermIDSlice    []string `json:"webperm_id_slice" binding:"required"`

	Role Role
}

type ReqReadRole struct {
	RoleID int `form:"role_id" binding:"required"`
}

type ReqDeleteRole struct {
	TenantID    string
	RoleID      int `json:"role_id" binding:"required"`
	RoleAccount string
	// force delete: 1
	Force int `json:"force"`
}

type ReqListRole struct {
	TenantID    string
	RoleAccount string `form:"role_account"`

	PageInfo
}
