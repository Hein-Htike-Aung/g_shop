/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package vo

type Compute struct {
	CouponPrice    float64 `json:"couponPrice"`
	DeductionPrice  float64                        `json:"deductionPrice"`
	PayPostage        float64              `json:"payPostage"`
	PayPrice    float64                  `json:"payPrice"`
	TotalPrice     float64                      `json:"totalPrice"`
	UseIntegral int                       `json:"useIntegral"`
	PayIntegral int                         `json:"payIntegral"`
}

