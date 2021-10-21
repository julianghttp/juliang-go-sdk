package common

//账户 -- 获取账户余额
type UsersGetBalance struct {
	key     string
	user_id string
}

func (u *UsersGetBalance) Key() string {
	return u.key
}

func (u *UsersGetBalance) SetKey(key string) {
	u.key = key
}

func (u *UsersGetBalance) User_id() string {
	return u.user_id
}

func (u *UsersGetBalance) SetUser_id(user_id string) {
	u.user_id = user_id
}
