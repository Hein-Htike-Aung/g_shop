/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package weixin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yixiang.co/go-mall/app/service/wechat_menu_service"
	dto2 "yixiang.co/go-mall/app/service/wechat_menu_service/dto"
	"yixiang.co/go-mall/pkg/app"
	"yixiang.co/go-mall/pkg/constant"
)

// menu API
type WechatMenuController struct {
}

// @Title get menu
// @Description get menu
// @Success 200 {object} app.Response
// @router / [get]
func (e *WechatMenuController) GetAll(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	meuService := wechat_menu_service.Menu{}
	vo := meuService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title update menu
// @Description update menu
// @Success 200 {object} app.Response
// @router / [post]
func (e *WechatMenuController) Post(c *gin.Context) {
	var (
		dto  dto2.WechatMenu
		appG = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &dto)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	meuService := wechat_menu_service.Menu{
		Dto: dto,
	}

	if err := meuService.Insert(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)

}
