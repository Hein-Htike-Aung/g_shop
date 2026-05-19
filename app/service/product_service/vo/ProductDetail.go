/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package vo

import (
	"yixiang.co/go-mall/app/models"
)

type ProductDetail struct {
	ProductAttr  []ProductAttr                                `json:"productAttr"`
	ProductValue map[string]models.YshopStoreProductAttrValue `json:"productValue"`
	Reply        models.YshopStoreProductReply                `json:"reply"`
	ReplyChance  string                                       `json:"replyChance"`
	ReplyCount   string                                       `json:"replyCount"`
	StoreInfo    Product                                      `json:"storeInfo"`
	Uid          int64                                        `json:"uid"`
	TempName     string                                       `json:"tempName"`
}



