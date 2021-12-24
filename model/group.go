package model

import "time"

type Group struct {
	ID           string `json:"id" gorm:"column:id; type:varchar(150); primaryKey; comment:主键,用户组ID"`
	TenantID     string `json:"tenant_id" gorm:"column:tenant_id; type:varchar(150); uniqueIndex:ui_tenant_id_group_account; not null; comment:租户ID"`
	GroupAccount string `json:"group_account" gorm:"column:group_account; type:varchar(150); uniqueIndex:ui_tenant_id_group_account; not null; comment:用户组账号"`
	GroupName    string `json:"group_name" gorm:"column:group_name; type:varchar(150); not null; comment:用户组名称"`
	Description  string `json:"description" gorm:"column:description; type:text; comment:用户组介绍"`

	CreateTime *time.Time `json:"create_time" gorm:"column:create_time; autoCreateTime; comment:创建时间"`
	UpdateTime *time.Time `json:"update_time" gorm:"column:update_time; autoUpdateTime; comment:更新时间"`

	// Group and Tenant association
	GroupTenant Tenant `gorm:"foreignKey:TenantID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// User and Group association
	UserGroup []*User `gorm:"many2many:user_group_ass; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Group and Cluster association
	GroupCluster []*Cluster `gorm:"many2many:group_cluster_ass; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Group) TableName() string {
	return "group"
}

type UserGroupAss struct {
	UserID  string `json:"user_id" gorm:"column:user_id; type:varchar(150);"`
	GroupID string `json:"group_id" gorm:"column:group_id; type:varchar(150);"`
}

func (UserGroupAss) TableName() string {
	return "user_group_ass"
}

type GroupClusterAss struct {
	GroupID   string `json:"group_id" gorm:"column:group_id; type:varchar(150);"`
	ClusterID int    `json:"cluster_id" gorm:"column:cluster_id;"`
}

func (GroupClusterAss) TableName() string {
	return "group_cluster_ass"
}
