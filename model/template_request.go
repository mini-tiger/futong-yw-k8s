package model

import "net/url"

// The content of this file:
// the structure of the request parameters associated with API

type ReqCreateTemplate struct {
	TenantID        string
	TemplateAccount string `json:"template_account" binding:"required,gt=1"`
	TemplateName    string `json:"template_name" binding:"required,gt=1"`
	TemplateKind    string `json:"template_kind" binding:"required"`
	Content         string `json:"content" binding:"required"`
	Description     string `json:"description"`
}

type ReqUpdateTemplate struct {
	TemplateID   int    `json:"template_id" binding:"required"`
	TemplateName string `json:"template_name" binding:"required,gt=1"`
	TemplateKind string `json:"template_kind" binding:"required"`
	Content      string `json:"content" binding:"required"`
	Description  string `json:"description"`
}

type ReqReadTemplate struct {
	TemplateID int `form:"template_id" binding:"required"`
}

type ReqDeleteTemplate struct {
	TemplateID int `json:"template_id" binding:"required"`
}

type ReqListTemplate struct {
	TenantID     string
	UrlQueryPara url.Values

	PageInfo
}

type ReqDeployTemplate struct {
	TenantID      string
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	TemplateID    int    `json:"template_id" binding:"required"`
	TemplateName  string `json:"template_name" binding:"required,gt=1"`
	TemplateKind  string `json:"template_kind" binding:"required"`
	Content       string `json:"content" binding:"required"`
}
