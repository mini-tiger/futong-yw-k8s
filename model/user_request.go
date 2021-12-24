package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqRegisterTenant struct {
	TenantID    string `json:"tenant_id" binding:"required,gt=1"`
	TenantName  string `json:"tenant_name" binding:"required,gt=1"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Description string `json:"description"`
}

type ReqLogin struct {
	UserType    int    `json:"user_type" binding:"required,min=1,max=2"`
	TenantID    string `json:"tenant_id" binding:"required"`
	UserAccount string `json:"user_account" binding:"omitempty"`
	Password    string `json:"password" binding:"required"`

	ClientIP string `json:"client_ip"`
}

type ReqCreateUser struct {
	TenantID    string
	UserAccount string `json:"user_account" binding:"required,gt=1"`
	Username    string `json:"username" binding:"required,gt=1"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Description string `json:"description"`
}

type ReqUpdateUser struct {
	UserID      string `json:"user_id" binding:"required"`
	Username    string `json:"username" binding:"required,gt=1"`
	Email       string `json:"email" binding:"required,email"`
	Description string `json:"description"`
}

type ReqResetPassword struct {
	UserType int
	UserID   string `json:"user_id" binding:"required"`
	Salt     string
	Password string `json:"password" binding:"required"`
	Tenant   Tenant
}

type ReqReadUser struct {
	UserID string `form:"user_id" binding:"required"`
}

type ReqDeleteUser struct {
	TenantID string
	UserID   string `json:"user_id" binding:"required"`
}

type ReqListUser struct {
	TenantID    string
	UserAccount string `form:"user_account"`

	PageInfo
}
