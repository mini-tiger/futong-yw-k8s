package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqUpdateWebpermTreeByRole struct {
	TenantID       string
	RoleID         int      `json:"role_id" binding:"required"`
	WebpermIDSlice []string `json:"webperm_id_slice" binding:"required"`
}

type ReqReadWebpermTreeByRole struct {
	TenantID string
	RoleID   int `form:"role_id" binding:"required"`
}

type ReqListWebpermFunsAndRights struct {
	UserType int
	TenantID string
	Tenant   Tenant
	User     User
}
