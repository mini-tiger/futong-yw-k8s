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

func CreateWebperm(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqCreateWebperm)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = service.CreateWebperm(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to create webperm, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataInsert, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully create webperm"))
}

func UpdateWebperm(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqUpdateWebperm)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = service.UpdateWebperm(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to update webperm, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataUpdate, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully update webperm"))
}

func DeleteWebperm(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqDeleteWebperm)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = service.DeleteWebperm(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to delete webperm, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataDelete, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully delete webperm"))
}

func ListWebpermTree(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	result, err := service.ListWebpermTree()
	if err != nil {
		cfg.Mlog.Error("failed to list webpermTree, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, result))
}

func ListWebpermTreeBindRole(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	result, err := service.ListWebpermTreeBindRole()
	if err != nil {
		cfg.Mlog.Error("failed to ListWebpermTreeBindRole, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, result))
}
