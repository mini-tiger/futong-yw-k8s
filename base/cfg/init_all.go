package cfg

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"hash"
	"log"
	"os"

	"ftk8s/model"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	AppConfObj   *AppConf
	LogConfObj   *LogConf
	MysqlConfObj *MysqlConf

	Gdb *gorm.DB
	Udb *sql.DB
)

// Project config info
type AppConf struct {
	AppRunMode string
	HttpIP     string
	HttpPort   string
	HttpAddr   string // HttpIp:HttpPort

	TypeUserService string

	TokenExpTime int

	RsaPrivateKey string
	RsaPublicKey  string

	// Initialize the administrator
	InitTenantID   string
	InitTenantName string
	InitPassword   string
	InitEmail      string
}

func InitAll() {
	// Load all config info
	loadSetting()

	// Set Gin run mode
	gin.SetMode(AppConfObj.AppRunMode)

	// Init the log object
	InitZapLogger(AppConfObj.AppRunMode, LogConfObj.LogLevel, LogConfObj.LogFile)

	// Init i18n config
	i18nPath := "base/i18n"
	i18nFilenameSlice := []string{"en_US.toml", "zh_CN.toml", "zh_TW.toml"}
	InitI18n(i18nPath, i18nFilenameSlice)

	// Make sure the database exists
	// ensureDatabase()
	// Get mysql client of *gorm.DB and *sql.DB
	Gdb, Udb = GetMysqlCliGormAndSql(MysqlConfObj, AppConfObj.AppRunMode)
	// Make sure the table exists
	// ensureTableAndInitData()

	// Init rsa key
	InitRsaKey()

	// Init casbin
	InitCasbin()
	// Init casbin rule
	// initCasbinRule()
}

// Load all config info
func loadSetting() {
	// Determine where to load config info
	flagSettingFromEnv := os.Getenv("FLAG_SETTING_FROM_ENV")
	if flagSettingFromEnv != "" {
		viper.AutomaticEnv()
		log.Println("load all config info by env")
	} else {
		// Specify file name (cannot write file extension)
		viper.SetConfigName("setting")
		// Specify the file search path
		viper.AddConfigPath("deploy/local")
		err := viper.ReadInConfig()
		if err != nil {
			log.Panicln("failed to read config file, error message: ", err.Error())
		}
		log.Println("load all config info by file")
	}

	// ######################### Project ###########################
	AppConfObj = new(AppConf)

	appRunMode := viper.GetString("APP_RUN_MODE")
	httpIP := viper.GetString("HTTP_IP")
	httpPort := viper.GetString("HTTP_PORT")
	typeUserService := viper.GetString("TYPE_USER_SERVICE")
	tokenExpTime := viper.GetInt("TOKEN_EXP_TIME")
	rsaPrivateKey := viper.GetString("RSA_PRIVATE_KEY")
	rsaPublicKey := viper.GetString("RSA_PUBLIC_KEY")
	initTenantID := viper.GetString("INIT_TENANT_ID")
	initTenantName := viper.GetString("INIT_TENANT_NAME")
	initPassword := viper.GetString("INIT_PASSWORD")
	initEmail := viper.GetString("INIT_EMAIL")

	AppConfObj.AppRunMode = appRunMode
	AppConfObj.HttpIP = httpIP
	AppConfObj.HttpPort = httpPort
	AppConfObj.HttpAddr = fmt.Sprintf("%s:%s", httpIP, httpPort)
	AppConfObj.TypeUserService = typeUserService
	AppConfObj.TokenExpTime = tokenExpTime
	AppConfObj.RsaPrivateKey = rsaPrivateKey
	AppConfObj.RsaPublicKey = rsaPublicKey
	AppConfObj.InitTenantID = initTenantID
	AppConfObj.InitTenantName = initTenantName
	AppConfObj.InitPassword = initPassword
	AppConfObj.InitEmail = initEmail

	// ########################### Log #############################
	LogConfObj = new(LogConf)

	logLevel := viper.GetString("LOG_LEVEL")
	logFile := viper.GetString("LOG_FILE")

	LogConfObj.LogLevel = logLevel
	LogConfObj.LogFile = logFile

	// ########################## MySQL #############################
	MysqlConfObj = new(MysqlConf)

	mysqlUsername := viper.GetString("MYSQL_USERNAME")
	mysqlPassword := viper.GetString("MYSQL_PASSWORD")
	mysqlHost := viper.GetString("MYSQL_HOST")
	mysqlPort := viper.GetInt("MYSQL_PORT")
	mysqlDatabase := viper.GetString("MYSQL_DATABASE")
	mysqlTimeout := viper.GetString("MYSQL_TIMEOUT")
	mysqlMaxIdleConns := viper.GetInt("MYSQL_MAX_IDLE_CONNS")
	mysqlMaxOpenConns := viper.GetInt("MYSQL_MAX_OPEN_CONNS")

	MysqlConfObj.MysqlUsername = mysqlUsername
	MysqlConfObj.MysqlPassword = mysqlPassword
	MysqlConfObj.MysqlHost = mysqlHost
	MysqlConfObj.MysqlPort = mysqlPort
	MysqlConfObj.MysqlDatabase = mysqlDatabase
	MysqlConfObj.MysqlTimeout = mysqlTimeout
	MysqlConfObj.MysqlMaxIdleConns = mysqlMaxIdleConns
	MysqlConfObj.MysqlMaxOpenConns = mysqlMaxOpenConns
}

