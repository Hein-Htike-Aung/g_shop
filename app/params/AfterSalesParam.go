/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package params

import "github.com/astaxie/beego/validation"


type AfterSalesParan struct {
	OrderCode    string `json:"orderCode"`
	ServiceType     int `json:"serviceType"`
	ReasonForApplication  string    `json:"reasonForApplication"`
	ApplicationInstructions        string `json:"applicationInstructions"`
	ApplicationDescriptionPicture  string    `json:"applicationDescriptionPicture"`
	ProductParamList []ProductParam `json:"productParamList"`
}

type ProductParam struct {
	ProductId int64 `json:"productId"`
}

func (p *AfterSalesParan) Valid(v *validation.Validation)  {
	if vv := v.Required(p.OrderCode,"yshop-warning"); !vv.Ok {
		vv.Message("invalid order number")
		return
	}
	if vv := v.Required(p.ServiceType,"yshop-warning"); !vv.Ok {
		vv.Message("please select service type")
		return
	}
	if vv := v.Required(p.ProductParamList,"yshop-warning"); !vv.Ok {
		vv.Message("please select items to return")
		return
	}
}


