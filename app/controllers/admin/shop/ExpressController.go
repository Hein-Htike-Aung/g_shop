/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package shop

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"yixiang.co/go-mall/app/models"
	"yixiang.co/go-mall/app/service/express_service"
	"yixiang.co/go-mall/pkg/app"
	"yixiang.co/go-mall/pkg/constant"
	"yixiang.co/go-mall/pkg/util"
)

// express API
type ExpressController struct {
}

// @Title express carrier list
// @Description express carrier list
// @Success 200 {object} app.Response
// @router / [get]
func (e *ExpressController) GetAll(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	expressService := express_service.Express{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := expressService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title add express carrier
// @Description add express carrier
// @Success 200 {object} app.Response
// @router / [post]
func (e *ExpressController) Post(c *gin.Context) {
	var (
		model models.YshopExpress
		appG  = app.Gin{C: c}
	)

	paramErr := app.BindAndValidate(c, &model)
	if paramErr != nil {
		appG.Response(http.StatusBadRequest, paramErr.Error(), nil)
		return
	}

	expressService := express_service.Express{
		M: &model,
	}

	if err := expressService.Insert(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)

}

// @Title update express carrier
// @Description update express carrier
// @Success 200 {object} app.Response
// @router / [put]
func (e *ExpressController) Put(c *gin.Context) {
	var (
		model models.YshopExpress
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	expressService := express_service.Express{
		M: &model,
	}

	if err := expressService.Save(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title delete express carrier
// @Description delete express carrier
// @Success 200 {object} app.Response
// @router /:id [delete]
func (e *ExpressController) Delete(c *gin.Context) {
	var (
		ids  []int64
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	ids = append(ids, id)
	expressService := express_service.Express{Ids: ids}

	if err := expressService.Del(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}
