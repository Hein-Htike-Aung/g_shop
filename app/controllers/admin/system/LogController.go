/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yixiang.co/go-mall/app/service/log_service"
	"yixiang.co/go-mall/pkg/app"
	"yixiang.co/go-mall/pkg/constant"
	"yixiang.co/go-mall/pkg/util"
)

// role API
type LogController struct {
}

// @Title log list
// @Description log list
// @Success 200 {object} app.Response
// @router / [get]
func (e *LogController) GetAll(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	blurry := c.DefaultQuery("blurry", "")
	logService := log_service.Log{
		Des:      blurry,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := logService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title delete logs
// @Description delete logs
// @Success 200 {object} app.Response
// @router / [delete]
func (e *LogController) Delete(c *gin.Context) {
	var (
		ids  []int64
		appG = app.Gin{C: c}
	)
	c.BindJSON(&ids)
	logService := log_service.Log{Ids: ids}

	if err := logService.Del(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}
