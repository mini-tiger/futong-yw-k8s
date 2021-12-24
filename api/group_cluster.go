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

// UpdateGroupByCluster
func UpdateGroupByCluster(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqUpdateGroupByCluster)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = service.UpdateGroupByCluster(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to update group by cluster, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataUpdate, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully update group by cluster"))
}

// ReadGroupByCluster
func ReadGroupByCluster(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqReadGroupByCluster)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	data, err := service.ReadGroupByCluster(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to read group by cluster, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, data))
}

// UpdateClusterByGroup
func UpdateClusterByGroup(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqUpdateClusterByGroup)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = service.UpdateClusterByGroup(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to update cluster by group, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataUpdate, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully update cluster by group"))
}

// ReadClusterByGroup
func ReadClusterByGroup(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqReadClusterByGroup)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	data, err := service.ReadClusterByGroup(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to read cluster by group, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, data))
}
