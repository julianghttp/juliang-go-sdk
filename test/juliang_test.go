package test

import (
	"fmt"
	"gitee.com/juliangip/juliang-go-sdk/com.juliang"
	"gitee.com/juliangip/juliang-go-sdk/com.juliang/common"
	"testing"
)

const (
	dynamic_trade_no  = "11351242352345234523453465735679"
	dynamic_trade_key = "0794dsfdsgsdgw3tg45t324502bb21cbcd"

	user_id  = "1215104"
	user_key = "514a5dfgsdgshwry45gergwerg2817a866328c"

	alone_trade_no = "405856456453623452345135232429"
	alone_key      = "7cc5a5ba63sgeuhfr9u2hihr30u3hq9ef84a82f1"
)

func TestLanguage_aloneReplaceWhiteIp(t *testing.T) {
	var whiteIp common.AloneReplaceWhiteIp
	whiteIp.SetTrade_no(alone_trade_no)
	whiteIp.SetKey(alone_key)
	whiteIp.SetOld_ip("")
	whiteIp.SetNew_ip("6.6.6.6,4.4.4.4,7.7.7.7")
	value := juliang.AloneReplaceWhiteIp(whiteIp)
	fmt.Println(value)
}

//独享代理 -- 获取代理IP白名单
func TestLanguage_aloneGetWhiteIp(t *testing.T) {
	var whiteIp common.AloneGetWhiteIp
	whiteIp.SetTrade_no(alone_trade_no)
	whiteIp.SetKey(alone_key)
	value := juliang.AloneGetWhiteIp(whiteIp)
	fmt.Println(value)
}

//独享代理 -- 设置代理IP白名单
func TestLanguage_aloneSetWhiteIp(t *testing.T) {
	var whiteIp common.AloneSetWhiteIp
	whiteIp.SetTrade_no(alone_trade_no)
	whiteIp.SetKey(alone_key)
	whiteIp.SetIps("10.10.10.10")
	value := juliang.AloneSetWhiteIp(whiteIp)
	fmt.Println(value)
}

//独享代理 -- 获取独享代理详情
func TestLanguage_aloneGetIps(t *testing.T) {
	var ips common.AloneGetIps
	ips.SetTrade_no(alone_trade_no)
	ips.SetKey(alone_key)
	value := juliang.AloneGetIps(ips)
	fmt.Println(value)
}

//账户 -- 获取账户余额
func TestLanguage_usersGetBalance(t *testing.T) {
	var balance common.UsersGetBalance
	balance.SetUser_id(user_id)
	balance.SetKey(user_key)
	value := juliang.UsersGetBalance(balance)
	fmt.Println(value)
}

//动态代理 -- 获取剩余可提取IP数量
func TestLanguage_dynamicBalance(t *testing.T) {
	var balance common.DynamicBalance
	balance.SetKey(dynamic_trade_key)
	balance.SetTrade_no(dynamic_trade_no)
	value := juliang.DynamicBalance(balance)
	fmt.Println(value)
}

//动态代理 -- 获取代理剩余可用时长
func TestLanguage_dynamicRemain(t *testing.T) {
	var remain common.DynamicRemain
	remain.SetKey(dynamic_trade_key)
	remain.SetTrade_no(dynamic_trade_no)
	remain.SetProxy("1.1.1.1:8088")
	value := juliang.DynamicRemain(remain)
	fmt.Println(value)
}

//动态代理 -- 获取代理IP白名单
func TestLanguage_dynamicGetWhiteIp(t *testing.T) {
	var getWhiteIp common.DynamicGetWhiteIp
	getWhiteIp.SetKey(dynamic_trade_key)
	getWhiteIp.SetTrade_no(dynamic_trade_no)
	value := juliang.DynamicGetWhiteIp(getWhiteIp)
	fmt.Println(value)
}

func TestLanguage_dynamicReplaceWhiteIp(t *testing.T) {
	var ip common.DynamicReplaceWhiteIp
	ip.SetTrade_no(dynamic_trade_no)
	ip.SetKey(dynamic_trade_key)
	//ip.SetOld_ip("10.10.10.10")
	ip.SetNew_ip("20.20.20.20,255.255.255.255")
	ip.SetReset("1")
	value := juliang.DynamicReplaceWhiteIp(ip)
	fmt.Println(value)
}

//动态代理 -- 设置代理IP白名单
func TestLanguage_dynamicSetWhiteIp(t *testing.T) {
	var setWhiteIp common.DynamicSetWhiteIp
	setWhiteIp.SetKey(dynamic_trade_key)
	setWhiteIp.SetTrade_no(dynamic_trade_no)
	setWhiteIp.SetIps("12.12.12.12")
	value := juliang.DynamicSetWhiteIp(setWhiteIp)
	fmt.Println(value)
}

//动态代理 -- 校验代理有效性
func TestLanguage_dynamicCheck(t *testing.T) {
	var dynamicCheck common.DynamicCheck
	dynamicCheck.SetKey(dynamic_trade_key)
	dynamicCheck.SetTrade_no(dynamic_trade_no)
	dynamicCheck.SetProxy("10.10.10.10:8088")
	value := juliang.DynamicCheck(dynamicCheck)
	fmt.Println(value)
}

//动态代理 -- 提取动态代理
func TestLanguage_dynamicGetIps(t *testing.T) {
	var dynamicGetIps common.DynamicGetIps
	dynamicGetIps.SetKey(dynamic_trade_key)
	dynamicGetIps.SetTrade_no(dynamic_trade_no)
	dynamicGetIps.SetNum("10")
	value := juliang.DynamicGetIps(dynamicGetIps)
	fmt.Println(value)
}
