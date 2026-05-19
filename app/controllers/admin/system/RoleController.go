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
	dto2 "yixiang.co/go-mall/app/service/menu_service/dto"
	"yixiang.co/go-mall/app/service/role_service"
	"yixiang.co/go-mall/pkg/app"
	"yixiang.co/go-mall/pkg/constant"
	"yixiang.co/go-mall/pkg/logging"
	"yixiang.co/go-mall/pkg/util"
)

// role API
type RoleController struct {
}

// @Title get role by ID
// @Description get role by ID
// @Param    id        path     int    true        "role ID"
// @Success 200 {object} app.Response
// @router /:id [get]
func (e *RoleController) GetOne(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	roleService := role_service.Role{
		Id: id,
	}
	vo := roleService.GetOneRole()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title role list
// @Description role list
// @Success 200 {object} app.Response
// @router / [get]
func (e *RoleController) GetAll(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	blurry := c.DefaultQuery("blurry", "")
	roleService := role_service.Role{
		Name:     blurry,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := roleService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title add role
// @Description add role
// @Success 200 {object} app.Response
// @router / [post]
func (e *RoleController) Post(c *gin.Context) {
	var (
		model models.SysRole
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	roleService := role_service.Role{
		M: &model,
	}

	if err := roleService.Insert(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @router / [put]
func (e *RoleController) Put(c *gin.Context) {
	var (
		model models.SysRole
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	roleService := role_service.Role{
		M: &model,
	}

	if err := roleService.Save(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title delete role
// @Description delete role
// @Success 200 {object} app.Response
// @router / [delete]
func (e *RoleController) Delete(c *gin.Context) {
	var (
		ids  []int64
		appG = app.Gin{C: c}
	)
	c.BindJSON(&ids)
	roleService := role_service.Role{Ids: ids}

	if err := roleService.Del(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title update role menus
// @Description update role menus
// @Success 200 {object} app.Response
// @router /menu [put]
func (e *RoleController) Menu(c *gin.Context) {
	var (
		model dto2.RoleMenu
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	logging.Info(model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	roleService := role_service.Role{Dto: model}
	if err := roleService.BatchRoleMenuAdd(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)

}
