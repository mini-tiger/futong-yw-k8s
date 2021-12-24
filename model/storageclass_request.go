package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqCreateStorageClass struct {
	ClusterID        int    `json:"cluster_id" binding:"required"`
	NamespaceName    string `json:"namespace_name" binding:"required"`
	StorageClassData string `json:"storageclass_data" binding:"required"`
}

type ReqUpdateStorageClass struct {
	ClusterID        int    `json:"cluster_id" binding:"required"`
	NamespaceName    string `json:"namespace_name" binding:"required"`
	StorageClassData string `json:"storageclass_data" binding:"required"`
}

type ReqReadStorageClass struct {
	ClusterID        int    `form:"cluster_id" binding:"required"`
	NamespaceName    string `form:"namespace_name" binding:"required"`
	StorageClassName string `form:"storageclass_name" binding:"required"`
}

type ReqDeleteStorageClass struct {
	ClusterID        int    `json:"cluster_id" binding:"required"`
	NamespaceName    string `json:"namespace_name" binding:"required"`
	StorageClassName string `json:"storageclass_name" binding:"required"`
}

type ReqListStorageClass struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`

	PageInfo
}
