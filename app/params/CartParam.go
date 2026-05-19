/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package params

import (
	"github.com/astaxie/beego/validation"
	"yixiang.co/go-mall/pkg/global"
)

type CartParam struct {
	ProductId    int64 `json:"productId"`
	UniqueId     string `json:"uniqueId"`
	CartNum  int    `json:"cartNum"`
	IsNew        int8 `json:"isNew"`
	CombinationId    int64 `json:"combinationId"`
	SeckillId     int64 `json:"seckillId"`
	BargainId  int64    `json:"bargainId"`
}

func (p *CartParam) Valid(v *validation.Validation)  {
	global.YSHOP_LOG.Info(p.CartNum)
	if vv := v.Range(p.CartNum,1,999,"cart quantity"); !vv.Ok {
		vv.Message("quantity must be between 1 and 999")
		return
	}
	if vv := v.Required(p.ProductId,"yshop-warning"); !vv.Ok {
		vv.Message("invalid parameters")
		return
	}
}

