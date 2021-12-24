package model

type ReqReadPod struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
	PodName       string `form:"pod_name" binding:"required"`
}

type ReqListPod struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`

	PageInfo
}

type ReqListPodByResource struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
	ResourceKind  string `form:"resource_kind" binding:"required"`
	ResourceName  string `form:"resource_name" binding:"required"`
}

type ReqReadPodLog struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
	PodName       string `form:"pod_name" binding:"required"`
	ContainerName string `form:"container_name" binding:"required"`
	LogLineNum    int64  `form:"log_line_num" binding:"required"`
}
