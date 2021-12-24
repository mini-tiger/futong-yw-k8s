package enum

// Some keys to pass context information
const (
	TenantInfo = "Tenant-Info"
	UserInfo   = "User-Info"
	TenantID   = "Tenant-ID"
)

// 常用的Header
const (
	HeaderKeyJson   = "Content-Type"
	HeaderValueJson = "application/json"
)

// 通用的请求字段
const (
	ReqFieldClusterID = "cluster_id"
)

// 数据库查询正序倒序
const (
	SortOrderAsc  = "asc"
	SortOrderDesc = "desc"
)

// 支持的资源类型
var SupportResourceKindSlice = []string{
	"Namespace",
	"Deployment",
	"Pod",
	"StatefulSet",
	"DaemonSet",
	"Job",
	"CronJob",
	"ConfigMap",
	"Secret",
	"Service",
	"Ingress",
	"PersistentVolumeClaim",
	"PersistentVolume",
	"StorageClass",
}
