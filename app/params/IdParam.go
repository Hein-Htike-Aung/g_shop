/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package params

import (
	"github.com/astaxie/beego/validation"
)

type IdParam struct {
    Id    int64 `json:"id"`
}

func (p *IdParam) Valid(v *validation.Validation)  {
	if vv := v.Required(p.Id,"yshop-warning"); !vv.Ok {
		vv.Message("invalid parameters")
		return
	}
}

