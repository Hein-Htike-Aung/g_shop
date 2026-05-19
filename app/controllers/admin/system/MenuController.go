/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package system

import (
	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"yixiang.co/go-mall/app/models"
	"yixiang.co/go-mall/app/service/menu_service"
	"yixiang.co/go-mall/pkg/app"
	"yixiang.co/go-mall/pkg/constant"
	"yixiang.co/go-mall/pkg/jwt"
)

// menu API
type MenuController struct {
}

// @Title menu list
// @Description menu list
// @Success 200 {object} app.Response
// @router / [get]
func (e *MenuController) GetAll(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	name := c.DefaultQuery("blurry", "")
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	menuService := menu_service.Menu{Name: name, Enabled: enabled}
	vo := menuService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title add menu
// @Description add menu
// @Success 200 {object} app.Response
// @router / [post]
func (e *MenuController) Post(c *gin.Context) {
	var (
		model models.SysMenu
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	menuService := menu_service.Menu{
		M: &model,
	}

	if err := menuService.Insert(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title update menu
// @Description update menu
// @Success 200 {object} app.Response
// @router / [put]
func (e *MenuController) Put(c *gin.Context) {
	var (
		model models.SysMenu
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	menuService := menu_service.Menu{
		M: &model,
	}

	if err := menuService.Save(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title delete menu
// @Description delete menu
// @Success 200 {object} app.Response
// @router / [delete]
func (e *MenuController) Delete(c *gin.Context) {
	var (
		ids  []int64
		appG = app.Gin{C: c}
	)
	c.BindJSON(&ids)
	menuService := menu_service.Menu{Ids: ids}

	if err := menuService.Del(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title build menus
// @Description build menus
// @Success 200 {object} app.Response
// @router /build [get]
func (e *MenuController) Build(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	uid, _ := jwt.GetAdminUserId(c)
	menuService := menu_service.Menu{Uid: uid}
	logs.Info(uid)
	menus := menuService.BuildMenus()
	appG.Response(http.StatusOK, constant.SUCCESS, menus)
}

// @Title menu tree
// @Description menu tree
// @Success 200 {object} app.Response
// @router /tree [get]
func (e *MenuController) GetTree(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	menuService := menu_service.Menu{}
	vo := menuService.GetTree()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}
