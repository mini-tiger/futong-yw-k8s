package api

import (
	"net/http"

	"ftk8s/base/cfg"
	"ftk8s/base/enum"
	"ftk8s/middle"
	"ftk8s/service"
	"ftk8s/util"

	"github.com/gin-gonic/gin"
)

func ListPermission(c *gin.Context) {
	langReq := c.Request.Header.Get(middle.HeaderFieldLang)

	permissions, err := service.ListPermission()
	if err != nil {
		cfg.Mlog.Error("failed to list permission, error message: ", err.Error())
		c.JSON(http.StatusBadRequest, util.Re(langReq, enum.MsgFailDataQuery, ""))
		return
	}

	c.JSON(http.StatusOK, util.ReOK(langReq, permissions))
}
