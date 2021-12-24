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

func CreateCronJob(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqCreateCronJob)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	cronJob, err := service.CreateCronJob(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to create cronJob, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataInsert, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, cronJob))
}

func UpdateCronJob(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqUpdateCronJob)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	cronJob, err := service.UpdateCronJob(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to update cronJob, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataUpdate, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, cronJob))
}

func ReadCronJob(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqReadCronJob)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	cronJob, err := service.ReadCronJob(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to read cronJob, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, cronJob))
}

func DeleteCronJob(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqDeleteCronJob)
	err := c.BindJSON(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindJSON, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	err = service.DeleteCronJob(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to delete cronJob, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataDelete, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, "successfully delete cronJob"))
}

func ListCronJob(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	reqObj := new(model.ReqListCronJob)
	err := c.BindQuery(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to BindQuery, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgInvalidParameter, ""))
		return
	}

	extra, cronJobs, err := service.ListCronJob(reqObj)
	if err != nil {
		cfg.Mlog.Error("failed to list cronJob, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOKExtra(langReq, cronJobs, extra))
}
