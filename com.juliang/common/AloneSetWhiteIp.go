package common

//独享代理 -- 设置代理IP白名单
type AloneSetWhiteIp struct {
	key      string
	trade_no string
	ips      string
}

func (a *AloneSetWhiteIp) Key() string {
	return a.key
}

func (a *AloneSetWhiteIp) SetKey(key string) {
	a.key = key
}

func (a *AloneSetWhiteIp) Trade_no() string {
	return a.trade_no
}

func (a *AloneSetWhiteIp) SetTrade_no(trade_no string) {
	a.trade_no = trade_no
}

func (a *AloneSetWhiteIp) Ips() string {
	return a.ips
}

func (a *AloneSetWhiteIp) SetIps(ips string) {
	a.ips = ips
}
