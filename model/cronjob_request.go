package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqCreateCronJob struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	CronJobData   string `json:"cronjob_data" binding:"required"`
}

type ReqUpdateCronJob struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	CronJobData   string `json:"cronjob_data" binding:"required"`
}

type ReqReadCronJob struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
	CronJobName   string `form:"cronjob_name" binding:"required"`
}

type ReqDeleteCronJob struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	CronJobName   string `json:"cronjob_name" binding:"required"`
}

type ReqListCronJob struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`

	PageInfo
}
