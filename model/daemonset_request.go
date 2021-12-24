package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqCreateDaemonSet struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	DaemonSetData string `json:"daemonset_data" binding:"required"`
}

type ReqUpdateDaemonSet struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	DaemonSetData string `json:"daemonset_data" binding:"required"`
}

type ReqReadDaemonSet struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
	DaemonSetName string `form:"daemonset_name" binding:"required"`
}

type ReqDeleteDaemonSet struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	DaemonSetName string `json:"daemonset_name" binding:"required"`
}

type ReqListDaemonSet struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`

	PageInfo
}
