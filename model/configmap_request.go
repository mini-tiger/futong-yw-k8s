package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqCreateConfigMap struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	ConfigMapData string `json:"configmap_data" binding:"required"`
}

type ReqUpdateConfigMap struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	ConfigMapData string `json:"configmap_data" binding:"required"`
}

type ReqReadConfigMap struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
	ConfigMapName string `form:"configmap_name" binding:"required"`
}

type ReqDeleteConfigMap struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	ConfigMapName string `json:"configmap_name" binding:"required"`
}

type ReqListConfigMap struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`

	PageInfo
}
