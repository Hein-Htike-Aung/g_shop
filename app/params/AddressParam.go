/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package params

import "github.com/astaxie/beego/validation"

type AddressParan struct {
	Id        int64 `json:"id"`
	RealName    string `json:"real_name"`
	Phone     string `json:"phone"`
	Detail  string    `json:"detail"`
	PostCode        string `json:"post_code"`
	IsDefault  bool    `json:"is_default"`
	Address AddressDetailParan `json:"address"`
}

func (p *AddressParan) Valid(v *validation.Validation)  {
	if vv := v.MaxSize(p.RealName,30,"name"); !vv.Ok {
		vv.Message("name must be at most 30 characters")
		return
	}
	if vv := v.MaxSize(p.Detail,60,"name"); !vv.Ok {
		vv.Message("address detail must be at most 60 characters")
		return
	}
}


