package common

//动态代理 -- 获取代理剩余可用时长
type DynamicRemain struct {
	key      string
	trade_no string
	proxy    string
}

func (d *DynamicRemain) Key() string {
	return d.key
}

func (d *DynamicRemain) SetKey(key string) {
	d.key = key
}

func (d *DynamicRemain) Trade_no() string {
	return d.trade_no
}

func (d *DynamicRemain) SetTrade_no(trade_no string) {
	d.trade_no = trade_no
}

func (d *DynamicRemain) Proxy() string {
	return d.proxy
}

func (d *DynamicRemain) SetProxy(proxy string) {
	d.proxy = proxy
}
