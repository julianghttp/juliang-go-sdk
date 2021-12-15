package common

type UsersGetCity struct {
	key      string //秘钥
	user_id  string //用户id
	province string //省份名称
}

func (u *UsersGetCity) Key() string {
	return u.key
}
func (u *UsersGetCity) SetKey(key string) {
	u.key = key
}
func (u *UsersGetCity) User_id() string {
	return u.user_id
}
func (u *UsersGetCity) SetUser_id(user_id string) {
	u.user_id = user_id
}
func (u *UsersGetCity) Province() string {
	return u.province
}
func (u *UsersGetCity) SetProvince(province string) {
	u.province = province
}
