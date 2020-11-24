package v2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//3.4订单查询接口   http://*:*/Order/QueryV2（请从供货商处获取）

func (c *Client) OrderQueryV2(mOrderId string) (result *OrderQueryV2Result, err error) {
	timestamp := Timestamp()
	sign := Md5Sign(c.Cfg.AppKey + timestamp + mOrderId + c.Cfg.AppSecret)
	err = c.Request("/Order/QueryV2",
		fmt.Sprintf("AppKey=%s&TimesTamp=%s&Sign=%s&MOrderID=%s&OrderID=", c.Cfg.AppKey, timestamp, sign, mOrderId), &result)
	if err != nil {
		return
	}
	if result.Code != 999 {
		err = errors.New(strconv.FormatInt(result.Code, 10) + "->" + result.Msg)
		return
	}
	if result.Sign != Md5Sign(c.Cfg.AppKey+strconv.FormatInt(result.TimesTamp, 10)+strconv.FormatInt(result.Code, 10)+strconv.FormatInt(int64(result.Data.OrderState), 10)+c.Cfg.AppSecret) {
		err = RES_SIGN_ERROR
		return
	}

	if result != nil && result.Data != nil && result.Data.ExtendParam != nil {
		if result.Data.ExtendParam.ChannelSerialNumber != "" {
			result.Data.ExtendParam.ChannelSerialNumber = strings.TrimSpace(RsaDecrypt(result.Data.ExtendParam.ChannelSerialNumber, c.Cfg.RsaPriKey))
		}
		if result.Data.ExtendParam.CardPwd != "" {
			result.Data.ExtendParam.CardPwd = strings.TrimSpace(RsaDecrypt(result.Data.ExtendParam.CardPwd, c.Cfg.RsaPriKey))
		}
		if result.Data.ExtendParam.CardNumber != "" {
			result.Data.ExtendParam.CardNumber = strings.TrimSpace(RsaDecrypt(result.Data.ExtendParam.CardNumber, c.Cfg.RsaPriKey))
		}
	}

	// 通过判断订单状态做相应处理
	return
}

type OrderQueryV2Result struct {
	Code      int64
	Msg       string
	TimesTamp int64
	Sign      string
	Data      *OrderQueryV2ResultData
}
type OrderQueryV2ResultData struct {
	OrderID        int64
	MOrderID       string
	OrderState     OrderStatus
	ChargeAccount  string
	BuyCount       int64
	Price          int64
	SellDebitAmout int64
	SellRebate     int64
	CreateTime     string
	ExtendParam    *OrderQueryV2ResultDataExtendParam
}

type OrderQueryV2ResultDataExtendParam struct {
	CardDeadline        string
	CardNumber          string
	CardPwd             string
	ChannelSerialNumber string
	FinishTime          string
	OfficialDes         string
	OfficialOrderID     string
}
