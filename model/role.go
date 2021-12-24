package model

import "time"

type Role struct {
	ID          int    `json:"id" gorm:"column:id; primaryKey; comment:主键,角色ID"`
	TenantID    string `json:"tenant_id" gorm:"column:tenant_id; type:varchar(150); uniqueIndex:ui_tenant_id_role_account; not null; comment:租户ID"`
	RoleAccount string `json:"role_account" gorm:"column:role_account; type:varchar(150); uniqueIndex:ui_tenant_id_role_account; not null; comment:角色账号"`
	RoleName    string `json:"role_name" gorm:"column:role_name; type:varchar(150); not null; comment:角色名称"`
	Description string `json:"description" gorm:"column:description; type:text; comment:角色介绍"`

	CreateTime *time.Time `json:"create_time" gorm:"column:create_time; autoCreateTime; comment:创建时间"`
	UpdateTime *time.Time `json:"update_time" gorm:"column:update_time; autoUpdateTime; comment:更新时间"`

	// Role and Tenant association
	RoleTenant Tenant `gorm:"foreignKey:TenantID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Role and Webperm association
	RoleWebperm []*Webperm `gorm:"many2many:role_webperm_ass; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Role) TableName() string {
	return "role"
}
