package common

//动态代理 -- 校验代理有效性
type DynamicCheck struct {
	key      string
	trade_no string
	proxy    string
}

func (d *DynamicCheck) Key() string {
	return d.key
}

func (d *DynamicCheck) SetKey(key string) {
	d.key = key
}

func (d *DynamicCheck) Trade_no() string {
	return d.trade_no
}

func (d *DynamicCheck) SetTrade_no(trade_no string) {
	d.trade_no = trade_no
}

func (d *DynamicCheck) Proxy() string {
	return d.proxy
}

func (d *DynamicCheck) SetProxy(proxy string) {
	d.proxy = proxy
}
