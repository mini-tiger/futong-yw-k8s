package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqUpdateGroupByCluster struct {
	ClusterID    int      `json:"cluster_id" binding:"required"`
	GroupIDSlice []string `json:"group_id_slice" binding:"required"`
}

type ReqReadGroupByCluster struct {
	ClusterID int `form:"cluster_id" binding:"required"`
}

type ReqUpdateClusterByGroup struct {
	GroupID        string `json:"group_id" binding:"required"`
	ClusterIDSlice []int  `json:"cluster_id_slice" binding:"required"`
}

type ReqReadClusterByGroup struct {
	GroupID string `form:"group_id" binding:"required"`
}