// Make sure the database exists
func ensureDatabase() {
	dsnNoDBName := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=UTC&timeout=%s",
		MysqlConfObj.MysqlUsername,
		MysqlConfObj.MysqlPassword,
		MysqlConfObj.MysqlHost,
		MysqlConfObj.MysqlPort,
		MysqlConfObj.MysqlTimeout,
	)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC&timeout=%s",
		MysqlConfObj.MysqlUsername,
		MysqlConfObj.MysqlPassword,
		MysqlConfObj.MysqlHost,
		MysqlConfObj.MysqlPort,
		MysqlConfObj.MysqlDatabase,
		MysqlConfObj.MysqlTimeout,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		Mlog.Panic("failed to exec sql.Open, error message: ", err.Error())
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		switch e := err.(type) {
		case *mysql.MySQLError:
			// mysql error unknown database
			if e.Number == 1049 {
				Mlog.Infof("database (%s) does not exist", MysqlConfObj.MysqlDatabase)
				dbForCreateDatabase, err := sql.Open("mysql", dsnNoDBName)
				if err != nil {
					Mlog.Panic("failed to exec sql.Open, error message: ", err.Error())
					return
				}
				defer dbForCreateDatabase.Close()
				_, err = dbForCreateDatabase.Exec(
					fmt.Sprintf("create database %s default character set utf8mb4", MysqlConfObj.MysqlDatabase),
				)
				if err != nil {
					Mlog.Panic("failed to create database, error message: ", err.Error())
					return
				}
				Mlog.Infof("successfully to create database (%s)", MysqlConfObj.MysqlDatabase)
			} else {
				Mlog.Panic("failed to exec db.Ping, error message: ", err.Error())
				return
			}
		default:
			Mlog.Panic("failed to exec db.Ping, error message: ", err.Error())
			return
		}
	}
}

