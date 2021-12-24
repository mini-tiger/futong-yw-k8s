package model

import "time"

type Template struct {
	ID              int    `json:"id" gorm:"column:id; primaryKey; comment:主键,模板ID"`
	TenantID        string `json:"tenant_id" gorm:"column:tenant_id; type:varchar(150); uniqueIndex:ui_tenant_id_template_account; not null; comment:租户ID"`
	TemplateAccount string `json:"template_account" gorm:"column:template_account; type:varchar(150); uniqueIndex:ui_tenant_id_template_account; not null; comment:模板账号"`
	TemplateName    string `json:"template_name" gorm:"column:template_name; type:varchar(150); not null; comment:模板名称"`
	TemplateKind    string `json:"template_kind" gorm:"column:template_kind; type:varchar(100); not null; comment:模板类型"`
	Content         string `json:"content" gorm:"column:content; type:longtext; not null; comment:模板内容"`
	Description     string `json:"description" gorm:"column:description; type:text; comment:模板介绍"`

	CreateTime *time.Time `json:"create_time" gorm:"column:create_time; autoCreateTime; comment:创建时间"`
	UpdateTime *time.Time `json:"update_time" gorm:"column:update_time; autoUpdateTime; comment:更新时间"`

	// Template and Tenant association
	TemplateTenant Tenant `gorm:"foreignKey:TenantID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Template) TableName() string {
	return "template"
}
