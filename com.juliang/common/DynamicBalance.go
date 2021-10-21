package common

//动态代理 -- 获取剩余可提取IP数量
type DynamicBalance struct {
	key      string
	trade_no string
}

func (d *DynamicBalance) Key() string {
	return d.key
}

func (d *DynamicBalance) SetKey(key string) {
	d.key = key
}

func (d *DynamicBalance) Trade_no() string {
	return d.trade_no
}

func (d *DynamicBalance) SetTrade_no(trade_no string) {
	d.trade_no = trade_no
}
