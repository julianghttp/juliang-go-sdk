package common

//动态代理 -- 获取代理IP白名单
type DynamicGetWhiteIp struct {
	key      string //密钥
	trade_no string //业务编号
}

func (d *DynamicGetWhiteIp) Key() string {
	return d.key
}

func (d *DynamicGetWhiteIp) SetKey(key string) {
	d.key = key
}

func (d *DynamicGetWhiteIp) Trade_no() string {
	return d.trade_no
}

func (d *DynamicGetWhiteIp) SetTrade_no(trade_no string) {
	d.trade_no = trade_no
}
