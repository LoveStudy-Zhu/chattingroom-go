package common
type User struct {
	//序列化反序列化成功
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
	UserStatus int `json:"userStatus"`
	Sex string `json:"sex"`
}