// Permission data to be initialized
var permissions = []model.Permission{
	// ####################### resource ##########################
	// template
	{ID: "DeployTemplate", PermissionName: "部署模板", PermissionUrl: "/api/resource/deploy-template", PermissionAction: "post", Description: "部署模板"},

	// namespace
	{ID: "CreateNamespace", PermissionName: "创建命名空间", PermissionUrl: "/api/resource/namespace", PermissionAction: "post", Description: "创建命名空间"},
	{ID: "UpdateNamespace", PermissionName: "更新命名空间", PermissionUrl: "/api/resource/namespace", PermissionAction: "put", Description: "更新命名空间"},
	{ID: "ReadNamespace", PermissionName: "获取命名空间", PermissionUrl: "/api/resource/namespace", PermissionAction: "get", Description: "获取命名空间"},
	{ID: "DeleteNamespace", PermissionName: "删除命名空间", PermissionUrl: "/api/resource/namespace", PermissionAction: "delete", Description: "删除命名空间"},
	{ID: "ListNamespace", PermissionName: "列表命名空间", PermissionUrl: "/api/resource/namespaces", PermissionAction: "get", Description: "列表命名空间"},

	// deployment
	{ID: "CreateDeployment", PermissionName: "创建无状态资源", PermissionUrl: "/api/resource/deployment", PermissionAction: "post", Description: "创建无状态资源"},
	{ID: "UpdateDeployment", PermissionName: "更新无状态资源", PermissionUrl: "/api/resource/deployment", PermissionAction: "put", Description: "更新无状态资源"},
	{ID: "ReadDeployment", PermissionName: "获取无状态资源", PermissionUrl: "/api/resource/deployment", PermissionAction: "get", Description: "获取无状态资源"},
	{ID: "DeleteDeployment", PermissionName: "删除无状态资源", PermissionUrl: "/api/resource/deployment", PermissionAction: "delete", Description: "删除无状态资源"},
	{ID: "ListDeployment", PermissionName: "列表无状态资源", PermissionUrl: "/api/resource/deployments", PermissionAction: "get", Description: "列表无状态资源"},

	// pod
	{ID: "ReadPod", PermissionName: "获取容器组资源", PermissionUrl: "/api/resource/pod", PermissionAction: "get", Description: "获取容器组资源"},
	{ID: "ListPod", PermissionName: "列表容器组资源", PermissionUrl: "/api/resource/pods", PermissionAction: "get", Description: "列表容器组资源"},
	{ID: "ReadPodLog", PermissionName: "读取容器组资源日志", PermissionUrl: "/api/resource/pod-log", PermissionAction: "get", Description: "读取容器组资源日志"},

	// statefulset
	{ID: "CreateStatefulSet", PermissionName: "创建有状态资源", PermissionUrl: "/api/resource/statefulset", PermissionAction: "post", Description: "创建有状态资源"},
	{ID: "UpdateStatefulSet", PermissionName: "更新有状态资源", PermissionUrl: "/api/resource/statefulset", PermissionAction: "put", Description: "更新有状态资源"},
	{ID: "ReadStatefulSet", PermissionName: "获取有状态资源", PermissionUrl: "/api/resource/statefulset", PermissionAction: "get", Description: "获取有状态资源"},
	{ID: "DeleteStatefulSet", PermissionName: "删除有状态资源", PermissionUrl: "/api/resource/statefulset", PermissionAction: "delete", Description: "删除有状态资源"},
	{ID: "ListStatefulSet", PermissionName: "列表有状态资源", PermissionUrl: "/api/resource/statefulsets", PermissionAction: "get", Description: "列表有状态资源"},

	// daemonset
	{ID: "CreateDaemonSet", PermissionName: "创建进程守护集资源", PermissionUrl: "/api/resource/daemonset", PermissionAction: "post", Description: "创建进程守护集资源"},
	{ID: "UpdateDaemonSet", PermissionName: "更新进程守护集资源", PermissionUrl: "/api/resource/daemonset", PermissionAction: "put", Description: "更新进程守护集资源"},
	{ID: "ReadDaemonSet", PermissionName: "获取进程守护集资源", PermissionUrl: "/api/resource/daemonset", PermissionAction: "get", Description: "获取进程守护集资源"},
	{ID: "DeleteDaemonSet", PermissionName: "删除进程守护集资源", PermissionUrl: "/api/resource/daemonset", PermissionAction: "delete", Description: "删除进程守护集资源"},
	{ID: "ListDaemonSet", PermissionName: "列表进程守护集资源", PermissionUrl: "/api/resource/daemonsets", PermissionAction: "get", Description: "列表进程守护集资源"},

	// job
	{ID: "CreateJob", PermissionName: "创建任务资源", PermissionUrl: "/api/resource/job", PermissionAction: "post", Description: "创建任务资源"},
	{ID: "UpdateJob", PermissionName: "更新任务资源", PermissionUrl: "/api/resource/job", PermissionAction: "put", Description: "更新任务资源"},
	{ID: "ReadJob", PermissionName: "获取任务资源", PermissionUrl: "/api/resource/job", PermissionAction: "get", Description: "获取任务资源"},
	{ID: "DeleteJob", PermissionName: "删除任务资源", PermissionUrl: "/api/resource/job", PermissionAction: "delete", Description: "删除任务资源"},
	{ID: "ListJob", PermissionName: "列表任务资源", PermissionUrl: "/api/resource/jobs", PermissionAction: "get", Description: "列表任务资源"},

	// cronjob
	{ID: "CreateCronJob", PermissionName: "创建定时任务资源", PermissionUrl: "/api/resource/cronjob", PermissionAction: "post", Description: "创建定时任务资源"},
	{ID: "UpdateCronJob", PermissionName: "更新定时任务资源", PermissionUrl: "/api/resource/cronjob", PermissionAction: "put", Description: "更新定时任务资源"},
	{ID: "ReadCronJob", PermissionName: "获取定时任务资源", PermissionUrl: "/api/resource/cronjob", PermissionAction: "get", Description: "获取定时任务资源"},
	{ID: "DeleteCronJob", PermissionName: "删除定时任务资源", PermissionUrl: "/api/resource/cronjob", PermissionAction: "delete", Description: "删除定时任务资源"},
	{ID: "ListCronJob", PermissionName: "列表定时任务资源", PermissionUrl: "/api/resource/cronjobs", PermissionAction: "get", Description: "列表定时任务资源"},

	// configmap
	{ID: "CreateConfigMap", PermissionName: "创建配置资源", PermissionUrl: "/api/resource/configmap", PermissionAction: "post", Description: "创建配置资源"},
	{ID: "UpdateConfigMap", PermissionName: "更新配置资源", PermissionUrl: "/api/resource/configmap", PermissionAction: "put", Description: "更新配置资源"},
	{ID: "ReadConfigMap", PermissionName: "获取配置资源", PermissionUrl: "/api/resource/configmap", PermissionAction: "get", Description: "获取配置资源"},
	{ID: "DeleteConfigMap", PermissionName: "删除配置资源", PermissionUrl: "/api/resource/configmap", PermissionAction: "delete", Description: "删除配置资源"},
	{ID: "ListConfigMap", PermissionName: "列表配置资源", PermissionUrl: "/api/resource/configmaps", PermissionAction: "get", Description: "列表配置资源"},

	// secret
	{ID: "CreateSecret", PermissionName: "创建保密字典资源", PermissionUrl: "/api/resource/secret", PermissionAction: "post", Description: "创建保密字典资源"},
	{ID: "UpdateSecret", PermissionName: "更新保密字典资源", PermissionUrl: "/api/resource/secret", PermissionAction: "put", Description: "更新保密字典资源"},
	{ID: "ReadSecret", PermissionName: "获取保密字典资源", PermissionUrl: "/api/resource/secret", PermissionAction: "get", Description: "获取保密字典资源"},
	{ID: "DeleteSecret", PermissionName: "删除保密字典资源", PermissionUrl: "/api/resource/secret", PermissionAction: "delete", Description: "删除保密字典资源"},
	{ID: "ListSecret", PermissionName: "列表保密字典资源", PermissionUrl: "/api/resource/secrets", PermissionAction: "get", Description: "列表保密字典资源"},

	// service
	{ID: "CreateService", PermissionName: "创建服务资源", PermissionUrl: "/api/resource/service", PermissionAction: "post", Description: "创建服务资源"},
	{ID: "UpdateService", PermissionName: "更新服务资源", PermissionUrl: "/api/resource/service", PermissionAction: "put", Description: "更新服务资源"},
	{ID: "ReadService", PermissionName: "获取服务资源", PermissionUrl: "/api/resource/service", PermissionAction: "get", Description: "获取服务资源"},
	{ID: "DeleteService", PermissionName: "删除服务资源", PermissionUrl: "/api/resource/service", PermissionAction: "delete", Description: "删除服务资源"},
	{ID: "ListService", PermissionName: "列表服务资源", PermissionUrl: "/api/resource/services", PermissionAction: "get", Description: "列表服务资源"},

	// ingress
	{ID: "CreateIngress", PermissionName: "创建路由资源", PermissionUrl: "/api/resource/ingress", PermissionAction: "post", Description: "创建路由资源"},
	{ID: "UpdateIngress", PermissionName: "更新路由资源", PermissionUrl: "/api/resource/ingress", PermissionAction: "put", Description: "更新路由资源"},
	{ID: "ReadIngress", PermissionName: "获取路由资源", PermissionUrl: "/api/resource/ingress", PermissionAction: "get", Description: "获取路由资源"},
	{ID: "DeleteIngress", PermissionName: "删除路由资源", PermissionUrl: "/api/resource/ingress", PermissionAction: "delete", Description: "删除路由资源"},
	{ID: "ListIngress", PermissionName: "列表路由资源", PermissionUrl: "/api/resource/ingresses", PermissionAction: "get", Description: "列表路由资源"},

	// pvc
	{ID: "CreatePVC", PermissionName: "创建存储声明资源", PermissionUrl: "/api/resource/pvc", PermissionAction: "post", Description: "创建存储声明资源"},
	{ID: "UpdatePVC", PermissionName: "更新存储声明资源", PermissionUrl: "/api/resource/pvc", PermissionAction: "put", Description: "更新存储声明资源"},
	{ID: "ReadPVC", PermissionName: "获取存储声明资源", PermissionUrl: "/api/resource/pvc", PermissionAction: "get", Description: "获取存储声明资源"},
	{ID: "DeletePVC", PermissionName: "删除存储声明资源", PermissionUrl: "/api/resource/pvc", PermissionAction: "delete", Description: "删除存储声明资源"},
	{ID: "ListPVC", PermissionName: "列表存储声明资源", PermissionUrl: "/api/resource/pvcs", PermissionAction: "get", Description: "列表存储声明资源"},

	// pv
	{ID: "CreatePV", PermissionName: "创建存储卷资源", PermissionUrl: "/api/resource/pv", PermissionAction: "post", Description: "创建存储卷资源"},
	{ID: "UpdatePV", PermissionName: "更新存储卷资源", PermissionUrl: "/api/resource/pv", PermissionAction: "put", Description: "更新存储卷资源"},
	{ID: "ReadPV", PermissionName: "获取存储卷资源", PermissionUrl: "/api/resource/pv", PermissionAction: "get", Description: "获取存储卷资源"},
	{ID: "DeletePV", PermissionName: "删除存储卷资源", PermissionUrl: "/api/resource/pv", PermissionAction: "delete", Description: "删除存储卷资源"},
	{ID: "ListPV", PermissionName: "列表存储卷资源", PermissionUrl: "/api/resource/pvs", PermissionAction: "get", Description: "列表存储卷资源"},

	// storageclass
	{ID: "CreateStorageClass", PermissionName: "创建存储类资源", PermissionUrl: "/api/resource/storageclass", PermissionAction: "post", Description: "创建存储类资源"},
	{ID: "UpdateStorageClass", PermissionName: "更新存储类资源", PermissionUrl: "/api/resource/storageclass", PermissionAction: "put", Description: "更新存储类资源"},
	{ID: "ReadStorageClass", PermissionName: "获取存储类资源", PermissionUrl: "/api/resource/storageclass", PermissionAction: "get", Description: "获取存储类资源"},
	{ID: "DeleteStorageClass", PermissionName: "删除存储类资源", PermissionUrl: "/api/resource/storageclass", PermissionAction: "delete", Description: "删除存储类资源"},
	{ID: "ListStorageClass", PermissionName: "列表存储类资源", PermissionUrl: "/api/resource/storageclasses", PermissionAction: "get", Description: "列表存储类资源"},
}

