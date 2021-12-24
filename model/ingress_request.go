package model

// The content of this file:
// the structure of the request parameters associated with API

type ReqCreateIngress struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	IngressData   string `json:"ingress_data" binding:"required"`
}

type ReqUpdateIngress struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	IngressData   string `json:"ingress_data" binding:"required"`
}

type ReqReadIngress struct {
	ClusterID     int    `form:"cluster_id" binding:"required"`
	NamespaceName string `form:"namespace_name" binding:"required"`
	IngressName   string `form:"ingress_name" binding:"required"`
}

type ReqDeleteIngress struct {
	ClusterID     int    `json:"cluster_id" binding:"required"`
	NamespaceName string `json:"namespace_name" binding:"required"`
	IngressName   string `json:"ingress_name" binding:"required"`
}

type ReqListIngress struct {
	ClusterID       int      `json:"cluster_id" binding:"required"`
	NamespaceName   string   `json:"namespace_name" binding:"required"`
	ServiceNameList []string `json:"service_name_list"`

	// Which page to display
	PageNum int `json:"page_num" binding:"omitempty,gt=0"`
	// How much data to display per page
	PageSize int `json:"page_size" binding:"omitempty,gt=0"`
	// The amount of data to skip
	SkipNum int
	// According to which field to sort
	SortField string `json:"sort_field" binding:"omitempty"`
	// Order of sort: ascending:asc descending:desc
	SortOrder string `json:"sort_order" binding:"omitempty"`
	// Is descending
	SortOrderIsDesc bool
}
