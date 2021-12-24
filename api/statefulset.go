package api

import (
	"ftk8s/base/cfg"
	"ftk8s/base/enum"
	"ftk8s/middle"
	"ftk8s/model"
	"ftk8s/service"
	"ftk8s/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStatefulSet(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqCreateStatefulSet)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	statefulset, err := service.CreateStatefulSet(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to create statefulset, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataInsert, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, statefulset))
}

func UpdateStatefulSet(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqUpdateStatefulSet)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	statefulset, err := service.UpdateStatefulSet(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to update statefulset, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataUpdate, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, statefulset))
}

func ReadStatefulSet(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqReadStatefulSet)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	statefulset, err := service.ReadStatefulSet(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to read statefulset, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, statefulset))
}

func DeleteStatefulSet(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqDeleteStatefulSet)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = service.DeleteStatefulSet(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to delete statefulset, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataDelete, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully delete statefulset"))
}

func ListStatefulSet(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqListStatefulSet)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	extra, statefulsets, err := service.ListStatefulSet(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to list statefulset, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOKExtra(langReq, statefulsets, extra))
}

func ReadStatefulSetHistory(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqReadStatefulSetHistory)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	data, err := service.ReadStatefulSetHistory(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to read statefulset, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, data))
}
