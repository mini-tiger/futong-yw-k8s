package model

import "time"

type Cluster struct {
	ID             int    `json:"id" gorm:"column:id; primaryKey; comment:主键,集群ID"`
	TenantID       string `json:"tenant_id" gorm:"column:tenant_id; type:varchar(150); uniqueIndex:ui_tenant_id_cluster_account; uniqueIndex:ui_tenant_id_cluster_api; not null; comment:租户ID"`
	ClusterAccount string `json:"cluster_account" gorm:"column:cluster_account; type:varchar(150); uniqueIndex:ui_tenant_id_cluster_account; not null; comment:集群账号"`
	ClusterName    string `json:"cluster_name" gorm:"column:cluster_name; type:varchar(150); not null; comment:集群名称"`
	ClusterAPI     string `json:"cluster_api" gorm:"column:cluster_api; type:varchar(150); uniqueIndex:ui_tenant_id_cluster_api; not null; comment:集群API"`
	K8sConfig      string `json:"k8s_config" gorm:"column:k8s_config; type:longtext; not null; comment:集群秘钥"`
	Description    string `json:"description" gorm:"column:description; type:text; comment:集群介绍"`

	CreateTime *time.Time `json:"create_time" gorm:"column:create_time; autoCreateTime; comment:创建时间"`
	UpdateTime *time.Time `json:"update_time" gorm:"column:update_time; autoUpdateTime; comment:更新时间"`

	// Cluster and Tenant association
	ClusterTenant Tenant `gorm:"foreignKey:TenantID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// Group and Cluster association
	GroupCluster []*Group `gorm:"many2many:group_cluster_ass; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (Cluster) TableName() string {
	return "cluster"
}
