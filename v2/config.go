package v2

/*
前置资料请从供货商相关业务负责人获取
URL（请求地址）
MemberAmountCode（扣款账户）
AppKey（账号标识）
AppSecret（账号密匙）
ProductCode（产品编号）

*/

// 码商SDK约定名称，不能修改
const NAME = "onfishes"

type Config struct {
	BaseServiceUrl string
	AppKey         string
	AppSecret      string
	RsaPriKey      string // 证书私钥，用于解密
	RsaPubKey      string // 证书公钥，提供给娱尚加密卡密
}