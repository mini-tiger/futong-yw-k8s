package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqCreateNamespace struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceData string `json:"namespace_data" binding:"required"`
}

type ReqUpdateNamespace struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceData string `json:"namespace_data" binding:"required"`
}

type ReqReadNamespace struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
}

type ReqDeleteNamespace struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
}

type ReqListNamespace struct {
	ClusterID int `form:"cluster_id" binding:"required"`

	PageInfo
}
