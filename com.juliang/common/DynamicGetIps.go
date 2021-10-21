package common

//动态代理 -- 提取动态代理
type DynamicGetIps struct {
	key         string //密钥
	trade_no    string //业务编号
	num         string //提取数量
	pt          string //代理类型
	result_type string //返回类型
	split       string //结果分隔符
	city_name   string //地区名称
	city_code   string //邮政编码
	ip_remain   string //剩余可用时长
	area        string //筛选地区
	no_area     string //排除地区
	ip_seg      string //筛选IP段
	no_ip_seg   string //排除IP段
	isp         string //运营商筛选
	filter      string //IP去重
}

func (d *DynamicGetIps) Key() string {
	return d.key
}

func (d *DynamicGetIps) SetKey(key string) {
	d.key = key
}

func (d *DynamicGetIps) Trade_no() string {
	return d.trade_no
}

func (d *DynamicGetIps) SetTrade_no(trade_no string) {
	d.trade_no = trade_no
}

func (d *DynamicGetIps) Num() string {
	return d.num
}

func (d *DynamicGetIps) SetNum(num string) {
	d.num = num
}

func (d *DynamicGetIps) Pt() string {
	return d.pt
}

func (d *DynamicGetIps) SetPt(pt string) {
	d.pt = pt
}

func (d *DynamicGetIps) Result_type() string {
	return d.result_type
}

func (d *DynamicGetIps) SetResult_type(result_type string) {
	d.result_type = result_type
}

func (d *DynamicGetIps) Split() string {
	return d.split
}

func (d *DynamicGetIps) SetSplit(split string) {
	d.split = split
}

func (d *DynamicGetIps) City_name() string {
	return d.city_name
}

func (d *DynamicGetIps) SetCity_name(city_name string) {
	d.city_name = city_name
}

func (d *DynamicGetIps) City_code() string {
	return d.city_code
}

func (d *DynamicGetIps) SetCity_code(city_code string) {
	d.city_code = city_code
}

func (d *DynamicGetIps) Ip_remain() string {
	return d.ip_remain
}

func (d *DynamicGetIps) SetIp_remain(ip_remain string) {
	d.ip_remain = ip_remain
}

func (d *DynamicGetIps) Area() string {
	return d.area
}

func (d *DynamicGetIps) SetArea(area string) {
	d.area = area
}

func (d *DynamicGetIps) No_area() string {
	return d.no_area
}

func (d *DynamicGetIps) SetNo_area(no_area string) {
	d.no_area = no_area
}

func (d *DynamicGetIps) Ip_seg() string {
	return d.ip_seg
}

func (d *DynamicGetIps) SetIp_seg(ip_seg string) {
	d.ip_seg = ip_seg
}

func (d *DynamicGetIps) No_ip_seg() string {
	return d.no_ip_seg
}

func (d *DynamicGetIps) SetNo_ip_seg(no_ip_seg string) {
	d.no_ip_seg = no_ip_seg
}

func (d *DynamicGetIps) Isp() string {
	return d.isp
}

func (d *DynamicGetIps) SetIsp(isp string) {
	d.isp = isp
}

func (d *DynamicGetIps) Filter() string {
	return d.filter
}

func (d *DynamicGetIps) SetFilter(filter string) {
	d.filter = filter
}
