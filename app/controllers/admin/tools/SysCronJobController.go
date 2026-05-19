/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */

package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"yixiang.co/go-mall/app/models"
	"yixiang.co/go-mall/app/service/cron_job_service"
	"yixiang.co/go-mall/pkg/app"
	"yixiang.co/go-mall/pkg/constant"
	"yixiang.co/go-mall/pkg/util"
)

type SysCronJobController struct {
}

// @Title cron job list
// @Description cron job list
// @Success 200 {object} app.Response
// @router / [get]
func (e *SysCronJobController) GetAll(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	service := cron_job_service.SysCronJob{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := service.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title add cron job
// @Description add cron job
// @Success 200 {object} app.Response
// @router / [post]
func (e *SysCronJobController) Post(c *gin.Context) {
	var (
		model models.SysCronJob
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	service := cron_job_service.SysCronJob{
		M: &model,
	}

	if err := service.Insert(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title update cron job
// @Description update cron job
// @Success 200 {object} app.Response
// @router / [put]
func (e *SysCronJobController) Put(c *gin.Context) {
	var (
		model models.SysCronJob
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	service := cron_job_service.SysCronJob{
		M: &model,
	}

	if err := service.Save(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title delete cron job
// @Description delete cron job
// @Success 200 {object} app.Response
// @router /:id [delete]
func (e *SysCronJobController) Delete(c *gin.Context) {
	var (
		ids  []int64
		appG = app.Gin{C: c}
	)

	if strId := c.Param("id"); strId != "" {
		id := com.StrTo(strId).MustInt64()
		ids = append(ids, id)
	} else {
		c.BindJSON(&ids)
	}

	service := cron_job_service.SysCronJob{Ids: ids}
	if err := service.Del(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title run cron job
// @Description run cron job
// @Success 200 {object} app.Response
// @router /exec/:id [put]
func (e *SysCronJobController) Exec(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	service := cron_job_service.SysCronJob{
		Id: id,
	}

	if err := service.Exec(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, err.Error())
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title stop cron job
// @Description stop cron job
// @Success 200 {object} app.Response
// @router /stop/:id [put]
func (e *SysCronJobController) Stop(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	service := cron_job_service.SysCronJob{
		Id: id,
	}

	if err := service.Stop(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, err.Error())
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}
