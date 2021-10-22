package test

import (
	"fmt"
	"gitee.com/juliangip/juliang-go-sdk/com.juliang"
	"gitee.com/juliangip/juliang-go-sdk/com.juliang/common"
	"testing"
)

//独享代理 -- 获取代理IP白名单
func TestLanguage_aloneGetWhiteIp(t *testing.T) {
	var whiteIp common.AloneGetWhiteIp
	whiteIp.SetTrade_no("405856456453623452345135232429")
	whiteIp.SetKey("7cc5a5ba63sgeuhfr9u2hihr30u3hq9ef84a82f1")
	value := juliang.AloneGetWhiteIp(whiteIp)
	fmt.Println(value)
}

//独享代理 -- 设置代理IP白名单
func TestLanguage_aloneSetWhiteIp(t *testing.T) {
	var whiteIp common.AloneSetWhiteIp
	whiteIp.SetTrade_no("405856456453623452345135232429")
	whiteIp.SetKey("7cc5a5ba63sgeuhfr9u2hihr30u3hq9ef84a82f1")
	whiteIp.SetIps("10.10.10.10")
	value := juliang.AloneSetWhiteIp(whiteIp)
	fmt.Println(value)
}

//独享代理 -- 获取独享代理详情
func TestLanguage_aloneGetIps(t *testing.T) {
	var ips common.AloneGetIps
	ips.SetTrade_no("405856456453623452345135232429")
	ips.SetKey("7cc5a5ba63sgeuhfr9u2hihr30u3hq9ef84a82f1")
	value := juliang.AloneGetIps(ips)
	fmt.Println(value)
}

//账户 -- 获取账户余额
func TestLanguage_usersGetBalance(t *testing.T) {
	var balance common.UsersGetBalance
	balance.SetUser_id("1215104")
	balance.SetKey("514a5dfgsdgshwry45gergwerg2817a866328c")
	value := juliang.UsersGetBalance(balance)
	fmt.Println(value)
}

//动态代理 -- 获取剩余可提取IP数量
func TestLanguage_dynamicBalance(t *testing.T) {
	var balance common.DynamicBalance
	balance.SetKey("0794dsfdsgsdgw3tg45t324502bb21cbcd")
	balance.SetTrade_no("11351242352345234523453465735679")
	value := juliang.DynamicBalance(balance)
	fmt.Println(value)
}

//动态代理 -- 获取代理剩余可用时长
func TestLanguage_dynamicRemain(t *testing.T) {
	var remain common.DynamicRemain
	remain.SetKey("0794dsfdsgsdgw3tg45t324502bb21cbcd")
	remain.SetTrade_no("11351242352345234523453465735679")
	remain.SetProxy("1.1.1.1:8088")
	value := juliang.DynamicRemain(remain)
	fmt.Println(value)
}

//动态代理 -- 获取代理IP白名单
func TestLanguage_dynamicGetWhiteIp(t *testing.T) {
	var getWhiteIp common.DynamicGetWhiteIp
	getWhiteIp.SetKey("0794dsfdsgsdgw3tg45t324502bb21cbcd")
	getWhiteIp.SetTrade_no("11351242352345234523453465735679")
	value := juliang.DynamicGetWhiteIp(getWhiteIp)
	fmt.Println(value)
}

//动态代理 -- 设置代理IP白名单
func TestLanguage_dynamicSetWhiteIp(t *testing.T) {
	var setWhiteIp common.DynamicSetWhiteIp
	setWhiteIp.SetTrade_no("11351242352345234523453465735679")
	setWhiteIp.SetKey("0794dsfdsgsdgw3tg45t324502bb21cbcd")
	setWhiteIp.SetIps("12.12.12.12")
	value := juliang.DynamicSetWhiteIp(setWhiteIp)
	fmt.Println(value)
}

//动态代理 -- 校验代理有效性
func TestLanguage_dynamicCheck(t *testing.T) {
	var dynamicCheck common.DynamicCheck
	dynamicCheck.SetKey("0794dsfdsgsdgw3tg45t324502bb21cbcd")
	dynamicCheck.SetTrade_no("11351242352345234523453465735679")
	dynamicCheck.SetProxy("10.10.10.10:8088")
	value := juliang.DynamicCheck(dynamicCheck)
	fmt.Println(value)
}

//动态代理 -- 提取动态代理
func TestLanguage_dynamicGetIps(t *testing.T) {
	var dynamicGetIps common.DynamicGetIps
	dynamicGetIps.SetKey("0794dsfdsgsdgw3tg45t324502bb21cbcd")
	dynamicGetIps.SetTrade_no("11351242352345234523453465735679")
	dynamicGetIps.SetNum("10")
	value := juliang.DynamicGetIps(dynamicGetIps)
	fmt.Println(value)
}