// Webperm data to be initialized
var webperms = []model.Webperm{
	// 集群
	{ID: "1", ParentID: "0", Name: "k8s_cluster", Path: "/k8s_cluster", ResourcesSort: 1, ResourcesType: "M", Title: "集群", Icon: "", Display: 1, OnlyBuiltinRoot: 2},
	// 集群-集群列表
	{ID: "1_1", ParentID: "1", Name: "k8s_clusterList", Path: "k8s_clusterList", ResourcesSort: 1, ResourcesType: "C", Title: "集群列表", Icon: "", Display: 1, OnlyBuiltinRoot: 2},

	// 系统设置
	{ID: "3", ParentID: "0", Name: "k8s_systemSetting", Path: "/k8s_systemSetting", ResourcesSort: 3, ResourcesType: "M", Title: "系统设置", Icon: "", Display: 1, OnlyBuiltinRoot: 2},
	// 系统设置-用户管理
	{ID: "3_1", ParentID: "3", Name: "k8s_userMg", Path: "k8s_userMg", ResourcesSort: 1, ResourcesType: "C", Title: "用户管理", Icon: "", Display: 1, OnlyBuiltinRoot: 2},
	{ID: "3_1_1", ParentID: "3_1", Name: "k8s_userMg:tb:detail", Path: "k8s_userMg/k8s_userDetail", ResourcesSort: 1, ResourcesType: "H", Title: "用户详情", Icon: "", Display: 1, OnlyBuiltinRoot: 2},
	{ID: "3_1_2", ParentID: "3_1", Name: "k8s_userMg:add", Path: "#", ResourcesSort: 2, ResourcesType: "F", Title: "创建", Icon: "", Display: 1, OnlyBuiltinRoot: 2},
	// 系统设置-用户组管理
	{ID: "3_2", ParentID: "3", Name: "k8s_userGroupMg", Path: "k8s_userGroupMg", ResourcesSort: 2, ResourcesType: "C", Title: "用户组管理", Icon: "", Display: 1, OnlyBuiltinRoot: 2},
	{ID: "3_2_1", ParentID: "3_2", Name: "k8s_userGroupMg:tb:detail", Path: "k8s_userGroupMg/k8s_userGroupDetail", ResourcesSort: 1, ResourcesType: "H", Title: "用户组详情", Icon: "", Display: 1, OnlyBuiltinRoot: 2},
	// 系统设置-角色管理
	{ID: "3_3", ParentID: "3", Name: "k8s_roleMg", Path: "k8s_roleMg", ResourcesSort: 3, ResourcesType: "C", Title: "角色管理", Icon: "", Display: 1, OnlyBuiltinRoot: 2},
	{ID: "3_3_1", ParentID: "3_3", Name: "k8s_roleMg:tb:detail", Path: "k8s_roleMg/k8s_roleMgDetail", ResourcesSort: 1, ResourcesType: "H", Title: "角色管理详情", Icon: "", Display: 1, OnlyBuiltinRoot: 2},
	// 系统设置-关联管理
	{ID: "3_4", ParentID: "3", Name: "k8s_associateMg", Path: "k8s_associateMg", ResourcesSort: 4, ResourcesType: "C", Title: "关联管理", Icon: "", Display: 1, OnlyBuiltinRoot: 2},
	// 系统设置-菜单管理
	{ID: "3_5", ParentID: "3", Name: "k8s_menuMg", Path: "k8s_menuMg", ResourcesSort: 5, ResourcesType: "C", Title: "菜单管理", Icon: "", Display: 1, OnlyBuiltinRoot: 1},
}

