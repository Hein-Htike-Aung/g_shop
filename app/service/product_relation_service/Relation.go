/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package product_relation_service

import (
	"errors"
	"yixiang.co/go-mall/app/models"
	"yixiang.co/go-mall/app/models/vo"
	"yixiang.co/go-mall/app/params"
	relationEnum "yixiang.co/go-mall/pkg/enums/relation"
	"yixiang.co/go-mall/pkg/global"
	"yixiang.co/go-mall/pkg/util"
)

type Relation struct {
	Id int64
	Name string

	Enabled int

	PageNum int
	PageSize int

	M *models.YshopStoreProductRelation

	Ids []int64

	Uid int64
	Param *params.RelationParam

}

// review list
func (d *Relation) GetUserCollectList() ([]models.YshopStoreProductRelation, int, int) {
	maps := make(map[string]interface{})
	maps["uid"] = d.Uid
	maps["type"] = relationEnum.COLLECT


	total, list := models.GetAllProductRelation(d.PageNum, d.PageSize, maps)
	totalNum := util.Int64ToInt(total)
	totalPage := util.GetTotalPage(totalNum, d.PageSize)
	return list, totalNum, totalPage
}

//add collect
func (d *Relation) AddRelation() error {
	//productId := com.StrTo(d.Param.Id).MustInt64()
	if IsRelation(d.Param.Id,d.Uid) {
		return errors.New("already favorited")
	}
	model := &models.YshopStoreProductRelation{
		Uid: d.Uid,
		ProductId: d.Param.Id,
		Type: d.Param.Category,
	}
	return models.AddStoreProductRelation(model)
}

//del collect
func (d *Relation) DelRelation() error {
	if !IsRelation(d.Param.Id,d.Uid) {
		return errors.New("already unfavorited")
	}
	err := global.YSHOP_DB.
		Where("uid = ?",d.Uid).
		Where("product_id = ?",d.Param.Id).
		Where("type = ?",relationEnum.COLLECT).
		Delete(&models.YshopStoreProductRelation{}).Error
	if err != nil {
		global.YSHOP_LOG.Error(err)
		return errors.New("cancellation failed")
	}
	return nil
}

// is favorited
func IsRelation(productId,uid int64) bool  {
	var (
		count int64
		error error
	)
	error = global.YSHOP_DB.Model(&models.YshopStoreProductRelation{}).
		Where("uid = ?",uid).
		Where("product_id = ?",productId).
		Where("type = ?",relationEnum.COLLECT).
		Count(&count).Error
	if error != nil {
		global.YSHOP_LOG.Error(error)
		return false
	}
	if count == 0 {
		return false
	}

	return true
}



func (d *Relation) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Name != "" {
		maps["name"] = d.Name
	}

	total,list := models.GetAllProductRelation(d.PageNum,d.PageSize,maps)
	return vo.ResultList{Content: list,TotalElements: total}
}



func (d *Relation) Insert() error {
	return models.AddStoreProductRelation(d.M)
}

func (d *Relation) Save() error {
	return models.UpdateByStoreProductRelation(d.M)
}

func (d *Relation) Del() error {
	return models.DelByStoreProductRelations(d.Ids)
}