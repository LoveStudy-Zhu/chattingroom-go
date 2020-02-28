package common

const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string `json:"type"`//消息类型
	Data string `json:"data"`//消息内容
}


//定义两个消息
type LoginMes struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

//登录返回消息
type LoginReMes struct {
	Code int `json:"code"`//返回状态码 500表示该用户未注册 200表示登录成功
	Error string `json:"error"`//返回错误信息
}

type RegisterMes struct {

}