package common

//独享代理 -- 获取独享代理详情
type AloneGetIps struct {
	key           string
	trade_no      string
	sock_port     string
	city_name     string
	city_code     string
	order_endtime string
}

func (a *AloneGetIps) Key() string {
	return a.key
}

func (a *AloneGetIps) SetKey(key string) {
	a.key = key
}

func (a *AloneGetIps) Trade_no() string {
	return a.trade_no
}

func (a *AloneGetIps) SetTrade_no(trade_no string) {
	a.trade_no = trade_no
}

func (a *AloneGetIps) Sock_port() string {
	return a.sock_port
}

func (a *AloneGetIps) SetSock_port(sock_port string) {
	a.sock_port = sock_port
}

func (a *AloneGetIps) City_name() string {
	return a.city_name
}

func (a *AloneGetIps) SetCity_name(city_name string) {
	a.city_name = city_name
}

func (a *AloneGetIps) City_code() string {
	return a.city_code
}

func (a *AloneGetIps) SetCity_code(city_code string) {
	a.city_code = city_code
}

func (a *AloneGetIps) Order_endtime() string {
	return a.order_endtime
}

func (a *AloneGetIps) SetOrder_endtime(order_endtime string) {
	a.order_endtime = order_endtime
}
