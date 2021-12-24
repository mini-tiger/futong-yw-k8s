package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqDeploymentByUI struct {
}

type ReqCreateDeployment struct {
	ClusterID      int    `json:"cluster_id" binding:"required"`
	NamespaceName  string `json:"namespace_name" binding:"required"`
	DeploymentData string `json:"deployment_data" binding:"required"`
}

type ReqCreateDeploymentByUI struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`

	ReqDeploymentByUI ReqDeploymentByUI
}

type ReqUpdateDeployment struct {
	ClusterID      int    `json:"cluster_id" binding:"required"`
	NamespaceName  string `json:"namespace_name" binding:"required"`
	DeploymentData string `json:"deployment_data" binding:"required"`
}

type ReqUpdateDeploymentByUI struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`

	ReqDeploymentByUI ReqDeploymentByUI
}

type ReqReadDeployment struct {
	ClusterID      int    `form:"cluster_id" binding:"required"`
	NamespaceName  string `form:"namespace_name" binding:"required"`
	DeploymentName string `form:"deployment_name" binding:"required"`
}

type ReqDeleteDeployment struct {
	ClusterID      int    `json:"cluster_id" binding:"required"`
	NamespaceName  string `json:"namespace_name" binding:"required"`
	DeploymentName string `json:"deployment_name" binding:"required"`
}

type ReqListDeployment struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`

	PageInfo
}

type ReqReadDeploymentHistory struct {
	ClusterID      int    `form:"cluster_id" binding:"required"`
	NamespaceName  string `form:"namespace_name" binding:"required"`
	DeploymentName string `form:"deployment_name" binding:"required"`
}
