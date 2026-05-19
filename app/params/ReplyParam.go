/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package params

type ReplyParam struct {
	Comment    string `json:"comment"`
	Pics     string `json:"pics"`
	ProductScore  int    `json:"productScore"`
	ServiceScore        int `json:"serviceScore"`
	Unique    string `json:"unique"`
	ProductId    int64 `json:"productId"`

}



