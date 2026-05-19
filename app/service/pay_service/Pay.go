/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* Note: This software was developed by www.yixiang.co
 */
package pay_service

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/wechat"
	"github.com/unknwon/com"
	"gorm.io/gorm"
	"time"
	"yixiang.co/go-mall/app/models"
	"yixiang.co/go-mall/app/service/order_service"
	orderDto "yixiang.co/go-mall/app/service/order_service/dto"
	"yixiang.co/go-mall/app/service/wechat_user_service"
	orderEnum "yixiang.co/go-mall/pkg/enums/order"
	"yixiang.co/go-mall/pkg/global"
)

type Pay struct {}

// start payment
func GoPay(returnMap map[string]interface{}, orderId,payType,from string,
			uid int64 ,dto *orderDto.OrderExtend) (map[string]interface{},error) {

	dto = &orderDto.OrderExtend{

	}
	switch payType {
		case "weixin":
			if from == "pc" {
				client := wechat.NewClient("", "", "", true)

				// set country
				client.SetCountry(wechat.China)
				client.DebugSwitch = gopay.DebugOn
				orderService := order_service.Order {
					Uid: uid,
					OrderId: orderId,
				}
				orderInfo,_ := orderService.GetOrderInfo()
				//expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
				bm := make(gopay.BodyMap)
				bm.Set("nonce_str", util.RandomString(32)).
					Set("body", "yshop-go PC payment").
					Set("out_trade_no", orderId).
					Set("total_fee", orderInfo.PayPrice * 100).
					Set("spbill_create_ip", "127.0.0.1").
					Set("notify_url", "https://www.yixiang.co/notify").
					Set("trade_type", wechat.TradeType_Native).
					Set("device_info", "WEB").
					Set("sign_type", wechat.SignType_MD5)
				var ctx = context.Background()
				wxRsp, err := client.UnifiedOrder(ctx, bm)
				if err != nil {
					global.YSHOP_LOG.Error(err)
					return nil,err
				}
				global.YSHOP_LOG.Info(wxRsp)

				jsConfig := gin.H{"codeUrl" : wxRsp.CodeUrl}
				dto.JsConfig = jsConfig
				returnMap["payMsg"] = "PC payment successful"
				returnMap["result"] = dto

			}
		case "yue":
			err := YuePay(orderId,uid)
			if err != nil {
				global.YSHOP_LOG.Error(err)
				return nil,err
			}
			returnMap["payMsg"] = "balance payment successful"

	}

	return returnMap,nil
}

//balance payment
func  YuePay(orderId string,uid int64) error  {
	var err error
	tx := global.YSHOP_DB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}else{
			tx.Commit()
		}
	}()
	orderService := order_service.Order {
		Uid: uid,
	    OrderId: orderId,
	}
	orderInfo,_ := orderService.GetOrderInfo()
	if orderInfo.Paid == 1 {
		return errors.New("order already paid")
	}
	userService := wechat_user_service.User{
		Id: uid,
	}
	userInfo := userService.GetUserInfo()
	global.YSHOP_LOG.Info(userInfo.NowMoney,orderInfo.PayPrice)
	if userInfo.NowMoney < orderInfo.PayPrice {
		return errors.New("insufficient balance")
	}
	err = tx.Exec("update yshop_user set now_money=now_money - ?" +
		" where id = ?",orderInfo.PayPrice,uid).Error
	if err != nil {
		global.YSHOP_LOG.Error(err)
		return errors.New("balance payment failed (0001)")
	}
	err = PaySuccess(tx,orderInfo.OrderId,"yue")
	if err != nil {
		global.YSHOP_LOG.Error(err)
		return errors.New("balance payment failed (0002)")
	}
	return nil
}

// handle payment success
func PaySuccess(tx *gorm.DB,orderId,payType string) error {
	var err error

	orderService := order_service.Order {
		OrderId: orderId,
	}
	orderInfo,_ := orderService.GetOrderInfo()

	// update order status
	updateOrder := &models.YshopStoreOrder{
		Paid: orderEnum.PAY_STATUS_1,
		PayType: payType,
		PayTime: time.Now(),
	}
	err = tx.Model(&models.YshopStoreOrder{}).Where("order_id = ?",orderId).Updates(updateOrder).Error

	if err != nil {
		global.YSHOP_LOG.Error(err)
		return err
	}

	// increment user purchase count
	err = tx.Exec("update yshop_user set pay_count = pay_count + 1" +
		" where id = ?",orderInfo.Uid).Error
	if err != nil {
		global.YSHOP_LOG.Error(err)
		return err
	}
	// add order status log
	err = models.AddStoreOrderStatus(tx,orderInfo.Id,"pay_success","user payment successful")
	if err != nil {
		global.YSHOP_LOG.Error(err)
		return err
	}

	userServie := wechat_user_service.User{Id: orderInfo.Uid}
	userInfo := userServie.GetUserInfo()
	payTypeMsg := "WeChat Pay"
	if payType == "yue" {
		payTypeMsg = "balance payment"
	}


	mark := payTypeMsg + com.ToStr(orderInfo.PayPrice) + " CNY for product purchase"
	err = models.Expend(tx,orderInfo.Uid,"purchase product","now_money","pay_product",
		mark,com.ToStr(orderInfo.Id),orderInfo.PayPrice,userInfo.NowMoney)
	if err != nil {
		global.YSHOP_LOG.Error(err)
		return err
	}

	return nil
	// TODO: notification
}

