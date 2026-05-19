/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package params

import (
	"github.com/astaxie/beego/validation"
)

type CartNumParam struct {
    Id    int64 `json:"id"`
	Number  int    `json:"number"`

}

func (p *CartNumParam) Valid(v *validation.Validation)  {
	if vv := v.Range(p.Number,1,999,"cart quantity"); !vv.Ok {
		vv.Message("quantity must be between 1 and 999")
		return
	}
	if vv := v.Required(p.Id,"yshop-warning"); !vv.Ok {
		vv.Message("invalid parameters")
		return
	}
}

