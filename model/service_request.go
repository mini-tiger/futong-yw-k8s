package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqCreateService struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	ServiceData   string `json:"service_data" binding:"required"`
}

type ReqUpdateService struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	ServiceData   string `json:"service_data" binding:"required"`
}

type ReqReadService struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
	ServiceName   string `form:"service_name" binding:"required"`
}

type ReqDeleteService struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	ServiceName   string `json:"service_name" binding:"required"`
}

type ReqListService struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
	ResourceKind  string `form:"resource_kind"`
	ResourceName  string `form:"resource_name"`

	PageInfo
}
