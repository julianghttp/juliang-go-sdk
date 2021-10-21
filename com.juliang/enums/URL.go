package enums

const (
	//主站地址
	DOMAIN = "http://v1.api.juliangip.com"
	// USERS_GETBALANCE DOMAIN             = "http://192.168.10.60:8087"
	//USERS_GETBALANCE 获取账户余额
	USERS_GETBALANCE = DOMAIN + "/v1/users/getbalance"
	// DYNAMIC_GETIPS 动态代理 -- 提取动态代理
	DYNAMIC_GETIPS = DOMAIN + "/v1/dynamic/getips"
	// DYNAMIC_CHECK 动态代理 -- 校验IP可用性
	DYNAMIC_CHECK = DOMAIN + "/v1/dynamic/check"
	// DYNAMIC_SETWHITEIP 动态代理 -- 设置代理IP白名单
	DYNAMIC_SETWHITEIP = DOMAIN + "/v1/dynamic/setwhiteip"
	// DYNAMIC_GETWHITEIP 动态代理 -- 获取IP白名单
	DYNAMIC_GETWHITEIP = DOMAIN + "/v1/dynamic/getwhiteip"
	// DYNAMIC_REMAIN 动态代理 -- 获取代理剩余可用时长
	DYNAMIC_REMAIN = DOMAIN + "/v1/dynamic/remain"
	// DYNAMIC_BALANCE 动态代理 -- 获取剩余可用时长
	DYNAMIC_BALANCE = DOMAIN + "/v1/dynamic/balance"
	// ALONE_GETIPS 独享代理 -- 获取代理详情
	ALONE_GETIPS = DOMAIN + "/v1/alone/getips"
	// ALONE_SETWHITEIP 独享代理 -- 设置代理IP白名单
	ALONE_SETWHITEIP = DOMAIN + "/v1/alone/setwhiteip"
	// ALONE_GETWHITEIP 独享代理 -- 获取代理IP白名单
	ALONE_GETWHITEIP = DOMAIN + "/v1/alone/getwhiteip"
)
