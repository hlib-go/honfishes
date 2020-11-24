package v2

import (
	"net/http"
	"strconv"
)

//3.2订单回调功能
//功能描述：
//在订单有结果后，供货商系统对商户（下游会员）进行订单状态通知请求，请求的url为会员提交订单时填写的CallBackUrl参数。
func (c *Client) OrderCallback(req *http.Request) (r *OrderCallbackResult, err error) {
	state, e := strconv.ParseInt(req.FormValue("State"), 64, 10)
	if e != nil {
		state = 4
	}
	r = &OrderCallbackResult{
		AppKey:        req.FormValue("AppKey"),
		TimesTamp:     req.FormValue("TimesTamp"),
		Sign:          req.FormValue("Sign"),
		MOrderID:      req.FormValue("MOrderID"),
		OrderID:       req.FormValue("OrderID"),
		State:         OrderStatus(state),
		ChargeAccount: req.FormValue("ChargeAccount"),
		ProductCode:   req.FormValue("ProductCode"),
		BuyCount:      req.FormValue("BuyCount"),
		ExtendParam:   req.FormValue("ExtendParam"),
	}
	VerifySign := Md5Sign(c.Cfg.AppKey + r.TimesTamp + r.OrderID + r.MOrderID + strconv.FormatInt(int64(r.State), 10) + c.Cfg.AppSecret)
	if r.Sign != VerifySign {
		err = RES_SIGN_ERROR
		return
	}
	return
}

type OrderCallbackResult struct {
	AppKey        string
	TimesTamp     string
	Sign          string
	MOrderID      string
	OrderID       string
	State         OrderStatus
	ChargeAccount string
	ProductCode   string
	BuyCount      string
	ExtendParam   string
}
