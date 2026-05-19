/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package models


type YshopShippingTemplates struct {
	Name    string `json:"name"`
	Type     int8 `json:"type"`
	RegionInfo     string `json:"region_info"`
	Appoint  int8    `json:"appoint"`
	AppointInfo    string `json:"appoint_info"`
	Sort     int8 `json:"sort"`
	BaseModel
}

func (YshopShippingTemplates) TableName() string {
	return "yshop_shipping_templates"
}



func AddShippingTemplates(m *YshopShippingTemplates) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByShippingTemplates(m *YshopShippingTemplates) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByShippingTemplatess(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&YshopShippingTemplates{}).Error
	if err != nil {
		return err
	}

	return err
}
