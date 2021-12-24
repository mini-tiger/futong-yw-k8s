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

func CreateRoleWithPermission(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqCreateRoleWithPermission)
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

	err = service.CreateRoleWithPermission(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to create role with permission, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataInsert, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully create role with permission"))
}

func UpdateRoleWithPermission(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqUpdateRoleWithPermission)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = service.UpdateRoleWithPermission(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to update role with permission, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataUpdate, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully update role with permission"))
}

func ReadRole(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqReadRole)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	role, err := service.ReadRole(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to read role, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, role))
}

func DeleteRole(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqDeleteRole)
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

	deleteFlag, err := service.DeleteRole(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to delete role, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataDelete, ""))
		return
	}

	if deleteFlag == 2 {
		cfg.Mlog.Warnf("this role(%s) has associated data", reqObj.RoleID)
		c.JSON(http.StatusOK, util.Re(langReq, enum.MsgHasAssociatedData, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully delete role"))
}

func ListRole(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqListRole)
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

	extra, roles, err := service.ListRole(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to list role, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOKExtra(langReq, roles, extra))
}
