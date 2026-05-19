/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package dto

type YshopUser struct {
	Id       int64   `json:"id"`
	RealName string  `json:"real_name"`
	Mark     string  `json:"mark"`
	Phone    string  `json:"phone"`
	Integral int `json:"integral"`
}
