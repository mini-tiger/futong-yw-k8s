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

func CreateConfigMap(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqCreateConfigMap)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	configMap, err := service.CreateConfigMap(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to create configMap, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataInsert, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, configMap))
}

func UpdateConfigMap(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqUpdateConfigMap)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	configMap, err := service.UpdateConfigMap(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to update configMap, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataUpdate, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, configMap))
}

func ReadConfigMap(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqReadConfigMap)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	configMap, err := service.ReadConfigMap(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to read configMap, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, configMap))
}

func DeleteConfigMap(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqDeleteConfigMap)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = service.DeleteConfigMap(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to delete configMap, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataDelete, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully delete configMap"))
}

func ListConfigMap(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqListConfigMap)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	extra, configMaps, err := service.ListConfigMap(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to list configMap, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOKExtra(langReq, configMaps, extra))
}