// Make sure the table exists and init data
func ensureTableAndInitData() {
	// tenant
	if !Gdb.Migrator().HasTable(model.Tenant{}.TableName()) {
		err := Gdb.AutoMigrate(&model.Tenant{})
		if err != nil {
			Mlog.Panic("failed to create table tenant, error message: ", err.Error())
			return
		}
		Mlog.Info("successfully to create table tenant")

		// initialize tenant data
		{
			salt := getRandomString(10)
			password := encodePassword("builtin_0!@", salt)
			builtInRoot := model.Tenant{
				ID:       "builtin_root",
				Password: password,
				Salt:     salt,
			}
			createResult := Gdb.Create(&builtInRoot)
			if createResult.Error != nil {
				Mlog.Panic("failed to initialize tenant data, error message: ", createResult.Error.Error())
				return
			}
			Mlog.Info("successfully initialized tenant data, the number of inserts is ", createResult.RowsAffected)
		}
		{
			salt := getRandomString(10)
			AppConfObj.InitPassword = encodePassword(AppConfObj.InitPassword, salt)
			tenant := model.Tenant{
				ID:         AppConfObj.InitTenantID,
				TenantName: AppConfObj.InitTenantName,
				Password:   AppConfObj.InitPassword,
				Salt:       salt,
				Email:      AppConfObj.InitEmail,
			}
			createResult := Gdb.Create(&tenant)
			if createResult.Error != nil {
				Mlog.Panic("failed to initialize tenant data, error message: ", createResult.Error.Error())
				return
			}
			Mlog.Info("successfully initialized tenant data, the number of inserts is ", createResult.RowsAffected)
		}
	}

	// user
	if !Gdb.Migrator().HasTable(model.User{}.TableName()) {
		err := Gdb.AutoMigrate(&model.User{})
		if err != nil {
			Mlog.Panic("failed to create table user, error message: ", err.Error())
			return
		}
		Mlog.Info("successfully to create table user")
	}

	// group
	if !Gdb.Migrator().HasTable(model.Group{}.TableName()) {
		err := Gdb.AutoMigrate(&model.Group{})
		if err != nil {
			Mlog.Panic("failed to create table group, error message: ", err.Error())
			return
		}
		Mlog.Info("successfully to create table group")
	}

	// webperm
	if !Gdb.Migrator().HasTable(model.Webperm{}.TableName()) {
		err := Gdb.AutoMigrate(&model.Webperm{})
		if err != nil {
			Mlog.Panic("failed to create table webperm, error message: ", err.Error())
			return
		}
		Mlog.Info("successfully to create table webperm")

		for _, webperm := range webperms {
			createResult := Gdb.Create(&webperm)
			if createResult.Error != nil {
				Mlog.Panic("failed to initialize webperm data, error message: ", createResult.Error.Error())
				return
			}
		}

		Mlog.Info("successfully initialized webperm data")
	}

	// role
	if !Gdb.Migrator().HasTable(model.Role{}.TableName()) {
		err := Gdb.AutoMigrate(&model.Role{})
		if err != nil {
			Mlog.Panic("failed to create table role, error message: ", err.Error())
			return
		}
		Mlog.Info("successfully to create table role")
	}

	// permission
	if !Gdb.Migrator().HasTable(model.Permission{}.TableName()) {
		err := Gdb.AutoMigrate(&model.Permission{})
		if err != nil {
			Mlog.Panic("failed to create table permission, error message: ", err.Error())
			return
		}
		Mlog.Info("successfully to create table permission")

		createResult := Gdb.Create(&permissions)
		if createResult.Error != nil {
			Mlog.Panic("failed to initialize permission data, error message: ", createResult.Error.Error())
			return
		}
		Mlog.Info("successfully initialized permission data, the number of inserts is ", createResult.RowsAffected)
	}

	// cluster
	if !Gdb.Migrator().HasTable(model.Cluster{}.TableName()) {
		err := Gdb.AutoMigrate(&model.Cluster{})
		if err != nil {
			Mlog.Panic("failed to create table cluster, error message: ", err.Error())
			return
		}
		Mlog.Info("successfully to create table cluster")
	}

	// template
	if !Gdb.Migrator().HasTable(model.Template{}.TableName()) {
		err := Gdb.AutoMigrate(&model.Template{})
		if err != nil {
			Mlog.Panic("failed to create table template, error message: ", err.Error())
			return
		}
		Mlog.Info("successfully to create table template")
	}
}

