package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqCreateGroup struct {
	TenantID     string
	GroupAccount string `json:"group_account" binding:"required,gt=1"`
	GroupName    string `json:"group_name" binding:"required,gt=1"`
	Description  string `json:"description"`
}

type ReqUpdateGroup struct {
	GroupID     string `json:"group_id" binding:"required"`
	GroupName   string `json:"group_name" binding:"required,gt=1"`
	Description string `json:"description"`
}

type ReqReadGroup struct {
	GroupID string `form:"group_id" binding:"required"`
}

type ReqDeleteGroup struct {
	TenantID     string
	GroupID      string `json:"group_id" binding:"required"`
	GroupAccount string
	// force delete: 1
	Force int `json:"force"`
}

type ReqListGroup struct {
	TenantID     string
	GroupAccount string `form:"group_account"`

	PageInfo
}
