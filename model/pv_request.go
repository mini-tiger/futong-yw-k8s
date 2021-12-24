package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqCreatePV struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	PVData        string `json:"pv_data" binding:"required"`
}

type ReqUpdatePV struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	PVData        string `json:"pv_data" binding:"required"`
}

type ReqReadPV struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
	PVName        string `form:"pv_name" binding:"required"`
}

type ReqDeletePV struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	PVName        string `json:"pv_name" binding:"required"`
}

type ReqListPV struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`

	PageInfo
}
