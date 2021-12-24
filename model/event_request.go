package model

type ReqListEvent struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
}

type ReqListEventByResource struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
	ResourceKind  string `form:"resource_kind" binding:"required"`
	ResourceName  string `form:"resource_name" binding:"required"`
}
