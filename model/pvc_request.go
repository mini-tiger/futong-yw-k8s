package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqCreatePVC struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	PVCData       string `json:"pvc_data" binding:"required"`
}

type ReqUpdatePVC struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	PVCData       string `json:"pvc_data" binding:"required"`
}

type ReqReadPVC struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
	PVCName       string `form:"pvc_name" binding:"required"`
}

type ReqDeletePVC struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	PVCName       string `json:"pvc_name" binding:"required"`
}

type ReqListPVC struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`

	PageInfo
}
