package api

import (
	"net/http"

	"ftk8s/base/cfg"
	"ftk8s/base/enum"
	"ftk8s/middle"
	"ftk8s/model"
	"ftk8s/service"
	"ftk8s/util"

	"github.com/gin-gonic/gin"
)

func UpdateWebpermTreeByRole(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqUpdateWebpermTreeByRole)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
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

	err = service.UpdateWebpermTreeByRole(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to service.UpdateWebpermTreeByRole, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.ReFail(langReq, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, ""))
}

func ReadWebpermTreeByRole(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqReadWebpermTreeByRole)
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

	data, err := service.ReadWebpermTreeByRole(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to service.ReadWebpermTreeByRole, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.ReFail(langReq, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, data))
}

func ListWebpermFunsAndRights(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)
	userType := c.GetInt(middle.HeaderFieldUserType)

	reqObj := new(model.ReqListWebpermFunsAndRights)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	reqObj.UserType = userType
	if userType == enum.UserTypeTenant {
		tenant, err := util.GetTenantFromContext(c)
		if err != nil {
			cfg.Mlog.Error("failed to GetTenantFromContext, error message: ", err.Error())
			c.JSON(http.StatusBadRequest, util.ReFail(langReq, ""))
			return
		}
		reqObj.Tenant = *tenant
		reqObj.TenantID = tenant.ID

	} else {
		user, err := util.GetUserFromContext(c)
		if err != nil {
			cfg.Mlog.Error("failed to GetUserFromContext, error message: ", err.Error())
			c.JSON(http.StatusBadRequest, util.ReFail(langReq, ""))
			return
		}
		reqObj.User = *user
		reqObj.TenantID = user.TenantID
	}

	result, err := service.ListWebpermFunsAndRights(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to service.ListWebpermFunsAndRights, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, result))
}
