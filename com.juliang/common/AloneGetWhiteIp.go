package common

//独享代理 -- 获取代理IP白名单
type AloneGetWhiteIp struct {
	key      string
	trade_no string
}

func (a *AloneGetWhiteIp) Key() string {
	return a.key
}

func (a *AloneGetWhiteIp) SetKey(key string) {
	a.key = key
}

func (a *AloneGetWhiteIp) Trade_no() string {
	return a.trade_no
}

func (a *AloneGetWhiteIp) SetTrade_no(trade_no string) {
	a.trade_no = trade_no
}
