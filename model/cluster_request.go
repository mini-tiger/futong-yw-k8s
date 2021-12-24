package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqCheckConnectCluster struct {
	ClusterAPI string `json:"cluster_api" binding:"required,url"`
	K8sConfig  string `json:"k8s_config" binding:"required"`
}

type ReqImportCluster struct {
	TenantID       string
	ClusterAccount string `json:"cluster_account" binding:"required,gt=1"`
	ClusterName    string `json:"cluster_name" binding:"required,gt=1"`
	ClusterAPI     string `json:"cluster_api" binding:"required,url"`
	K8sConfig      string `json:"k8s_config" binding:"required"`
	Description    string `json:"description"`
}

type ReqUpdateCluster struct {
	ClusterID   int    `json:"cluster_id" binding:"required"`
	ClusterName string `json:"cluster_name" binding:"required,gt=1"`
	ClusterAPI  string `json:"cluster_api" binding:"required,url"`
	K8sConfig   string `json:"k8s_config" binding:"required"`
	Description string `json:"description"`
}

type ReqReadCluster struct {
	ClusterID int `form:"cluster_id" binding:"required"`
}

type ReqDeleteCluster struct {
	ClusterID int `json:"cluster_id" binding:"required"`
}

type ReqListCluster struct {
	UserType       int
	TenantID       string
	UserID         string
	ClusterAccount string `form:"cluster_account"`

	PageInfo
}