// initCasbinRule init casbin rule data
func initCasbinRule() {
	result := make([]model.CasbinRule, 0)
	queryResult := Gdb.Find(&result)
	if queryResult.Error != nil {
		Mlog.Panic("failed to query data in table casbin_rule, error message: ", queryResult.Error.Error())
		return
	}
	if queryResult.RowsAffected == 0 {
		casbinRules := make([]model.CasbinRule, 0)

		for _, permission := range permissions {
			casbinRule := model.CasbinRule{
				PType: "p",
				V0:    "role-admin",
				V1:    AppConfObj.InitTenantID,
				V2:    permission.PermissionUrl,
				V3:    permission.PermissionAction,
			}
			casbinRules = append(casbinRules, casbinRule)
		}

		createResult := Gdb.Create(&casbinRules)
		if createResult.Error != nil {
			Mlog.Panic("failed to initialize casbin_rule data, error message: ", createResult.Error.Error())
			return
		}
		Mlog.Info("successfully initialized casbin_rule data, the number of inserts is ", createResult.RowsAffected)
	}

	err := CasbinSE.LoadPolicy()
	if err != nil {
		Mlog.Panic("failed to CasbinSE.LoadPolicy, error message: ", err.Error())
		return
	}
}

// source: https://github.com/gogs/gogs/blob/9ee80e3e5426821f03a4e99fad34418f5c736413/modules/base/tool.go#L58
// getRandomString generate random string by specify chars.
func getRandomString(n int, alphabets ...byte) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		if len(alphabets) == 0 {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return string(bytes)
}

