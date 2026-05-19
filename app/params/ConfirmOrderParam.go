/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package params

import (
	"github.com/astaxie/beego/validation"
)

type ConfirmOrderParam struct {
    CartId    string `json:"cartId"`
}

func (p *ConfirmOrderParam) Valid(v *validation.Validation)  {
	if vv := v.Required(p.CartId,"yshop-warning"); !vv.Ok {
		vv.Message("please submit items to purchase")
		return
	}
}

