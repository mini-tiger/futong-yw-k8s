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

// CreateGroup
func CreateGroup(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqCreateGroup)
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

	err = service.CreateGroup(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to create group, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataInsert, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully create group"))
}

// UpdateGroup
func UpdateGroup(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqUpdateGroup)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = service.UpdateGroup(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to update group, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataUpdate, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully update group"))
}

// ReadGroup
func ReadGroup(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqReadGroup)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	group, err := service.ReadGroup(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to read group, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, group))
}

// DeleteGroup
func DeleteGroup(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqDeleteGroup)
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

	deleteFlag, err := service.DeleteGroup(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to delete group, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataDelete, ""))
		return
	}

	if deleteFlag == 2 {
		cfg.Mlog.Warnf("this group(%s) has associated data", reqObj.GroupID)
		c.JSON(http.StatusOK, util.Re(langReq, enum.MsgHasAssociatedData, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully delete group"))
}

// ListGroup
func ListGroup(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqListGroup)
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

	extra, groups, err := service.ListGroup(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to list group, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOKExtra(langReq, groups, extra))
}
