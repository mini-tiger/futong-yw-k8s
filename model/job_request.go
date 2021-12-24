package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqCreateJob struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	JobData       string `json:"job_data" binding:"required"`
}

type ReqUpdateJob struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	JobData       string `json:"job_data" binding:"required"`
}

type ReqReadJob struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
	JobName       string `form:"job_name" binding:"required"`
}

type ReqDeleteJob struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	JobName       string `json:"job_name" binding:"required"`
}

type ReqListJob struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
	CronJobName   string `form:"cronjob_name"`

	PageInfo
}
