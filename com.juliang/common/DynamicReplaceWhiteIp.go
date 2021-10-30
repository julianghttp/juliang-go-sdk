package common

//动态代理 -- 替换IP白名单
type DynamicReplaceWhiteIp struct {
	key      string //密钥
	trade_no string //业务编号
	old_ip   string //已存在,需要被替换的白名单IP
	new_ip   string //替换的新的白名单IP
	reset    string //是否重置已存在的白名单IP
}

func (d *DynamicReplaceWhiteIp) Key() string {
	return d.key
}

func (d *DynamicReplaceWhiteIp) SetKey(key string) {
	d.key = key
}

func (d *DynamicReplaceWhiteIp) Trade_no() string {
	return d.trade_no
}

func (d *DynamicReplaceWhiteIp) SetTrade_no(trade_no string) {
	d.trade_no = trade_no
}

func (d *DynamicReplaceWhiteIp) Old_ip() string {
	return d.old_ip
}

func (d *DynamicReplaceWhiteIp) SetOld_ip(old_ip string) {
	d.old_ip = old_ip
}

func (d *DynamicReplaceWhiteIp) New_ip() string {
	return d.new_ip
}

func (d *DynamicReplaceWhiteIp) SetNew_ip(new_ip string) {
	d.new_ip = new_ip
}

func (d *DynamicReplaceWhiteIp) Reset() string {
	return d.reset
}

func (d *DynamicReplaceWhiteIp) SetReset(reset string) {
	d.reset = reset
}
