package common

//动态代理 -- 设置代理IP白名单
type DynamicSetWhiteIp struct {
	key      string //密钥
	trade_no string //业务编号
	ips      string //IP列表
}

func (d *DynamicSetWhiteIp) Key() string {
	return d.key
}

func (d *DynamicSetWhiteIp) SetKey(key string) {
	d.key = key
}

func (d *DynamicSetWhiteIp) Trade_no() string {
	return d.trade_no
}

func (d *DynamicSetWhiteIp) SetTrade_no(trade_no string) {
	d.trade_no = trade_no
}

func (d *DynamicSetWhiteIp) Ips() string {
	return d.ips
}

func (d *DynamicSetWhiteIp) SetIps(ips string) {
	d.ips = ips
}
