package common

const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResMesType ="RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType  = "SmsMes"
)

//这里定义几个用户状态的常量
const(
	UserOnline = iota
	UserOffLine
	UserBusyStatus
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
	UserIds [] int		//增加一个用户的切片
	Error string `json:"error"`//返回错误信息
}

type RegisterMes struct {
	User User `json:"user"`
}

type  RegisterResMes struct {
	Code int `json:"code"`//返回状态码
	Error string `json:"error"`//返回错误信息
}


//为了配合服务器推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`//用户id
	Status int `json:"status"`//用户状态
}


type SmsMes struct {
	Content string `json:"content"`//消息内容
	User //匿名结构体


}