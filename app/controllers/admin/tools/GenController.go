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
	"yixiang.co/go-mall/app/params/admin"
	"yixiang.co/go-mall/app/service/gen_service"
	"yixiang.co/go-mall/pkg/app"
	"yixiang.co/go-mall/pkg/constant"
	"yixiang.co/go-mall/pkg/util"
)

// code generator API
type GenController struct {
}

// @Title list all tables
// @Description list all tables
// @Success 200 {object} app.Response
// @router /tools/gen/tables [get]
func (e *GenController) GetAllDBTables(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	genService := gen_service.Gen{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := genService.GetDBTablesAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title import database table
// @Description import database table
// @Success 200 {object} app.Response
// @router /tools/gen/import [post]
func (e *GenController) ImportTable(c *gin.Context) {
	var (
		param admin.GenTableParan
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &param)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	genService := gen_service.Gen{
		GenTableParan: &param,
	}

	if err := genService.Insert(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)

}

// @Title list imported tables
// @Description list imported tables
// @Success 200 {object} app.Response
// @router /tools/gen/systables [get]
func (e *GenController) GetAllTables(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	genService := gen_service.Gen{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := genService.GetTablesAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title get table info
// @Description get table info
// @Success 200 {object} app.Response
// @router /tools/gen/config/:name[get]
func (e *GenController) GetTableInfo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	name := c.Param("name")
	genService := gen_service.Gen{
		Name: name,
	}
	vo := genService.GetTableInfo()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title get table columns
// @Description get table columns
// @Success 200 {object} app.Response
// @router /tools/gen/columns[get]
func (e *GenController) GetTableColumns(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	name := c.DefaultQuery("tableName", "")
	genService := gen_service.Gen{
		Name: name,
	}
	vo := genService.GetTableColumns()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title save configuration
// @Description save configuration
// @Success 200 {object} app.Response
// @router /gen/config [put]
func (e *GenController) ConfigPut(c *gin.Context) {
	var (
		model models.SysTables
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	genService := gen_service.Gen{
		Table: &model,
	}

	if err := genService.TableSave(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title save column configuration
// @Description save column configuration
// @Success 200 {object} app.Response
// @router /gen/columns [put]
func (e *GenController) ColumnsPut(c *gin.Context) {
	var (
		model []models.SysColumns
		appG  = app.Gin{C: c}
	)
	if err := c.ShouldBindJSON(&model); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	genService := gen_service.Gen{
		Columns: model,
	}

	if err := genService.ColumnSave(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title code preview
// @Description code preview
// @Success 200 {object} app.Response
// @router /tools/gen/preview [get]
func (e *GenController) Preview(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	name := c.Param("name")
	genService := gen_service.Gen{
		Name: name,
	}
	vo := genService.Preview()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title generate code
// @Description generate code
// @Success 200 {object} app.Response
// @router /tools/gen/code [get]
func (e *GenController) GenCode(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	name := c.Param("name")
	genService := gen_service.Gen{
		Name: name,
	}
	genService.GenCode()
	appG.Response(http.StatusOK, "code generated successfully under template/", nil)
}
