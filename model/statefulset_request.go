package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqCreateStatefulSet struct {
	ClusterID       int    `json:"cluster_id" binding:"required"`
	NamespaceName   string `json:"namespace_name" binding:"required"`
	StatefulSetData string `json:"statefulset_data" binding:"required"`
}

type ReqUpdateStatefulSet struct {
	ClusterID       int    `json:"cluster_id" binding:"required"`
	NamespaceName   string `json:"namespace_name" binding:"required"`
	StatefulSetData string `json:"statefulset_data" binding:"required"`
}

type ReqReadStatefulSet struct {
	ClusterID       int    `form:"cluster_id" binding:"required"`
	NamespaceName   string `form:"namespace_name" binding:"required"`
	StatefulSetName string `form:"statefulset_name" binding:"required"`
}

type ReqDeleteStatefulSet struct {
	ClusterID       int    `json:"cluster_id" binding:"required"`
	NamespaceName   string `json:"namespace_name" binding:"required"`
	StatefulSetName string `json:"statefulset_name" binding:"required"`
}

type ReqListStatefulSet struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`

	PageInfo
}

type ReqReadStatefulSetHistory struct {
	ClusterID       int    `form:"cluster_id" binding:"required"`
	NamespaceName   string `form:"namespace_name" binding:"required"`
	StatefulSetName string `form:"statefulset_name" binding:"required"`
}
