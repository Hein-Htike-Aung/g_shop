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
	"yixiang.co/go-mall/app/service/material_group_service"
	"yixiang.co/go-mall/pkg/app"
	"yixiang.co/go-mall/pkg/constant"
	"yixiang.co/go-mall/pkg/jwt"
)

// material group API
type MaterialGroupController struct {
}

// @Title material group list
// @Description material group list
// @Success 200 {object} app.Response
// @router / [get]
func (e *MaterialGroupController) GetAll(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	name := c.DefaultQuery("blurry", "")
	materialGroupService := material_group_service.MaterialGroup{
		Name: name,
	}
	vo := materialGroupService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Titleadd material group
// @Descriptionadd material group
// @Success 200 {object} app.Response
// @router / [post]
func (e *MaterialGroupController) Post(c *gin.Context) {
	var (
		model models.SysMaterialGroup
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	uid, _ := jwt.GetAdminUserId(c)
	model.CreateId = uid
	materialGroupService := material_group_service.MaterialGroup{
		M: &model,
	}

	if err := materialGroupService.Insert(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)

}

// @Title update material group
// @Description update material group
// @Success 200 {object} app.Response
// @router / [put]
func (e *MaterialGroupController) Put(c *gin.Context) {
	var (
		model models.SysMaterialGroup
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	uid, _ := jwt.GetAdminUserId(c)
	model.CreateId = uid
	materialGroupService := material_group_service.MaterialGroup{
		M: &model,
	}

	if err := materialGroupService.Save(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title delete material group
// @Description delete material group
// @Success 200 {object} app.Response
// @router /:id [delete]
func (e *MaterialGroupController) Delete(c *gin.Context) {
	var (
		ids  []int64
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	ids = append(ids, id)
	materialGroupService := material_group_service.MaterialGroup{Ids: ids}

	if err := materialGroupService.Del(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}
