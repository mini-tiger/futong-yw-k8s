package middle

import (
	"net/http"
	"strconv"
	"strings"

	"ftk8s/base/cfg"
	"ftk8s/base/enum"
	"ftk8s/storage"
	"ftk8s/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type reqClusterID struct {
	ClusterID int `json:"cluster_id"`
}

// CheckPermOnlyBuiltinRoot only BuiltinRoot can access
func CheckPermOnlyBuiltinRoot() gin.HandlerFunc {
	return func(c *gin.Context) {
		langReq := c.Request.Header.Get(HeaderFieldLang)

		tenantID, err := util.GetTenantIDFromContext(c)
		if err != nil {
			cfg.Mlog.Error("failed to GetTenantIDFromContext, error message: ", err.Error())
			c.JSON(http.StatusBadRequest, util.ReFail(langReq, ""))
			c.Abort()
			return
		}

		if tenantID == enum.BuiltinRoot {
			c.Next()

		} else {
			cfg.Mlog.Error("only BuiltinRoot can access")
			c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgPermissionDenied, "Permission denied"))
			c.Abort()
			return
		}
	}
}

// CheckPermOnlyTenant only tenant has permission
func CheckPermOnlyTenant() gin.HandlerFunc {
	return func(c *gin.Context) {
		langReq := c.Request.Header.Get(HeaderFieldLang)
		userType := c.GetInt(HeaderFieldUserType)

		if userType == enum.UserTypeTenant {
			c.Next()

		} else {
			cfg.Mlog.Error("only tenant has permission")
			c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgPermissionDenied, "only tenant has permission"))
			c.Abort()
			return
		}
	}
}

// CheckPermAccessCluster check permission to access the cluster
func CheckPermAccessCluster() gin.HandlerFunc {
	return func(c *gin.Context) {
		langReq := c.Request.Header.Get(HeaderFieldLang)
		userType := c.GetInt(HeaderFieldUserType)

		// If it is a tenant, can access any cluster
		if userType == enum.UserTypeTenant {
			c.Next()

			// If it is a user, check the permission to cluster
		} else if userType == enum.UserTypeNormal {
			var clusterID int
			clusterIDString := c.Query(enum.ReqFieldClusterID)
			if clusterIDString == "" {
				clusterIDObj := new(reqClusterID)
				if err := c.ShouldBindBodyWith(clusterIDObj, binding.JSON); err != nil {
					cfg.Mlog.Error("failed to ShouldBindBodyWith, error message: ", err.Error())
					c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
					c.Abort()
					return
				} else {
					clusterID = clusterIDObj.ClusterID
				}

			} else {
				var err error
				clusterID, err = strconv.Atoi(clusterIDString)
				if err != nil {
					cfg.Mlog.Error("failed to strconv.Atoi, error message: ", err.Error())
					c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
					c.Abort()
					return
				}
			}

			user, err := util.GetUserFromContext(c)
			if err != nil {
				cfg.Mlog.Error("failed to GetUserFromContext, error message: ", err.Error())
				c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgPermissionDenied, "No permission to access this cluster"))
				c.Abort()
				return
			}

			clusters, err := storage.ListClusterIDByUserID(user.ID)
			if err != nil {
				cfg.Mlog.Error("failed to ListClusterIDByUserID, error message: ", err.Error())
				c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgPermissionDenied, "No permission to access this cluster"))
				c.Abort()
				return
			}

			flag := false
			for _, cluster := range clusters {
				if cluster.ID == clusterID {
					flag = true
					break
				}
			}

			if !flag {
				cfg.Mlog.Errorf("the user(%s) no permission to access this cluster", user.ID)
				c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgPermissionDenied, "No permission to access this cluster"))
				c.Abort()
				return
			}

			c.Next()
		}
	}
}

// CheckPermOperateResource check permission to operate the resource
func CheckPermOperateResource() gin.HandlerFunc {
	return func(c *gin.Context) {
		langReq := c.Request.Header.Get(HeaderFieldLang)
		userType := c.GetInt(HeaderFieldUserType)
		reqMethodTemp := c.Request.Method
		reqMethod := strings.ToLower(reqMethodTemp)
		permissionUrl := c.Request.URL.EscapedPath()

		// If it is a tenant, there is no permission limit
		if userType == enum.UserTypeTenant {
			c.Next()

			// If it is a user, check the permission
		} else if userType == enum.UserTypeNormal {
			user, err := util.GetUserFromContext(c)
			if err != nil {
				cfg.Mlog.Error("failed to check user permission, error message: ", err.Error())
				c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgPermissionDenied, "Permission denied"))
				c.Abort()
				return
			}

			tenantID := user.TenantID

			roles := make([]string, 0)
			rolesOfUser := cfg.CasbinSE.GetRolesForUserInDomain(user.ID, tenantID)
			roles = append(roles, rolesOfUser...)
			groupIDSlice, err := storage.GetGroupIDSliceByUserID(user)
			if err != nil {
				cfg.Mlog.Error(err.Error())
				c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgPermissionDenied, "Permission denied"))
				c.Abort()
				return
			}
			for _, groupID := range groupIDSlice {
				rolesOfGroup := cfg.CasbinSE.GetRolesForUserInDomain(groupID, tenantID)
				roles = append(roles, rolesOfGroup...)
			}

			if len(roles) == 0 {
				c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgPermissionDenied, "Permission denied"))
				c.Abort()
				return
			}

			for _, roleAccount := range roles {
				ok, err := cfg.CasbinSE.Enforce(roleAccount, tenantID, permissionUrl, reqMethod)
				if err != nil {
					cfg.Mlog.Error("failed to check user permission, error message: ", err.Error())
					c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgPermissionDenied, "Permission denied"))
					c.Abort()
					return
				}
				if !ok {
					c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgPermissionDenied, "Permission denied"))
					c.Abort()
					return
				}
			}

			c.Next()
		}
	}
}
