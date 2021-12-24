package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqUpdateRoleByPermission struct {
	TenantID         string
	PermissionID     string   `json:"permission_id" binding:"required"`
	RoleAccountSlice []string `json:"role_account_slice" binding:"required"`
}

type ReqReadRoleByPermission struct {
	TenantID     string
	PermissionID string `form:"permission_id" binding:"required"`
}

type ReqUpdatePermissionByRole struct {
	TenantID          string
	RoleAccount       string   `json:"role_account" binding:"required"`
	PermissionIDSlice []string `json:"permission_id_slice" binding:"required"`
}

type ReqReadPermissionByRole struct {
	TenantID    string
	RoleAccount string `form:"role_account" binding:"required"`
}
