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

// RegisterTenant register a tenant
func RegisterTenant(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqRegisterTenant)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = service.RegisterTenant(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to register tenant, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataInsert, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully registered tenant"))
}

// Login
func Login(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqLogin)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}
	reqObj.ClientIP = c.ClientIP()

	token, err := service.Login(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to login, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgLoginFailed, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, token))
}

// CreateUser create a user
func CreateUser(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqCreateUser)
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

	err = service.CreateUser(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to create user, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataInsert, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully create user"))
}

// UpdateUser
func UpdateUser(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqUpdateUser)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = service.UpdateUser(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to update user, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataInsert, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully update user"))
}

// ResetPassword
func ResetPassword(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)
	userType := c.GetInt(middle.HeaderFieldUserType)

	reqObj := new(model.ReqResetPassword)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
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
	}

	err = service.ResetPassword(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to reset password, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataUpdate, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully reset password"))
}

// ReadUser
func ReadUser(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqReadUser)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	result, err := service.ReadUser(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to read user, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, result))
}

// DeleteUser
func DeleteUser(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqDeleteUser)
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

	err = service.DeleteUser(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to delete user, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataDelete, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully delete user"))
}

// ListUser
func ListUser(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqListUser)
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

	extra, users, err := service.ListUser(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to list user, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOKExtra(langReq, users, extra))
}
