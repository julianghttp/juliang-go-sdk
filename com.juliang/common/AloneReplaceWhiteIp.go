package common

//独享代理 -- 替换IP白名单
type AloneReplaceWhiteIp struct {
	key      string //密钥
	trade_no string //业务编号
	old_ip   string //已存在,需要被替换的白名单IP
	new_ip   string //替换的新的白名单IP
	reset    string //是否重置已存在的白名单IP
}

func (a *AloneReplaceWhiteIp) Key() string {
	return a.key
}

func (a *AloneReplaceWhiteIp) SetKey(key string) {
	a.key = key
}

func (a *AloneReplaceWhiteIp) Trade_no() string {
	return a.trade_no
}

func (a *AloneReplaceWhiteIp) SetTrade_no(trade_no string) {
	a.trade_no = trade_no
}

func (a *AloneReplaceWhiteIp) Old_ip() string {
	return a.old_ip
}

func (a *AloneReplaceWhiteIp) SetOld_ip(old_ip string) {
	a.old_ip = old_ip
}

func (a *AloneReplaceWhiteIp) New_ip() string {
	return a.new_ip
}

func (a *AloneReplaceWhiteIp) SetNew_ip(new_ip string) {
	a.new_ip = new_ip
}

func (a *AloneReplaceWhiteIp) Reset() string {
	return a.reset
}

func (a *AloneReplaceWhiteIp) SetReset(reset string) {
	a.reset = reset
}