func encodePassword(password string, salt string) string {
	newPasswd := pbkdf2([]byte(password), []byte(salt), 10000, 50, sha256.New)
	return hex.EncodeToString(newPasswd)
}

// source: https://github.com/gogs/gogs/blob/9ee80e3e5426821f03a4e99fad34418f5c736413/modules/base/tool.go#L73
func pbkdf2(password, salt []byte, iter, keyLen int, h func() hash.Hash) []byte {
	prf := hmac.New(h, password)
	hashLen := prf.Size()
	numBlocks := (keyLen + hashLen - 1) / hashLen

	var buf [4]byte
	dk := make([]byte, 0, numBlocks*hashLen)
	U := make([]byte, hashLen)
	for block := 1; block <= numBlocks; block++ {
		// N.B.: || means concatenation, ^ means XOR
		// for each block T_i = U_1 ^ U_2 ^ ... ^ U_iter
		// U_1 = PRF(password, salt || uint(i))
		prf.Reset()
		prf.Write(salt)
		buf[0] = byte(block >> 24)
		buf[1] = byte(block >> 16)
		buf[2] = byte(block >> 8)
		buf[3] = byte(block)
		prf.Write(buf[:4])
		dk = prf.Sum(dk)
		T := dk[len(dk)-hashLen:]
		copy(U, T)

		// U_n = PRF(password, U_(n-1))
		for n := 2; n <= iter; n++ {
			prf.Reset()
			prf.Write(U)
			U = U[:0]
			U = prf.Sum(U)
			for x := range U {
				T[x] ^= U[x]
			}
		}
	}
	return dk[:keyLen]
}
