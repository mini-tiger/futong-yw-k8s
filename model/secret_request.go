package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqCreateSecret struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	SecretData    string `json:"secret_data" binding:"required"`
}

type ReqUpdateSecret struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	SecretData    string `json:"secret_data" binding:"required"`
}

type ReqReadSecret struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
	SecretName    string `form:"secret_name" binding:"required"`
}

type ReqDeleteSecret struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	SecretName    string `json:"secret_name" binding:"required"`
}

type ReqListSecret struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`

	PageInfo
}
