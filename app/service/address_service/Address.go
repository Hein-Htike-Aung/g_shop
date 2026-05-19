/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package address_service

import (
	"encoding/json"
	"errors"
	"yixiang.co/go-mall/app/models"
	"yixiang.co/go-mall/app/models/vo"
	"yixiang.co/go-mall/app/params"
	"yixiang.co/go-mall/pkg/constant"
	"yixiang.co/go-mall/pkg/global"
	"yixiang.co/go-mall/pkg/redis"
	"yixiang.co/go-mall/pkg/util"
)

type Address struct {
	Id int64
	Name string

	Enabled int

	PageNum int
	PageSize int

	M *models.YshopUserAddress

	Ids []int64

	Param *params.AddressParan
	Uid int64

}

// delete address
func (d *Address) DelAddress() error {
	err := global.YSHOP_DB.
		Where("uid = ?",d.Uid).
		Where("id = ?",d.Id).
		Delete(&models.YshopUserAddress{}).Error
	if err != nil {
		global.YSHOP_LOG.Error(err)
		return  errors.New("operation failed")
	}

	return nil
}

// set default address
func (d *Address) SetDefault() error {
	var err error
	tx := global.YSHOP_DB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}else{
			tx.Commit()
		}
	}()
	err = tx.Model(&models.YshopUserAddress{}).
		Where("uid = ?",d.Uid).Update("is_default",0).Error
	if err != nil {
		global.YSHOP_LOG.Error(err)
		return  errors.New("operation failed")
	}
	err = tx.Model(&models.YshopUserAddress{}).
		Where("id = ?",d.Id).Update("is_default",1).Error
	if err != nil {
		global.YSHOP_LOG.Error(err)
		return  errors.New("operation failed")
	}
	return nil
}

//get list
func (d *Address) GetList() ([]models.YshopUserAddress,int,int){
	maps := make(map[string]interface{})
	maps["uid"] = d.Uid
	total,list := models.GetAllUserAddress(d.PageNum,d.PageSize,maps)

	totalNum := util.Int64ToInt(total)
	totalPage := util.GetTotalPage(totalNum,d.PageSize)
	return  list,totalNum,totalPage
}

//add or update
func (d *Address) AddOrUpdate() (int64,error) {
	var err error
	tx := global.YSHOP_DB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}else{
			tx.Commit()
		}
	}()
	userAddress := &models.YshopUserAddress{
		City: d.Param.Address.City,
		CityId: d.Param.Address.CityId,
		District: d.Param.Address.District,
		Province: d.Param.Address.Province,
		Detail: d.Param.Detail,
		Uid: d.Uid,
		Phone: d.Param.Phone,
		PostCode: d.Param.PostCode,
		RealName: d.Param.RealName,
	}
	if d.Param.IsDefault {
		userAddress.IsDefault = 1
		err = tx.Model(&models.YshopUserAddress{}).
			Where("uid = ?",d.Uid).Update("is_default",0).Error
		if err != nil {
			global.YSHOP_LOG.Error(err)
			return  0,errors.New("operation failed")
		}
	}
	if d.Param.Id == 0 {
		err = models.AddUserAddress(userAddress)
		if err != nil {
			global.YSHOP_LOG.Error(err)
			return  0,errors.New("operation failed")
		}
	}else {
		err = tx.Model(&models.YshopUserAddress{}).
			Where("id = ?",d.Param.Id).
			Updates(userAddress).Error
		if err != nil {
			global.YSHOP_LOG.Error(err)
			return  0,errors.New("operation failed")
		}
	}
	return userAddress.Id,nil
}


//get city list
func (d *Address) GetCitys() []models.YshopSystemCity {
	key := constant.CITY_LIST
	if b,err := redis.Get(key);err == nil {
		var city []models.YshopSystemCity
		err = json.Unmarshal(b, &city)
		return city
	}
	maps := make(map[string]interface{})
	maps["is_show"] = 1
	list := models.GetAllSystemCity(maps)
	redis.Set(key,list,0)
	return list
}

func (d *Address) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Name != "" {
		maps["name"] = d.Name
	}

	total,list := models.GetAllUserAddress(d.PageNum,d.PageSize,maps)
	return vo.ResultList{Content: list,TotalElements: total}
}


func (d *Address) Insert() error {
	return models.AddUserAddress(d.M)
}

func (d *Address) Save() error {
	return models.UpdateByUserAddress(d.M)
}

func (d *Address) Del() error {
	return models.DelByUserAddress(d.Ids)
}