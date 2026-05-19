/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package models

import "gorm.io/gorm"

type YshopUserBill struct {
	Uid        int64 `json:"uid"`
	LinkId    string `json:"link_id"`
	Pm     int8 `json:"pm"`
	Title  string    `json:"title"`
	Category        string `json:"category"`
	Type    string `json:"type"`
	Number     float64 `json:"number"`
	Balance  float64    `json:"balance"`
	Mark        string `json:"mark"`
	Status    int8 `json:"status"`
	BaseModel
}

func (YshopUserBill) TableName() string {
	return "yshop_user_bill"
}


// add expense ledger entry
func Expend(tx *gorm.DB,uid int64,title,category,typestr,mark,linkId string,number,balance float64) error {
	data := &YshopUserBill{
		Uid: uid,
		Title: title,
		Category: category,
		Type: typestr,
		Number: number,
		Balance: balance,
		Mark: mark,
		Pm: 0,
		LinkId: linkId,
	}
	return tx.Model(&YshopUserBill{}).Create(data).Error
}

// add income ledger entry
func Income(tx *gorm.DB,uid int64,title,category,typestr,mark,linkId string,number,balance float64) error {
	data := &YshopUserBill{
		Uid: uid,
		Title: title,
		Category: category,
		Type: typestr,
		Number: number,
		Balance: balance,
		Mark: mark,
		Pm: 1,
		LinkId: linkId,
	}
	return tx.Model(&YshopUserBill{}).Create(data).Error
}

func AddUserBill(m *YshopUserBill) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}

	return err
}

func UpdateByUserBill(m *YshopUserBill) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}

	return err
}

func DelByUserBill(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&YshopUserBill{}).Error
	if err != nil {
		return err
	}

	return err
}
