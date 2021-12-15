package common

//账户 -- 获取账户下对应类型的所有正常状态订单
type UsersGetAllOrders struct {
	key          string //密钥
	user_id      string //用户id
	product_type string //产品类型
	show         string //是否返回产品秘钥
}

func (u *UsersGetAllOrders) Key() string {
	return u.key
}

func (u *UsersGetAllOrders) SetKey(key string) {
	u.key = key
}

func (u *UsersGetAllOrders) User_id() string {
	return u.user_id
}

func (u *UsersGetAllOrders) SetUser_id(user_id string) {
	u.user_id = user_id
}

func (u *UsersGetAllOrders) Product_type() string {
	return u.product_type
}

func (u *UsersGetAllOrders) SetProduct_type(product_type string) {
	u.product_type = product_type
}

func (u *UsersGetAllOrders) Show() string {
	return u.show
}

func (u *UsersGetAllOrders) SetShow(show string) {
	u.show = show
}
