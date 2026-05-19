/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package system

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"yixiang.co/go-mall/app/models"
	"yixiang.co/go-mall/app/service/dict_detail_service"
	"yixiang.co/go-mall/pkg/app"
	"yixiang.co/go-mall/pkg/constant"
	"yixiang.co/go-mall/pkg/util"
)

// dictionary detail API
type DictDetailController struct {
}

// @Title dictionary detail list
// @Description dictionary detail list
// @Success 200 {object} app.Response
// @router / [get]
func (e *DictDetailController) GetAll(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	dictName := c.DefaultQuery("dictName", "")
	dictId := com.StrTo(c.DefaultQuery("dictId", "-1")).MustInt64()
	detailService := dict_detail_service.DictDetail{
		DictName: dictName,
		DictId:   dictId,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := detailService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title add dictionary detail
// @Description add dictionary detail
// @Success 200 {object} app.Response
// @router / [post]
func (e *DictDetailController) Post(c *gin.Context) {
	var (
		model models.SysDictDetail
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	dictDetailService := dict_detail_service.DictDetail{
		M: &model,
	}

	if err := dictDetailService.Insert(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title update dictionary detail
// @Description update dictionary detail
// @Success 200 {object} app.Response
// @router / [put]
func (e *DictDetailController) Put(c *gin.Context) {
	var (
		model models.SysDictDetail
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	dictDetailService := dict_detail_service.DictDetail{
		M: &model,
	}

	if err := dictDetailService.Save(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title delete dictionary detail
// @Description delete dictionary detail
// @Success 200 {object} app.Response
// @router /:id [delete]
func (e *DictDetailController) Delete(c *gin.Context) {
	var (
		ids  []int64
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	ids = append(ids, id)

	dictDetailService := dict_detail_service.DictDetail{Ids: ids}
	if err := dictDetailService.Del(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}
