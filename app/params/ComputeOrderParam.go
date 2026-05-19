/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package params

type ComputeOrderParam struct {
    AddressId    int64 `json:"addressId"`
	CouponId    int64 `json:"couponId"`
	PayType    int `json:"payType"`
	UseIntegral    int `json:"useIntegral"`
	ShippingType    int `json:"shippingType"`
	BargainId    int64 `json:"bargainId"`
	PinkId    int64 `json:"pinkId"`
	CombinationId    int64 `json:"combinationId"`
}


