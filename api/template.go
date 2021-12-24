package api

import (
	"net/http"

	"ftk8s/base/cfg"
	"ftk8s/base/enum"
	"ftk8s/ksc"
	"ftk8s/middle"
	"ftk8s/model"
	"ftk8s/service"
	"ftk8s/storage"
	"ftk8s/util"

	"github.com/gin-gonic/gin"
)

func CreateTemplate(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqCreateTemplate)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	isSupport := util.IsSupportResourceKind(reqObj.TemplateKind)
	if !isSupport {
		cfg.Mlog.Error("Invalid resource kind: ", reqObj.TemplateKind)
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	tenantID, err := util.GetTenantIDFromContext(c)
	if err != nil {
		cfg.Mlog.Error("failed to GetTenantIDFromContext, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.ReFail(langReq, ""))
		return
	}
	reqObj.TenantID = tenantID

	err = service.CreateTemplate(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to create template, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataInsert, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully create template"))
}

func UpdateTemplate(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqUpdateTemplate)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	isSupport := util.IsSupportResourceKind(reqObj.TemplateKind)
	if !isSupport {
		cfg.Mlog.Error("Invalid resource kind: ", reqObj.TemplateKind)
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = service.UpdateTemplate(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to update template, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataUpdate, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully update template"))
}

func ReadTemplate(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqReadTemplate)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	template, err := service.ReadTemplate(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to read template, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, template))
}

func DeleteTemplate(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqDeleteTemplate)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = service.DeleteTemplate(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to delete template, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataDelete, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully delete template"))
}

func ListTemplate(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqListTemplate)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	tenantID, err := util.GetTenantIDFromContext(c)
	if err != nil {
		cfg.Mlog.Error("failed to GetTenantIDFromContext, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.ReFail(langReq, ""))
		return
	}
	reqObj.TenantID = tenantID

	reqObj.UrlQueryPara = c.Request.URL.Query()

	extra, templates, err := service.ListTemplate(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to list template, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOKExtra(langReq, templates, extra))
}

func DeployTemplate(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)
	userType := c.GetInt(middle.HeaderFieldUserType)

	reqObj := new(model.ReqDeployTemplate)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	isSupport := util.IsSupportResourceKind(reqObj.TemplateKind)
	if !isSupport {
		cfg.Mlog.Error("Invalid resource kind: ", reqObj.TemplateKind)
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	HasPerm := false
	// If it is a tenant, there is no permission limit
	if userType == enum.UserTypeTenant {
		HasPerm = true

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
			return
		}
		for _, groupID := range groupIDSlice {
			rolesOfGroup := cfg.CasbinSE.GetRolesForUserInDomain(groupID, tenantID)
			roles = append(roles, rolesOfGroup...)
		}

		if len(roles) == 0 {
			c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgPermissionDenied, "Permission denied"))
			return
		}

		for _, roleAccount := range roles {
			permissionUrl := "/api/resource/"
			switch reqObj.TemplateKind {
			case ksc.KindNameNamespace:
				permissionUrl += "namespace"
			case ksc.KindNameDeployment:
				permissionUrl += "deployment"
			case ksc.KindNameConfigMap:
				permissionUrl += "configmap"
			case ksc.KindNameService:
				permissionUrl += "service"
			default:
				permissionUrl = "unknown"
			}

			ok, err := cfg.CasbinSE.Enforce(roleAccount, tenantID, permissionUrl, "post")
			if err != nil {
				cfg.Mlog.Error("failed to check user permission, error message: ", err.Error())
				c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgPermissionDenied, "Permission denied"))
				return
			}
			if !ok {
				c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgPermissionDenied, "Permission denied"))
				return
			}
		}

	}

	if !HasPerm {
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgPermissionDenied, "Permission denied"))
		return
	}

	switch reqObj.TemplateKind {
	case ksc.KindNameNamespace:
		obj := new(model.ReqCreateNamespace)
		obj.ClusterID = reqObj.ClusterID
		obj.NamespaceData = reqObj.Content
		_, err := service.CreateNamespace(obj)
		if err != nil {
			cfg.Mlog.Error("failed to deploy template of namespace, error message: ", err.Error())
			c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataInsert, ""))
			return
		}
	case ksc.KindNameDeployment:
		obj := new(model.ReqCreateDeployment)
		obj.ClusterID = reqObj.ClusterID
		obj.NamespaceName = reqObj.NamespaceName
		obj.DeploymentData = reqObj.Content
		_, err := service.CreateDeployment(obj)
		if err != nil {
			cfg.Mlog.Error("failed to deploy template of deployment, error message: ", err.Error())
			c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataInsert, ""))
			return
		}
	case ksc.KindNameConfigMap:
		obj := new(model.ReqCreateConfigMap)
		obj.ClusterID = reqObj.ClusterID
		obj.NamespaceName = reqObj.NamespaceName
		obj.ConfigMapData = reqObj.Content
		_, err := service.CreateConfigMap(obj)
		if err != nil {
			cfg.Mlog.Error("failed to deploy template of configmap, error message: ", err.Error())
			c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataInsert, ""))
			return
		}
	case ksc.KindNameService:
		obj := new(model.ReqCreateService)
		obj.ClusterID = reqObj.ClusterID
		obj.NamespaceName = reqObj.NamespaceName
		obj.ServiceData = reqObj.Content
		_, err := service.CreateService(obj)
		if err != nil {
			cfg.Mlog.Error("failed to deploy template of service, error message: ", err.Error())
			c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataInsert, ""))
			return
		}
	default:
		cfg.Mlog.Error("failed to deploy template, unsupported template_kind")
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataInsert, ""))
		return
	}
}
