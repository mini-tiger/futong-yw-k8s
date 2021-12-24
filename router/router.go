package router

import (
	"net/http"

	"ftk8s/api"
	"ftk8s/middle"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	// Use i18n middleware
	r.Use(middle.HandLang())

	// Service health check
	r.GET("/check/ftk8s", func(c *gin.Context) {
		c.JSON(http.StatusOK, "The service(futong-yw-k8s) is OK!")
	})

	// ####################### project #########################
	// Project API routing prefix
	g := r.Group("/api")
	// Register a tenant
	g.POST("/register", api.RegisterTenant)
	// Login
	g.POST("/login", api.Login)

	// ####################### system ##########################
	// 查看当前用户树形菜单按钮
	g.GET("/system/webperm-role-ass", middle.Auth(), api.ListWebpermFunsAndRights)

	systemR := g.Group("/system")
	systemR.Use(middle.Auth(), middle.CheckPermOnlyTenant())
	{
		// Manage user
		systemR.POST("/user", api.CreateUser)
		systemR.PUT("/user", api.UpdateUser)
		systemR.PUT("/user/reset", api.ResetPassword)
		systemR.GET("/user", api.ReadUser)
		systemR.DELETE("/user", api.DeleteUser)
		systemR.GET("/users", api.ListUser)

		// Manage group
		systemR.POST("/group", api.CreateGroup)
		systemR.PUT("/group", api.UpdateGroup)
		systemR.GET("/group", api.ReadGroup)
		systemR.DELETE("/group", api.DeleteGroup)
		systemR.GET("/groups", api.ListGroup)

		// Manage association: user - group
		systemR.POST("/user-group", api.UpdateUserByGroup)
		systemR.GET("/user-group", api.ReadUserByGroup)
		systemR.POST("/group-user", api.UpdateGroupByUser)
		systemR.GET("/group-user", api.ReadGroupByUser)

		// Manage association: group - cluster
		systemR.POST("/group-cluster", api.UpdateGroupByCluster)
		systemR.GET("/group-cluster", api.ReadGroupByCluster)
		systemR.POST("/cluster-group", api.UpdateClusterByGroup)
		systemR.GET("/cluster-group", api.ReadClusterByGroup)

		// Manage role
		systemR.POST("/role", api.CreateRoleWithPermission)
		systemR.PUT("/role", api.UpdateRoleWithPermission)
		systemR.GET("/role", api.ReadRole)
		systemR.DELETE("/role", api.DeleteRole)
		systemR.GET("/roles", api.ListRole)

		// Manage association: role - user,group
		systemR.POST("/role-user-group", api.UpdateRoleByUserGroup)
		systemR.GET("/role-user-group", api.ReadRoleByUserGroup)
		systemR.POST("/user-group-role", api.UpdateUserGroupByRole)
		systemR.GET("/user-group-role", api.ReadUserGroupByRole)

		// Manage permission
		systemR.GET("/permissions", api.ListPermission)

		// Manage association: role - permission
		systemR.POST("/role-permission", api.UpdateRoleByPermission)
		systemR.GET("/role-permission", api.ReadRoleByPermission)
		systemR.POST("/permission-role", api.UpdatePermissionByRole)
		systemR.GET("/permission-role", api.ReadPermissionByRole)

		// Manage webperm
		systemR.POST("/webperm", middle.CheckPermOnlyBuiltinRoot(), api.CreateWebperm)
		systemR.PUT("/webperm", middle.CheckPermOnlyBuiltinRoot(), api.UpdateWebperm)
		systemR.DELETE("/webperm", middle.CheckPermOnlyBuiltinRoot(), api.DeleteWebperm)
		systemR.GET("/webperms", middle.CheckPermOnlyBuiltinRoot(), api.ListWebpermTree)
		systemR.GET("/webperms-bindrole", api.ListWebpermTreeBindRole)

		// Manage association: role - webperm
		systemR.POST("/webpermtree-role", api.UpdateWebpermTreeByRole)
		systemR.GET("/webpermtree-role", api.ReadWebpermTreeByRole)
	}

	// ####################### resource ##########################
	resTenantR := g.Group("/resource")
	resTenantR.Use(middle.Auth(), middle.CheckPermOnlyTenant())
	{
		// Manage cluster only tenant
		resTenantR.POST("/cluster/check", api.CheckConnectCluster)
		resTenantR.POST("/cluster", api.ImportCluster)
		resTenantR.PUT("/cluster", api.UpdateCluster)
		resTenantR.GET("/cluster", api.ReadCluster)
		resTenantR.DELETE("/cluster", api.DeleteCluster)
	}

	resAuthR := g.Group("/resource")
	resAuthR.Use(middle.Auth())
	{
		// List cluster
		resAuthR.GET("/clusters", api.ListCluster)

		// Manage template
		resAuthR.POST("/template", api.CreateTemplate)
		resAuthR.PUT("/template", api.UpdateTemplate)
		resAuthR.GET("/template", api.ReadTemplate)
		resAuthR.DELETE("/template", api.DeleteTemplate)
		resAuthR.GET("/templates", api.ListTemplate)
	}

	resR := g.Group("/resource")
	resR.Use(middle.Auth(), middle.CheckPermAccessCluster(), middle.CheckPermOperateResource())
	{
		// Deploy template
		resR.POST("/deploy-template", api.DeployTemplate)

		// Manage namespace
		resR.POST("/namespace", api.CreateNamespace)
		resR.PUT("/namespace", api.UpdateNamespace)
		resR.GET("/namespace", api.ReadNamespace)
		resR.DELETE("/namespace", api.DeleteNamespace)
		resR.GET("/namespaces", api.ListNamespace)

		// Manage event
		resR.GET("/events", api.ListEvent)
		resR.GET("/events-resource", api.ListEventByResource)

		// Manage pod
		resR.GET("/pod", api.ReadPod)
		resR.GET("/pods", api.ListPod)
		resR.GET("/pods-resource", api.ListPodByResource)
		resR.GET("/pod-log", api.ReadPodLog)

		// Manage deployment
		resR.POST("/deployment", api.CreateDeployment)
		// resR.POST("/deployment-ui", api.CreateDeploymentByUI)
		resR.PUT("/deployment", api.UpdateDeployment)
		// resR.PUT("/deployment-ui", api.UpdateDeploymentByUI)
		resR.GET("/deployment", api.ReadDeployment)
		resR.DELETE("/deployment", api.DeleteDeployment)
		resR.GET("/deployments", api.ListDeployment)
		resR.GET("/deployment-history", api.ReadDeploymentHistory)

		// Manage statefulset: appsv1
		resR.POST("/statefulset", api.CreateStatefulSet)
		resR.PUT("/statefulset", api.UpdateStatefulSet)
		resR.GET("/statefulset", api.ReadStatefulSet)
		resR.DELETE("/statefulset", api.DeleteStatefulSet)
		resR.GET("/statefulsets", api.ListStatefulSet)
		resR.GET("/statefulset-history", api.ReadStatefulSetHistory)

		// Manage daemonset: appsv1
		resR.POST("/daemonset", api.CreateDaemonSet)
		resR.PUT("/daemonset", api.UpdateDaemonSet)
		resR.GET("/daemonset", api.ReadDaemonSet)
		resR.DELETE("/daemonset", api.DeleteDaemonSet)
		resR.GET("/daemonsets", api.ListDaemonSet)

		// Manage job: batchv1
		resR.POST("/job", api.CreateJob)
		resR.PUT("/job", api.UpdateJob)
		resR.GET("/job", api.ReadJob)
		resR.DELETE("/job", api.DeleteJob)
		resR.GET("/jobs", api.ListJob)

		// Manage cronjob: batchv1beta1
		resR.POST("/cronjob", api.CreateCronJob)
		resR.PUT("/cronjob", api.UpdateCronJob)
		resR.GET("/cronjob", api.ReadCronJob)
		resR.DELETE("/cronjob", api.DeleteCronJob)
		resR.GET("/cronjobs", api.ListCronJob)

		// Manage configmap
		resR.POST("/configmap", api.CreateConfigMap)
		resR.PUT("/configmap", api.UpdateConfigMap)
		resR.GET("/configmap", api.ReadConfigMap)
		resR.DELETE("/configmap", api.DeleteConfigMap)
		resR.GET("/configmaps", api.ListConfigMap)

		// Manage secret
		resR.POST("/secret", api.CreateSecret)
		resR.PUT("/secret", api.UpdateSecret)
		resR.GET("/secret", api.ReadSecret)
		resR.DELETE("/secret", api.DeleteSecret)
		resR.GET("/secrets", api.ListSecret)

		// Manage service
		resR.POST("/service", api.CreateService)
		resR.PUT("/service", api.UpdateService)
		resR.GET("/service", api.ReadService)
		resR.DELETE("/service", api.DeleteService)
		resR.GET("/services", api.ListService)

		// Manage ingress: extensionsv1beta1
		resR.POST("/ingress", api.CreateIngress)
		resR.PUT("/ingress", api.UpdateIngress)
		resR.GET("/ingress", api.ReadIngress)
		resR.DELETE("/ingress", api.DeleteIngress)
		resR.POST("/ingresses", api.ListIngress)

		// Manage pvc(PersistentVolumeClaim): corev1
		resR.POST("/pvc", api.CreatePVC)
		resR.PUT("/pvc", api.UpdatePVC)
		resR.GET("/pvc", api.ReadPVC)
		resR.DELETE("/pvc", api.DeletePVC)
		resR.GET("/pvcs", api.ListPVC)

		// Manage pv(PersistentVolume): corev1
		resR.POST("/pv", api.CreatePV)
		resR.PUT("/pv", api.UpdatePV)
		resR.GET("/pv", api.ReadPV)
		resR.DELETE("/pv", api.DeletePV)
		resR.GET("/pvs", api.ListPV)

		// Manage storageclass: storagev1
		resR.POST("/storageclass", api.CreateStorageClass)
		resR.PUT("/storageclass", api.UpdateStorageClass)
		resR.GET("/storageclass", api.ReadStorageClass)
		resR.DELETE("/storageclass", api.DeleteStorageClass)
		resR.GET("/storageclasses", api.ListStorageClass)
	}

	return r
}
