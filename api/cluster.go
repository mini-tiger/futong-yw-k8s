package api

import (
	"net/http"

	"ftk8s/base/cfg"
	"ftk8s/base/enum"
	"ftk8s/ksc"
	"ftk8s/middle"
	"ftk8s/model"
	"ftk8s/service"
	"ftk8s/util"

	"github.com/gin-gonic/gin"
)

// CheckConnectCluster
func CheckConnectCluster(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqCheckConnectCluster)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = ksc.CheckConnectCluster(reqObj.ClusterAPI, reqObj.K8sConfig)
	if err != nil {
		cfg.Mlog.Error("failed to CheckConnectCluster, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.ReFail(langReq, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, ""))
}

// ImportCluster import k8s cluster key
func ImportCluster(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqImportCluster)
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

	err = service.ImportCluster(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to import cluster, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataInsert, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully import cluster"))
}

// UpdateCluster
func UpdateCluster(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqUpdateCluster)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = service.UpdateCluster(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to update cluster, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataUpdate, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully update cluster"))
}

// ReadCluster
func ReadCluster(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqReadCluster)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	cluster, err := service.ReadCluster(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to read cluster, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, cluster))
}

// DeleteCluster
func DeleteCluster(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqDeleteCluster)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = service.DeleteCluster(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to delete cluster, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataDelete, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully delete cluster"))
}

// ListCluster
func ListCluster(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)
	userType := c.GetInt(middle.HeaderFieldUserType)

	reqObj := new(model.ReqListCluster)
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
		reqObj.TenantID = tenant.ID
	} else {
		user, err := util.GetUserFromContext(c)
		if err != nil {
			cfg.Mlog.Error("failed to GetUserFromContext, error message: ", err.Error())
			c.JSON(http.StatusBadRequest, util.ReFail(langReq, ""))
			return
		}
		reqObj.UserID = user.ID
	}

	extra, clusters, err := service.ListCluster(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to list cluster, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	result := make([]model.Cluster, 0)
	for _, cluster := range clusters {
		cluster.K8sConfig = ""
		result = append(result, cluster)
	}

	c.JSON(http.StatusOK, util.ReOKExtra(langReq, clusters, extra))
}
