package model

import "time"

type Permission struct {
	ID               string `json:"id" gorm:"column:id; type:varchar(150); primaryKey; comment:主键,权限ID"`
	PermissionName   string `json:"permission_name" gorm:"column:permission_name; type:varchar(191); uniqueIndex; not null; comment:权限名称"`
	PermissionUrl    string `json:"permission_url" gorm:"column:permission_url; type:varchar(191); uniqueIndex:ui_permission_url_permission_action; not null; comment:权限URL"`
	PermissionAction string `json:"permission_action" gorm:"column:permission_action; type:varchar(100); uniqueIndex:ui_permission_url_permission_action; not null; comment:权限动作"`
	Description      string `json:"description" gorm:"column:description; type:text; comment:权限介绍"`

	CreateTime *time.Time `json:"create_time" gorm:"column:create_time; autoCreateTime; comment:创建时间"`
	UpdateTime *time.Time `json:"update_time" gorm:"column:update_time; autoUpdateTime; comment:更新时间"`
}

func (Permission) TableName() string {
	return "permission"
}

type CasbinRule struct {
	PType string `json:"p_type" gorm:"column:p_type; type:varchar(100);"`
	V0    string `json:"v0" gorm:"column:v0; type:varchar(100);"`
	V1    string `json:"v1" gorm:"column:v1; type:varchar(100);"`
	V2    string `json:"v2" gorm:"column:v2; type:varchar(100);"`
	V3    string `json:"v3" gorm:"column:v3; type:varchar(100);"`
	V4    string `json:"v4" gorm:"column:v4; type:varchar(100);"`
	V5    string `json:"v5" gorm:"column:v5; type:varchar(100);"`
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}
