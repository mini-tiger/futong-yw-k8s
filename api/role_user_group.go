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

// UpdateRoleByUserGroup
func UpdateRoleByUserGroup(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqUpdateRoleByUserGroup)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	var roleAssObj string
	if reqObj.UserID == "" && reqObj.GroupID != "" {
		roleAssObj = reqObj.GroupID
	} else if reqObj.UserID != "" && reqObj.GroupID == "" {
		roleAssObj = reqObj.UserID
	} else {
		cfg.Mlog.Error("invalid parameter user_id and group_id")
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}
	reqObj.RoleAssObj = roleAssObj

	tenantID, err := util.GetTenantIDFromContext(c)
	if err != nil {
		cfg.Mlog.Error("failed to GetTenantIDFromContext, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.ReFail(langReq, ""))
		return
	}
	reqObj.TenantID = tenantID

	err = service.UpdateRoleByUserGroup(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to service.UpdateRoleByUserGroup, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.ReFail(langReq, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, ""))
}

// ReadRoleByUserGroup
func ReadRoleByUserGroup(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqReadRoleByUserGroup)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	var roleAssObj string
	if reqObj.UseID == "" && reqObj.GroupID != "" {
		roleAssObj = reqObj.GroupID
	} else if reqObj.UseID != "" && reqObj.GroupID == "" {
		roleAssObj = reqObj.UseID
	} else {
		cfg.Mlog.Error("invalid parameter user_id and group_id")
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}
	reqObj.RoleAssObj = roleAssObj

	tenantID, err := util.GetTenantIDFromContext(c)
	if err != nil {
		cfg.Mlog.Error("failed to GetTenantIDFromContext, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.ReFail(langReq, ""))
		return
	}
	reqObj.TenantID = tenantID

	data, err := service.ReadRoleByUserGroup(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to service.ReadRoleByUserGroup, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.ReFail(langReq, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, data))
}

// UpdateUserGroupByRole
func UpdateUserGroupByRole(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqUpdateUserGroupByRole)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	var roleAssObjSlice []string
	var roleAssObjPrefix string
	if len(reqObj.UserIDSlice) == 0 && len(reqObj.GroupIDSlice) != 0 {
		roleAssObjSlice = reqObj.GroupIDSlice
		roleAssObjPrefix = enum.PrefixGroup
	} else if len(reqObj.UserIDSlice) != 0 && len(reqObj.GroupIDSlice) == 0 {
		roleAssObjSlice = reqObj.UserIDSlice
		roleAssObjPrefix = enum.PrefixUser
	} else {
		cfg.Mlog.Error("invalid parameter user_id_slice and group_id_slice")
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}
	reqObj.RoleAssObjSlice = roleAssObjSlice
	reqObj.RoleAssObjPrefix = roleAssObjPrefix

	tenantID, err := util.GetTenantIDFromContext(c)
	if err != nil {
		cfg.Mlog.Error("failed to GetTenantIDFromContext, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.ReFail(langReq, ""))
		return
	}
	reqObj.TenantID = tenantID

	err = service.UpdateUserGroupByRole(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to service.UpdateUserGroupByRole, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.ReFail(langReq, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, ""))
}

// ReadUserGroupByRole
func ReadUserGroupByRole(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqReadUserGroupByRole)
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

	data, err := service.ReadUserGroupByRole(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to service.ReadUserGroupByRole, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.ReFail(langReq, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, data))
}
